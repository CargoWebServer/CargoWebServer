//+build BPMN
package BPMS

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/BPMN20"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/BPMNDI"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DC"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DI"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"golang.org/x/net/html/charset"
)

type BPMSXmlFactory struct {
	m_references map[string]interface{}
	m_object     map[string]map[string][]string
}

/** Initialization function from xml file **/
func (this *BPMSXmlFactory) InitXml(inputPath string, object *BPMN20.Definitions) error {
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
	var xmlElement *BPMN20.XsdDefinitions
	xmlElement = new(BPMN20.XsdDefinitions)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(xmlElement); err != nil {
		return err
	}
	this.InitDefinitions(xmlElement, object)
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
func (this *BPMSXmlFactory) SerializeXml(outputPath string, toSerialize *BPMN20.Definitions) error {
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
	var xmlElement *BPMN20.XsdDefinitions
	xmlElement = new(BPMN20.XsdDefinitions)

	this.SerialyzeDefinitions(xmlElement, toSerialize)
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

/** inititialisation of CorrelationKey **/
func (this *BPMSXmlFactory) InitCorrelationKey(xmlElement *BPMN20.XsdCorrelationKey, object *BPMN20.CorrelationKey) {
	log.Println("Initialize CorrelationKey")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CorrelationKey%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CorrelationKey **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref correlationPropertyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_correlationPropertyRef); i++ {
		if _, ok := this.m_object[object.M_id]["correlationPropertyRef"]; !ok {
			this.m_object[object.M_id]["correlationPropertyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["correlationPropertyRef"] = append(this.m_object[object.M_id]["correlationPropertyRef"], xmlElement.M_correlationPropertyRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of Rendering **/
func (this *BPMSXmlFactory) InitRendering(xmlElement *BPMN20.XsdRendering, object *BPMN20.Rendering) {
	log.Println("Initialize Rendering")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Rendering%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Rendering **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of CorrelationPropertyRetrievalExpression **/
func (this *BPMSXmlFactory) InitCorrelationPropertyRetrievalExpression(xmlElement *BPMN20.XsdCorrelationPropertyRetrievalExpression, object *BPMN20.CorrelationPropertyRetrievalExpression) {
	log.Println("Initialize CorrelationPropertyRetrievalExpression")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CorrelationPropertyRetrievalExpression%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CorrelationPropertyRetrievalExpression **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if object.M_messagePath == nil {
		object.M_messagePath = new(BPMN20.FormalExpression)
	}
	this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(&xmlElement.M_messagePath)), object.M_messagePath)

	/** association initialisation **/

	/** Init ref messageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_messageRef) > 0 {
		if _, ok := this.m_object[object.M_id]["messageRef"]; !ok {
			this.m_object[object.M_id]["messageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageRef"] = append(this.m_object[object.M_id]["messageRef"], xmlElement.M_messageRef)
	}
}

/** inititialisation of GlobalScriptTask **/
func (this *BPMSXmlFactory) InitGlobalScriptTask(xmlElement *BPMN20.XsdGlobalScriptTask, object *BPMN20.GlobalScriptTask) {
	log.Println("Initialize GlobalScriptTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalScriptTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalScriptTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref supportedInterfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supportedInterfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["supportedInterfaceRef"]; !ok {
			this.m_object[object.M_id]["supportedInterfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supportedInterfaceRef"] = append(this.m_object[object.M_id]["supportedInterfaceRef"], xmlElement.M_supportedInterfaceRef[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetCallableElementPtr(object)
	}

	/** Init ioBinding **/
	object.M_ioBinding = make([]*BPMN20.InputOutputBinding, 0)
	for i := 0; i < len(xmlElement.M_ioBinding); i++ {
		val := new(BPMN20.InputOutputBinding)
		this.InitInputOutputBinding(xmlElement.M_ioBinding[i], val)
		object.M_ioBinding = append(object.M_ioBinding, val)

		/** association initialisation **/
		val.SetCallableElementPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init script **/
	if xmlElement.M_script != nil {
		object.M_script = new(BPMN20.Script)
		this.InitScript(xmlElement.M_script, object.M_script)

		/** association initialisation **/
		object.M_script.SetGlobalScriptTaskPtr(object)
	}

	/** GlobalTask **/
	object.M_scriptLanguage = xmlElement.M_scriptLanguage
}

/** inititialisation of Script **/
func (this *BPMSXmlFactory) InitScript(xmlElement *BPMN20.XsdScript, object *BPMN20.Script) {
	log.Println("Initialize Script")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Script%" + Utility.RandomUUID()
	}
	object.M_script = xmlElement.M_script
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of EventBasedGateway **/
func (this *BPMSXmlFactory) InitEventBasedGateway(xmlElement *BPMN20.XsdEventBasedGateway, object *BPMN20.EventBasedGateway) {
	log.Println("Initialize EventBasedGateway")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.EventBasedGateway%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** EventBasedGateway **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** GatewayDirection **/
	if xmlElement.M_gatewayDirection == "##Unspecified" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	} else if xmlElement.M_gatewayDirection == "##Converging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Converging
	} else if xmlElement.M_gatewayDirection == "##Diverging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Diverging
	} else if xmlElement.M_gatewayDirection == "##Mixed" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Mixed
	} else {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	}

	/** Gateway **/
	object.M_instantiate = xmlElement.M_instantiate

	/** EventBasedGatewayType **/
	if xmlElement.M_eventGatewayType == "##Exclusive" {
		object.M_eventGatewayType = BPMN20.EventBasedGatewayType_Exclusive
	} else if xmlElement.M_eventGatewayType == "##Parallel" {
		object.M_eventGatewayType = BPMN20.EventBasedGatewayType_Parallel
	} else {
		object.M_eventGatewayType = BPMN20.EventBasedGatewayType_Exclusive
	}
}

/** inititialisation of Text **/
func (this *BPMSXmlFactory) InitText(xmlElement *BPMN20.XsdText, object *BPMN20.Text) {
	log.Println("Initialize Text")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Text%" + Utility.RandomUUID()
	}
	object.M_text = xmlElement.M_text
	this.m_references["Text"] = object
}

/** inititialisation of ConversationAssociation **/
func (this *BPMSXmlFactory) InitConversationAssociation(xmlElement *BPMN20.XsdConversationAssociation, object *BPMN20.ConversationAssociation) {
	log.Println("Initialize ConversationAssociation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ConversationAssociation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ConversationAssociation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref innerConversationNodeRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_innerConversationNodeRef) > 0 {
		if _, ok := this.m_object[object.M_id]["innerConversationNodeRef"]; !ok {
			this.m_object[object.M_id]["innerConversationNodeRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["innerConversationNodeRef"] = append(this.m_object[object.M_id]["innerConversationNodeRef"], xmlElement.M_innerConversationNodeRef)
	}

	/** Init ref outerConversationNodeRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_outerConversationNodeRef) > 0 {
		if _, ok := this.m_object[object.M_id]["outerConversationNodeRef"]; !ok {
			this.m_object[object.M_id]["outerConversationNodeRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outerConversationNodeRef"] = append(this.m_object[object.M_id]["outerConversationNodeRef"], xmlElement.M_outerConversationNodeRef)
	}
}

/** inititialisation of StartEvent **/
func (this *BPMSXmlFactory) InitStartEvent(xmlElement *BPMN20.XsdStartEvent, object *BPMN20.StartEvent) {
	log.Println("Initialize StartEvent")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.StartEvent%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** StartEvent **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetEventPtr(object)
	}

	/** Init dataOutput **/
	object.M_dataOutput = make([]*BPMN20.DataOutput, 0)
	for i := 0; i < len(xmlElement.M_dataOutput); i++ {
		val := new(BPMN20.DataOutput)
		this.InitDataOutput(xmlElement.M_dataOutput[i], val)
		object.M_dataOutput = append(object.M_dataOutput, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init outputSet **/
	if xmlElement.M_outputSet != nil {
		object.M_outputSet = new(BPMN20.OutputSet)
		this.InitOutputSet(xmlElement.M_outputSet, object.M_outputSet)

		/** association initialisation **/
		object.M_outputSet.SetCatchEventPtr(object)
	}

	/** Init cancelEventDefinition **/
	object.M_eventDefinition = make([]BPMN20.EventDefinition, 0)
	for i := 0; i < len(xmlElement.M_eventDefinition_0); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_eventDefinition_0[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_1); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_eventDefinition_1[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_2); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_eventDefinition_2[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_3); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_eventDefinition_3[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_4); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_eventDefinition_4[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_5); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_eventDefinition_5[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_6); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_eventDefinition_6[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_7); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_eventDefinition_7[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_8); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_eventDefinition_8[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_9); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_eventDefinition_9[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init ref eventDefinitionRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_eventDefinitionRef); i++ {
		if _, ok := this.m_object[object.M_id]["eventDefinitionRef"]; !ok {
			this.m_object[object.M_id]["eventDefinitionRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["eventDefinitionRef"] = append(this.m_object[object.M_id]["eventDefinitionRef"], xmlElement.M_eventDefinitionRef[i])
	}

	/** Event **/
	object.M_parallelMultiple = xmlElement.M_parallelMultiple

	/** CatchEvent **/
	object.M_isInterrupting = xmlElement.M_isInterrupting
}

/** inititialisation of Participant **/
func (this *BPMSXmlFactory) InitParticipant(xmlElement *BPMN20.XsdParticipant, object *BPMN20.Participant) {
	log.Println("Initialize Participant")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Participant%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Participant **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref interfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_interfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["interfaceRef"]; !ok {
			this.m_object[object.M_id]["interfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["interfaceRef"] = append(this.m_object[object.M_id]["interfaceRef"], xmlElement.M_interfaceRef[i])
	}

	/** Init ref endPointRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_endPointRef); i++ {
		if _, ok := this.m_object[object.M_id]["endPointRef"]; !ok {
			this.m_object[object.M_id]["endPointRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["endPointRef"] = append(this.m_object[object.M_id]["endPointRef"], xmlElement.M_endPointRef[i])
	}

	/** Init participantMultiplicity **/
	if xmlElement.M_participantMultiplicity != nil {
		object.M_participantMultiplicity = new(BPMN20.ParticipantMultiplicity)
		this.InitParticipantMultiplicity(xmlElement.M_participantMultiplicity, object.M_participantMultiplicity)

		/** association initialisation **/
		object.M_participantMultiplicity.SetParticipantPtr(object)
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref processRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_processRef) > 0 {
		if _, ok := this.m_object[object.M_id]["processRef"]; !ok {
			this.m_object[object.M_id]["processRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["processRef"] = append(this.m_object[object.M_id]["processRef"], xmlElement.M_processRef)
	}
}

/** inititialisation of Conversation **/
func (this *BPMSXmlFactory) InitConversation(xmlElement *BPMN20.XsdConversation, object *BPMN20.Conversation) {
	log.Println("Initialize Conversation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Conversation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Conversation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** Init ref messageFlowRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_messageFlowRef); i++ {
		if _, ok := this.m_object[object.M_id]["messageFlowRef"]; !ok {
			this.m_object[object.M_id]["messageFlowRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageFlowRef"] = append(this.m_object[object.M_id]["messageFlowRef"], xmlElement.M_messageFlowRef[i])
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetConversationNodePtr(object)
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of DataInput **/
func (this *BPMSXmlFactory) InitDataInput(xmlElement *BPMN20.XsdDataInput, object *BPMN20.DataInput) {
	log.Println("Initialize DataInput")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataInput%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataInput **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}

	/** BaseElement **/
	object.M_isCollection = xmlElement.M_isCollection
}

/** inititialisation of BPMNDiagram **/
func (this *BPMSXmlFactory) InitBPMNDiagram(xmlElement *BPMNDI.XsdBPMNDiagram, object *BPMNDI.BPMNDiagram) {
	log.Println("Initialize BPMNDiagram")
	if len(object.UUID) == 0 {
		object.UUID = "BPMNDI.BPMNDiagram%" + Utility.RandomUUID()
	}

	/** BPMNDiagram **/
	object.M_name = xmlElement.M_name

	/** BPMNDiagram **/
	object.M_documentation = xmlElement.M_documentation

	/** BPMNDiagram **/
	object.M_resolution = xmlElement.M_resolution

	/** BPMNDiagram **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init BPMNPlane **/
	if xmlElement.M_BPMNPlane != nil {
		object.M_BPMNPlane = new(BPMNDI.BPMNPlane)
		this.InitBPMNPlane(xmlElement.M_BPMNPlane, object.M_BPMNPlane)

		/** association initialisation **/
	}

	/** Init BPMNLabelStyle **/
	object.M_BPMNLabelStyle = make([]*BPMNDI.BPMNLabelStyle, 0)
	for i := 0; i < len(xmlElement.M_BPMNLabelStyle); i++ {
		val := new(BPMNDI.BPMNLabelStyle)
		this.InitBPMNLabelStyle(xmlElement.M_BPMNLabelStyle[i], val)
		object.M_BPMNLabelStyle = append(object.M_BPMNLabelStyle, val)

		/** association initialisation **/
	}
}

/** inititialisation of TextAnnotation **/
func (this *BPMSXmlFactory) InitTextAnnotation(xmlElement *BPMN20.XsdTextAnnotation, object *BPMN20.TextAnnotation) {
	log.Println("Initialize TextAnnotation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.TextAnnotation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** TextAnnotation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init text **/
	if xmlElement.M_text != nil {
		object.M_text = new(BPMN20.Text)
		this.InitText(xmlElement.M_text, object.M_text)

		/** association initialisation **/
		object.M_text.SetTextAnnotationPtr(object)
	}

	/** Artifact **/
	object.M_textFormat = xmlElement.M_textFormat
}

/** inititialisation of ErrorEventDefinition **/
func (this *BPMSXmlFactory) InitErrorEventDefinition(xmlElement *BPMN20.XsdErrorEventDefinition, object *BPMN20.ErrorEventDefinition) {
	log.Println("Initialize ErrorEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ErrorEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ErrorEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref errorRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_errorRef) > 0 {
		if _, ok := this.m_object[object.M_id]["errorRef"]; !ok {
			this.m_object[object.M_id]["errorRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["errorRef"] = append(this.m_object[object.M_id]["errorRef"], xmlElement.M_errorRef)
	}
}

/** inititialisation of GlobalTask **/
func (this *BPMSXmlFactory) InitGlobalTask(xmlElement *BPMN20.XsdGlobalTask, object *BPMN20.GlobalTask_impl) {
	log.Println("Initialize GlobalTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalTask_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref supportedInterfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supportedInterfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["supportedInterfaceRef"]; !ok {
			this.m_object[object.M_id]["supportedInterfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supportedInterfaceRef"] = append(this.m_object[object.M_id]["supportedInterfaceRef"], xmlElement.M_supportedInterfaceRef[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetCallableElementPtr(object)
	}

	/** Init ioBinding **/
	object.M_ioBinding = make([]*BPMN20.InputOutputBinding, 0)
	for i := 0; i < len(xmlElement.M_ioBinding); i++ {
		val := new(BPMN20.InputOutputBinding)
		this.InitInputOutputBinding(xmlElement.M_ioBinding[i], val)
		object.M_ioBinding = append(object.M_ioBinding, val)

		/** association initialisation **/
		val.SetCallableElementPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}
}

/** inititialisation of ComplexGateway **/
func (this *BPMSXmlFactory) InitComplexGateway(xmlElement *BPMN20.XsdComplexGateway, object *BPMN20.ComplexGateway) {
	log.Println("Initialize ComplexGateway")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ComplexGateway%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ComplexGateway **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** GatewayDirection **/
	if xmlElement.M_gatewayDirection == "##Unspecified" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	} else if xmlElement.M_gatewayDirection == "##Converging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Converging
	} else if xmlElement.M_gatewayDirection == "##Diverging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Diverging
	} else if xmlElement.M_gatewayDirection == "##Mixed" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Mixed
	} else {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	}

	/** Init formalExpression **/
	if xmlElement.M_activationCondition != nil {
		if object.M_activationCondition == nil {
			object.M_activationCondition = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_activationCondition)), object.M_activationCondition)

		/** association initialisation **/
		object.M_activationCondition.SetComplexGatewayPtr(object)
	}

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
}

/** inititialisation of DataObject **/
func (this *BPMSXmlFactory) InitDataObject(xmlElement *BPMN20.XsdDataObject, object *BPMN20.DataObject) {
	log.Println("Initialize DataObject")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataObject%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataObject **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}

	/** FlowElement **/
	object.M_isCollection = xmlElement.M_isCollection
}

/** inititialisation of BPMNLabelStyle **/
func (this *BPMSXmlFactory) InitBPMNLabelStyle(xmlElement *BPMNDI.XsdBPMNLabelStyle, object *BPMNDI.BPMNLabelStyle) {
	log.Println("Initialize BPMNLabelStyle")
	if len(object.UUID) == 0 {
		object.UUID = "BPMNDI.BPMNLabelStyle%" + Utility.RandomUUID()
	}

	/** BPMNLabelStyle **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init Font **/
	if object.M_Font == nil {
		object.M_Font = new(DC.Font)
	}
	this.InitFont(&xmlElement.M_Font, object.M_Font)

	/** association initialisation **/
}

/** inititialisation of CategoryValue **/
func (this *BPMSXmlFactory) InitCategoryValue(xmlElement *BPMN20.XsdCategoryValue, object *BPMN20.CategoryValue) {
	log.Println("Initialize CategoryValue")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CategoryValue%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CategoryValue **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_value = xmlElement.M_value
}

/** inititialisation of Group **/
func (this *BPMSXmlFactory) InitGroup(xmlElement *BPMN20.XsdGroup, object *BPMN20.Group) {
	log.Println("Initialize Group")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Group%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Group **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_categoryValueRef) > 0 {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef)
	}
}

/** inititialisation of ImplicitThrowEvent **/
func (this *BPMSXmlFactory) InitImplicitThrowEvent(xmlElement *BPMN20.XsdImplicitThrowEvent, object *BPMN20.ImplicitThrowEvent) {
	log.Println("Initialize ImplicitThrowEvent")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ImplicitThrowEvent%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ImplicitThrowEvent **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetEventPtr(object)
	}

	/** Init dataInput **/
	object.M_dataInput = make([]*BPMN20.DataInput, 0)
	for i := 0; i < len(xmlElement.M_dataInput); i++ {
		val := new(BPMN20.DataInput)
		this.InitDataInput(xmlElement.M_dataInput[i], val)
		object.M_dataInput = append(object.M_dataInput, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init inputSet **/
	if xmlElement.M_inputSet != nil {
		object.M_inputSet = new(BPMN20.InputSet)
		this.InitInputSet(xmlElement.M_inputSet, object.M_inputSet)

		/** association initialisation **/
		object.M_inputSet.SetThrowEventPtr(object)
	}

	/** Init cancelEventDefinition **/
	object.M_eventDefinition = make([]BPMN20.EventDefinition, 0)
	for i := 0; i < len(xmlElement.M_eventDefinition_0); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_eventDefinition_0[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_1); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_eventDefinition_1[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_2); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_eventDefinition_2[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_3); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_eventDefinition_3[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_4); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_eventDefinition_4[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_5); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_eventDefinition_5[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_6); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_eventDefinition_6[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_7); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_eventDefinition_7[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_8); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_eventDefinition_8[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_9); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_eventDefinition_9[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init ref eventDefinitionRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_eventDefinitionRef); i++ {
		if _, ok := this.m_object[object.M_id]["eventDefinitionRef"]; !ok {
			this.m_object[object.M_id]["eventDefinitionRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["eventDefinitionRef"] = append(this.m_object[object.M_id]["eventDefinitionRef"], xmlElement.M_eventDefinitionRef[i])
	}

}

/** inititialisation of ResourceParameter **/
func (this *BPMSXmlFactory) InitResourceParameter(xmlElement *BPMN20.XsdResourceParameter, object *BPMN20.ResourceParameter) {
	log.Println("Initialize ResourceParameter")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ResourceParameter%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ResourceParameter **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref type **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_type) > 0 {
		if _, ok := this.m_object[object.M_id]["type"]; !ok {
			this.m_object[object.M_id]["type"] = make([]string, 0)
		}
		this.m_object[object.M_id]["type"] = append(this.m_object[object.M_id]["type"], xmlElement.M_type)
	}

	/** BaseElement **/
	object.M_isRequired = xmlElement.M_isRequired
}

/** inititialisation of StandardLoopCharacteristics **/
func (this *BPMSXmlFactory) InitStandardLoopCharacteristics(xmlElement *BPMN20.XsdStandardLoopCharacteristics, object *BPMN20.StandardLoopCharacteristics) {
	log.Println("Initialize StandardLoopCharacteristics")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.StandardLoopCharacteristics%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** StandardLoopCharacteristics **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if xmlElement.M_loopCondition != nil {
		if object.M_loopCondition == nil {
			object.M_loopCondition = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_loopCondition)), object.M_loopCondition)

		/** association initialisation **/
		object.M_loopCondition.SetStandardLoopCharacteristicsPtr(object)
	}

	/** LoopCharacteristics **/
	object.M_testBefore = xmlElement.M_testBefore

	/** LoopCharacteristics **/
	object.M_loopMaximum = xmlElement.M_loopMaximum
}

/** inititialisation of IntermediateThrowEvent **/
func (this *BPMSXmlFactory) InitIntermediateThrowEvent(xmlElement *BPMN20.XsdIntermediateThrowEvent, object *BPMN20.IntermediateThrowEvent) {
	log.Println("Initialize IntermediateThrowEvent")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.IntermediateThrowEvent%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** IntermediateThrowEvent **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetEventPtr(object)
	}

	/** Init dataInput **/
	object.M_dataInput = make([]*BPMN20.DataInput, 0)
	for i := 0; i < len(xmlElement.M_dataInput); i++ {
		val := new(BPMN20.DataInput)
		this.InitDataInput(xmlElement.M_dataInput[i], val)
		object.M_dataInput = append(object.M_dataInput, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init inputSet **/
	if xmlElement.M_inputSet != nil {
		object.M_inputSet = new(BPMN20.InputSet)
		this.InitInputSet(xmlElement.M_inputSet, object.M_inputSet)

		/** association initialisation **/
		object.M_inputSet.SetThrowEventPtr(object)
	}

	/** Init cancelEventDefinition **/
	object.M_eventDefinition = make([]BPMN20.EventDefinition, 0)
	for i := 0; i < len(xmlElement.M_eventDefinition_0); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_eventDefinition_0[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_1); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_eventDefinition_1[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_2); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_eventDefinition_2[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_3); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_eventDefinition_3[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_4); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_eventDefinition_4[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_5); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_eventDefinition_5[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_6); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_eventDefinition_6[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_7); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_eventDefinition_7[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_8); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_eventDefinition_8[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_9); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_eventDefinition_9[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init ref eventDefinitionRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_eventDefinitionRef); i++ {
		if _, ok := this.m_object[object.M_id]["eventDefinitionRef"]; !ok {
			this.m_object[object.M_id]["eventDefinitionRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["eventDefinitionRef"] = append(this.m_object[object.M_id]["eventDefinitionRef"], xmlElement.M_eventDefinitionRef[i])
	}

}

/** inititialisation of Association **/
func (this *BPMSXmlFactory) InitAssociation(xmlElement *BPMN20.XsdAssociation, object *BPMN20.Association) {
	log.Println("Initialize Association")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Association%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Association **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_sourceRef) > 0 {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], xmlElement.M_sourceRef)
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetRef) > 0 {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], xmlElement.M_targetRef)
	}

	/** AssociationDirection **/
	if xmlElement.M_associationDirection == "##None" {
		object.M_associationDirection = BPMN20.AssociationDirection_None
	} else if xmlElement.M_associationDirection == "##One" {
		object.M_associationDirection = BPMN20.AssociationDirection_One
	} else if xmlElement.M_associationDirection == "##Both" {
		object.M_associationDirection = BPMN20.AssociationDirection_Both
	} else {
		object.M_associationDirection = BPMN20.AssociationDirection_None
	}
}

/** inititialisation of ParticipantAssociation **/
func (this *BPMSXmlFactory) InitParticipantAssociation(xmlElement *BPMN20.XsdParticipantAssociation, object *BPMN20.ParticipantAssociation) {
	log.Println("Initialize ParticipantAssociation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ParticipantAssociation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ParticipantAssociation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref innerParticipantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_innerParticipantRef) > 0 {
		if _, ok := this.m_object[object.M_id]["innerParticipantRef"]; !ok {
			this.m_object[object.M_id]["innerParticipantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["innerParticipantRef"] = append(this.m_object[object.M_id]["innerParticipantRef"], xmlElement.M_innerParticipantRef)
	}

	/** Init ref outerParticipantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_outerParticipantRef) > 0 {
		if _, ok := this.m_object[object.M_id]["outerParticipantRef"]; !ok {
			this.m_object[object.M_id]["outerParticipantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outerParticipantRef"] = append(this.m_object[object.M_id]["outerParticipantRef"], xmlElement.M_outerParticipantRef)
	}

}

/** inititialisation of MessageFlowAssociation **/
func (this *BPMSXmlFactory) InitMessageFlowAssociation(xmlElement *BPMN20.XsdMessageFlowAssociation, object *BPMN20.MessageFlowAssociation) {
	log.Println("Initialize MessageFlowAssociation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.MessageFlowAssociation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** MessageFlowAssociation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref innerMessageFlowRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_innerMessageFlowRef) > 0 {
		if _, ok := this.m_object[object.M_id]["innerMessageFlowRef"]; !ok {
			this.m_object[object.M_id]["innerMessageFlowRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["innerMessageFlowRef"] = append(this.m_object[object.M_id]["innerMessageFlowRef"], xmlElement.M_innerMessageFlowRef)
	}

	/** Init ref outerMessageFlowRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_outerMessageFlowRef) > 0 {
		if _, ok := this.m_object[object.M_id]["outerMessageFlowRef"]; !ok {
			this.m_object[object.M_id]["outerMessageFlowRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outerMessageFlowRef"] = append(this.m_object[object.M_id]["outerMessageFlowRef"], xmlElement.M_outerMessageFlowRef)
	}
}

/** inititialisation of GlobalBusinessRuleTask **/
func (this *BPMSXmlFactory) InitGlobalBusinessRuleTask(xmlElement *BPMN20.XsdGlobalBusinessRuleTask, object *BPMN20.GlobalBusinessRuleTask) {
	log.Println("Initialize GlobalBusinessRuleTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalBusinessRuleTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalBusinessRuleTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref supportedInterfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supportedInterfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["supportedInterfaceRef"]; !ok {
			this.m_object[object.M_id]["supportedInterfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supportedInterfaceRef"] = append(this.m_object[object.M_id]["supportedInterfaceRef"], xmlElement.M_supportedInterfaceRef[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetCallableElementPtr(object)
	}

	/** Init ioBinding **/
	object.M_ioBinding = make([]*BPMN20.InputOutputBinding, 0)
	for i := 0; i < len(xmlElement.M_ioBinding); i++ {
		val := new(BPMN20.InputOutputBinding)
		this.InitInputOutputBinding(xmlElement.M_ioBinding[i], val)
		object.M_ioBinding = append(object.M_ioBinding, val)

		/** association initialisation **/
		val.SetCallableElementPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_Unspecified
		}
	}
}

/** inititialisation of HumanPerformer **/
func (this *BPMSXmlFactory) InitHumanPerformer(xmlElement *BPMN20.XsdHumanPerformer, object *BPMN20.HumanPerformer_impl) {
	log.Println("Initialize HumanPerformer")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.HumanPerformer_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** HumanPerformer **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref resourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_resourceRef != nil {
		if _, ok := this.m_object[object.M_id]["resourceRef"]; !ok {
			this.m_object[object.M_id]["resourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["resourceRef"] = append(this.m_object[object.M_id]["resourceRef"], *xmlElement.M_resourceRef)
	}

	/** Init resourceParameterBinding **/
	object.M_resourceParameterBinding = make([]*BPMN20.ResourceParameterBinding, 0)
	for i := 0; i < len(xmlElement.M_resourceParameterBinding); i++ {
		val := new(BPMN20.ResourceParameterBinding)
		this.InitResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], val)
		object.M_resourceParameterBinding = append(object.M_resourceParameterBinding, val)

		/** association initialisation **/
	}

	/** Init resourceAssignmentExpression **/
	if xmlElement.M_resourceAssignmentExpression != nil {
		object.M_resourceAssignmentExpression = new(BPMN20.ResourceAssignmentExpression)
		this.InitResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)

		/** association initialisation **/
	}
}

/** inititialisation of ItemDefinition **/
func (this *BPMSXmlFactory) InitItemDefinition(xmlElement *BPMN20.XsdItemDefinition, object *BPMN20.ItemDefinition) {
	log.Println("Initialize ItemDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ItemDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ItemDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref structureRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_structureRef) > 0 {
		if _, ok := this.m_object[object.M_id]["structureRef"]; !ok {
			this.m_object[object.M_id]["structureRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["structureRef"] = append(this.m_object[object.M_id]["structureRef"], xmlElement.M_structureRef)
	}

	/** RootElement **/
	object.M_isCollection = xmlElement.M_isCollection

	/** ItemKind **/
	if xmlElement.M_itemKind == "##Information" {
		object.M_itemKind = BPMN20.ItemKind_Information
	} else if xmlElement.M_itemKind == "##Physical" {
		object.M_itemKind = BPMN20.ItemKind_Physical
	} else {
		object.M_itemKind = BPMN20.ItemKind_Information
	}
}

/** inititialisation of Import **/
func (this *BPMSXmlFactory) InitImport(xmlElement *BPMN20.XsdImport, object *BPMN20.Import) {
	log.Println("Initialize Import")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Import%" + Utility.RandomUUID()
	}

	/** Import **/
	object.M_namespace = xmlElement.M_namespace

	/** Import **/
	object.M_location = xmlElement.M_location

	/** Import **/
	object.M_importType = xmlElement.M_importType
	this.m_references["Import"] = object
}

/** inititialisation of AdHocSubProcess **/
func (this *BPMSXmlFactory) InitAdHocSubProcess(xmlElement *BPMN20.XsdAdHocSubProcess, object *BPMN20.AdHocSubProcess) {
	log.Println("Initialize AdHocSubProcess")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.AdHocSubProcess%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** AdHocSubProcess **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}

	/** Init laneSet **/
	object.M_laneSet = make([]*BPMN20.LaneSet, 0)
	for i := 0; i < len(xmlElement.M_laneSet); i++ {
		val := new(BPMN20.LaneSet)
		this.InitLaneSet(xmlElement.M_laneSet[i], val)
		object.M_laneSet = append(object.M_laneSet, val)

		/** association initialisation **/
	}

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Activity **/
	object.M_triggeredByEvent = xmlElement.M_triggeredByEvent

	/** Init formalExpression **/
	if xmlElement.M_completionCondition != nil {
		if object.M_completionCondition == nil {
			object.M_completionCondition = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_completionCondition)), object.M_completionCondition)

		/** association initialisation **/
		object.M_completionCondition.SetAdHocSubProcessPtr(object)
	}

	/** SubProcess **/
	object.M_cancelRemainingInstances = xmlElement.M_cancelRemainingInstances

	/** AdHocOrdering **/
	if xmlElement.M_ordering == "##Parallel" {
		object.M_ordering = BPMN20.AdHocOrdering_Parallel
	} else if xmlElement.M_ordering == "##Sequential" {
		object.M_ordering = BPMN20.AdHocOrdering_Sequential
	}
}

/** inititialisation of PartnerRole **/
func (this *BPMSXmlFactory) InitPartnerRole(xmlElement *BPMN20.XsdPartnerRole, object *BPMN20.PartnerRole) {
	log.Println("Initialize PartnerRole")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.PartnerRole%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** PartnerRole **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of BusinessRuleTask **/
func (this *BPMSXmlFactory) InitBusinessRuleTask(xmlElement *BPMN20.XsdBusinessRuleTask, object *BPMN20.BusinessRuleTask) {
	log.Println("Initialize BusinessRuleTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.BusinessRuleTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** BusinessRuleTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_Unspecified
		}
	}
}

/** inititialisation of IntermediateCatchEvent **/
func (this *BPMSXmlFactory) InitIntermediateCatchEvent(xmlElement *BPMN20.XsdIntermediateCatchEvent, object *BPMN20.IntermediateCatchEvent) {
	log.Println("Initialize IntermediateCatchEvent")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.IntermediateCatchEvent%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** IntermediateCatchEvent **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetEventPtr(object)
	}

	/** Init dataOutput **/
	object.M_dataOutput = make([]*BPMN20.DataOutput, 0)
	for i := 0; i < len(xmlElement.M_dataOutput); i++ {
		val := new(BPMN20.DataOutput)
		this.InitDataOutput(xmlElement.M_dataOutput[i], val)
		object.M_dataOutput = append(object.M_dataOutput, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init outputSet **/
	if xmlElement.M_outputSet != nil {
		object.M_outputSet = new(BPMN20.OutputSet)
		this.InitOutputSet(xmlElement.M_outputSet, object.M_outputSet)

		/** association initialisation **/
		object.M_outputSet.SetCatchEventPtr(object)
	}

	/** Init cancelEventDefinition **/
	object.M_eventDefinition = make([]BPMN20.EventDefinition, 0)
	for i := 0; i < len(xmlElement.M_eventDefinition_0); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_eventDefinition_0[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_1); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_eventDefinition_1[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_2); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_eventDefinition_2[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_3); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_eventDefinition_3[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_4); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_eventDefinition_4[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_5); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_eventDefinition_5[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_6); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_eventDefinition_6[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_7); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_eventDefinition_7[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_8); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_eventDefinition_8[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_9); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_eventDefinition_9[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init ref eventDefinitionRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_eventDefinitionRef); i++ {
		if _, ok := this.m_object[object.M_id]["eventDefinitionRef"]; !ok {
			this.m_object[object.M_id]["eventDefinitionRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["eventDefinitionRef"] = append(this.m_object[object.M_id]["eventDefinitionRef"], xmlElement.M_eventDefinitionRef[i])
	}

	/** Event **/
	object.M_parallelMultiple = xmlElement.M_parallelMultiple
}

/** inititialisation of SubProcess **/
func (this *BPMSXmlFactory) InitSubProcess(xmlElement *BPMN20.XsdSubProcess, object *BPMN20.SubProcess_impl) {
	log.Println("Initialize SubProcess")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.SubProcess_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** SubProcess **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}

	/** Init laneSet **/
	object.M_laneSet = make([]*BPMN20.LaneSet, 0)
	for i := 0; i < len(xmlElement.M_laneSet); i++ {
		val := new(BPMN20.LaneSet)
		this.InitLaneSet(xmlElement.M_laneSet[i], val)
		object.M_laneSet = append(object.M_laneSet, val)

		/** association initialisation **/
	}

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Activity **/
	object.M_triggeredByEvent = xmlElement.M_triggeredByEvent
}

/** inititialisation of Task **/
func (this *BPMSXmlFactory) InitTask(xmlElement *BPMN20.XsdTask, object *BPMN20.Task_impl) {
	log.Println("Initialize Task")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Task_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Task **/
	object.M_id = xmlElement.M_id
	log.Println("--------------------------------> Task Id is ", object.M_id)
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
}

/** inititialisation of Transaction **/
func (this *BPMSXmlFactory) InitTransaction(xmlElement *BPMN20.XsdTransaction, object *BPMN20.Transaction) {
	log.Println("Initialize Transaction")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Transaction%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Transaction **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}

	/** Init laneSet **/
	object.M_laneSet = make([]*BPMN20.LaneSet, 0)
	for i := 0; i < len(xmlElement.M_laneSet); i++ {
		val := new(BPMN20.LaneSet)
		this.InitLaneSet(xmlElement.M_laneSet[i], val)
		object.M_laneSet = append(object.M_laneSet, val)

		/** association initialisation **/
	}

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubProcessPtr(object)
	}

	/** Activity **/
	object.M_triggeredByEvent = xmlElement.M_triggeredByEvent
	if !strings.HasPrefix(xmlElement.M_method, "##") {
		object.M_methodStr = xmlElement.M_method
	} else {

		/** TransactionMethod **/
		if xmlElement.M_method == "##Compensate" {
			object.M_method = BPMN20.TransactionMethod_Compensate
		} else if xmlElement.M_method == "##Image" {
			object.M_method = BPMN20.TransactionMethod_Image
		} else if xmlElement.M_method == "##Store" {
			object.M_method = BPMN20.TransactionMethod_Store
		} else {
			object.M_method = BPMN20.TransactionMethod_Compensate
		}
	}
}

/** inititialisation of EscalationEventDefinition **/
func (this *BPMSXmlFactory) InitEscalationEventDefinition(xmlElement *BPMN20.XsdEscalationEventDefinition, object *BPMN20.EscalationEventDefinition) {
	log.Println("Initialize EscalationEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.EscalationEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** EscalationEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref escalationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_escalationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["escalationRef"]; !ok {
			this.m_object[object.M_id]["escalationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["escalationRef"] = append(this.m_object[object.M_id]["escalationRef"], xmlElement.M_escalationRef)
	}
}

/** inititialisation of GlobalUserTask **/
func (this *BPMSXmlFactory) InitGlobalUserTask(xmlElement *BPMN20.XsdGlobalUserTask, object *BPMN20.GlobalUserTask) {
	log.Println("Initialize GlobalUserTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalUserTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalUserTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref supportedInterfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supportedInterfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["supportedInterfaceRef"]; !ok {
			this.m_object[object.M_id]["supportedInterfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supportedInterfaceRef"] = append(this.m_object[object.M_id]["supportedInterfaceRef"], xmlElement.M_supportedInterfaceRef[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetCallableElementPtr(object)
	}

	/** Init ioBinding **/
	object.M_ioBinding = make([]*BPMN20.InputOutputBinding, 0)
	for i := 0; i < len(xmlElement.M_ioBinding); i++ {
		val := new(BPMN20.InputOutputBinding)
		this.InitInputOutputBinding(xmlElement.M_ioBinding[i], val)
		object.M_ioBinding = append(object.M_ioBinding, val)

		/** association initialisation **/
		val.SetCallableElementPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init rendering **/
	object.M_rendering = make([]*BPMN20.Rendering, 0)
	for i := 0; i < len(xmlElement.M_rendering); i++ {
		val := new(BPMN20.Rendering)
		this.InitRendering(xmlElement.M_rendering[i], val)
		object.M_rendering = append(object.M_rendering, val)

		/** association initialisation **/
		val.SetGlobalUserTaskPtr(object)
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_Unspecified
		}
	}
}

/** inititialisation of BPMNLabel **/
func (this *BPMSXmlFactory) InitBPMNLabel(xmlElement *BPMNDI.XsdBPMNLabel, object *BPMNDI.BPMNLabel) {
	log.Println("Initialize BPMNLabel")
	if len(object.UUID) == 0 {
		object.UUID = "BPMNDI.BPMNLabel%" + Utility.RandomUUID()
	}

	/** Init extension **/

	/** BPMNLabel **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init Bounds **/
	if xmlElement.M_Bounds != nil {
		object.M_Bounds = new(DC.Bounds)
		this.InitBounds(xmlElement.M_Bounds, object.M_Bounds)

		/** association initialisation **/
	}

	/** Init ref labelStyle **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_labelStyle) > 0 {
		if _, ok := this.m_object[object.M_id]["labelStyle"]; !ok {
			this.m_object[object.M_id]["labelStyle"] = make([]string, 0)
		}
		this.m_object[object.M_id]["labelStyle"] = append(this.m_object[object.M_id]["labelStyle"], xmlElement.M_labelStyle)
	}
}

/** inititialisation of ConditionalEventDefinition **/
func (this *BPMSXmlFactory) InitConditionalEventDefinition(xmlElement *BPMN20.XsdConditionalEventDefinition, object *BPMN20.ConditionalEventDefinition) {
	log.Println("Initialize ConditionalEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ConditionalEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ConditionalEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if xmlElement.M_condition != nil {
		if object.M_condition == nil {
			object.M_condition = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_condition)), object.M_condition)

		/** association initialisation **/
		object.M_condition.SetConditionalEventDefinitionPtr(object)
	}
}

/** inititialisation of TimerEventDefinition **/
func (this *BPMSXmlFactory) InitTimerEventDefinition(xmlElement *BPMN20.XsdTimerEventDefinition, object *BPMN20.TimerEventDefinition) {
	log.Println("Initialize TimerEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.TimerEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** TimerEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of InputOutputSpecification **/
func (this *BPMSXmlFactory) InitInputOutputSpecification(xmlElement *BPMN20.XsdInputOutputSpecification, object *BPMN20.InputOutputSpecification) {
	log.Println("Initialize InputOutputSpecification")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.InputOutputSpecification%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** InputOutputSpecification **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init dataInput **/
	object.M_dataInput = make([]*BPMN20.DataInput, 0)
	for i := 0; i < len(xmlElement.M_dataInput); i++ {
		val := new(BPMN20.DataInput)
		this.InitDataInput(xmlElement.M_dataInput[i], val)
		object.M_dataInput = append(object.M_dataInput, val)

		/** association initialisation **/
		val.SetInputOutputSpecificationPtr(object)
	}

	/** Init dataOutput **/
	object.M_dataOutput = make([]*BPMN20.DataOutput, 0)
	for i := 0; i < len(xmlElement.M_dataOutput); i++ {
		val := new(BPMN20.DataOutput)
		this.InitDataOutput(xmlElement.M_dataOutput[i], val)
		object.M_dataOutput = append(object.M_dataOutput, val)

		/** association initialisation **/
		val.SetInputOutputSpecificationPtr(object)
	}

	/** Init inputSet **/
	object.M_inputSet = make([]*BPMN20.InputSet, 0)
	for i := 0; i < len(xmlElement.M_inputSet); i++ {
		val := new(BPMN20.InputSet)
		this.InitInputSet(xmlElement.M_inputSet[i], val)
		object.M_inputSet = append(object.M_inputSet, val)

		/** association initialisation **/
		val.SetInputOutputSpecificationPtr(object)
	}

	/** Init outputSet **/
	object.M_outputSet = make([]*BPMN20.OutputSet, 0)
	for i := 0; i < len(xmlElement.M_outputSet); i++ {
		val := new(BPMN20.OutputSet)
		this.InitOutputSet(xmlElement.M_outputSet[i], val)
		object.M_outputSet = append(object.M_outputSet, val)

		/** association initialisation **/
		val.SetInputOutputSpecificationPtr(object)
	}
}

/** inititialisation of Message **/
func (this *BPMSXmlFactory) InitMessage(xmlElement *BPMN20.XsdMessage, object *BPMN20.Message) {
	log.Println("Initialize Message")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Message%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Message **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init ref itemRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemRef"]; !ok {
			this.m_object[object.M_id]["itemRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemRef"] = append(this.m_object[object.M_id]["itemRef"], xmlElement.M_itemRef)
	}
}

/** inititialisation of ComplexBehaviorDefinition **/
func (this *BPMSXmlFactory) InitComplexBehaviorDefinition(xmlElement *BPMN20.XsdComplexBehaviorDefinition, object *BPMN20.ComplexBehaviorDefinition) {
	log.Println("Initialize ComplexBehaviorDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ComplexBehaviorDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ComplexBehaviorDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if object.M_condition == nil {
		object.M_condition = new(BPMN20.FormalExpression)
	}
	this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(&xmlElement.M_condition)), object.M_condition)

	/** association initialisation **/
	object.M_condition.SetComplexBehaviorDefinitionPtr(object)

	/** Init implicitThrowEvent **/
	if xmlElement.M_event != nil {
		if object.M_event == nil {
			object.M_event = new(BPMN20.ImplicitThrowEvent)
		}
		this.InitImplicitThrowEvent((*BPMN20.XsdImplicitThrowEvent)(unsafe.Pointer(xmlElement.M_event)), object.M_event)

		/** association initialisation **/
	}
}

/** inititialisation of DataObjectReference **/
func (this *BPMSXmlFactory) InitDataObjectReference(xmlElement *BPMN20.XsdDataObjectReference, object *BPMN20.DataObjectReference) {
	log.Println("Initialize DataObjectReference")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataObjectReference%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataObjectReference **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}

	/** Init ref dataObjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_dataObjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["dataObjectRef"]; !ok {
			this.m_object[object.M_id]["dataObjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["dataObjectRef"] = append(this.m_object[object.M_id]["dataObjectRef"], xmlElement.M_dataObjectRef)
	}
}

/** inititialisation of OutputSet **/
func (this *BPMSXmlFactory) InitOutputSet(xmlElement *BPMN20.XsdOutputSet, object *BPMN20.OutputSet) {
	log.Println("Initialize OutputSet")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.OutputSet%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** OutputSet **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref dataOutputRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_dataOutputRefs); i++ {
		if _, ok := this.m_object[object.M_id]["dataOutputRefs"]; !ok {
			this.m_object[object.M_id]["dataOutputRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["dataOutputRefs"] = append(this.m_object[object.M_id]["dataOutputRefs"], xmlElement.M_dataOutputRefs[i])
	}

	/** Init ref optionalOutputRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_optionalOutputRefs); i++ {
		if _, ok := this.m_object[object.M_id]["optionalOutputRefs"]; !ok {
			this.m_object[object.M_id]["optionalOutputRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["optionalOutputRefs"] = append(this.m_object[object.M_id]["optionalOutputRefs"], xmlElement.M_optionalOutputRefs[i])
	}

	/** Init ref whileExecutingOutputRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_whileExecutingOutputRefs); i++ {
		if _, ok := this.m_object[object.M_id]["whileExecutingOutputRefs"]; !ok {
			this.m_object[object.M_id]["whileExecutingOutputRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["whileExecutingOutputRefs"] = append(this.m_object[object.M_id]["whileExecutingOutputRefs"], xmlElement.M_whileExecutingOutputRefs[i])
	}

	/** Init ref inputSetRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_inputSetRefs); i++ {
		if _, ok := this.m_object[object.M_id]["inputSetRefs"]; !ok {
			this.m_object[object.M_id]["inputSetRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inputSetRefs"] = append(this.m_object[object.M_id]["inputSetRefs"], xmlElement.M_inputSetRefs[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of ResourceRole **/
func (this *BPMSXmlFactory) InitResourceRole(xmlElement *BPMN20.XsdResourceRole, object *BPMN20.ResourceRole_impl) {
	log.Println("Initialize ResourceRole")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ResourceRole_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ResourceRole **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref resourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_resourceRef != nil {
		if _, ok := this.m_object[object.M_id]["resourceRef"]; !ok {
			this.m_object[object.M_id]["resourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["resourceRef"] = append(this.m_object[object.M_id]["resourceRef"], *xmlElement.M_resourceRef)
	}

	/** Init resourceParameterBinding **/
	object.M_resourceParameterBinding = make([]*BPMN20.ResourceParameterBinding, 0)
	for i := 0; i < len(xmlElement.M_resourceParameterBinding); i++ {
		val := new(BPMN20.ResourceParameterBinding)
		this.InitResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], val)
		object.M_resourceParameterBinding = append(object.M_resourceParameterBinding, val)

		/** association initialisation **/
	}

	/** Init resourceAssignmentExpression **/
	if xmlElement.M_resourceAssignmentExpression != nil {
		object.M_resourceAssignmentExpression = new(BPMN20.ResourceAssignmentExpression)
		this.InitResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)

		/** association initialisation **/
	}
}

/** inititialisation of ResourceParameterBinding **/
func (this *BPMSXmlFactory) InitResourceParameterBinding(xmlElement *BPMN20.XsdResourceParameterBinding, object *BPMN20.ResourceParameterBinding) {
	log.Println("Initialize ResourceParameterBinding")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ResourceParameterBinding%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ResourceParameterBinding **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if xmlElement.M_expression_0 != nil {
		if object.M_expression == nil {
			object.M_expression = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression(xmlElement.M_expression_0, object.M_expression.(*BPMN20.FormalExpression))

		/** association initialisation **/
		object.M_expression.(*BPMN20.FormalExpression).SetResourceParameterBindingPtr(object)
	}

	/** Init expression **/
	if xmlElement.M_expression_1 != nil {
		if object.M_expression == nil {
			object.M_expression = new(BPMN20.FormalExpression)
		}
		this.InitExpression(xmlElement.M_expression_1, object.M_expression.(*BPMN20.Expression_impl))

		/** association initialisation **/
		object.M_expression.(*BPMN20.Expression_impl).SetResourceParameterBindingPtr(object)
	}

	/** Init ref parameterRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_parameterRef) > 0 {
		if _, ok := this.m_object[object.M_id]["parameterRef"]; !ok {
			this.m_object[object.M_id]["parameterRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["parameterRef"] = append(this.m_object[object.M_id]["parameterRef"], xmlElement.M_parameterRef)
	}
}

/** inititialisation of Performer **/
func (this *BPMSXmlFactory) InitPerformer(xmlElement *BPMN20.XsdPerformer, object *BPMN20.Performer_impl) {
	log.Println("Initialize Performer")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Performer_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Performer **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref resourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_resourceRef != nil {
		if _, ok := this.m_object[object.M_id]["resourceRef"]; !ok {
			this.m_object[object.M_id]["resourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["resourceRef"] = append(this.m_object[object.M_id]["resourceRef"], *xmlElement.M_resourceRef)
	}

	/** Init resourceParameterBinding **/
	object.M_resourceParameterBinding = make([]*BPMN20.ResourceParameterBinding, 0)
	for i := 0; i < len(xmlElement.M_resourceParameterBinding); i++ {
		val := new(BPMN20.ResourceParameterBinding)
		this.InitResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], val)
		object.M_resourceParameterBinding = append(object.M_resourceParameterBinding, val)

		/** association initialisation **/
	}

	/** Init resourceAssignmentExpression **/
	if xmlElement.M_resourceAssignmentExpression != nil {
		object.M_resourceAssignmentExpression = new(BPMN20.ResourceAssignmentExpression)
		this.InitResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)

		/** association initialisation **/
	}
}

/** inititialisation of ChoreographyTask **/
func (this *BPMSXmlFactory) InitChoreographyTask(xmlElement *BPMN20.XsdChoreographyTask, object *BPMN20.ChoreographyTask) {
	log.Println("Initialize ChoreographyTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ChoreographyTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ChoreographyTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetChoreographyActivityPtr(object)
	}

	/** Init ref initiatingParticipantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_initiatingParticipantRef) > 0 {
		if _, ok := this.m_object[object.M_id]["initiatingParticipantRef"]; !ok {
			this.m_object[object.M_id]["initiatingParticipantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["initiatingParticipantRef"] = append(this.m_object[object.M_id]["initiatingParticipantRef"], xmlElement.M_initiatingParticipantRef)
	}

	/** ChoreographyLoopType **/
	if xmlElement.M_loopType == "##None" {
		object.M_loopType = BPMN20.ChoreographyLoopType_None
	} else if xmlElement.M_loopType == "##Standard" {
		object.M_loopType = BPMN20.ChoreographyLoopType_Standard
	} else if xmlElement.M_loopType == "##MultiInstanceSequential" {
		object.M_loopType = BPMN20.ChoreographyLoopType_MultiInstanceSequential
	} else if xmlElement.M_loopType == "##MultiInstanceParallel" {
		object.M_loopType = BPMN20.ChoreographyLoopType_MultiInstanceParallel
	} else {
		object.M_loopType = BPMN20.ChoreographyLoopType_None
	}

	/** Init ref messageFlowRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_messageFlowRef) > 0 {
		if _, ok := this.m_object[object.M_id]["messageFlowRef"]; !ok {
			this.m_object[object.M_id]["messageFlowRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageFlowRef"] = append(this.m_object[object.M_id]["messageFlowRef"], xmlElement.M_messageFlowRef)
	}

}

/** inititialisation of Interface **/
func (this *BPMSXmlFactory) InitInterface(xmlElement *BPMN20.XsdInterface, object *BPMN20.Interface) {
	log.Println("Initialize Interface")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Interface%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Interface **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init operation **/
	object.M_operation = make([]*BPMN20.Operation, 0)
	for i := 0; i < len(xmlElement.M_operation); i++ {
		val := new(BPMN20.Operation)
		this.InitOperation(xmlElement.M_operation[i], val)
		object.M_operation = append(object.M_operation, val)

		/** association initialisation **/
		val.SetInterfacePtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init ref implementationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_implementationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["implementationRef"]; !ok {
			this.m_object[object.M_id]["implementationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["implementationRef"] = append(this.m_object[object.M_id]["implementationRef"], xmlElement.M_implementationRef)
	}
}

/** inititialisation of Process **/
func (this *BPMSXmlFactory) InitProcess(xmlElement *BPMN20.XsdProcess, object *BPMN20.Process) {
	log.Println("Initialize Process")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Process%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Process **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref supportedInterfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supportedInterfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["supportedInterfaceRef"]; !ok {
			this.m_object[object.M_id]["supportedInterfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supportedInterfaceRef"] = append(this.m_object[object.M_id]["supportedInterfaceRef"], xmlElement.M_supportedInterfaceRef[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetCallableElementPtr(object)
	}

	/** Init ioBinding **/
	object.M_ioBinding = make([]*BPMN20.InputOutputBinding, 0)
	for i := 0; i < len(xmlElement.M_ioBinding); i++ {
		val := new(BPMN20.InputOutputBinding)
		this.InitInputOutputBinding(xmlElement.M_ioBinding[i], val)
		object.M_ioBinding = append(object.M_ioBinding, val)

		/** association initialisation **/
		val.SetCallableElementPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetProcessPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetProcessPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init laneSet **/
	object.M_laneSet = make([]*BPMN20.LaneSet, 0)
	for i := 0; i < len(xmlElement.M_laneSet); i++ {
		val := new(BPMN20.LaneSet)
		this.InitLaneSet(xmlElement.M_laneSet[i], val)
		object.M_laneSet = append(object.M_laneSet, val)

		/** association initialisation **/
	}

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init correlationSubscription **/
	object.M_correlationSubscription = make([]*BPMN20.CorrelationSubscription, 0)
	for i := 0; i < len(xmlElement.M_correlationSubscription); i++ {
		val := new(BPMN20.CorrelationSubscription)
		this.InitCorrelationSubscription(xmlElement.M_correlationSubscription[i], val)
		object.M_correlationSubscription = append(object.M_correlationSubscription, val)

		/** association initialisation **/
		val.SetProcessPtr(object)
	}

	/** Init ref supports **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supports); i++ {
		if _, ok := this.m_object[object.M_id]["supports"]; !ok {
			this.m_object[object.M_id]["supports"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supports"] = append(this.m_object[object.M_id]["supports"], xmlElement.M_supports[i])
	}

	/** ProcessType **/
	if xmlElement.M_processType == "##None" {
		object.M_processType = BPMN20.ProcessType_None
	} else if xmlElement.M_processType == "##Public" {
		object.M_processType = BPMN20.ProcessType_Public
	} else if xmlElement.M_processType == "##Private" {
		object.M_processType = BPMN20.ProcessType_Private
	} else {
		object.M_processType = BPMN20.ProcessType_None
	}

	/** CallableElement **/
	object.M_isClosed = xmlElement.M_isClosed

	/** CallableElement **/
	object.M_isExecutable = xmlElement.M_isExecutable

	/** Init ref definitionalCollaborationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_definitionalCollaborationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["definitionalCollaborationRef"]; !ok {
			this.m_object[object.M_id]["definitionalCollaborationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["definitionalCollaborationRef"] = append(this.m_object[object.M_id]["definitionalCollaborationRef"], xmlElement.M_definitionalCollaborationRef)
	}
}

/** inititialisation of BoundaryEvent **/
func (this *BPMSXmlFactory) InitBoundaryEvent(xmlElement *BPMN20.XsdBoundaryEvent, object *BPMN20.BoundaryEvent) {
	log.Println("Initialize BoundaryEvent")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.BoundaryEvent%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** BoundaryEvent **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetEventPtr(object)
	}

	/** Init dataOutput **/
	object.M_dataOutput = make([]*BPMN20.DataOutput, 0)
	for i := 0; i < len(xmlElement.M_dataOutput); i++ {
		val := new(BPMN20.DataOutput)
		this.InitDataOutput(xmlElement.M_dataOutput[i], val)
		object.M_dataOutput = append(object.M_dataOutput, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init outputSet **/
	if xmlElement.M_outputSet != nil {
		object.M_outputSet = new(BPMN20.OutputSet)
		this.InitOutputSet(xmlElement.M_outputSet, object.M_outputSet)

		/** association initialisation **/
		object.M_outputSet.SetCatchEventPtr(object)
	}

	/** Init cancelEventDefinition **/
	object.M_eventDefinition = make([]BPMN20.EventDefinition, 0)
	for i := 0; i < len(xmlElement.M_eventDefinition_0); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_eventDefinition_0[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_1); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_eventDefinition_1[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_2); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_eventDefinition_2[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_3); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_eventDefinition_3[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_4); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_eventDefinition_4[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_5); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_eventDefinition_5[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_6); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_eventDefinition_6[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_7); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_eventDefinition_7[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_8); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_eventDefinition_8[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_9); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_eventDefinition_9[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetCatchEventPtr(object)
	}

	/** Init ref eventDefinitionRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_eventDefinitionRef); i++ {
		if _, ok := this.m_object[object.M_id]["eventDefinitionRef"]; !ok {
			this.m_object[object.M_id]["eventDefinitionRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["eventDefinitionRef"] = append(this.m_object[object.M_id]["eventDefinitionRef"], xmlElement.M_eventDefinitionRef[i])
	}

	/** Event **/
	object.M_parallelMultiple = xmlElement.M_parallelMultiple

	/** CatchEvent **/
	object.M_cancelActivity = xmlElement.M_cancelActivity

	/** Init ref attachedToRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_attachedToRef) > 0 {
		if _, ok := this.m_object[object.M_id]["attachedToRef"]; !ok {
			this.m_object[object.M_id]["attachedToRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["attachedToRef"] = append(this.m_object[object.M_id]["attachedToRef"], xmlElement.M_attachedToRef)
	}
}

/** inititialisation of ServiceTask **/
func (this *BPMSXmlFactory) InitServiceTask(xmlElement *BPMN20.XsdServiceTask, object *BPMN20.ServiceTask) {
	log.Println("Initialize ServiceTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ServiceTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ServiceTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_WebService
		}
	}

	/** Init ref operationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_operationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["operationRef"]; !ok {
			this.m_object[object.M_id]["operationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["operationRef"] = append(this.m_object[object.M_id]["operationRef"], xmlElement.M_operationRef)
	}
}

/** inititialisation of CorrelationSubscription **/
func (this *BPMSXmlFactory) InitCorrelationSubscription(xmlElement *BPMN20.XsdCorrelationSubscription, object *BPMN20.CorrelationSubscription) {
	log.Println("Initialize CorrelationSubscription")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CorrelationSubscription%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CorrelationSubscription **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init correlationPropertyBinding **/
	object.M_correlationPropertyBinding = make([]*BPMN20.CorrelationPropertyBinding, 0)
	for i := 0; i < len(xmlElement.M_correlationPropertyBinding); i++ {
		val := new(BPMN20.CorrelationPropertyBinding)
		this.InitCorrelationPropertyBinding(xmlElement.M_correlationPropertyBinding[i], val)
		object.M_correlationPropertyBinding = append(object.M_correlationPropertyBinding, val)

		/** association initialisation **/
		val.SetCorrelationSubscriptionPtr(object)
	}

	/** Init ref correlationKeyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_correlationKeyRef) > 0 {
		if _, ok := this.m_object[object.M_id]["correlationKeyRef"]; !ok {
			this.m_object[object.M_id]["correlationKeyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["correlationKeyRef"] = append(this.m_object[object.M_id]["correlationKeyRef"], xmlElement.M_correlationKeyRef)
	}
}

/** inititialisation of Extension **/
func (this *BPMSXmlFactory) InitExtension(xmlElement *BPMN20.XsdExtension, object *BPMN20.Extension) {
	log.Println("Initialize Extension")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Extension%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
	}

	/** Init ref extensionDefinition **/
	/*if len("xsd:QName") == 0 {
		"xsd:QName"=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object["xsd:QName"]; !ok {
		this.m_object["xsd:QName"]=make(map[string][]string)
	}
	if len(xmlElement.M_extensionDefinition) > 0 {
		if _, ok:= this.m_object["xsd:QName"]["extensionDefinition"]; !ok {
			this.m_object["xsd:QName"]["extensionDefinition"]=make([]string,0)
		}
		this.m_object["xsd:QName"]["extensionDefinition"] = append(this.m_object["xsd:QName"]["extensionDefinition"], xmlElement.M_extensionDefinition)
	}*/

	/** Extension **/
	object.M_mustUnderstand = xmlElement.M_mustUnderstand
	this.m_references["Extension"] = object
}

/** inititialisation of Expression **/
func (this *BPMSXmlFactory) InitExpression(xmlElement *BPMN20.XsdExpression, object *BPMN20.Expression_impl) {
	log.Println("Initialize Expression")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Expression_impl%" + Utility.RandomUUID()
	}
	/** other content **/
	exprStr := xmlElement.M_other
	object.SetOther(exprStr)
}

/** inititialisation of Monitoring **/
func (this *BPMSXmlFactory) InitMonitoring(xmlElement *BPMN20.XsdMonitoring, object *BPMN20.Monitoring) {
	log.Println("Initialize Monitoring")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Monitoring%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Monitoring **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Property **/
func (this *BPMSXmlFactory) InitProperty(xmlElement *BPMN20.XsdProperty, object *BPMN20.Property) {
	log.Println("Initialize Property")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Property%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Property **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}
}

/** inititialisation of Lane **/
func (this *BPMSXmlFactory) InitLane(xmlElement *BPMN20.XsdLane, object *BPMN20.Lane) {
	log.Println("Initialize Lane")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Lane%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Lane **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref flowNodeRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_flowNodeRef); i++ {
		if _, ok := this.m_object[object.M_id]["flowNodeRef"]; !ok {
			this.m_object[object.M_id]["flowNodeRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["flowNodeRef"] = append(this.m_object[object.M_id]["flowNodeRef"], xmlElement.M_flowNodeRef[i])
	}

	/** Init laneSet **/
	if xmlElement.M_childLaneSet != nil {
		if object.M_childLaneSet == nil {
			object.M_childLaneSet = new(BPMN20.LaneSet)
		}
		this.InitLaneSet((*BPMN20.XsdLaneSet)(unsafe.Pointer(xmlElement.M_childLaneSet)), object.M_childLaneSet)

		/** association initialisation **/
		object.M_childLaneSet.SetLanePtr(object)
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref partitionElementRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_partitionElementRef) > 0 {
		if _, ok := this.m_object[object.M_id]["partitionElementRef"]; !ok {
			this.m_object[object.M_id]["partitionElementRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["partitionElementRef"] = append(this.m_object[object.M_id]["partitionElementRef"], xmlElement.M_partitionElementRef)
	}
}

/** inititialisation of CallActivity **/
func (this *BPMSXmlFactory) InitCallActivity(xmlElement *BPMN20.XsdCallActivity, object *BPMN20.CallActivity) {
	log.Println("Initialize CallActivity")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CallActivity%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CallActivity **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}

	/** Init ref calledElement **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_calledElement) > 0 {
		if _, ok := this.m_object[object.M_id]["calledElement"]; !ok {
			this.m_object[object.M_id]["calledElement"] = make([]string, 0)
		}
		this.m_object[object.M_id]["calledElement"] = append(this.m_object[object.M_id]["calledElement"], xmlElement.M_calledElement)
	}
}

/** inititialisation of Category **/
func (this *BPMSXmlFactory) InitCategory(xmlElement *BPMN20.XsdCategory, object *BPMN20.Category) {
	log.Println("Initialize Category")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Category%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Category **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init categoryValue **/
	object.M_categoryValue = make([]*BPMN20.CategoryValue, 0)
	for i := 0; i < len(xmlElement.M_categoryValue); i++ {
		val := new(BPMN20.CategoryValue)
		this.InitCategoryValue(xmlElement.M_categoryValue[i], val)
		object.M_categoryValue = append(object.M_categoryValue, val)

		/** association initialisation **/
		val.SetCategoryPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of MessageEventDefinition **/
func (this *BPMSXmlFactory) InitMessageEventDefinition(xmlElement *BPMN20.XsdMessageEventDefinition, object *BPMN20.MessageEventDefinition) {
	log.Println("Initialize MessageEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.MessageEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** MessageEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref operationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_operationRef != nil {
		if _, ok := this.m_object[object.M_id]["operationRef"]; !ok {
			this.m_object[object.M_id]["operationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["operationRef"] = append(this.m_object[object.M_id]["operationRef"], *xmlElement.M_operationRef)
	}

	/** Init ref messageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_messageRef) > 0 {
		if _, ok := this.m_object[object.M_id]["messageRef"]; !ok {
			this.m_object[object.M_id]["messageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageRef"] = append(this.m_object[object.M_id]["messageRef"], xmlElement.M_messageRef)
	}
}

/** inititialisation of DataOutput **/
func (this *BPMSXmlFactory) InitDataOutput(xmlElement *BPMN20.XsdDataOutput, object *BPMN20.DataOutput) {
	log.Println("Initialize DataOutput")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataOutput%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataOutput **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}

	/** BaseElement **/
	object.M_isCollection = xmlElement.M_isCollection
}

/** inititialisation of SequenceFlow **/
func (this *BPMSXmlFactory) InitSequenceFlow(xmlElement *BPMN20.XsdSequenceFlow, object *BPMN20.SequenceFlow) {
	log.Println("Initialize SequenceFlow")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.SequenceFlow%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** SequenceFlow **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init formalExpression **/
	if xmlElement.M_conditionExpression != nil {
		if object.M_conditionExpression == nil {
			object.M_conditionExpression = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_conditionExpression)), object.M_conditionExpression)

		/** association initialisation **/
		object.M_conditionExpression.SetSequenceFlowPtr(object)
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_sourceRef) > 0 {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], xmlElement.M_sourceRef)
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetRef) > 0 {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], xmlElement.M_targetRef)
	}

	/** FlowElement **/
	object.M_isImmediate = xmlElement.M_isImmediate
}

/** inititialisation of Bounds **/
func (this *BPMSXmlFactory) InitBounds(xmlElement *DC.XsdBounds, object *DC.Bounds) {
	log.Println("Initialize Bounds")
	if len(object.UUID) == 0 {
		object.UUID = "DC.Bounds%" + Utility.RandomUUID()
	}

	/** Bounds **/
	object.M_x = xmlElement.M_x

	/** Bounds **/
	object.M_y = xmlElement.M_y

	/** Bounds **/
	object.M_width = xmlElement.M_width

	/** Bounds **/
	object.M_height = xmlElement.M_height
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of DataStore **/
func (this *BPMSXmlFactory) InitDataStore(xmlElement *BPMN20.XsdDataStore, object *BPMN20.DataStore) {
	log.Println("Initialize DataStore")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataStore%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataStore **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_capacity = xmlElement.M_capacity

	/** RootElement **/
	object.M_isUnlimited = xmlElement.M_isUnlimited

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}
}

/** inititialisation of CancelEventDefinition **/
func (this *BPMSXmlFactory) InitCancelEventDefinition(xmlElement *BPMN20.XsdCancelEventDefinition, object *BPMN20.CancelEventDefinition) {
	log.Println("Initialize CancelEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CancelEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CancelEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Operation **/
func (this *BPMSXmlFactory) InitOperation(xmlElement *BPMN20.XsdOperation, object *BPMN20.Operation) {
	log.Println("Initialize Operation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Operation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Operation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref inMessageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_inMessageRef) > 0 {
		if _, ok := this.m_object[object.M_id]["inMessageRef"]; !ok {
			this.m_object[object.M_id]["inMessageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inMessageRef"] = append(this.m_object[object.M_id]["inMessageRef"], xmlElement.M_inMessageRef)
	}

	/** Init ref outMessageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_outMessageRef != nil {
		if _, ok := this.m_object[object.M_id]["outMessageRef"]; !ok {
			this.m_object[object.M_id]["outMessageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outMessageRef"] = append(this.m_object[object.M_id]["outMessageRef"], *xmlElement.M_outMessageRef)
	}

	/** Init ref errorRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_errorRef); i++ {
		if _, ok := this.m_object[object.M_id]["errorRef"]; !ok {
			this.m_object[object.M_id]["errorRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["errorRef"] = append(this.m_object[object.M_id]["errorRef"], xmlElement.M_errorRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref implementationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_implementationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["implementationRef"]; !ok {
			this.m_object[object.M_id]["implementationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["implementationRef"] = append(this.m_object[object.M_id]["implementationRef"], xmlElement.M_implementationRef)
	}
}

/** inititialisation of BPMNEdge **/
func (this *BPMSXmlFactory) InitBPMNEdge(xmlElement *BPMNDI.XsdBPMNEdge, object *BPMNDI.BPMNEdge) {
	log.Println("Initialize BPMNEdge")
	if len(object.UUID) == 0 {
		object.UUID = "BPMNDI.BPMNEdge%" + Utility.RandomUUID()
	}

	/** Init extension **/

	/** BPMNEdge **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init Point **/
	object.M_waypoint = make([]*DC.Point, 0)
	for i := 0; i < len(xmlElement.M_waypoint); i++ {
		val := new(DC.Point)
		this.InitPoint(xmlElement.M_waypoint[i], val)
		object.M_waypoint = append(object.M_waypoint, val)

		/** association initialisation **/
	}

	/** Init BPMNLabel **/
	if xmlElement.M_BPMNLabel != nil {
		object.M_BPMNLabel = new(BPMNDI.BPMNLabel)
		this.InitBPMNLabel(xmlElement.M_BPMNLabel, object.M_BPMNLabel)

		/** association initialisation **/
	}
	/** Init bpmnElement **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if len(xmlElement.M_bpmnElement) > 0 {
		if _, ok := this.m_object[object.M_id]; !ok {
			this.m_object[object.M_id] = make(map[string][]string)
		}
		if _, ok := this.m_object[object.M_id]["bpmnElement"]; !ok {
			this.m_object[object.M_id]["bpmnElement"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElement"] = append(this.m_object[object.M_id]["bpmnElement"], xmlElement.M_bpmnElement)
		object.M_bpmnElement = xmlElement.M_bpmnElement
	}

	/** Init ref sourceElement **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_sourceElement) > 0 {
		if _, ok := this.m_object[object.M_id]["sourceElement"]; !ok {
			this.m_object[object.M_id]["sourceElement"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceElement"] = append(this.m_object[object.M_id]["sourceElement"], xmlElement.M_sourceElement)
	}

	/** Init ref targetElement **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetElement) > 0 {
		if _, ok := this.m_object[object.M_id]["targetElement"]; !ok {
			this.m_object[object.M_id]["targetElement"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetElement"] = append(this.m_object[object.M_id]["targetElement"], xmlElement.M_targetElement)
	}

	/** MessageVisibleKind **/
	if xmlElement.M_messageVisibleKind == "##initiating" {
		object.M_messageVisibleKind = BPMNDI.MessageVisibleKind_Initiating
	} else if xmlElement.M_messageVisibleKind == "##non_initiating" {
		object.M_messageVisibleKind = BPMNDI.MessageVisibleKind_Non_initiating
	}
}

/** inititialisation of Error **/
func (this *BPMSXmlFactory) InitError(xmlElement *BPMN20.XsdError, object *BPMN20.Error) {
	log.Println("Initialize Error")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Error%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Error **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_errorCode = xmlElement.M_errorCode

	/** Init ref structureRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_structureRef) > 0 {
		if _, ok := this.m_object[object.M_id]["structureRef"]; !ok {
			this.m_object[object.M_id]["structureRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["structureRef"] = append(this.m_object[object.M_id]["structureRef"], xmlElement.M_structureRef)
	}
}

/** inititialisation of Assignment **/
func (this *BPMSXmlFactory) InitAssignment(xmlElement *BPMN20.XsdAssignment, object *BPMN20.Assignment) {
	log.Println("Initialize Assignment")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Assignment%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Assignment **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if xmlElement.M_from != nil {
		if object.M_from == nil {
			object.M_from = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_from)), object.M_from)

		/** association initialisation **/
		object.M_from.SetAssignmentPtr(object)
	}

	/** Init formalExpression **/
	if xmlElement.M_to != nil {
		if object.M_to == nil {
			object.M_to = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_to)), object.M_to)

		/** association initialisation **/
		object.M_to.SetAssignmentPtr(object)
	}
}

/** inititialisation of CallChoreography **/
func (this *BPMSXmlFactory) InitCallChoreography(xmlElement *BPMN20.XsdCallChoreography, object *BPMN20.CallChoreography) {
	log.Println("Initialize CallChoreography")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CallChoreography%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CallChoreography **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetChoreographyActivityPtr(object)
	}

	/** Init ref initiatingParticipantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_initiatingParticipantRef) > 0 {
		if _, ok := this.m_object[object.M_id]["initiatingParticipantRef"]; !ok {
			this.m_object[object.M_id]["initiatingParticipantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["initiatingParticipantRef"] = append(this.m_object[object.M_id]["initiatingParticipantRef"], xmlElement.M_initiatingParticipantRef)
	}

	/** ChoreographyLoopType **/
	if xmlElement.M_loopType == "##None" {
		object.M_loopType = BPMN20.ChoreographyLoopType_None
	} else if xmlElement.M_loopType == "##Standard" {
		object.M_loopType = BPMN20.ChoreographyLoopType_Standard
	} else if xmlElement.M_loopType == "##MultiInstanceSequential" {
		object.M_loopType = BPMN20.ChoreographyLoopType_MultiInstanceSequential
	} else if xmlElement.M_loopType == "##MultiInstanceParallel" {
		object.M_loopType = BPMN20.ChoreographyLoopType_MultiInstanceParallel
	} else {
		object.M_loopType = BPMN20.ChoreographyLoopType_None
	}

	/** Init participantAssociation **/
	object.M_participantAssociation = make([]*BPMN20.ParticipantAssociation, 0)
	for i := 0; i < len(xmlElement.M_participantAssociation); i++ {
		val := new(BPMN20.ParticipantAssociation)
		this.InitParticipantAssociation(xmlElement.M_participantAssociation[i], val)
		object.M_participantAssociation = append(object.M_participantAssociation, val)

		/** association initialisation **/
	}

	/** Init ref calledChoreographyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_calledChoreographyRef) > 0 {
		if _, ok := this.m_object[object.M_id]["calledChoreographyRef"]; !ok {
			this.m_object[object.M_id]["calledChoreographyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["calledChoreographyRef"] = append(this.m_object[object.M_id]["calledChoreographyRef"], xmlElement.M_calledChoreographyRef)
	}
}

/** inititialisation of EndEvent **/
func (this *BPMSXmlFactory) InitEndEvent(xmlElement *BPMN20.XsdEndEvent, object *BPMN20.EndEvent) {
	log.Println("Initialize EndEvent")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.EndEvent%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** EndEvent **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetEventPtr(object)
	}

	/** Init dataInput **/
	object.M_dataInput = make([]*BPMN20.DataInput, 0)
	for i := 0; i < len(xmlElement.M_dataInput); i++ {
		val := new(BPMN20.DataInput)
		this.InitDataInput(xmlElement.M_dataInput[i], val)
		object.M_dataInput = append(object.M_dataInput, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init inputSet **/
	if xmlElement.M_inputSet != nil {
		object.M_inputSet = new(BPMN20.InputSet)
		this.InitInputSet(xmlElement.M_inputSet, object.M_inputSet)

		/** association initialisation **/
		object.M_inputSet.SetThrowEventPtr(object)
	}

	/** Init cancelEventDefinition **/
	object.M_eventDefinition = make([]BPMN20.EventDefinition, 0)
	for i := 0; i < len(xmlElement.M_eventDefinition_0); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_eventDefinition_0[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_1); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_eventDefinition_1[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_2); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_eventDefinition_2[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_3); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_eventDefinition_3[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_4); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_eventDefinition_4[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_5); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_eventDefinition_5[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_6); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_eventDefinition_6[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_7); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_eventDefinition_7[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_8); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_eventDefinition_8[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_eventDefinition_9); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_eventDefinition_9[i], val)
		object.M_eventDefinition = append(object.M_eventDefinition, val)

		/** association initialisation **/
		val.SetThrowEventPtr(object)
	}

	/** Init ref eventDefinitionRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_eventDefinitionRef); i++ {
		if _, ok := this.m_object[object.M_id]["eventDefinitionRef"]; !ok {
			this.m_object[object.M_id]["eventDefinitionRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["eventDefinitionRef"] = append(this.m_object[object.M_id]["eventDefinitionRef"], xmlElement.M_eventDefinitionRef[i])
	}

}

/** inititialisation of ParallelGateway **/
func (this *BPMSXmlFactory) InitParallelGateway(xmlElement *BPMN20.XsdParallelGateway, object *BPMN20.ParallelGateway) {
	log.Println("Initialize ParallelGateway")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ParallelGateway%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ParallelGateway **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** GatewayDirection **/
	if xmlElement.M_gatewayDirection == "##Unspecified" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	} else if xmlElement.M_gatewayDirection == "##Converging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Converging
	} else if xmlElement.M_gatewayDirection == "##Diverging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Diverging
	} else if xmlElement.M_gatewayDirection == "##Mixed" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Mixed
	} else {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	}
}

/** inititialisation of SubChoreography **/
func (this *BPMSXmlFactory) InitSubChoreography(xmlElement *BPMN20.XsdSubChoreography, object *BPMN20.SubChoreography) {
	log.Println("Initialize SubChoreography")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.SubChoreography%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** SubChoreography **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetChoreographyActivityPtr(object)
	}

	/** Init ref initiatingParticipantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_initiatingParticipantRef) > 0 {
		if _, ok := this.m_object[object.M_id]["initiatingParticipantRef"]; !ok {
			this.m_object[object.M_id]["initiatingParticipantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["initiatingParticipantRef"] = append(this.m_object[object.M_id]["initiatingParticipantRef"], xmlElement.M_initiatingParticipantRef)
	}

	/** ChoreographyLoopType **/
	if xmlElement.M_loopType == "##None" {
		object.M_loopType = BPMN20.ChoreographyLoopType_None
	} else if xmlElement.M_loopType == "##Standard" {
		object.M_loopType = BPMN20.ChoreographyLoopType_Standard
	} else if xmlElement.M_loopType == "##MultiInstanceSequential" {
		object.M_loopType = BPMN20.ChoreographyLoopType_MultiInstanceSequential
	} else if xmlElement.M_loopType == "##MultiInstanceParallel" {
		object.M_loopType = BPMN20.ChoreographyLoopType_MultiInstanceParallel
	} else {
		object.M_loopType = BPMN20.ChoreographyLoopType_None
	}

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubChoreographyPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubChoreographyPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetSubChoreographyPtr(object)
	}
}

/** inititialisation of CorrelationProperty **/
func (this *BPMSXmlFactory) InitCorrelationProperty(xmlElement *BPMN20.XsdCorrelationProperty, object *BPMN20.CorrelationProperty) {
	log.Println("Initialize CorrelationProperty")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CorrelationProperty%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CorrelationProperty **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init correlationPropertyRetrievalExpression **/
	object.M_correlationPropertyRetrievalExpression = make([]*BPMN20.CorrelationPropertyRetrievalExpression, 0)
	for i := 0; i < len(xmlElement.M_correlationPropertyRetrievalExpression); i++ {
		val := new(BPMN20.CorrelationPropertyRetrievalExpression)
		this.InitCorrelationPropertyRetrievalExpression(xmlElement.M_correlationPropertyRetrievalExpression[i], val)
		object.M_correlationPropertyRetrievalExpression = append(object.M_correlationPropertyRetrievalExpression, val)

		/** association initialisation **/
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init ref type **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_type) > 0 {
		if _, ok := this.m_object[object.M_id]["type"]; !ok {
			this.m_object[object.M_id]["type"] = make([]string, 0)
		}
		this.m_object[object.M_id]["type"] = append(this.m_object[object.M_id]["type"], xmlElement.M_type)
	}
}

/** inititialisation of LinkEventDefinition **/
func (this *BPMSXmlFactory) InitLinkEventDefinition(xmlElement *BPMN20.XsdLinkEventDefinition, object *BPMN20.LinkEventDefinition) {
	log.Println("Initialize LinkEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.LinkEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** LinkEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref source **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_source); i++ {
		if _, ok := this.m_object[object.M_id]["source"]; !ok {
			this.m_object[object.M_id]["source"] = make([]string, 0)
		}
		this.m_object[object.M_id]["source"] = append(this.m_object[object.M_id]["source"], xmlElement.M_source[i])
	}

	/** Init ref target **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_target != nil {
		if _, ok := this.m_object[object.M_id]["target"]; !ok {
			this.m_object[object.M_id]["target"] = make([]string, 0)
		}
		this.m_object[object.M_id]["target"] = append(this.m_object[object.M_id]["target"], *xmlElement.M_target)
	}

	/** EventDefinition **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of DataOutputAssociation **/
func (this *BPMSXmlFactory) InitDataOutputAssociation(xmlElement *BPMN20.XsdDataOutputAssociation, object *BPMN20.DataOutputAssociation) {
	log.Println("Initialize DataOutputAssociation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataOutputAssociation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataOutputAssociation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_sourceRef); i++ {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], xmlElement.M_sourceRef[i])
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetRef) > 0 {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], xmlElement.M_targetRef)
	}

	/** Init formalExpression **/
	if xmlElement.M_transformation != nil {
		if object.M_transformation == nil {
			object.M_transformation = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_transformation)), object.M_transformation)

		/** association initialisation **/
		object.M_transformation.SetDataAssociationPtr(object)
	}

	/** Init assignment **/
	object.M_assignment = make([]*BPMN20.Assignment, 0)
	for i := 0; i < len(xmlElement.M_assignment); i++ {
		val := new(BPMN20.Assignment)
		this.InitAssignment(xmlElement.M_assignment[i], val)
		object.M_assignment = append(object.M_assignment, val)

		/** association initialisation **/
		val.SetDataAssociationPtr(object)
	}
}

/** inititialisation of DataStoreReference **/
func (this *BPMSXmlFactory) InitDataStoreReference(xmlElement *BPMN20.XsdDataStoreReference, object *BPMN20.DataStoreReference) {
	log.Println("Initialize DataStoreReference")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataStoreReference%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataStoreReference **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init dataState **/
	if xmlElement.M_dataState != nil {
		object.M_dataState = new(BPMN20.DataState)
		this.InitDataState(xmlElement.M_dataState, object.M_dataState)

		/** association initialisation **/
	}

	/** Init ref itemSubjectRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_itemSubjectRef) > 0 {
		if _, ok := this.m_object[object.M_id]["itemSubjectRef"]; !ok {
			this.m_object[object.M_id]["itemSubjectRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["itemSubjectRef"] = append(this.m_object[object.M_id]["itemSubjectRef"], xmlElement.M_itemSubjectRef)
	}

	/** Init ref dataStoreRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_dataStoreRef) > 0 {
		if _, ok := this.m_object[object.M_id]["dataStoreRef"]; !ok {
			this.m_object[object.M_id]["dataStoreRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["dataStoreRef"] = append(this.m_object[object.M_id]["dataStoreRef"], xmlElement.M_dataStoreRef)
	}
}

/** inititialisation of ReceiveTask **/
func (this *BPMSXmlFactory) InitReceiveTask(xmlElement *BPMN20.XsdReceiveTask, object *BPMN20.ReceiveTask) {
	log.Println("Initialize ReceiveTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ReceiveTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ReceiveTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_WebService
		}
	}

	/** Task **/
	object.M_instantiate = xmlElement.M_instantiate

	/** Init ref messageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_messageRef) > 0 {
		if _, ok := this.m_object[object.M_id]["messageRef"]; !ok {
			this.m_object[object.M_id]["messageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageRef"] = append(this.m_object[object.M_id]["messageRef"], xmlElement.M_messageRef)
	}

	/** Init ref operationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_operationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["operationRef"]; !ok {
			this.m_object[object.M_id]["operationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["operationRef"] = append(this.m_object[object.M_id]["operationRef"], xmlElement.M_operationRef)
	}
}

/** inititialisation of Collaboration **/
func (this *BPMSXmlFactory) InitCollaboration(xmlElement *BPMN20.XsdCollaboration, object *BPMN20.Collaboration_impl) {
	log.Println("Initialize Collaboration")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Collaboration_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Collaboration **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init participant **/
	object.M_participant = make([]*BPMN20.Participant, 0)
	for i := 0; i < len(xmlElement.M_participant); i++ {
		val := new(BPMN20.Participant)
		this.InitParticipant(xmlElement.M_participant[i], val)
		object.M_participant = append(object.M_participant, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlow **/
	object.M_messageFlow = make([]*BPMN20.MessageFlow, 0)
	for i := 0; i < len(xmlElement.M_messageFlow); i++ {
		val := new(BPMN20.MessageFlow)
		this.InitMessageFlow(xmlElement.M_messageFlow[i], val)
		object.M_messageFlow = append(object.M_messageFlow, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init callConversation **/
	object.M_conversationNode = make([]BPMN20.ConversationNode, 0)
	for i := 0; i < len(xmlElement.M_conversationNode_0); i++ {
		val := new(BPMN20.CallConversation)
		this.InitCallConversation(xmlElement.M_conversationNode_0[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_1); i++ {
		val := new(BPMN20.Conversation)
		this.InitConversation(xmlElement.M_conversationNode_1[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init subConversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_2); i++ {
		val := new(BPMN20.SubConversation)
		this.InitSubConversation(xmlElement.M_conversationNode_2[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversationAssociation **/
	object.M_conversationAssociation = make([]*BPMN20.ConversationAssociation, 0)
	for i := 0; i < len(xmlElement.M_conversationAssociation); i++ {
		val := new(BPMN20.ConversationAssociation)
		this.InitConversationAssociation(xmlElement.M_conversationAssociation[i], val)
		object.M_conversationAssociation = append(object.M_conversationAssociation, val)

		/** association initialisation **/
	}

	/** Init participantAssociation **/
	object.M_participantAssociation = make([]*BPMN20.ParticipantAssociation, 0)
	for i := 0; i < len(xmlElement.M_participantAssociation); i++ {
		val := new(BPMN20.ParticipantAssociation)
		this.InitParticipantAssociation(xmlElement.M_participantAssociation[i], val)
		object.M_participantAssociation = append(object.M_participantAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlowAssociation **/
	object.M_messageFlowAssociation = make([]*BPMN20.MessageFlowAssociation, 0)
	for i := 0; i < len(xmlElement.M_messageFlowAssociation); i++ {
		val := new(BPMN20.MessageFlowAssociation)
		this.InitMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], val)
		object.M_messageFlowAssociation = append(object.M_messageFlowAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init ref choreographyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_choreographyRef); i++ {
		if _, ok := this.m_object[object.M_id]["choreographyRef"]; !ok {
			this.m_object[object.M_id]["choreographyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["choreographyRef"] = append(this.m_object[object.M_id]["choreographyRef"], xmlElement.M_choreographyRef[i])
	}

	/** Init conversationLink **/
	object.M_conversationLink = make([]*BPMN20.ConversationLink, 0)
	for i := 0; i < len(xmlElement.M_conversationLink); i++ {
		val := new(BPMN20.ConversationLink)
		this.InitConversationLink(xmlElement.M_conversationLink[i], val)
		object.M_conversationLink = append(object.M_conversationLink, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_isClosed = xmlElement.M_isClosed
}

/** inititialisation of ParticipantMultiplicity **/
func (this *BPMSXmlFactory) InitParticipantMultiplicity(xmlElement *BPMN20.XsdParticipantMultiplicity, object *BPMN20.ParticipantMultiplicity) {
	log.Println("Initialize ParticipantMultiplicity")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ParticipantMultiplicity%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ParticipantMultiplicity **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_minimum = xmlElement.M_minimum

	/** BaseElement **/
	object.M_maximum = xmlElement.M_maximum
}

/** inititialisation of DataState **/
func (this *BPMSXmlFactory) InitDataState(xmlElement *BPMN20.XsdDataState, object *BPMN20.DataState) {
	log.Println("Initialize DataState")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataState%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataState **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of ResourceAssignmentExpression **/
func (this *BPMSXmlFactory) InitResourceAssignmentExpression(xmlElement *BPMN20.XsdResourceAssignmentExpression, object *BPMN20.ResourceAssignmentExpression) {
	log.Println("Initialize ResourceAssignmentExpression")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ResourceAssignmentExpression%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ResourceAssignmentExpression **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if xmlElement.M_expression_0 != nil {
		if object.M_expression == nil {
			object.M_expression = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression(xmlElement.M_expression_0, object.M_expression.(*BPMN20.FormalExpression))

		/** association initialisation **/
		object.M_expression.(*BPMN20.FormalExpression).SetResourceAssignmentExpressionPtr(object)
	}

	/** Init expression **/
	if xmlElement.M_expression_1 != nil {
		if object.M_expression == nil {
			object.M_expression = new(BPMN20.FormalExpression)
		}
		this.InitExpression(xmlElement.M_expression_1, object.M_expression.(*BPMN20.Expression_impl))

		/** association initialisation **/
		object.M_expression.(*BPMN20.Expression_impl).SetResourceAssignmentExpressionPtr(object)
	}
}

/** inititialisation of PartnerEntity **/
func (this *BPMSXmlFactory) InitPartnerEntity(xmlElement *BPMN20.XsdPartnerEntity, object *BPMN20.PartnerEntity) {
	log.Println("Initialize PartnerEntity")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.PartnerEntity%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** PartnerEntity **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of GlobalChoreographyTask **/
func (this *BPMSXmlFactory) InitGlobalChoreographyTask(xmlElement *BPMN20.XsdGlobalChoreographyTask, object *BPMN20.GlobalChoreographyTask) {
	log.Println("Initialize GlobalChoreographyTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalChoreographyTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalChoreographyTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init participant **/
	object.M_participant = make([]*BPMN20.Participant, 0)
	for i := 0; i < len(xmlElement.M_participant); i++ {
		val := new(BPMN20.Participant)
		this.InitParticipant(xmlElement.M_participant[i], val)
		object.M_participant = append(object.M_participant, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlow **/
	object.M_messageFlow = make([]*BPMN20.MessageFlow, 0)
	for i := 0; i < len(xmlElement.M_messageFlow); i++ {
		val := new(BPMN20.MessageFlow)
		this.InitMessageFlow(xmlElement.M_messageFlow[i], val)
		object.M_messageFlow = append(object.M_messageFlow, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init callConversation **/
	object.M_conversationNode = make([]BPMN20.ConversationNode, 0)
	for i := 0; i < len(xmlElement.M_conversationNode_0); i++ {
		val := new(BPMN20.CallConversation)
		this.InitCallConversation(xmlElement.M_conversationNode_0[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_1); i++ {
		val := new(BPMN20.Conversation)
		this.InitConversation(xmlElement.M_conversationNode_1[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init subConversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_2); i++ {
		val := new(BPMN20.SubConversation)
		this.InitSubConversation(xmlElement.M_conversationNode_2[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversationAssociation **/
	object.M_conversationAssociation = make([]*BPMN20.ConversationAssociation, 0)
	for i := 0; i < len(xmlElement.M_conversationAssociation); i++ {
		val := new(BPMN20.ConversationAssociation)
		this.InitConversationAssociation(xmlElement.M_conversationAssociation[i], val)
		object.M_conversationAssociation = append(object.M_conversationAssociation, val)

		/** association initialisation **/
	}

	/** Init participantAssociation **/
	object.M_participantAssociation = make([]*BPMN20.ParticipantAssociation, 0)
	for i := 0; i < len(xmlElement.M_participantAssociation); i++ {
		val := new(BPMN20.ParticipantAssociation)
		this.InitParticipantAssociation(xmlElement.M_participantAssociation[i], val)
		object.M_participantAssociation = append(object.M_participantAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlowAssociation **/
	object.M_messageFlowAssociation = make([]*BPMN20.MessageFlowAssociation, 0)
	for i := 0; i < len(xmlElement.M_messageFlowAssociation); i++ {
		val := new(BPMN20.MessageFlowAssociation)
		this.InitMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], val)
		object.M_messageFlowAssociation = append(object.M_messageFlowAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init ref choreographyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_choreographyRef); i++ {
		if _, ok := this.m_object[object.M_id]["choreographyRef"]; !ok {
			this.m_object[object.M_id]["choreographyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["choreographyRef"] = append(this.m_object[object.M_id]["choreographyRef"], xmlElement.M_choreographyRef[i])
	}

	/** Init conversationLink **/
	object.M_conversationLink = make([]*BPMN20.ConversationLink, 0)
	for i := 0; i < len(xmlElement.M_conversationLink); i++ {
		val := new(BPMN20.ConversationLink)
		this.InitConversationLink(xmlElement.M_conversationLink[i], val)
		object.M_conversationLink = append(object.M_conversationLink, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_isClosed = xmlElement.M_isClosed

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init ref initiatingParticipantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_initiatingParticipantRef) > 0 {
		if _, ok := this.m_object[object.M_id]["initiatingParticipantRef"]; !ok {
			this.m_object[object.M_id]["initiatingParticipantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["initiatingParticipantRef"] = append(this.m_object[object.M_id]["initiatingParticipantRef"], xmlElement.M_initiatingParticipantRef)
	}
}

/** inititialisation of Relationship **/
func (this *BPMSXmlFactory) InitRelationship(xmlElement *BPMN20.XsdRelationship, object *BPMN20.Relationship) {
	log.Println("Initialize Relationship")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Relationship%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Relationship **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref source **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_source); i++ {
		if _, ok := this.m_object[object.M_id]["source"]; !ok {
			this.m_object[object.M_id]["source"] = make([]string, 0)
		}
		this.m_object[object.M_id]["source"] = append(this.m_object[object.M_id]["source"], xmlElement.M_source[i])
	}

	/** Init ref target **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_target); i++ {
		if _, ok := this.m_object[object.M_id]["target"]; !ok {
			this.m_object[object.M_id]["target"] = make([]string, 0)
		}
		this.m_object[object.M_id]["target"] = append(this.m_object[object.M_id]["target"], xmlElement.M_target[i])
	}

	/** BaseElement **/
	object.M_type = xmlElement.M_type

	/** RelationshipDirection **/
	if xmlElement.M_direction == "##None" {
		object.M_direction = BPMN20.RelationshipDirection_None
	} else if xmlElement.M_direction == "##Forward" {
		object.M_direction = BPMN20.RelationshipDirection_Forward
	} else if xmlElement.M_direction == "##Backward" {
		object.M_direction = BPMN20.RelationshipDirection_Backward
	} else if xmlElement.M_direction == "##Both" {
		object.M_direction = BPMN20.RelationshipDirection_Both
	}
}

/** inititialisation of Definitions **/
func (this *BPMSXmlFactory) InitDefinitions(xmlElement *BPMN20.XsdDefinitions, object *BPMN20.Definitions) {
	log.Println("Initialize Definitions")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Definitions%" + Utility.RandomUUID()
	}

	/** Init import **/
	object.M_import = make([]*BPMN20.Import, 0)
	for i := 0; i < len(xmlElement.M_import); i++ {
		val := new(BPMN20.Import)
		this.InitImport(xmlElement.M_import[i], val)
		object.M_import = append(object.M_import, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init extension **/
	object.M_extension = make([]*BPMN20.Extension, 0)
	for i := 0; i < len(xmlElement.M_extension); i++ {
		val := new(BPMN20.Extension)
		this.InitExtension(xmlElement.M_extension[i], val)
		object.M_extension = append(object.M_extension, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init category **/
	object.M_rootElement = make([]BPMN20.RootElement, 0)
	for i := 0; i < len(xmlElement.M_rootElement_0); i++ {
		val := new(BPMN20.Category)
		this.InitCategory(xmlElement.M_rootElement_0[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalChoreographyTask **/
	for i := 0; i < len(xmlElement.M_rootElement_1); i++ {
		val := new(BPMN20.GlobalChoreographyTask)
		this.InitGlobalChoreographyTask(xmlElement.M_rootElement_1[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init choreography **/
	for i := 0; i < len(xmlElement.M_rootElement_2); i++ {
		val := new(BPMN20.Choreography_impl)
		this.InitChoreography(xmlElement.M_rootElement_2[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalConversation **/
	for i := 0; i < len(xmlElement.M_rootElement_3); i++ {
		val := new(BPMN20.GlobalConversation)
		this.InitGlobalConversation(xmlElement.M_rootElement_3[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init collaboration **/
	for i := 0; i < len(xmlElement.M_rootElement_4); i++ {
		val := new(BPMN20.Collaboration_impl)
		this.InitCollaboration(xmlElement.M_rootElement_4[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init correlationProperty **/
	for i := 0; i < len(xmlElement.M_rootElement_5); i++ {
		val := new(BPMN20.CorrelationProperty)
		this.InitCorrelationProperty(xmlElement.M_rootElement_5[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init dataStore **/
	for i := 0; i < len(xmlElement.M_rootElement_6); i++ {
		val := new(BPMN20.DataStore)
		this.InitDataStore(xmlElement.M_rootElement_6[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init endPoint **/
	for i := 0; i < len(xmlElement.M_rootElement_7); i++ {
		val := new(BPMN20.EndPoint)
		this.InitEndPoint(xmlElement.M_rootElement_7[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init error **/
	for i := 0; i < len(xmlElement.M_rootElement_8); i++ {
		val := new(BPMN20.Error)
		this.InitError(xmlElement.M_rootElement_8[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init escalation **/
	for i := 0; i < len(xmlElement.M_rootElement_9); i++ {
		val := new(BPMN20.Escalation)
		this.InitEscalation(xmlElement.M_rootElement_9[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init cancelEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_10); i++ {
		val := new(BPMN20.CancelEventDefinition)
		this.InitCancelEventDefinition(xmlElement.M_rootElement_10[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init compensateEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_11); i++ {
		val := new(BPMN20.CompensateEventDefinition)
		this.InitCompensateEventDefinition(xmlElement.M_rootElement_11[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init conditionalEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_12); i++ {
		val := new(BPMN20.ConditionalEventDefinition)
		this.InitConditionalEventDefinition(xmlElement.M_rootElement_12[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init errorEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_13); i++ {
		val := new(BPMN20.ErrorEventDefinition)
		this.InitErrorEventDefinition(xmlElement.M_rootElement_13[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init escalationEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_14); i++ {
		val := new(BPMN20.EscalationEventDefinition)
		this.InitEscalationEventDefinition(xmlElement.M_rootElement_14[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init linkEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_15); i++ {
		val := new(BPMN20.LinkEventDefinition)
		this.InitLinkEventDefinition(xmlElement.M_rootElement_15[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init messageEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_16); i++ {
		val := new(BPMN20.MessageEventDefinition)
		this.InitMessageEventDefinition(xmlElement.M_rootElement_16[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init signalEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_17); i++ {
		val := new(BPMN20.SignalEventDefinition)
		this.InitSignalEventDefinition(xmlElement.M_rootElement_17[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init terminateEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_18); i++ {
		val := new(BPMN20.TerminateEventDefinition)
		this.InitTerminateEventDefinition(xmlElement.M_rootElement_18[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init timerEventDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_19); i++ {
		val := new(BPMN20.TimerEventDefinition)
		this.InitTimerEventDefinition(xmlElement.M_rootElement_19[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalBusinessRuleTask **/
	for i := 0; i < len(xmlElement.M_rootElement_20); i++ {
		val := new(BPMN20.GlobalBusinessRuleTask)
		this.InitGlobalBusinessRuleTask(xmlElement.M_rootElement_20[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalManualTask **/
	for i := 0; i < len(xmlElement.M_rootElement_21); i++ {
		val := new(BPMN20.GlobalManualTask)
		this.InitGlobalManualTask(xmlElement.M_rootElement_21[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalScriptTask **/
	for i := 0; i < len(xmlElement.M_rootElement_22); i++ {
		val := new(BPMN20.GlobalScriptTask)
		this.InitGlobalScriptTask(xmlElement.M_rootElement_22[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalTask **/
	for i := 0; i < len(xmlElement.M_rootElement_23); i++ {
		val := new(BPMN20.GlobalTask_impl)
		this.InitGlobalTask(xmlElement.M_rootElement_23[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init globalUserTask **/
	for i := 0; i < len(xmlElement.M_rootElement_24); i++ {
		val := new(BPMN20.GlobalUserTask)
		this.InitGlobalUserTask(xmlElement.M_rootElement_24[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init interface **/
	for i := 0; i < len(xmlElement.M_rootElement_25); i++ {
		val := new(BPMN20.Interface)
		this.InitInterface(xmlElement.M_rootElement_25[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init itemDefinition **/
	for i := 0; i < len(xmlElement.M_rootElement_26); i++ {
		val := new(BPMN20.ItemDefinition)
		this.InitItemDefinition(xmlElement.M_rootElement_26[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init message **/
	for i := 0; i < len(xmlElement.M_rootElement_27); i++ {
		val := new(BPMN20.Message)
		this.InitMessage(xmlElement.M_rootElement_27[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init partnerEntity **/
	for i := 0; i < len(xmlElement.M_rootElement_28); i++ {
		val := new(BPMN20.PartnerEntity)
		this.InitPartnerEntity(xmlElement.M_rootElement_28[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init partnerRole **/
	for i := 0; i < len(xmlElement.M_rootElement_29); i++ {
		val := new(BPMN20.PartnerRole)
		this.InitPartnerRole(xmlElement.M_rootElement_29[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init process **/
	for i := 0; i < len(xmlElement.M_rootElement_30); i++ {
		val := new(BPMN20.Process)
		this.InitProcess(xmlElement.M_rootElement_30[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init resource **/
	for i := 0; i < len(xmlElement.M_rootElement_31); i++ {
		val := new(BPMN20.Resource)
		this.InitResource(xmlElement.M_rootElement_31[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init signal **/
	for i := 0; i < len(xmlElement.M_rootElement_32); i++ {
		val := new(BPMN20.Signal)
		this.InitSignal(xmlElement.M_rootElement_32[i], val)
		object.M_rootElement = append(object.M_rootElement, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Init BPMNDiagram **/
	object.M_BPMNDiagram = make([]*BPMNDI.BPMNDiagram, 0)
	for i := 0; i < len(xmlElement.M_BPMNDiagram); i++ {
		val := new(BPMNDI.BPMNDiagram)
		this.InitBPMNDiagram(xmlElement.M_BPMNDiagram[i], val)
		object.M_BPMNDiagram = append(object.M_BPMNDiagram, val)

		/** association initialisation **/
	}

	/** Init relationship **/
	object.M_relationship = make([]*BPMN20.Relationship, 0)
	for i := 0; i < len(xmlElement.M_relationship); i++ {
		val := new(BPMN20.Relationship)
		this.InitRelationship(xmlElement.M_relationship[i], val)
		object.M_relationship = append(object.M_relationship, val)

		/** association initialisation **/
		val.SetDefinitionsPtr(object)
	}

	/** Definitions **/
	object.M_id = xmlElement.M_id

	object.M_name = xmlElement.M_name

	/** Definitions **/
	object.M_targetNamespace = xmlElement.M_targetNamespace

	/** Definitions **/
	object.M_expressionLanguage = xmlElement.M_expressionLanguage

	/** Definitions **/
	object.M_typeLanguage = xmlElement.M_typeLanguage

	/** Definitions **/
	object.M_exporter = xmlElement.M_exporter

	/** Definitions **/
	object.M_exporterVersion = xmlElement.M_exporterVersion
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of EndPoint **/
func (this *BPMSXmlFactory) InitEndPoint(xmlElement *BPMN20.XsdEndPoint, object *BPMN20.EndPoint) {
	log.Println("Initialize EndPoint")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.EndPoint%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** EndPoint **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of TerminateEventDefinition **/
func (this *BPMSXmlFactory) InitTerminateEventDefinition(xmlElement *BPMN20.XsdTerminateEventDefinition, object *BPMN20.TerminateEventDefinition) {
	log.Println("Initialize TerminateEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.TerminateEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** TerminateEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of SendTask **/
func (this *BPMSXmlFactory) InitSendTask(xmlElement *BPMN20.XsdSendTask, object *BPMN20.SendTask) {
	log.Println("Initialize SendTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.SendTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** SendTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_WebService
		}
	}

	/** Init ref messageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_messageRef) > 0 {
		if _, ok := this.m_object[object.M_id]["messageRef"]; !ok {
			this.m_object[object.M_id]["messageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageRef"] = append(this.m_object[object.M_id]["messageRef"], xmlElement.M_messageRef)
	}

	/** Init ref operationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_operationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["operationRef"]; !ok {
			this.m_object[object.M_id]["operationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["operationRef"] = append(this.m_object[object.M_id]["operationRef"], xmlElement.M_operationRef)
	}
}

/** inititialisation of CorrelationPropertyBinding **/
func (this *BPMSXmlFactory) InitCorrelationPropertyBinding(xmlElement *BPMN20.XsdCorrelationPropertyBinding, object *BPMN20.CorrelationPropertyBinding) {
	log.Println("Initialize CorrelationPropertyBinding")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CorrelationPropertyBinding%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CorrelationPropertyBinding **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if object.M_dataPath == nil {
		object.M_dataPath = new(BPMN20.FormalExpression)
	}
	this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(&xmlElement.M_dataPath)), object.M_dataPath)

	/** association initialisation **/
	object.M_dataPath.SetCorrelationPropertyBindingPtr(object)

	/** Init ref correlationPropertyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_correlationPropertyRef) > 0 {
		if _, ok := this.m_object[object.M_id]["correlationPropertyRef"]; !ok {
			this.m_object[object.M_id]["correlationPropertyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["correlationPropertyRef"] = append(this.m_object[object.M_id]["correlationPropertyRef"], xmlElement.M_correlationPropertyRef)
	}
}

/** inititialisation of Point **/
func (this *BPMSXmlFactory) InitPoint(xmlElement *DC.XsdPoint, object *DC.Point) {
	log.Println("Initialize Point")
	if len(object.UUID) == 0 {
		object.UUID = "DC.Point%" + Utility.RandomUUID()
	}

	/** Point **/
	object.M_x = xmlElement.M_x

	/** Point **/
	object.M_y = xmlElement.M_y
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ConversationLink **/
func (this *BPMSXmlFactory) InitConversationLink(xmlElement *BPMN20.XsdConversationLink, object *BPMN20.ConversationLink) {
	log.Println("Initialize ConversationLink")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ConversationLink%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ConversationLink **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_sourceRef) > 0 {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], xmlElement.M_sourceRef)
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetRef) > 0 {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], xmlElement.M_targetRef)
	}
}

/** inititialisation of PotentialOwner **/
func (this *BPMSXmlFactory) InitPotentialOwner(xmlElement *BPMN20.XsdPotentialOwner, object *BPMN20.PotentialOwner) {
	log.Println("Initialize PotentialOwner")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.PotentialOwner%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** PotentialOwner **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref resourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_resourceRef != nil {
		if _, ok := this.m_object[object.M_id]["resourceRef"]; !ok {
			this.m_object[object.M_id]["resourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["resourceRef"] = append(this.m_object[object.M_id]["resourceRef"], *xmlElement.M_resourceRef)
	}

	/** Init resourceParameterBinding **/
	object.M_resourceParameterBinding = make([]*BPMN20.ResourceParameterBinding, 0)
	for i := 0; i < len(xmlElement.M_resourceParameterBinding); i++ {
		val := new(BPMN20.ResourceParameterBinding)
		this.InitResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], val)
		object.M_resourceParameterBinding = append(object.M_resourceParameterBinding, val)

		/** association initialisation **/
	}

	/** Init resourceAssignmentExpression **/
	if xmlElement.M_resourceAssignmentExpression != nil {
		object.M_resourceAssignmentExpression = new(BPMN20.ResourceAssignmentExpression)
		this.InitResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)

		/** association initialisation **/
	}
}

/** inititialisation of GlobalManualTask **/
func (this *BPMSXmlFactory) InitGlobalManualTask(xmlElement *BPMN20.XsdGlobalManualTask, object *BPMN20.GlobalManualTask) {
	log.Println("Initialize GlobalManualTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalManualTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalManualTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref supportedInterfaceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_supportedInterfaceRef); i++ {
		if _, ok := this.m_object[object.M_id]["supportedInterfaceRef"]; !ok {
			this.m_object[object.M_id]["supportedInterfaceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["supportedInterfaceRef"] = append(this.m_object[object.M_id]["supportedInterfaceRef"], xmlElement.M_supportedInterfaceRef[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetCallableElementPtr(object)
	}

	/** Init ioBinding **/
	object.M_ioBinding = make([]*BPMN20.InputOutputBinding, 0)
	for i := 0; i < len(xmlElement.M_ioBinding); i++ {
		val := new(BPMN20.InputOutputBinding)
		this.InitInputOutputBinding(xmlElement.M_ioBinding[i], val)
		object.M_ioBinding = append(object.M_ioBinding, val)

		/** association initialisation **/
		val.SetCallableElementPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetGlobalTaskPtr(object)
	}
}

/** inititialisation of LaneSet **/
func (this *BPMSXmlFactory) InitLaneSet(xmlElement *BPMN20.XsdLaneSet, object *BPMN20.LaneSet) {
	log.Println("Initialize LaneSet")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.LaneSet%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** LaneSet **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init lane **/
	object.M_lane = make([]*BPMN20.Lane, 0)
	for i := 0; i < len(xmlElement.M_lane); i++ {
		val := new(BPMN20.Lane)
		this.InitLane(xmlElement.M_lane[i], val)
		object.M_lane = append(object.M_lane, val)

		/** association initialisation **/
		val.SetLaneSetPtr(object)
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of Signal **/
func (this *BPMSXmlFactory) InitSignal(xmlElement *BPMN20.XsdSignal, object *BPMN20.Signal) {
	log.Println("Initialize Signal")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Signal%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Signal **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** Init ref structureRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_structureRef) > 0 {
		if _, ok := this.m_object[object.M_id]["structureRef"]; !ok {
			this.m_object[object.M_id]["structureRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["structureRef"] = append(this.m_object[object.M_id]["structureRef"], xmlElement.M_structureRef)
	}
}

/** inititialisation of BPMNPlane **/
func (this *BPMSXmlFactory) InitBPMNPlane(xmlElement *BPMNDI.XsdBPMNPlane, object *BPMNDI.BPMNPlane) {
	log.Println("Initialize BPMNPlane")
	if len(object.UUID) == 0 {
		object.UUID = "BPMNDI.BPMNPlane%" + Utility.RandomUUID()
	}

	/** Init extension **/

	/** BPMNPlane **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init BPMNShape **/
	object.M_DiagramElement = make([]DI.DiagramElement, 0)
	for i := 0; i < len(xmlElement.M_DiagramElement_0); i++ {
		val := new(BPMNDI.BPMNShape)
		this.InitBPMNShape(xmlElement.M_DiagramElement_0[i], val)
		object.M_DiagramElement = append(object.M_DiagramElement, val)

		/** association initialisation **/
		val.SetPlanePtr(object)
	}

	/** Init BPMNEdge **/
	for i := 0; i < len(xmlElement.M_DiagramElement_1); i++ {
		val := new(BPMNDI.BPMNEdge)
		this.InitBPMNEdge(xmlElement.M_DiagramElement_1[i], val)
		object.M_DiagramElement = append(object.M_DiagramElement, val)

		/** association initialisation **/
		val.SetPlanePtr(object)
	}
	/** Init bpmnElement **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if len(xmlElement.M_bpmnElement) > 0 {
		if _, ok := this.m_object[object.M_id]; !ok {
			this.m_object[object.M_id] = make(map[string][]string)
		}
		if _, ok := this.m_object[object.M_id]["bpmnElement"]; !ok {
			this.m_object[object.M_id]["bpmnElement"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElement"] = append(this.m_object[object.M_id]["bpmnElement"], xmlElement.M_bpmnElement)
		object.M_bpmnElement = xmlElement.M_bpmnElement
	}
}

/** inititialisation of ExtensionElements **/
func (this *BPMSXmlFactory) InitExtensionElements(xmlElement *BPMN20.XsdExtensionElements, object *BPMN20.ExtensionElements) {
	log.Println("Initialize ExtensionElements")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ExtensionElements%" + Utility.RandomUUID()
	}
	object.M_value = xmlElement.M_value
	this.m_references["ExtensionElements"] = object
}

/** inititialisation of MultiInstanceLoopCharacteristics **/
func (this *BPMSXmlFactory) InitMultiInstanceLoopCharacteristics(xmlElement *BPMN20.XsdMultiInstanceLoopCharacteristics, object *BPMN20.MultiInstanceLoopCharacteristics) {
	log.Println("Initialize MultiInstanceLoopCharacteristics")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.MultiInstanceLoopCharacteristics%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** MultiInstanceLoopCharacteristics **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init formalExpression **/
	if xmlElement.M_loopCardinality != nil {
		if object.M_loopCardinality == nil {
			object.M_loopCardinality = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_loopCardinality)), object.M_loopCardinality)

		/** association initialisation **/
		object.M_loopCardinality.SetMultiInstanceLoopCharacteristicsPtr(object)
	}

	/** Init ref loopDataInputRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_loopDataInputRef != nil {
		if _, ok := this.m_object[object.M_id]["loopDataInputRef"]; !ok {
			this.m_object[object.M_id]["loopDataInputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["loopDataInputRef"] = append(this.m_object[object.M_id]["loopDataInputRef"], *xmlElement.M_loopDataInputRef)
	}

	/** Init ref loopDataOutputRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_loopDataOutputRef != nil {
		if _, ok := this.m_object[object.M_id]["loopDataOutputRef"]; !ok {
			this.m_object[object.M_id]["loopDataOutputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["loopDataOutputRef"] = append(this.m_object[object.M_id]["loopDataOutputRef"], *xmlElement.M_loopDataOutputRef)
	}

	/** Init dataInput **/
	if xmlElement.M_inputDataItem != nil {
		if object.M_inputDataItem == nil {
			object.M_inputDataItem = new(BPMN20.DataInput)
		}
		this.InitDataInput((*BPMN20.XsdDataInput)(unsafe.Pointer(xmlElement.M_inputDataItem)), object.M_inputDataItem)

		/** association initialisation **/
		object.M_inputDataItem.SetMultiInstanceLoopCharacteristicsPtr(object)
	}

	/** Init dataOutput **/
	if xmlElement.M_outputDataItem != nil {
		if object.M_outputDataItem == nil {
			object.M_outputDataItem = new(BPMN20.DataOutput)
		}
		this.InitDataOutput((*BPMN20.XsdDataOutput)(unsafe.Pointer(xmlElement.M_outputDataItem)), object.M_outputDataItem)

		/** association initialisation **/
		object.M_outputDataItem.SetMultiInstanceLoopCharacteristicsPtr(object)
	}

	/** Init complexBehaviorDefinition **/
	object.M_complexBehaviorDefinition = make([]*BPMN20.ComplexBehaviorDefinition, 0)
	for i := 0; i < len(xmlElement.M_complexBehaviorDefinition); i++ {
		val := new(BPMN20.ComplexBehaviorDefinition)
		this.InitComplexBehaviorDefinition(xmlElement.M_complexBehaviorDefinition[i], val)
		object.M_complexBehaviorDefinition = append(object.M_complexBehaviorDefinition, val)

		/** association initialisation **/
		val.SetMultiInstanceLoopCharacteristicsPtr(object)
	}

	/** Init formalExpression **/
	if xmlElement.M_completionCondition != nil {
		if object.M_completionCondition == nil {
			object.M_completionCondition = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_completionCondition)), object.M_completionCondition)

		/** association initialisation **/
		object.M_completionCondition.SetMultiInstanceLoopCharacteristicsPtr(object)
	}

	/** LoopCharacteristics **/
	object.M_isSequential = xmlElement.M_isSequential

	/** MultiInstanceFlowCondition **/
	if xmlElement.M_behavior == "##None" {
		object.M_behavior = BPMN20.MultiInstanceFlowCondition_None
	} else if xmlElement.M_behavior == "##One" {
		object.M_behavior = BPMN20.MultiInstanceFlowCondition_One
	} else if xmlElement.M_behavior == "##All" {
		object.M_behavior = BPMN20.MultiInstanceFlowCondition_All
	} else if xmlElement.M_behavior == "##Complex" {
		object.M_behavior = BPMN20.MultiInstanceFlowCondition_Complex
	} else {
		object.M_behavior = BPMN20.MultiInstanceFlowCondition_All
	}

	/** Init ref oneBehaviorEventRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_oneBehaviorEventRef) > 0 {
		if _, ok := this.m_object[object.M_id]["oneBehaviorEventRef"]; !ok {
			this.m_object[object.M_id]["oneBehaviorEventRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["oneBehaviorEventRef"] = append(this.m_object[object.M_id]["oneBehaviorEventRef"], xmlElement.M_oneBehaviorEventRef)
	}

	/** Init ref noneBehaviorEventRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_noneBehaviorEventRef) > 0 {
		if _, ok := this.m_object[object.M_id]["noneBehaviorEventRef"]; !ok {
			this.m_object[object.M_id]["noneBehaviorEventRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["noneBehaviorEventRef"] = append(this.m_object[object.M_id]["noneBehaviorEventRef"], xmlElement.M_noneBehaviorEventRef)
	}
}

/** inititialisation of UserTask **/
func (this *BPMSXmlFactory) InitUserTask(xmlElement *BPMN20.XsdUserTask, object *BPMN20.UserTask) {
	log.Println("Initialize UserTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.UserTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** UserTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}

	/** Init rendering **/
	object.M_rendering = make([]*BPMN20.Rendering, 0)
	for i := 0; i < len(xmlElement.M_rendering); i++ {
		val := new(BPMN20.Rendering)
		this.InitRendering(xmlElement.M_rendering[i], val)
		object.M_rendering = append(object.M_rendering, val)

		/** association initialisation **/
	}
	if !strings.HasPrefix(xmlElement.M_implementation, "##") {
		object.M_implementationStr = xmlElement.M_implementation
	} else {

		/** Implementation **/
		if xmlElement.M_implementation == "##unspecified" {
			object.M_implementation = BPMN20.Implementation_Unspecified
		} else if xmlElement.M_implementation == "##WebService" {
			object.M_implementation = BPMN20.Implementation_WebService
		} else {
			object.M_implementation = BPMN20.Implementation_Unspecified
		}
	}
}

/** inititialisation of GlobalConversation **/
func (this *BPMSXmlFactory) InitGlobalConversation(xmlElement *BPMN20.XsdGlobalConversation, object *BPMN20.GlobalConversation) {
	log.Println("Initialize GlobalConversation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.GlobalConversation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** GlobalConversation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init participant **/
	object.M_participant = make([]*BPMN20.Participant, 0)
	for i := 0; i < len(xmlElement.M_participant); i++ {
		val := new(BPMN20.Participant)
		this.InitParticipant(xmlElement.M_participant[i], val)
		object.M_participant = append(object.M_participant, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlow **/
	object.M_messageFlow = make([]*BPMN20.MessageFlow, 0)
	for i := 0; i < len(xmlElement.M_messageFlow); i++ {
		val := new(BPMN20.MessageFlow)
		this.InitMessageFlow(xmlElement.M_messageFlow[i], val)
		object.M_messageFlow = append(object.M_messageFlow, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init callConversation **/
	object.M_conversationNode = make([]BPMN20.ConversationNode, 0)
	for i := 0; i < len(xmlElement.M_conversationNode_0); i++ {
		val := new(BPMN20.CallConversation)
		this.InitCallConversation(xmlElement.M_conversationNode_0[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_1); i++ {
		val := new(BPMN20.Conversation)
		this.InitConversation(xmlElement.M_conversationNode_1[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init subConversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_2); i++ {
		val := new(BPMN20.SubConversation)
		this.InitSubConversation(xmlElement.M_conversationNode_2[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversationAssociation **/
	object.M_conversationAssociation = make([]*BPMN20.ConversationAssociation, 0)
	for i := 0; i < len(xmlElement.M_conversationAssociation); i++ {
		val := new(BPMN20.ConversationAssociation)
		this.InitConversationAssociation(xmlElement.M_conversationAssociation[i], val)
		object.M_conversationAssociation = append(object.M_conversationAssociation, val)

		/** association initialisation **/
	}

	/** Init participantAssociation **/
	object.M_participantAssociation = make([]*BPMN20.ParticipantAssociation, 0)
	for i := 0; i < len(xmlElement.M_participantAssociation); i++ {
		val := new(BPMN20.ParticipantAssociation)
		this.InitParticipantAssociation(xmlElement.M_participantAssociation[i], val)
		object.M_participantAssociation = append(object.M_participantAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlowAssociation **/
	object.M_messageFlowAssociation = make([]*BPMN20.MessageFlowAssociation, 0)
	for i := 0; i < len(xmlElement.M_messageFlowAssociation); i++ {
		val := new(BPMN20.MessageFlowAssociation)
		this.InitMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], val)
		object.M_messageFlowAssociation = append(object.M_messageFlowAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init ref choreographyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_choreographyRef); i++ {
		if _, ok := this.m_object[object.M_id]["choreographyRef"]; !ok {
			this.m_object[object.M_id]["choreographyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["choreographyRef"] = append(this.m_object[object.M_id]["choreographyRef"], xmlElement.M_choreographyRef[i])
	}

	/** Init conversationLink **/
	object.M_conversationLink = make([]*BPMN20.ConversationLink, 0)
	for i := 0; i < len(xmlElement.M_conversationLink); i++ {
		val := new(BPMN20.ConversationLink)
		this.InitConversationLink(xmlElement.M_conversationLink[i], val)
		object.M_conversationLink = append(object.M_conversationLink, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_isClosed = xmlElement.M_isClosed
}

/** inititialisation of SignalEventDefinition **/
func (this *BPMSXmlFactory) InitSignalEventDefinition(xmlElement *BPMN20.XsdSignalEventDefinition, object *BPMN20.SignalEventDefinition) {
	log.Println("Initialize SignalEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.SignalEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** SignalEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref signalRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_signalRef) > 0 {
		if _, ok := this.m_object[object.M_id]["signalRef"]; !ok {
			this.m_object[object.M_id]["signalRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["signalRef"] = append(this.m_object[object.M_id]["signalRef"], xmlElement.M_signalRef)
	}
}

/** inititialisation of CallConversation **/
func (this *BPMSXmlFactory) InitCallConversation(xmlElement *BPMN20.XsdCallConversation, object *BPMN20.CallConversation) {
	log.Println("Initialize CallConversation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CallConversation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CallConversation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** Init ref messageFlowRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_messageFlowRef); i++ {
		if _, ok := this.m_object[object.M_id]["messageFlowRef"]; !ok {
			this.m_object[object.M_id]["messageFlowRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageFlowRef"] = append(this.m_object[object.M_id]["messageFlowRef"], xmlElement.M_messageFlowRef[i])
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetConversationNodePtr(object)
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init participantAssociation **/
	object.M_participantAssociation = make([]*BPMN20.ParticipantAssociation, 0)
	for i := 0; i < len(xmlElement.M_participantAssociation); i++ {
		val := new(BPMN20.ParticipantAssociation)
		this.InitParticipantAssociation(xmlElement.M_participantAssociation[i], val)
		object.M_participantAssociation = append(object.M_participantAssociation, val)

		/** association initialisation **/
		val.SetCallConversationPtr(object)
	}

	/** Init ref calledCollaborationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_calledCollaborationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["calledCollaborationRef"]; !ok {
			this.m_object[object.M_id]["calledCollaborationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["calledCollaborationRef"] = append(this.m_object[object.M_id]["calledCollaborationRef"], xmlElement.M_calledCollaborationRef)
	}
}

/** inititialisation of InclusiveGateway **/
func (this *BPMSXmlFactory) InitInclusiveGateway(xmlElement *BPMN20.XsdInclusiveGateway, object *BPMN20.InclusiveGateway) {
	log.Println("Initialize InclusiveGateway")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.InclusiveGateway%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** InclusiveGateway **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** GatewayDirection **/
	if xmlElement.M_gatewayDirection == "##Unspecified" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	} else if xmlElement.M_gatewayDirection == "##Converging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Converging
	} else if xmlElement.M_gatewayDirection == "##Diverging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Diverging
	} else if xmlElement.M_gatewayDirection == "##Mixed" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Mixed
	} else {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	}

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
}

/** inititialisation of ScriptTask **/
func (this *BPMSXmlFactory) InitScriptTask(xmlElement *BPMN20.XsdScriptTask, object *BPMN20.ScriptTask) {
	log.Println("Initialize ScriptTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ScriptTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ScriptTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}

	/** Init script **/
	if xmlElement.M_script != nil {
		object.M_script = new(BPMN20.Script)
		this.InitScript(xmlElement.M_script, object.M_script)

		/** association initialisation **/
		object.M_script.SetScriptTaskPtr(object)
	}

	/** Task **/
	object.M_scriptFormat = xmlElement.M_scriptFormat
}

/** inititialisation of MessageFlow **/
func (this *BPMSXmlFactory) InitMessageFlow(xmlElement *BPMN20.XsdMessageFlow, object *BPMN20.MessageFlow) {
	log.Println("Initialize MessageFlow")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.MessageFlow%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** MessageFlow **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_sourceRef) > 0 {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], xmlElement.M_sourceRef)
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetRef) > 0 {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], xmlElement.M_targetRef)
	}

	/** Init ref messageRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_messageRef) > 0 {
		if _, ok := this.m_object[object.M_id]["messageRef"]; !ok {
			this.m_object[object.M_id]["messageRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageRef"] = append(this.m_object[object.M_id]["messageRef"], xmlElement.M_messageRef)
	}
}

/** inititialisation of CompensateEventDefinition **/
func (this *BPMSXmlFactory) InitCompensateEventDefinition(xmlElement *BPMN20.XsdCompensateEventDefinition, object *BPMN20.CompensateEventDefinition) {
	log.Println("Initialize CompensateEventDefinition")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.CompensateEventDefinition%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** CompensateEventDefinition **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** EventDefinition **/
	object.M_waitForCompletion = xmlElement.M_waitForCompletion

	/** Init ref activityRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_activityRef) > 0 {
		if _, ok := this.m_object[object.M_id]["activityRef"]; !ok {
			this.m_object[object.M_id]["activityRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["activityRef"] = append(this.m_object[object.M_id]["activityRef"], xmlElement.M_activityRef)
	}
}

/** inititialisation of Auditing **/
func (this *BPMSXmlFactory) InitAuditing(xmlElement *BPMN20.XsdAuditing, object *BPMN20.Auditing) {
	log.Println("Initialize Auditing")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Auditing%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Auditing **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Resource **/
func (this *BPMSXmlFactory) InitResource(xmlElement *BPMN20.XsdResource, object *BPMN20.Resource) {
	log.Println("Initialize Resource")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Resource%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Resource **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init resourceParameter **/
	object.M_resourceParameter = make([]*BPMN20.ResourceParameter, 0)
	for i := 0; i < len(xmlElement.M_resourceParameter); i++ {
		val := new(BPMN20.ResourceParameter)
		this.InitResourceParameter(xmlElement.M_resourceParameter[i], val)
		object.M_resourceParameter = append(object.M_resourceParameter, val)

		/** association initialisation **/
		val.SetResourcePtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of Font **/
func (this *BPMSXmlFactory) InitFont(xmlElement *DC.XsdFont, object *DC.Font) {
	log.Println("Initialize Font")
	if len(object.UUID) == 0 {
		object.UUID = "DC.Font%" + Utility.RandomUUID()
	}

	/** Font **/
	object.M_name = xmlElement.M_name

	/** Font **/
	object.M_size = xmlElement.M_size

	/** Font **/
	object.M_isBold = xmlElement.M_isBold

	/** Font **/
	object.M_isItalic = xmlElement.M_isItalic

	/** Font **/
	object.M_isUnderline = xmlElement.M_isUnderline

	/** Font **/
	object.M_isStrikeThrough = xmlElement.M_isStrikeThrough
	if len(object.M_name) > 0 {
		this.m_references[object.M_name] = object
	}
}

/** inititialisation of SubConversation **/
func (this *BPMSXmlFactory) InitSubConversation(xmlElement *BPMN20.XsdSubConversation, object *BPMN20.SubConversation) {
	log.Println("Initialize SubConversation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.SubConversation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** SubConversation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref participantRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_participantRef); i++ {
		if _, ok := this.m_object[object.M_id]["participantRef"]; !ok {
			this.m_object[object.M_id]["participantRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["participantRef"] = append(this.m_object[object.M_id]["participantRef"], xmlElement.M_participantRef[i])
	}

	/** Init ref messageFlowRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_messageFlowRef); i++ {
		if _, ok := this.m_object[object.M_id]["messageFlowRef"]; !ok {
			this.m_object[object.M_id]["messageFlowRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["messageFlowRef"] = append(this.m_object[object.M_id]["messageFlowRef"], xmlElement.M_messageFlowRef[i])
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetConversationNodePtr(object)
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init callConversation **/
	object.M_conversationNode = make([]BPMN20.ConversationNode, 0)
	for i := 0; i < len(xmlElement.M_conversationNode_0); i++ {
		val := new(BPMN20.CallConversation)
		this.InitCallConversation(xmlElement.M_conversationNode_0[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetSubConversationPtr(object)
	}

	/** Init conversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_1); i++ {
		val := new(BPMN20.Conversation)
		this.InitConversation(xmlElement.M_conversationNode_1[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetSubConversationPtr(object)
	}

	/** Init subConversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_2); i++ {
		val := new(BPMN20.SubConversation)
		this.InitSubConversation(xmlElement.M_conversationNode_2[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetSubConversationPtr(object)
	}
}

/** inititialisation of InputSet **/
func (this *BPMSXmlFactory) InitInputSet(xmlElement *BPMN20.XsdInputSet, object *BPMN20.InputSet) {
	log.Println("Initialize InputSet")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.InputSet%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** InputSet **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref dataInputRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_dataInputRefs); i++ {
		if _, ok := this.m_object[object.M_id]["dataInputRefs"]; !ok {
			this.m_object[object.M_id]["dataInputRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["dataInputRefs"] = append(this.m_object[object.M_id]["dataInputRefs"], xmlElement.M_dataInputRefs[i])
	}

	/** Init ref optionalInputRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_optionalInputRefs); i++ {
		if _, ok := this.m_object[object.M_id]["optionalInputRefs"]; !ok {
			this.m_object[object.M_id]["optionalInputRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["optionalInputRefs"] = append(this.m_object[object.M_id]["optionalInputRefs"], xmlElement.M_optionalInputRefs[i])
	}

	/** Init ref whileExecutingInputRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_whileExecutingInputRefs); i++ {
		if _, ok := this.m_object[object.M_id]["whileExecutingInputRefs"]; !ok {
			this.m_object[object.M_id]["whileExecutingInputRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["whileExecutingInputRefs"] = append(this.m_object[object.M_id]["whileExecutingInputRefs"], xmlElement.M_whileExecutingInputRefs[i])
	}

	/** Init ref outputSetRefs **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outputSetRefs); i++ {
		if _, ok := this.m_object[object.M_id]["outputSetRefs"]; !ok {
			this.m_object[object.M_id]["outputSetRefs"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outputSetRefs"] = append(this.m_object[object.M_id]["outputSetRefs"], xmlElement.M_outputSetRefs[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name
}

/** inititialisation of DataInputAssociation **/
func (this *BPMSXmlFactory) InitDataInputAssociation(xmlElement *BPMN20.XsdDataInputAssociation, object *BPMN20.DataInputAssociation) {
	log.Println("Initialize DataInputAssociation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.DataInputAssociation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** DataInputAssociation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_sourceRef); i++ {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], xmlElement.M_sourceRef[i])
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_targetRef) > 0 {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], xmlElement.M_targetRef)
	}

	/** Init formalExpression **/
	if xmlElement.M_transformation != nil {
		if object.M_transformation == nil {
			object.M_transformation = new(BPMN20.FormalExpression)
		}
		this.InitFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_transformation)), object.M_transformation)

		/** association initialisation **/
		object.M_transformation.SetDataAssociationPtr(object)
	}

	/** Init assignment **/
	object.M_assignment = make([]*BPMN20.Assignment, 0)
	for i := 0; i < len(xmlElement.M_assignment); i++ {
		val := new(BPMN20.Assignment)
		this.InitAssignment(xmlElement.M_assignment[i], val)
		object.M_assignment = append(object.M_assignment, val)

		/** association initialisation **/
		val.SetDataAssociationPtr(object)
	}
}

/** inititialisation of ManualTask **/
func (this *BPMSXmlFactory) InitManualTask(xmlElement *BPMN20.XsdManualTask, object *BPMN20.ManualTask) {
	log.Println("Initialize ManualTask")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ManualTask%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ManualTask **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** Init ioSpecification **/
	if xmlElement.M_ioSpecification != nil {
		object.M_ioSpecification = new(BPMN20.InputOutputSpecification)
		this.InitInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)

		/** association initialisation **/
		object.M_ioSpecification.SetActivityPtr(object)
	}

	/** Init property **/
	object.M_property = make([]*BPMN20.Property, 0)
	for i := 0; i < len(xmlElement.M_property); i++ {
		val := new(BPMN20.Property)
		this.InitProperty(xmlElement.M_property[i], val)
		object.M_property = append(object.M_property, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataInputAssociation **/
	object.M_dataInputAssociation = make([]*BPMN20.DataInputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataInputAssociation); i++ {
		val := new(BPMN20.DataInputAssociation)
		this.InitDataInputAssociation(xmlElement.M_dataInputAssociation[i], val)
		object.M_dataInputAssociation = append(object.M_dataInputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init dataOutputAssociation **/
	object.M_dataOutputAssociation = make([]*BPMN20.DataOutputAssociation, 0)
	for i := 0; i < len(xmlElement.M_dataOutputAssociation); i++ {
		val := new(BPMN20.DataOutputAssociation)
		this.InitDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], val)
		object.M_dataOutputAssociation = append(object.M_dataOutputAssociation, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init humanPerformer **/
	object.M_resourceRole = make([]BPMN20.ResourceRole, 0)
	for i := 0; i < len(xmlElement.M_resourceRole_0); i++ {
		val := new(BPMN20.HumanPerformer_impl)
		this.InitHumanPerformer(xmlElement.M_resourceRole_0[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init potentialOwner **/
	for i := 0; i < len(xmlElement.M_resourceRole_1); i++ {
		val := new(BPMN20.PotentialOwner)
		this.InitPotentialOwner(xmlElement.M_resourceRole_1[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init performer **/
	for i := 0; i < len(xmlElement.M_resourceRole_2); i++ {
		val := new(BPMN20.Performer_impl)
		this.InitPerformer(xmlElement.M_resourceRole_2[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init resourceRole **/
	for i := 0; i < len(xmlElement.M_resourceRole_3); i++ {
		val := new(BPMN20.ResourceRole_impl)
		this.InitResourceRole(xmlElement.M_resourceRole_3[i], val)
		object.M_resourceRole = append(object.M_resourceRole, val)

		/** association initialisation **/
		val.SetActivityPtr(object)
	}

	/** Init multiInstanceLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_0 != nil {
		object.M_loopCharacteristics = new(BPMN20.MultiInstanceLoopCharacteristics)
		this.InitMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.MultiInstanceLoopCharacteristics).SetActivityPtr(object)
	}

	/** Init standardLoopCharacteristics **/
	if xmlElement.M_loopCharacteristics_1 != nil {
		object.M_loopCharacteristics = new(BPMN20.StandardLoopCharacteristics)
		this.InitStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics))

		/** association initialisation **/
		object.M_loopCharacteristics.(*BPMN20.StandardLoopCharacteristics).SetActivityPtr(object)
	}

	/** FlowNode **/
	object.M_isForCompensation = xmlElement.M_isForCompensation

	/** FlowNode **/
	object.M_startQuantity = xmlElement.M_startQuantity

	/** FlowNode **/
	object.M_completionQuantity = xmlElement.M_completionQuantity

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
}

/** inititialisation of Choreography **/
func (this *BPMSXmlFactory) InitChoreography(xmlElement *BPMN20.XsdChoreography, object *BPMN20.Choreography_impl) {
	log.Println("Initialize Choreography")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Choreography_impl%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Choreography **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init participant **/
	object.M_participant = make([]*BPMN20.Participant, 0)
	for i := 0; i < len(xmlElement.M_participant); i++ {
		val := new(BPMN20.Participant)
		this.InitParticipant(xmlElement.M_participant[i], val)
		object.M_participant = append(object.M_participant, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlow **/
	object.M_messageFlow = make([]*BPMN20.MessageFlow, 0)
	for i := 0; i < len(xmlElement.M_messageFlow); i++ {
		val := new(BPMN20.MessageFlow)
		this.InitMessageFlow(xmlElement.M_messageFlow[i], val)
		object.M_messageFlow = append(object.M_messageFlow, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init association **/
	object.M_artifact = make([]BPMN20.Artifact, 0)
	for i := 0; i < len(xmlElement.M_artifact_0); i++ {
		val := new(BPMN20.Association)
		this.InitAssociation(xmlElement.M_artifact_0[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init group **/
	for i := 0; i < len(xmlElement.M_artifact_1); i++ {
		val := new(BPMN20.Group)
		this.InitGroup(xmlElement.M_artifact_1[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init textAnnotation **/
	for i := 0; i < len(xmlElement.M_artifact_2); i++ {
		val := new(BPMN20.TextAnnotation)
		this.InitTextAnnotation(xmlElement.M_artifact_2[i], val)
		object.M_artifact = append(object.M_artifact, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init callConversation **/
	object.M_conversationNode = make([]BPMN20.ConversationNode, 0)
	for i := 0; i < len(xmlElement.M_conversationNode_0); i++ {
		val := new(BPMN20.CallConversation)
		this.InitCallConversation(xmlElement.M_conversationNode_0[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_1); i++ {
		val := new(BPMN20.Conversation)
		this.InitConversation(xmlElement.M_conversationNode_1[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init subConversation **/
	for i := 0; i < len(xmlElement.M_conversationNode_2); i++ {
		val := new(BPMN20.SubConversation)
		this.InitSubConversation(xmlElement.M_conversationNode_2[i], val)
		object.M_conversationNode = append(object.M_conversationNode, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init conversationAssociation **/
	object.M_conversationAssociation = make([]*BPMN20.ConversationAssociation, 0)
	for i := 0; i < len(xmlElement.M_conversationAssociation); i++ {
		val := new(BPMN20.ConversationAssociation)
		this.InitConversationAssociation(xmlElement.M_conversationAssociation[i], val)
		object.M_conversationAssociation = append(object.M_conversationAssociation, val)

		/** association initialisation **/
	}

	/** Init participantAssociation **/
	object.M_participantAssociation = make([]*BPMN20.ParticipantAssociation, 0)
	for i := 0; i < len(xmlElement.M_participantAssociation); i++ {
		val := new(BPMN20.ParticipantAssociation)
		this.InitParticipantAssociation(xmlElement.M_participantAssociation[i], val)
		object.M_participantAssociation = append(object.M_participantAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init messageFlowAssociation **/
	object.M_messageFlowAssociation = make([]*BPMN20.MessageFlowAssociation, 0)
	for i := 0; i < len(xmlElement.M_messageFlowAssociation); i++ {
		val := new(BPMN20.MessageFlowAssociation)
		this.InitMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], val)
		object.M_messageFlowAssociation = append(object.M_messageFlowAssociation, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init correlationKey **/
	object.M_correlationKey = make([]*BPMN20.CorrelationKey, 0)
	for i := 0; i < len(xmlElement.M_correlationKey); i++ {
		val := new(BPMN20.CorrelationKey)
		this.InitCorrelationKey(xmlElement.M_correlationKey[i], val)
		object.M_correlationKey = append(object.M_correlationKey, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** Init ref choreographyRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_choreographyRef); i++ {
		if _, ok := this.m_object[object.M_id]["choreographyRef"]; !ok {
			this.m_object[object.M_id]["choreographyRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["choreographyRef"] = append(this.m_object[object.M_id]["choreographyRef"], xmlElement.M_choreographyRef[i])
	}

	/** Init conversationLink **/
	object.M_conversationLink = make([]*BPMN20.ConversationLink, 0)
	for i := 0; i < len(xmlElement.M_conversationLink); i++ {
		val := new(BPMN20.ConversationLink)
		this.InitConversationLink(xmlElement.M_conversationLink[i], val)
		object.M_conversationLink = append(object.M_conversationLink, val)

		/** association initialisation **/
		val.SetCollaborationPtr(object)
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_isClosed = xmlElement.M_isClosed

	/** Init adHocSubProcess **/
	object.M_flowElement = make([]BPMN20.FlowElement, 0)
	for i := 0; i < len(xmlElement.M_flowElement_0); i++ {
		val := new(BPMN20.AdHocSubProcess)
		this.InitAdHocSubProcess(xmlElement.M_flowElement_0[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init boundaryEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_1); i++ {
		val := new(BPMN20.BoundaryEvent)
		this.InitBoundaryEvent(xmlElement.M_flowElement_1[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init businessRuleTask **/
	for i := 0; i < len(xmlElement.M_flowElement_2); i++ {
		val := new(BPMN20.BusinessRuleTask)
		this.InitBusinessRuleTask(xmlElement.M_flowElement_2[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callActivity **/
	for i := 0; i < len(xmlElement.M_flowElement_3); i++ {
		val := new(BPMN20.CallActivity)
		this.InitCallActivity(xmlElement.M_flowElement_3[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init callChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_4); i++ {
		val := new(BPMN20.CallChoreography)
		this.InitCallChoreography(xmlElement.M_flowElement_4[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init choreographyTask **/
	for i := 0; i < len(xmlElement.M_flowElement_5); i++ {
		val := new(BPMN20.ChoreographyTask)
		this.InitChoreographyTask(xmlElement.M_flowElement_5[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init complexGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_6); i++ {
		val := new(BPMN20.ComplexGateway)
		this.InitComplexGateway(xmlElement.M_flowElement_6[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObject **/
	for i := 0; i < len(xmlElement.M_flowElement_7); i++ {
		val := new(BPMN20.DataObject)
		this.InitDataObject(xmlElement.M_flowElement_7[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataObjectReference **/
	for i := 0; i < len(xmlElement.M_flowElement_8); i++ {
		val := new(BPMN20.DataObjectReference)
		this.InitDataObjectReference(xmlElement.M_flowElement_8[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init dataStoreReference **/
	for i := 0; i < len(xmlElement.M_flowElement_9); i++ {
		val := new(BPMN20.DataStoreReference)
		this.InitDataStoreReference(xmlElement.M_flowElement_9[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init endEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_10); i++ {
		val := new(BPMN20.EndEvent)
		this.InitEndEvent(xmlElement.M_flowElement_10[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init eventBasedGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_11); i++ {
		val := new(BPMN20.EventBasedGateway)
		this.InitEventBasedGateway(xmlElement.M_flowElement_11[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init exclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_12); i++ {
		val := new(BPMN20.ExclusiveGateway)
		this.InitExclusiveGateway(xmlElement.M_flowElement_12[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init implicitThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_13); i++ {
		val := new(BPMN20.ImplicitThrowEvent)
		this.InitImplicitThrowEvent(xmlElement.M_flowElement_13[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init inclusiveGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_14); i++ {
		val := new(BPMN20.InclusiveGateway)
		this.InitInclusiveGateway(xmlElement.M_flowElement_14[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateCatchEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_15); i++ {
		val := new(BPMN20.IntermediateCatchEvent)
		this.InitIntermediateCatchEvent(xmlElement.M_flowElement_15[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init intermediateThrowEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_16); i++ {
		val := new(BPMN20.IntermediateThrowEvent)
		this.InitIntermediateThrowEvent(xmlElement.M_flowElement_16[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init manualTask **/
	for i := 0; i < len(xmlElement.M_flowElement_17); i++ {
		val := new(BPMN20.ManualTask)
		this.InitManualTask(xmlElement.M_flowElement_17[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init parallelGateway **/
	for i := 0; i < len(xmlElement.M_flowElement_18); i++ {
		val := new(BPMN20.ParallelGateway)
		this.InitParallelGateway(xmlElement.M_flowElement_18[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init receiveTask **/
	for i := 0; i < len(xmlElement.M_flowElement_19); i++ {
		val := new(BPMN20.ReceiveTask)
		this.InitReceiveTask(xmlElement.M_flowElement_19[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init scriptTask **/
	for i := 0; i < len(xmlElement.M_flowElement_20); i++ {
		val := new(BPMN20.ScriptTask)
		this.InitScriptTask(xmlElement.M_flowElement_20[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sendTask **/
	for i := 0; i < len(xmlElement.M_flowElement_21); i++ {
		val := new(BPMN20.SendTask)
		this.InitSendTask(xmlElement.M_flowElement_21[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init sequenceFlow **/
	for i := 0; i < len(xmlElement.M_flowElement_22); i++ {
		val := new(BPMN20.SequenceFlow)
		this.InitSequenceFlow(xmlElement.M_flowElement_22[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init serviceTask **/
	for i := 0; i < len(xmlElement.M_flowElement_23); i++ {
		val := new(BPMN20.ServiceTask)
		this.InitServiceTask(xmlElement.M_flowElement_23[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init startEvent **/
	for i := 0; i < len(xmlElement.M_flowElement_24); i++ {
		val := new(BPMN20.StartEvent)
		this.InitStartEvent(xmlElement.M_flowElement_24[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subChoreography **/
	for i := 0; i < len(xmlElement.M_flowElement_25); i++ {
		val := new(BPMN20.SubChoreography)
		this.InitSubChoreography(xmlElement.M_flowElement_25[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init subProcess **/
	for i := 0; i < len(xmlElement.M_flowElement_26); i++ {
		val := new(BPMN20.SubProcess_impl)
		this.InitSubProcess(xmlElement.M_flowElement_26[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init task **/
	for i := 0; i < len(xmlElement.M_flowElement_27); i++ {
		val := new(BPMN20.Task_impl)
		this.InitTask(xmlElement.M_flowElement_27[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init transaction **/
	for i := 0; i < len(xmlElement.M_flowElement_28); i++ {
		val := new(BPMN20.Transaction)
		this.InitTransaction(xmlElement.M_flowElement_28[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}

	/** Init userTask **/
	for i := 0; i < len(xmlElement.M_flowElement_29); i++ {
		val := new(BPMN20.UserTask)
		this.InitUserTask(xmlElement.M_flowElement_29[i], val)
		object.M_flowElement = append(object.M_flowElement, val)

		/** association initialisation **/
	}
}

/** inititialisation of BPMNShape **/
func (this *BPMSXmlFactory) InitBPMNShape(xmlElement *BPMNDI.XsdBPMNShape, object *BPMNDI.BPMNShape) {
	log.Println("Initialize BPMNShape")
	if len(object.UUID) == 0 {
		object.UUID = "BPMNDI.BPMNShape%" + Utility.RandomUUID()
	}

	/** Init extension **/

	/** BPMNShape **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init Bounds **/
	if object.M_Bounds == nil {
		object.M_Bounds = new(DC.Bounds)
	}
	this.InitBounds(&xmlElement.M_Bounds, object.M_Bounds)

	/** association initialisation **/

	/** Init BPMNLabel **/
	if xmlElement.M_BPMNLabel != nil {
		object.M_BPMNLabel = new(BPMNDI.BPMNLabel)
		this.InitBPMNLabel(xmlElement.M_BPMNLabel, object.M_BPMNLabel)

		/** association initialisation **/
	}
	/** Init bpmnElement **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if len(xmlElement.M_bpmnElement) > 0 {
		if _, ok := this.m_object[object.M_id]; !ok {
			this.m_object[object.M_id] = make(map[string][]string)
		}
		if _, ok := this.m_object[object.M_id]["bpmnElement"]; !ok {
			this.m_object[object.M_id]["bpmnElement"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElement"] = append(this.m_object[object.M_id]["bpmnElement"], xmlElement.M_bpmnElement)
		object.M_bpmnElement = xmlElement.M_bpmnElement
	}

	/** LabeledShape **/
	object.M_isHorizontal = xmlElement.M_isHorizontal

	/** LabeledShape **/
	object.M_isExpanded = xmlElement.M_isExpanded

	/** LabeledShape **/
	object.M_isMarkerVisible = xmlElement.M_isMarkerVisible

	/** LabeledShape **/
	object.M_isMessageVisible = xmlElement.M_isMessageVisible

	/** ParticipantBandKind **/
	if xmlElement.M_participantBandKind == "##top_initiating" {
		object.M_participantBandKind = BPMNDI.ParticipantBandKind_Top_initiating
	} else if xmlElement.M_participantBandKind == "##middle_initiating" {
		object.M_participantBandKind = BPMNDI.ParticipantBandKind_Middle_initiating
	} else if xmlElement.M_participantBandKind == "##bottom_initiating" {
		object.M_participantBandKind = BPMNDI.ParticipantBandKind_Bottom_initiating
	} else if xmlElement.M_participantBandKind == "##top_non_initiating" {
		object.M_participantBandKind = BPMNDI.ParticipantBandKind_Top_non_initiating
	} else if xmlElement.M_participantBandKind == "##middle_non_initiating" {
		object.M_participantBandKind = BPMNDI.ParticipantBandKind_Middle_non_initiating
	} else if xmlElement.M_participantBandKind == "##bottom_non_initiating" {
		object.M_participantBandKind = BPMNDI.ParticipantBandKind_Bottom_non_initiating
	}

	/** Init ref choreographyActivityShape **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_choreographyActivityShape) > 0 {
		if _, ok := this.m_object[object.M_id]["choreographyActivityShape"]; !ok {
			this.m_object[object.M_id]["choreographyActivityShape"] = make([]string, 0)
		}
		this.m_object[object.M_id]["choreographyActivityShape"] = append(this.m_object[object.M_id]["choreographyActivityShape"], xmlElement.M_choreographyActivityShape)
	}
}

/** inititialisation of Documentation **/
func (this *BPMSXmlFactory) InitDocumentation(xmlElement *BPMN20.XsdDocumentation, object *BPMN20.Documentation) {
	log.Println("Initialize Documentation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Documentation%" + Utility.RandomUUID()
	}
	object.M_text = xmlElement.M_text

	/** Documentation **/
	object.M_textFormat = xmlElement.M_textFormat
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of FormalExpression **/
func (this *BPMSXmlFactory) InitFormalExpression(xmlElement *BPMN20.XsdFormalExpression, object *BPMN20.FormalExpression) {
	log.Println("Initialize FormalExpression")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.FormalExpression%" + Utility.RandomUUID()
	}

	/** Expression **/
	object.M_language = xmlElement.M_language

	/** Init ref evaluatesToTypeRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_evaluatesToTypeRef) > 0 {
		if _, ok := this.m_object[object.M_id]["evaluatesToTypeRef"]; !ok {
			this.m_object[object.M_id]["evaluatesToTypeRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["evaluatesToTypeRef"] = append(this.m_object[object.M_id]["evaluatesToTypeRef"], xmlElement.M_evaluatesToTypeRef)
	}
	/** other content **/
	exprStr := xmlElement.M_other
	object.SetOther(exprStr)
}

/** inititialisation of Escalation **/
func (this *BPMSXmlFactory) InitEscalation(xmlElement *BPMN20.XsdEscalation, object *BPMN20.Escalation) {
	log.Println("Initialize Escalation")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.Escalation%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** Escalation **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	object.M_name = xmlElement.M_name

	/** RootElement **/
	object.M_escalationCode = xmlElement.M_escalationCode

	/** Init ref structureRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_structureRef) > 0 {
		if _, ok := this.m_object[object.M_id]["structureRef"]; !ok {
			this.m_object[object.M_id]["structureRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["structureRef"] = append(this.m_object[object.M_id]["structureRef"], xmlElement.M_structureRef)
	}
}

/** inititialisation of InputOutputBinding **/
func (this *BPMSXmlFactory) InitInputOutputBinding(xmlElement *BPMN20.XsdInputOutputBinding, object *BPMN20.InputOutputBinding) {
	log.Println("Initialize InputOutputBinding")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.InputOutputBinding%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** InputOutputBinding **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref operationRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_operationRef) > 0 {
		if _, ok := this.m_object[object.M_id]["operationRef"]; !ok {
			this.m_object[object.M_id]["operationRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["operationRef"] = append(this.m_object[object.M_id]["operationRef"], xmlElement.M_operationRef)
	}

	/** Init ref inputDataRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_inputDataRef) > 0 {
		if _, ok := this.m_object[object.M_id]["inputDataRef"]; !ok {
			this.m_object[object.M_id]["inputDataRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inputDataRef"] = append(this.m_object[object.M_id]["inputDataRef"], xmlElement.M_inputDataRef)
	}

	/** Init ref outputDataRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_outputDataRef) > 0 {
		if _, ok := this.m_object[object.M_id]["outputDataRef"]; !ok {
			this.m_object[object.M_id]["outputDataRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outputDataRef"] = append(this.m_object[object.M_id]["outputDataRef"], xmlElement.M_outputDataRef)
	}
}

/** inititialisation of ExclusiveGateway **/
func (this *BPMSXmlFactory) InitExclusiveGateway(xmlElement *BPMN20.XsdExclusiveGateway, object *BPMN20.ExclusiveGateway) {
	log.Println("Initialize ExclusiveGateway")
	if len(object.UUID) == 0 {
		object.UUID = "BPMN20.ExclusiveGateway%" + Utility.RandomUUID()
	}

	/** Init documentation **/
	object.M_documentation = make([]*BPMN20.Documentation, 0)
	for i := 0; i < len(xmlElement.M_documentation); i++ {
		val := new(BPMN20.Documentation)
		this.InitDocumentation(xmlElement.M_documentation[i], val)
		object.M_documentation = append(object.M_documentation, val)

		/** association initialisation **/
		val.SetBaseElementPtr(object)
	}

	/** Init extensionElements **/
	if xmlElement.M_extensionElements != nil {
		object.M_extensionElements = new(BPMN20.ExtensionElements)
		this.InitExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)

		/** association initialisation **/
		object.M_extensionElements.SetBaseElementPtr(object)
	}

	/** ExclusiveGateway **/
	object.M_id = xmlElement.M_id
	//	object.M_other= xmlElement.M_other
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init auditing **/
	if xmlElement.M_auditing != nil {
		object.M_auditing = new(BPMN20.Auditing)
		this.InitAuditing(xmlElement.M_auditing, object.M_auditing)

		/** association initialisation **/
		object.M_auditing.SetFlowElementPtr(object)
	}

	/** Init monitoring **/
	if xmlElement.M_monitoring != nil {
		object.M_monitoring = new(BPMN20.Monitoring)
		this.InitMonitoring(xmlElement.M_monitoring, object.M_monitoring)

		/** association initialisation **/
		object.M_monitoring.SetFlowElementPtr(object)
	}

	/** Init ref categoryValueRef **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_categoryValueRef); i++ {
		if _, ok := this.m_object[object.M_id]["categoryValueRef"]; !ok {
			this.m_object[object.M_id]["categoryValueRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["categoryValueRef"] = append(this.m_object[object.M_id]["categoryValueRef"], xmlElement.M_categoryValueRef[i])
	}

	/** BaseElement **/
	object.M_name = xmlElement.M_name

	/** Init ref incoming **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_incoming); i++ {
		if _, ok := this.m_object[object.M_id]["incoming"]; !ok {
			this.m_object[object.M_id]["incoming"] = make([]string, 0)
		}
		this.m_object[object.M_id]["incoming"] = append(this.m_object[object.M_id]["incoming"], xmlElement.M_incoming[i])
	}

	/** Init ref outgoing **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outgoing); i++ {
		if _, ok := this.m_object[object.M_id]["outgoing"]; !ok {
			this.m_object[object.M_id]["outgoing"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outgoing"] = append(this.m_object[object.M_id]["outgoing"], xmlElement.M_outgoing[i])
	}

	/** GatewayDirection **/
	if xmlElement.M_gatewayDirection == "##Unspecified" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	} else if xmlElement.M_gatewayDirection == "##Converging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Converging
	} else if xmlElement.M_gatewayDirection == "##Diverging" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Diverging
	} else if xmlElement.M_gatewayDirection == "##Mixed" {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Mixed
	} else {
		object.M_gatewayDirection = BPMN20.GatewayDirection_Unspecified
	}

	/** Init ref default **/
	if len(object.M_id) == 0 {
		object.M_id = uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_default) > 0 {
		if _, ok := this.m_object[object.M_id]["default"]; !ok {
			this.m_object[object.M_id]["default"] = make([]string, 0)
		}
		this.m_object[object.M_id]["default"] = append(this.m_object[object.M_id]["default"], xmlElement.M_default)
	}
}

/** serialysation of Escalation **/
func (this *BPMSXmlFactory) SerialyzeEscalation(xmlElement *BPMN20.XsdEscalation, object *BPMN20.Escalation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Escalation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_escalationCode = object.M_escalationCode

	/** Serialyze ref structureRef **/
	xmlElement.M_structureRef = object.M_structureRef
}

/** serialysation of Font **/
func (this *BPMSXmlFactory) SerialyzeFont(xmlElement *DC.XsdFont, object *DC.Font) {
	if xmlElement == nil {
		return
	}

	/** Font **/
	xmlElement.M_name = object.M_name

	/** Font **/
	xmlElement.M_size = object.M_size

	/** Font **/
	xmlElement.M_isBold = object.M_isBold

	/** Font **/
	xmlElement.M_isItalic = object.M_isItalic

	/** Font **/
	xmlElement.M_isUnderline = object.M_isUnderline

	/** Font **/
	xmlElement.M_isStrikeThrough = object.M_isStrikeThrough
	if len(object.M_name) > 0 {
		this.m_references[object.M_name] = object
	}
}

/** serialysation of Group **/
func (this *BPMSXmlFactory) SerialyzeGroup(xmlElement *BPMN20.XsdGroup, object *BPMN20.Group) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Group **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef
}

/** serialysation of MessageFlowAssociation **/
func (this *BPMSXmlFactory) SerialyzeMessageFlowAssociation(xmlElement *BPMN20.XsdMessageFlowAssociation, object *BPMN20.MessageFlowAssociation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** MessageFlowAssociation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref innerMessageFlowRef **/
	xmlElement.M_innerMessageFlowRef = object.M_innerMessageFlowRef

	/** Serialyze ref outerMessageFlowRef **/
	xmlElement.M_outerMessageFlowRef = object.M_outerMessageFlowRef
}

/** serialysation of GlobalUserTask **/
func (this *BPMSXmlFactory) SerialyzeGlobalUserTask(xmlElement *BPMN20.XsdGlobalUserTask, object *BPMN20.GlobalUserTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalUserTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref supportedInterfaceRef **/
	xmlElement.M_supportedInterfaceRef = object.M_supportedInterfaceRef

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze InputOutputBinding **/
	if len(object.M_ioBinding) > 0 {
		xmlElement.M_ioBinding = make([]*BPMN20.XsdInputOutputBinding, 0)
	}

	/** Now I will save the value of ioBinding **/
	for i := 0; i < len(object.M_ioBinding); i++ {
		xmlElement.M_ioBinding = append(xmlElement.M_ioBinding, new(BPMN20.XsdInputOutputBinding))
		this.SerialyzeInputOutputBinding(xmlElement.M_ioBinding[i], object.M_ioBinding[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:ResourceRole")
		}
	}

	/** Serialyze Rendering **/
	if len(object.M_rendering) > 0 {
		xmlElement.M_rendering = make([]*BPMN20.XsdRendering, 0)
	}

	/** Now I will save the value of rendering **/
	for i := 0; i < len(object.M_rendering); i++ {
		xmlElement.M_rendering = append(xmlElement.M_rendering, new(BPMN20.XsdRendering))
		this.SerialyzeRendering(xmlElement.M_rendering[i], object.M_rendering[i])
	}
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##unspecified"
		}
	}
}

/** serialysation of Operation **/
func (this *BPMSXmlFactory) SerialyzeOperation(xmlElement *BPMN20.XsdOperation, object *BPMN20.Operation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Operation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref inMessageRef **/
	xmlElement.M_inMessageRef = object.M_inMessageRef

	/** Serialyze ref outMessageRef **/
	xmlElement.M_outMessageRef = &object.M_outMessageRef

	/** Serialyze ref errorRef **/
	xmlElement.M_errorRef = object.M_errorRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref implementationRef **/
	xmlElement.M_implementationRef = object.M_implementationRef
}

/** serialysation of CallConversation **/
func (this *BPMSXmlFactory) SerialyzeCallConversation(xmlElement *BPMN20.XsdCallConversation, object *BPMN20.CallConversation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CallConversation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** Serialyze ref messageFlowRef **/
	xmlElement.M_messageFlowRef = object.M_messageFlowRef

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ParticipantAssociation **/
	if len(object.M_participantAssociation) > 0 {
		xmlElement.M_participantAssociation = make([]*BPMN20.XsdParticipantAssociation, 0)
	}

	/** Now I will save the value of participantAssociation **/
	for i := 0; i < len(object.M_participantAssociation); i++ {
		xmlElement.M_participantAssociation = append(xmlElement.M_participantAssociation, new(BPMN20.XsdParticipantAssociation))
		this.SerialyzeParticipantAssociation(xmlElement.M_participantAssociation[i], object.M_participantAssociation[i])
	}

	/** Serialyze ref calledCollaborationRef **/
	xmlElement.M_calledCollaborationRef = object.M_calledCollaborationRef
}

/** serialysation of Lane **/
func (this *BPMSXmlFactory) SerialyzeLane(xmlElement *BPMN20.XsdLane, object *BPMN20.Lane) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Lane **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze BaseElement **/

	/** Now I will save the value of partitionElement **/
	if object.M_partitionElement != nil {
		switch v := object.M_partitionElement.(type) {
		}
	}

	/** Serialyze ref flowNodeRef **/
	xmlElement.M_flowNodeRef = object.M_flowNodeRef

	/** Serialyze LaneSet **/
	if object.M_childLaneSet != nil {
		xmlElement.M_childLaneSet = new(BPMN20.XsdChildLaneSet)
	}

	/** Now I will save the value of childLaneSet **/
	if object.M_childLaneSet != nil {
		this.SerialyzeLaneSet((*BPMN20.XsdLaneSet)(unsafe.Pointer(xmlElement.M_childLaneSet)), object.M_childLaneSet)
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref partitionElementRef **/
	xmlElement.M_partitionElementRef = object.M_partitionElementRef
}

/** serialysation of CorrelationSubscription **/
func (this *BPMSXmlFactory) SerialyzeCorrelationSubscription(xmlElement *BPMN20.XsdCorrelationSubscription, object *BPMN20.CorrelationSubscription) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CorrelationSubscription **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze CorrelationPropertyBinding **/
	if len(object.M_correlationPropertyBinding) > 0 {
		xmlElement.M_correlationPropertyBinding = make([]*BPMN20.XsdCorrelationPropertyBinding, 0)
	}

	/** Now I will save the value of correlationPropertyBinding **/
	for i := 0; i < len(object.M_correlationPropertyBinding); i++ {
		xmlElement.M_correlationPropertyBinding = append(xmlElement.M_correlationPropertyBinding, new(BPMN20.XsdCorrelationPropertyBinding))
		this.SerialyzeCorrelationPropertyBinding(xmlElement.M_correlationPropertyBinding[i], object.M_correlationPropertyBinding[i])
	}

	/** Serialyze ref correlationKeyRef **/
	xmlElement.M_correlationKeyRef = object.M_correlationKeyRef
}

/** serialysation of ExclusiveGateway **/
func (this *BPMSXmlFactory) SerialyzeExclusiveGateway(xmlElement *BPMN20.XsdExclusiveGateway, object *BPMN20.ExclusiveGateway) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ExclusiveGateway **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** GatewayDirection **/
	if object.M_gatewayDirection == BPMN20.GatewayDirection_Unspecified {
		xmlElement.M_gatewayDirection = "##Unspecified"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Converging {
		xmlElement.M_gatewayDirection = "##Converging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Diverging {
		xmlElement.M_gatewayDirection = "##Diverging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Mixed {
		xmlElement.M_gatewayDirection = "##Mixed"
	} else {
		xmlElement.M_gatewayDirection = "##Unspecified"
	}

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
}

/** serialysation of ManualTask **/
func (this *BPMSXmlFactory) SerialyzeManualTask(xmlElement *BPMN20.XsdManualTask, object *BPMN20.ManualTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ManualTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
}

/** serialysation of Choreography **/
func (this *BPMSXmlFactory) SerialyzeChoreography(xmlElement *BPMN20.XsdChoreography, object *BPMN20.Choreography_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Choreography **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Participant **/
	if len(object.M_participant) > 0 {
		xmlElement.M_participant = make([]*BPMN20.XsdParticipant, 0)
	}

	/** Now I will save the value of participant **/
	for i := 0; i < len(object.M_participant); i++ {
		xmlElement.M_participant = append(xmlElement.M_participant, new(BPMN20.XsdParticipant))
		this.SerialyzeParticipant(xmlElement.M_participant[i], object.M_participant[i])
	}

	/** Serialyze MessageFlow **/
	if len(object.M_messageFlow) > 0 {
		xmlElement.M_messageFlow = make([]*BPMN20.XsdMessageFlow, 0)
	}

	/** Now I will save the value of messageFlow **/
	for i := 0; i < len(object.M_messageFlow); i++ {
		xmlElement.M_messageFlow = append(xmlElement.M_messageFlow, new(BPMN20.XsdMessageFlow))
		this.SerialyzeMessageFlow(xmlElement.M_messageFlow[i], object.M_messageFlow[i])
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze Collaboration:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze Collaboration:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze Collaboration:artifact:TextAnnotation")
		}
	}

	/** Serialyze ConversationNode **/
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_0 = make([]*BPMN20.XsdCallConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_1 = make([]*BPMN20.XsdConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_2 = make([]*BPMN20.XsdSubConversation, 0)
	}

	/** Now I will save the value of conversationNode **/
	for i := 0; i < len(object.M_conversationNode); i++ {
		switch v := object.M_conversationNode[i].(type) {
		case *BPMN20.CallConversation:
			xmlElement.M_conversationNode_0 = append(xmlElement.M_conversationNode_0, new(BPMN20.XsdCallConversation))
			this.SerialyzeCallConversation(xmlElement.M_conversationNode_0[len(xmlElement.M_conversationNode_0)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:CallConversation")
		case *BPMN20.Conversation:
			xmlElement.M_conversationNode_1 = append(xmlElement.M_conversationNode_1, new(BPMN20.XsdConversation))
			this.SerialyzeConversation(xmlElement.M_conversationNode_1[len(xmlElement.M_conversationNode_1)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:Conversation")
		case *BPMN20.SubConversation:
			xmlElement.M_conversationNode_2 = append(xmlElement.M_conversationNode_2, new(BPMN20.XsdSubConversation))
			this.SerialyzeSubConversation(xmlElement.M_conversationNode_2[len(xmlElement.M_conversationNode_2)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:SubConversation")
		}
	}

	/** Serialyze ConversationAssociation **/
	if len(object.M_conversationAssociation) > 0 {
		xmlElement.M_conversationAssociation = make([]*BPMN20.XsdConversationAssociation, 0)
	}

	/** Now I will save the value of conversationAssociation **/
	for i := 0; i < len(object.M_conversationAssociation); i++ {
		xmlElement.M_conversationAssociation = append(xmlElement.M_conversationAssociation, new(BPMN20.XsdConversationAssociation))
		this.SerialyzeConversationAssociation(xmlElement.M_conversationAssociation[i], object.M_conversationAssociation[i])
	}

	/** Serialyze ParticipantAssociation **/
	if len(object.M_participantAssociation) > 0 {
		xmlElement.M_participantAssociation = make([]*BPMN20.XsdParticipantAssociation, 0)
	}

	/** Now I will save the value of participantAssociation **/
	for i := 0; i < len(object.M_participantAssociation); i++ {
		xmlElement.M_participantAssociation = append(xmlElement.M_participantAssociation, new(BPMN20.XsdParticipantAssociation))
		this.SerialyzeParticipantAssociation(xmlElement.M_participantAssociation[i], object.M_participantAssociation[i])
	}

	/** Serialyze MessageFlowAssociation **/
	if len(object.M_messageFlowAssociation) > 0 {
		xmlElement.M_messageFlowAssociation = make([]*BPMN20.XsdMessageFlowAssociation, 0)
	}

	/** Now I will save the value of messageFlowAssociation **/
	for i := 0; i < len(object.M_messageFlowAssociation); i++ {
		xmlElement.M_messageFlowAssociation = append(xmlElement.M_messageFlowAssociation, new(BPMN20.XsdMessageFlowAssociation))
		this.SerialyzeMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], object.M_messageFlowAssociation[i])
	}

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref choreographyRef **/
	xmlElement.M_choreographyRef = object.M_choreographyRef

	/** Serialyze ConversationLink **/
	if len(object.M_conversationLink) > 0 {
		xmlElement.M_conversationLink = make([]*BPMN20.XsdConversationLink, 0)
	}

	/** Now I will save the value of conversationLink **/
	for i := 0; i < len(object.M_conversationLink); i++ {
		xmlElement.M_conversationLink = append(xmlElement.M_conversationLink, new(BPMN20.XsdConversationLink))
		this.SerialyzeConversationLink(xmlElement.M_conversationLink[i], object.M_conversationLink[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_isClosed = object.M_isClosed

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze Choreography:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze Choreography:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze Choreography:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze Choreography:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze Choreography:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze Choreography:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze Choreography:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze Choreography:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze Choreography:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze Choreography:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze Choreography:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze Choreography:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze Choreography:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze Choreography:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze Choreography:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze Choreography:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze Choreography:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze Choreography:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze Choreography:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze Choreography:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze Choreography:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze Choreography:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze Choreography:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze Choreography:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze Choreography:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze Choreography:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze Choreography:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze Choreography:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze Choreography:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze Choreography:flowElement:UserTask")
		}
	}
}

/** serialysation of Extension **/
func (this *BPMSXmlFactory) SerialyzeExtension(xmlElement *BPMN20.XsdExtension, object *BPMN20.Extension) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ref extensionDefinition **/
	if object.M_extensionDefinition != nil {
	}

	/** Extension **/
	xmlElement.M_mustUnderstand = object.M_mustUnderstand
	this.m_references["Extension"] = object
}

/** serialysation of LinkEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeLinkEventDefinition(xmlElement *BPMN20.XsdLinkEventDefinition, object *BPMN20.LinkEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** LinkEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref source **/
	xmlElement.M_source = object.M_source

	/** Serialyze ref target **/
	xmlElement.M_target = &object.M_target

	/** EventDefinition **/
	xmlElement.M_name = object.M_name
}

/** serialysation of ReceiveTask **/
func (this *BPMSXmlFactory) SerialyzeReceiveTask(xmlElement *BPMN20.XsdReceiveTask, object *BPMN20.ReceiveTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ReceiveTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##WebService"
		}
	}

	/** Task **/
	xmlElement.M_instantiate = object.M_instantiate

	/** Serialyze ref messageRef **/
	xmlElement.M_messageRef = object.M_messageRef

	/** Serialyze ref operationRef **/
	xmlElement.M_operationRef = object.M_operationRef
}

/** serialysation of SubProcess **/
func (this *BPMSXmlFactory) SerialyzeSubProcess(xmlElement *BPMN20.XsdSubProcess, object *BPMN20.SubProcess_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** SubProcess **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default

	/** Serialyze LaneSet **/
	if len(object.M_laneSet) > 0 {
		xmlElement.M_laneSet = make([]*BPMN20.XsdLaneSet, 0)
	}

	/** Now I will save the value of laneSet **/
	for i := 0; i < len(object.M_laneSet); i++ {
		xmlElement.M_laneSet = append(xmlElement.M_laneSet, new(BPMN20.XsdLaneSet))
		this.SerialyzeLaneSet(xmlElement.M_laneSet[i], object.M_laneSet[i])
	}

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze SubProcess:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze SubProcess:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze SubProcess:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze SubProcess:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze SubProcess:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze SubProcess:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze SubProcess:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze SubProcess:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze SubProcess:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze SubProcess:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze SubProcess:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze SubProcess:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze SubProcess:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze SubProcess:flowElement:UserTask")
		}
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze SubProcess:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze SubProcess:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze SubProcess:artifact:TextAnnotation")
		}
	}

	/** Activity **/
	xmlElement.M_triggeredByEvent = object.M_triggeredByEvent
}

/** serialysation of DataObject **/
func (this *BPMSXmlFactory) SerialyzeDataObject(xmlElement *BPMN20.XsdDataObject, object *BPMN20.DataObject) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataObject **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef

	/** FlowElement **/
	xmlElement.M_isCollection = object.M_isCollection
}

/** serialysation of ScriptTask **/
func (this *BPMSXmlFactory) SerialyzeScriptTask(xmlElement *BPMN20.XsdScriptTask, object *BPMN20.ScriptTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ScriptTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default

	/** Serialyze Script **/
	if object.M_script != nil {
		xmlElement.M_script = new(BPMN20.XsdScript)
	}

	/** Now I will save the value of script **/
	if object.M_script != nil {
		this.SerialyzeScript(xmlElement.M_script, object.M_script)
	}

	/** Task **/
	xmlElement.M_scriptFormat = object.M_scriptFormat
}

/** serialysation of BPMNLabelStyle **/
func (this *BPMSXmlFactory) SerialyzeBPMNLabelStyle(xmlElement *BPMNDI.XsdBPMNLabelStyle, object *BPMNDI.BPMNLabelStyle) {
	if xmlElement == nil {
		return
	}

	/** BPMNLabelStyle **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Font **/

	/** Now I will save the value of Font **/
	if object.M_Font != nil {
		this.SerialyzeFont(&xmlElement.M_Font, object.M_Font)
	}
}

/** serialysation of DataOutput **/
func (this *BPMSXmlFactory) SerialyzeDataOutput(xmlElement *BPMN20.XsdDataOutput, object *BPMN20.DataOutput) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataOutput **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef

	/** BaseElement **/
	xmlElement.M_isCollection = object.M_isCollection
}

/** serialysation of ParallelGateway **/
func (this *BPMSXmlFactory) SerialyzeParallelGateway(xmlElement *BPMN20.XsdParallelGateway, object *BPMN20.ParallelGateway) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ParallelGateway **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** GatewayDirection **/
	if object.M_gatewayDirection == BPMN20.GatewayDirection_Unspecified {
		xmlElement.M_gatewayDirection = "##Unspecified"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Converging {
		xmlElement.M_gatewayDirection = "##Converging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Diverging {
		xmlElement.M_gatewayDirection = "##Diverging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Mixed {
		xmlElement.M_gatewayDirection = "##Mixed"
	} else {
		xmlElement.M_gatewayDirection = "##Unspecified"
	}
}

/** serialysation of SendTask **/
func (this *BPMSXmlFactory) SerialyzeSendTask(xmlElement *BPMN20.XsdSendTask, object *BPMN20.SendTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** SendTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##WebService"
		}
	}

	/** Serialyze ref messageRef **/
	xmlElement.M_messageRef = object.M_messageRef

	/** Serialyze ref operationRef **/
	xmlElement.M_operationRef = object.M_operationRef
}

/** serialysation of Point **/
func (this *BPMSXmlFactory) SerialyzePoint(xmlElement *DC.XsdPoint, object *DC.Point) {
	if xmlElement == nil {
		return
	}

	/** Point **/
	xmlElement.M_x = object.M_x

	/** Point **/
	xmlElement.M_y = object.M_y
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of CancelEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeCancelEventDefinition(xmlElement *BPMN20.XsdCancelEventDefinition, object *BPMN20.CancelEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CancelEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of TimerEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeTimerEventDefinition(xmlElement *BPMN20.XsdTimerEventDefinition, object *BPMN20.TimerEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** TimerEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of BPMNEdge **/
func (this *BPMSXmlFactory) SerialyzeBPMNEdge(xmlElement *BPMNDI.XsdBPMNEdge, object *BPMNDI.BPMNEdge) {
	if xmlElement == nil {
		return
	}

	/** BPMNEdge **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Point **/
	if len(object.M_waypoint) > 0 {
		xmlElement.M_waypoint = make([]*DC.XsdPoint, 0)
	}

	/** Now I will save the value of waypoint **/
	for i := 0; i < len(object.M_waypoint); i++ {
		xmlElement.M_waypoint = append(xmlElement.M_waypoint, new(DC.XsdPoint))
		this.SerialyzePoint(xmlElement.M_waypoint[i], object.M_waypoint[i])
	}

	/** Serialyze BPMNLabel **/
	if object.M_BPMNLabel != nil {
		xmlElement.M_BPMNLabel = new(BPMNDI.XsdBPMNLabel)
	}

	/** Now I will save the value of BPMNLabel **/
	if object.M_BPMNLabel != nil {
		this.SerialyzeBPMNLabel(xmlElement.M_BPMNLabel, object.M_BPMNLabel)
	}
	/** bpmnElement **/
	xmlElement.M_bpmnElement = object.M_bpmnElement

	/** Serialyze ref sourceElement **/
	if object.M_sourceElement != nil {
		id := object.M_sourceElement.(DI.DiagramElement).GetId()
		xmlElement.M_sourceElement = id
	}

	/** Serialyze ref targetElement **/
	xmlElement.M_targetElement = object.M_targetElement

	/** MessageVisibleKind **/
	if object.M_messageVisibleKind == BPMNDI.MessageVisibleKind_Initiating {
		xmlElement.M_messageVisibleKind = "##initiating"
	} else if object.M_messageVisibleKind == BPMNDI.MessageVisibleKind_Non_initiating {
		xmlElement.M_messageVisibleKind = "##non_initiating"
	}
}

/** serialysation of ComplexGateway **/
func (this *BPMSXmlFactory) SerialyzeComplexGateway(xmlElement *BPMN20.XsdComplexGateway, object *BPMN20.ComplexGateway) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ComplexGateway **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** GatewayDirection **/
	if object.M_gatewayDirection == BPMN20.GatewayDirection_Unspecified {
		xmlElement.M_gatewayDirection = "##Unspecified"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Converging {
		xmlElement.M_gatewayDirection = "##Converging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Diverging {
		xmlElement.M_gatewayDirection = "##Diverging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Mixed {
		xmlElement.M_gatewayDirection = "##Mixed"
	} else {
		xmlElement.M_gatewayDirection = "##Unspecified"
	}

	/** Serialyze FormalExpression **/
	if object.M_activationCondition != nil {
		xmlElement.M_activationCondition = new(BPMN20.XsdActivationCondition)
	}

	/** Now I will save the value of activationCondition **/
	if object.M_activationCondition != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_activationCondition)), object.M_activationCondition)
	}

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
}

/** serialysation of ParticipantMultiplicity **/
func (this *BPMSXmlFactory) SerialyzeParticipantMultiplicity(xmlElement *BPMN20.XsdParticipantMultiplicity, object *BPMN20.ParticipantMultiplicity) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ParticipantMultiplicity **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_minimum = object.M_minimum

	/** BaseElement **/
	xmlElement.M_maximum = object.M_maximum
}

/** serialysation of GlobalBusinessRuleTask **/
func (this *BPMSXmlFactory) SerialyzeGlobalBusinessRuleTask(xmlElement *BPMN20.XsdGlobalBusinessRuleTask, object *BPMN20.GlobalBusinessRuleTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalBusinessRuleTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref supportedInterfaceRef **/
	xmlElement.M_supportedInterfaceRef = object.M_supportedInterfaceRef

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze InputOutputBinding **/
	if len(object.M_ioBinding) > 0 {
		xmlElement.M_ioBinding = make([]*BPMN20.XsdInputOutputBinding, 0)
	}

	/** Now I will save the value of ioBinding **/
	for i := 0; i < len(object.M_ioBinding); i++ {
		xmlElement.M_ioBinding = append(xmlElement.M_ioBinding, new(BPMN20.XsdInputOutputBinding))
		this.SerialyzeInputOutputBinding(xmlElement.M_ioBinding[i], object.M_ioBinding[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:ResourceRole")
		}
	}
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##unspecified"
		}
	}
}

/** serialysation of ItemDefinition **/
func (this *BPMSXmlFactory) SerialyzeItemDefinition(xmlElement *BPMN20.XsdItemDefinition, object *BPMN20.ItemDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ItemDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref structureRef **/
	xmlElement.M_structureRef = object.M_structureRef

	/** RootElement **/
	xmlElement.M_isCollection = object.M_isCollection

	/** ItemKind **/
	if object.M_itemKind == BPMN20.ItemKind_Information {
		xmlElement.M_itemKind = "##Information"
	} else if object.M_itemKind == BPMN20.ItemKind_Physical {
		xmlElement.M_itemKind = "##Physical"
	} else {
		xmlElement.M_itemKind = "##Information"
	}
}

/** serialysation of Monitoring **/
func (this *BPMSXmlFactory) SerialyzeMonitoring(xmlElement *BPMN20.XsdMonitoring, object *BPMN20.Monitoring) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Monitoring **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Collaboration **/
func (this *BPMSXmlFactory) SerialyzeCollaboration(xmlElement *BPMN20.XsdCollaboration, object *BPMN20.Collaboration_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Collaboration **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Participant **/
	if len(object.M_participant) > 0 {
		xmlElement.M_participant = make([]*BPMN20.XsdParticipant, 0)
	}

	/** Now I will save the value of participant **/
	for i := 0; i < len(object.M_participant); i++ {
		xmlElement.M_participant = append(xmlElement.M_participant, new(BPMN20.XsdParticipant))
		this.SerialyzeParticipant(xmlElement.M_participant[i], object.M_participant[i])
	}

	/** Serialyze MessageFlow **/
	if len(object.M_messageFlow) > 0 {
		xmlElement.M_messageFlow = make([]*BPMN20.XsdMessageFlow, 0)
	}

	/** Now I will save the value of messageFlow **/
	for i := 0; i < len(object.M_messageFlow); i++ {
		xmlElement.M_messageFlow = append(xmlElement.M_messageFlow, new(BPMN20.XsdMessageFlow))
		this.SerialyzeMessageFlow(xmlElement.M_messageFlow[i], object.M_messageFlow[i])
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze Collaboration:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze Collaboration:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze Collaboration:artifact:TextAnnotation")
		}
	}

	/** Serialyze ConversationNode **/
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_0 = make([]*BPMN20.XsdCallConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_1 = make([]*BPMN20.XsdConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_2 = make([]*BPMN20.XsdSubConversation, 0)
	}

	/** Now I will save the value of conversationNode **/
	for i := 0; i < len(object.M_conversationNode); i++ {
		switch v := object.M_conversationNode[i].(type) {
		case *BPMN20.CallConversation:
			xmlElement.M_conversationNode_0 = append(xmlElement.M_conversationNode_0, new(BPMN20.XsdCallConversation))
			this.SerialyzeCallConversation(xmlElement.M_conversationNode_0[len(xmlElement.M_conversationNode_0)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:CallConversation")
		case *BPMN20.Conversation:
			xmlElement.M_conversationNode_1 = append(xmlElement.M_conversationNode_1, new(BPMN20.XsdConversation))
			this.SerialyzeConversation(xmlElement.M_conversationNode_1[len(xmlElement.M_conversationNode_1)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:Conversation")
		case *BPMN20.SubConversation:
			xmlElement.M_conversationNode_2 = append(xmlElement.M_conversationNode_2, new(BPMN20.XsdSubConversation))
			this.SerialyzeSubConversation(xmlElement.M_conversationNode_2[len(xmlElement.M_conversationNode_2)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:SubConversation")
		}
	}

	/** Serialyze ConversationAssociation **/
	if len(object.M_conversationAssociation) > 0 {
		xmlElement.M_conversationAssociation = make([]*BPMN20.XsdConversationAssociation, 0)
	}

	/** Now I will save the value of conversationAssociation **/
	for i := 0; i < len(object.M_conversationAssociation); i++ {
		xmlElement.M_conversationAssociation = append(xmlElement.M_conversationAssociation, new(BPMN20.XsdConversationAssociation))
		this.SerialyzeConversationAssociation(xmlElement.M_conversationAssociation[i], object.M_conversationAssociation[i])
	}

	/** Serialyze ParticipantAssociation **/
	if len(object.M_participantAssociation) > 0 {
		xmlElement.M_participantAssociation = make([]*BPMN20.XsdParticipantAssociation, 0)
	}

	/** Now I will save the value of participantAssociation **/
	for i := 0; i < len(object.M_participantAssociation); i++ {
		xmlElement.M_participantAssociation = append(xmlElement.M_participantAssociation, new(BPMN20.XsdParticipantAssociation))
		this.SerialyzeParticipantAssociation(xmlElement.M_participantAssociation[i], object.M_participantAssociation[i])
	}

	/** Serialyze MessageFlowAssociation **/
	if len(object.M_messageFlowAssociation) > 0 {
		xmlElement.M_messageFlowAssociation = make([]*BPMN20.XsdMessageFlowAssociation, 0)
	}

	/** Now I will save the value of messageFlowAssociation **/
	for i := 0; i < len(object.M_messageFlowAssociation); i++ {
		xmlElement.M_messageFlowAssociation = append(xmlElement.M_messageFlowAssociation, new(BPMN20.XsdMessageFlowAssociation))
		this.SerialyzeMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], object.M_messageFlowAssociation[i])
	}

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref choreographyRef **/
	xmlElement.M_choreographyRef = object.M_choreographyRef

	/** Serialyze ConversationLink **/
	if len(object.M_conversationLink) > 0 {
		xmlElement.M_conversationLink = make([]*BPMN20.XsdConversationLink, 0)
	}

	/** Now I will save the value of conversationLink **/
	for i := 0; i < len(object.M_conversationLink); i++ {
		xmlElement.M_conversationLink = append(xmlElement.M_conversationLink, new(BPMN20.XsdConversationLink))
		this.SerialyzeConversationLink(xmlElement.M_conversationLink[i], object.M_conversationLink[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_isClosed = object.M_isClosed
}

/** serialysation of Interface **/
func (this *BPMSXmlFactory) SerialyzeInterface(xmlElement *BPMN20.XsdInterface, object *BPMN20.Interface) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Interface **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Operation **/
	if len(object.M_operation) > 0 {
		xmlElement.M_operation = make([]*BPMN20.XsdOperation, 0)
	}

	/** Now I will save the value of operation **/
	for i := 0; i < len(object.M_operation); i++ {
		xmlElement.M_operation = append(xmlElement.M_operation, new(BPMN20.XsdOperation))
		this.SerialyzeOperation(xmlElement.M_operation[i], object.M_operation[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref implementationRef **/
	xmlElement.M_implementationRef = object.M_implementationRef
}

/** serialysation of Resource **/
func (this *BPMSXmlFactory) SerialyzeResource(xmlElement *BPMN20.XsdResource, object *BPMN20.Resource) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Resource **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ResourceParameter **/
	if len(object.M_resourceParameter) > 0 {
		xmlElement.M_resourceParameter = make([]*BPMN20.XsdResourceParameter, 0)
	}

	/** Now I will save the value of resourceParameter **/
	for i := 0; i < len(object.M_resourceParameter); i++ {
		xmlElement.M_resourceParameter = append(xmlElement.M_resourceParameter, new(BPMN20.XsdResourceParameter))
		this.SerialyzeResourceParameter(xmlElement.M_resourceParameter[i], object.M_resourceParameter[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of Text **/
func (this *BPMSXmlFactory) SerialyzeText(xmlElement *BPMN20.XsdText, object *BPMN20.Text) {
	if xmlElement == nil {
		return
	}
	xmlElement.M_text = object.M_text
	this.m_references["Text"] = object
}

/** serialysation of GlobalManualTask **/
func (this *BPMSXmlFactory) SerialyzeGlobalManualTask(xmlElement *BPMN20.XsdGlobalManualTask, object *BPMN20.GlobalManualTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalManualTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref supportedInterfaceRef **/
	xmlElement.M_supportedInterfaceRef = object.M_supportedInterfaceRef

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze InputOutputBinding **/
	if len(object.M_ioBinding) > 0 {
		xmlElement.M_ioBinding = make([]*BPMN20.XsdInputOutputBinding, 0)
	}

	/** Now I will save the value of ioBinding **/
	for i := 0; i < len(object.M_ioBinding); i++ {
		xmlElement.M_ioBinding = append(xmlElement.M_ioBinding, new(BPMN20.XsdInputOutputBinding))
		this.SerialyzeInputOutputBinding(xmlElement.M_ioBinding[i], object.M_ioBinding[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:ResourceRole")
		}
	}
}

/** serialysation of BoundaryEvent **/
func (this *BPMSXmlFactory) SerialyzeBoundaryEvent(xmlElement *BPMN20.XsdBoundaryEvent, object *BPMN20.BoundaryEvent) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** BoundaryEvent **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataOutput **/
	if len(object.M_dataOutput) > 0 {
		xmlElement.M_dataOutput = make([]*BPMN20.XsdDataOutput, 0)
	}

	/** Now I will save the value of dataOutput **/
	for i := 0; i < len(object.M_dataOutput); i++ {
		xmlElement.M_dataOutput = append(xmlElement.M_dataOutput, new(BPMN20.XsdDataOutput))
		this.SerialyzeDataOutput(xmlElement.M_dataOutput[i], object.M_dataOutput[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze OutputSet **/
	if object.M_outputSet != nil {
		xmlElement.M_outputSet = new(BPMN20.XsdOutputSet)
	}

	/** Now I will save the value of outputSet **/
	if object.M_outputSet != nil {
		this.SerialyzeOutputSet(xmlElement.M_outputSet, object.M_outputSet)
	}

	/** Serialyze EventDefinition **/
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_0 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_1 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_2 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_3 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_4 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_5 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_6 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_7 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_8 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_9 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}

	/** Now I will save the value of eventDefinition **/
	for i := 0; i < len(object.M_eventDefinition); i++ {
		switch v := object.M_eventDefinition[i].(type) {
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_eventDefinition_0 = append(xmlElement.M_eventDefinition_0, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_eventDefinition_0[len(xmlElement.M_eventDefinition_0)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_eventDefinition_1 = append(xmlElement.M_eventDefinition_1, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_eventDefinition_1[len(xmlElement.M_eventDefinition_1)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_eventDefinition_2 = append(xmlElement.M_eventDefinition_2, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_eventDefinition_2[len(xmlElement.M_eventDefinition_2)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_eventDefinition_3 = append(xmlElement.M_eventDefinition_3, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_eventDefinition_3[len(xmlElement.M_eventDefinition_3)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_eventDefinition_4 = append(xmlElement.M_eventDefinition_4, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_eventDefinition_4[len(xmlElement.M_eventDefinition_4)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_eventDefinition_5 = append(xmlElement.M_eventDefinition_5, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_eventDefinition_5[len(xmlElement.M_eventDefinition_5)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_eventDefinition_6 = append(xmlElement.M_eventDefinition_6, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_eventDefinition_6[len(xmlElement.M_eventDefinition_6)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_eventDefinition_7 = append(xmlElement.M_eventDefinition_7, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_eventDefinition_7[len(xmlElement.M_eventDefinition_7)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_eventDefinition_8 = append(xmlElement.M_eventDefinition_8, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_eventDefinition_8[len(xmlElement.M_eventDefinition_8)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_eventDefinition_9 = append(xmlElement.M_eventDefinition_9, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_eventDefinition_9[len(xmlElement.M_eventDefinition_9)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:TimerEventDefinition")
		}
	}

	/** Serialyze ref eventDefinitionRef **/
	xmlElement.M_eventDefinitionRef = object.M_eventDefinitionRef

	/** Event **/
	xmlElement.M_parallelMultiple = object.M_parallelMultiple

	/** CatchEvent **/
	xmlElement.M_cancelActivity = object.M_cancelActivity

	/** Serialyze ref attachedToRef **/
	xmlElement.M_attachedToRef = object.M_attachedToRef
}

/** serialysation of ChoreographyTask **/
func (this *BPMSXmlFactory) SerialyzeChoreographyTask(xmlElement *BPMN20.XsdChoreographyTask, object *BPMN20.ChoreographyTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ChoreographyTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref initiatingParticipantRef **/
	xmlElement.M_initiatingParticipantRef = object.M_initiatingParticipantRef

	/** ChoreographyLoopType **/
	if object.M_loopType == BPMN20.ChoreographyLoopType_None {
		xmlElement.M_loopType = "##None"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_Standard {
		xmlElement.M_loopType = "##Standard"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_MultiInstanceSequential {
		xmlElement.M_loopType = "##MultiInstanceSequential"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_MultiInstanceParallel {
		xmlElement.M_loopType = "##MultiInstanceParallel"
	} else {
		xmlElement.M_loopType = "##None"
	}

	/** Serialyze ref messageFlowRef **/
	xmlElement.M_messageFlowRef = object.M_messageFlowRef

}

/** serialysation of GlobalScriptTask **/
func (this *BPMSXmlFactory) SerialyzeGlobalScriptTask(xmlElement *BPMN20.XsdGlobalScriptTask, object *BPMN20.GlobalScriptTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalScriptTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref supportedInterfaceRef **/
	xmlElement.M_supportedInterfaceRef = object.M_supportedInterfaceRef

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze InputOutputBinding **/
	if len(object.M_ioBinding) > 0 {
		xmlElement.M_ioBinding = make([]*BPMN20.XsdInputOutputBinding, 0)
	}

	/** Now I will save the value of ioBinding **/
	for i := 0; i < len(object.M_ioBinding); i++ {
		xmlElement.M_ioBinding = append(xmlElement.M_ioBinding, new(BPMN20.XsdInputOutputBinding))
		this.SerialyzeInputOutputBinding(xmlElement.M_ioBinding[i], object.M_ioBinding[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:ResourceRole")
		}
	}

	/** Serialyze Script **/
	if object.M_script != nil {
		xmlElement.M_script = new(BPMN20.XsdScript)
	}

	/** Now I will save the value of script **/
	if object.M_script != nil {
		this.SerialyzeScript(xmlElement.M_script, object.M_script)
	}

	/** GlobalTask **/
	xmlElement.M_scriptLanguage = object.M_scriptLanguage
}

/** serialysation of ImplicitThrowEvent **/
func (this *BPMSXmlFactory) SerialyzeImplicitThrowEvent(xmlElement *BPMN20.XsdImplicitThrowEvent, object *BPMN20.ImplicitThrowEvent) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ImplicitThrowEvent **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInput **/
	if len(object.M_dataInput) > 0 {
		xmlElement.M_dataInput = make([]*BPMN20.XsdDataInput, 0)
	}

	/** Now I will save the value of dataInput **/
	for i := 0; i < len(object.M_dataInput); i++ {
		xmlElement.M_dataInput = append(xmlElement.M_dataInput, new(BPMN20.XsdDataInput))
		this.SerialyzeDataInput(xmlElement.M_dataInput[i], object.M_dataInput[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze InputSet **/
	if object.M_inputSet != nil {
		xmlElement.M_inputSet = new(BPMN20.XsdInputSet)
	}

	/** Now I will save the value of inputSet **/
	if object.M_inputSet != nil {
		this.SerialyzeInputSet(xmlElement.M_inputSet, object.M_inputSet)
	}

	/** Serialyze EventDefinition **/
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_0 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_1 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_2 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_3 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_4 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_5 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_6 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_7 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_8 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_9 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}

	/** Now I will save the value of eventDefinition **/
	for i := 0; i < len(object.M_eventDefinition); i++ {
		switch v := object.M_eventDefinition[i].(type) {
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_eventDefinition_0 = append(xmlElement.M_eventDefinition_0, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_eventDefinition_0[len(xmlElement.M_eventDefinition_0)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_eventDefinition_1 = append(xmlElement.M_eventDefinition_1, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_eventDefinition_1[len(xmlElement.M_eventDefinition_1)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_eventDefinition_2 = append(xmlElement.M_eventDefinition_2, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_eventDefinition_2[len(xmlElement.M_eventDefinition_2)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_eventDefinition_3 = append(xmlElement.M_eventDefinition_3, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_eventDefinition_3[len(xmlElement.M_eventDefinition_3)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_eventDefinition_4 = append(xmlElement.M_eventDefinition_4, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_eventDefinition_4[len(xmlElement.M_eventDefinition_4)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_eventDefinition_5 = append(xmlElement.M_eventDefinition_5, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_eventDefinition_5[len(xmlElement.M_eventDefinition_5)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_eventDefinition_6 = append(xmlElement.M_eventDefinition_6, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_eventDefinition_6[len(xmlElement.M_eventDefinition_6)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_eventDefinition_7 = append(xmlElement.M_eventDefinition_7, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_eventDefinition_7[len(xmlElement.M_eventDefinition_7)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_eventDefinition_8 = append(xmlElement.M_eventDefinition_8, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_eventDefinition_8[len(xmlElement.M_eventDefinition_8)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_eventDefinition_9 = append(xmlElement.M_eventDefinition_9, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_eventDefinition_9[len(xmlElement.M_eventDefinition_9)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:TimerEventDefinition")
		}
	}

	/** Serialyze ref eventDefinitionRef **/
	xmlElement.M_eventDefinitionRef = object.M_eventDefinitionRef

}

/** serialysation of SubChoreography **/
func (this *BPMSXmlFactory) SerialyzeSubChoreography(xmlElement *BPMN20.XsdSubChoreography, object *BPMN20.SubChoreography) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** SubChoreography **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref initiatingParticipantRef **/
	xmlElement.M_initiatingParticipantRef = object.M_initiatingParticipantRef

	/** ChoreographyLoopType **/
	if object.M_loopType == BPMN20.ChoreographyLoopType_None {
		xmlElement.M_loopType = "##None"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_Standard {
		xmlElement.M_loopType = "##Standard"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_MultiInstanceSequential {
		xmlElement.M_loopType = "##MultiInstanceSequential"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_MultiInstanceParallel {
		xmlElement.M_loopType = "##MultiInstanceParallel"
	} else {
		xmlElement.M_loopType = "##None"
	}

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze SubChoreography:flowElement:UserTask")
		}
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze SubChoreography:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze SubChoreography:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze SubChoreography:artifact:TextAnnotation")
		}
	}
}

/** serialysation of ParticipantAssociation **/
func (this *BPMSXmlFactory) SerialyzeParticipantAssociation(xmlElement *BPMN20.XsdParticipantAssociation, object *BPMN20.ParticipantAssociation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ParticipantAssociation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref innerParticipantRef **/
	xmlElement.M_innerParticipantRef = object.M_innerParticipantRef

	/** Serialyze ref outerParticipantRef **/
	xmlElement.M_outerParticipantRef = object.M_outerParticipantRef

}

/** serialysation of CorrelationPropertyRetrievalExpression **/
func (this *BPMSXmlFactory) SerialyzeCorrelationPropertyRetrievalExpression(xmlElement *BPMN20.XsdCorrelationPropertyRetrievalExpression, object *BPMN20.CorrelationPropertyRetrievalExpression) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CorrelationPropertyRetrievalExpression **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/

	/** Now I will save the value of messagePath **/
	if object.M_messagePath != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(&xmlElement.M_messagePath)), object.M_messagePath)
	}

	/** Serialyze ref messageRef **/
	xmlElement.M_messageRef = object.M_messageRef
}

/** serialysation of SignalEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeSignalEventDefinition(xmlElement *BPMN20.XsdSignalEventDefinition, object *BPMN20.SignalEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** SignalEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref signalRef **/
	xmlElement.M_signalRef = object.M_signalRef
}

/** serialysation of DataInput **/
func (this *BPMSXmlFactory) SerialyzeDataInput(xmlElement *BPMN20.XsdDataInput, object *BPMN20.DataInput) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataInput **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef

	/** BaseElement **/
	xmlElement.M_isCollection = object.M_isCollection
}

/** serialysation of Documentation **/
func (this *BPMSXmlFactory) SerialyzeDocumentation(xmlElement *BPMN20.XsdDocumentation, object *BPMN20.Documentation) {
	if xmlElement == nil {
		return
	}
	xmlElement.M_text = object.M_text

	/** Documentation **/
	xmlElement.M_textFormat = object.M_textFormat
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ConditionalEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeConditionalEventDefinition(xmlElement *BPMN20.XsdConditionalEventDefinition, object *BPMN20.ConditionalEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ConditionalEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/
	if object.M_condition != nil {
		xmlElement.M_condition = new(BPMN20.XsdCondition)
	}

	/** Now I will save the value of condition **/
	if object.M_condition != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_condition)), object.M_condition)
	}
}

/** serialysation of Task **/
func (this *BPMSXmlFactory) SerialyzeTask(xmlElement *BPMN20.XsdTask, object *BPMN20.Task_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Task **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
}

/** serialysation of EscalationEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeEscalationEventDefinition(xmlElement *BPMN20.XsdEscalationEventDefinition, object *BPMN20.EscalationEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** EscalationEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref escalationRef **/
	xmlElement.M_escalationRef = object.M_escalationRef
}

/** serialysation of Performer **/
func (this *BPMSXmlFactory) SerialyzePerformer(xmlElement *BPMN20.XsdPerformer, object *BPMN20.Performer_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Performer **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref resourceRef **/
	xmlElement.M_resourceRef = &object.M_resourceRef

	/** Serialyze ResourceParameterBinding **/
	if len(object.M_resourceParameterBinding) > 0 {
		xmlElement.M_resourceParameterBinding = make([]*BPMN20.XsdResourceParameterBinding, 0)
	}

	/** Now I will save the value of resourceParameterBinding **/
	for i := 0; i < len(object.M_resourceParameterBinding); i++ {
		xmlElement.M_resourceParameterBinding = append(xmlElement.M_resourceParameterBinding, new(BPMN20.XsdResourceParameterBinding))
		this.SerialyzeResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], object.M_resourceParameterBinding[i])
	}

	/** Serialyze ResourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		xmlElement.M_resourceAssignmentExpression = new(BPMN20.XsdResourceAssignmentExpression)
	}

	/** Now I will save the value of resourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		this.SerialyzeResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)
	}
}

/** serialysation of MultiInstanceLoopCharacteristics **/
func (this *BPMSXmlFactory) SerialyzeMultiInstanceLoopCharacteristics(xmlElement *BPMN20.XsdMultiInstanceLoopCharacteristics, object *BPMN20.MultiInstanceLoopCharacteristics) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** MultiInstanceLoopCharacteristics **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/
	if object.M_loopCardinality != nil {
		xmlElement.M_loopCardinality = new(BPMN20.XsdLoopCardinality)
	}

	/** Now I will save the value of loopCardinality **/
	if object.M_loopCardinality != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_loopCardinality)), object.M_loopCardinality)
	}

	/** Serialyze ref loopDataInputRef **/
	xmlElement.M_loopDataInputRef = &object.M_loopDataInputRef

	/** Serialyze ref loopDataOutputRef **/
	xmlElement.M_loopDataOutputRef = &object.M_loopDataOutputRef

	/** Serialyze DataInput **/
	if object.M_inputDataItem != nil {
		xmlElement.M_inputDataItem = new(BPMN20.XsdInputDataItem)
	}

	/** Now I will save the value of inputDataItem **/
	if object.M_inputDataItem != nil {
		this.SerialyzeDataInput((*BPMN20.XsdDataInput)(unsafe.Pointer(xmlElement.M_inputDataItem)), object.M_inputDataItem)
	}

	/** Serialyze DataOutput **/
	if object.M_outputDataItem != nil {
		xmlElement.M_outputDataItem = new(BPMN20.XsdOutputDataItem)
	}

	/** Now I will save the value of outputDataItem **/
	if object.M_outputDataItem != nil {
		this.SerialyzeDataOutput((*BPMN20.XsdDataOutput)(unsafe.Pointer(xmlElement.M_outputDataItem)), object.M_outputDataItem)
	}

	/** Serialyze ComplexBehaviorDefinition **/
	if len(object.M_complexBehaviorDefinition) > 0 {
		xmlElement.M_complexBehaviorDefinition = make([]*BPMN20.XsdComplexBehaviorDefinition, 0)
	}

	/** Now I will save the value of complexBehaviorDefinition **/
	for i := 0; i < len(object.M_complexBehaviorDefinition); i++ {
		xmlElement.M_complexBehaviorDefinition = append(xmlElement.M_complexBehaviorDefinition, new(BPMN20.XsdComplexBehaviorDefinition))
		this.SerialyzeComplexBehaviorDefinition(xmlElement.M_complexBehaviorDefinition[i], object.M_complexBehaviorDefinition[i])
	}

	/** Serialyze FormalExpression **/
	if object.M_completionCondition != nil {
		xmlElement.M_completionCondition = new(BPMN20.XsdCompletionCondition)
	}

	/** Now I will save the value of completionCondition **/
	if object.M_completionCondition != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_completionCondition)), object.M_completionCondition)
	}

	/** LoopCharacteristics **/
	xmlElement.M_isSequential = object.M_isSequential

	/** MultiInstanceFlowCondition **/
	if object.M_behavior == BPMN20.MultiInstanceFlowCondition_None {
		xmlElement.M_behavior = "##None"
	} else if object.M_behavior == BPMN20.MultiInstanceFlowCondition_One {
		xmlElement.M_behavior = "##One"
	} else if object.M_behavior == BPMN20.MultiInstanceFlowCondition_All {
		xmlElement.M_behavior = "##All"
	} else if object.M_behavior == BPMN20.MultiInstanceFlowCondition_Complex {
		xmlElement.M_behavior = "##Complex"
	} else {
		xmlElement.M_behavior = "##All"
	}

	/** Serialyze ref oneBehaviorEventRef **/
	xmlElement.M_oneBehaviorEventRef = object.M_oneBehaviorEventRef

	/** Serialyze ref noneBehaviorEventRef **/
	xmlElement.M_noneBehaviorEventRef = object.M_noneBehaviorEventRef
}

/** serialysation of IntermediateThrowEvent **/
func (this *BPMSXmlFactory) SerialyzeIntermediateThrowEvent(xmlElement *BPMN20.XsdIntermediateThrowEvent, object *BPMN20.IntermediateThrowEvent) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** IntermediateThrowEvent **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInput **/
	if len(object.M_dataInput) > 0 {
		xmlElement.M_dataInput = make([]*BPMN20.XsdDataInput, 0)
	}

	/** Now I will save the value of dataInput **/
	for i := 0; i < len(object.M_dataInput); i++ {
		xmlElement.M_dataInput = append(xmlElement.M_dataInput, new(BPMN20.XsdDataInput))
		this.SerialyzeDataInput(xmlElement.M_dataInput[i], object.M_dataInput[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze InputSet **/
	if object.M_inputSet != nil {
		xmlElement.M_inputSet = new(BPMN20.XsdInputSet)
	}

	/** Now I will save the value of inputSet **/
	if object.M_inputSet != nil {
		this.SerialyzeInputSet(xmlElement.M_inputSet, object.M_inputSet)
	}

	/** Serialyze EventDefinition **/
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_0 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_1 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_2 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_3 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_4 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_5 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_6 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_7 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_8 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_9 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}

	/** Now I will save the value of eventDefinition **/
	for i := 0; i < len(object.M_eventDefinition); i++ {
		switch v := object.M_eventDefinition[i].(type) {
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_eventDefinition_0 = append(xmlElement.M_eventDefinition_0, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_eventDefinition_0[len(xmlElement.M_eventDefinition_0)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_eventDefinition_1 = append(xmlElement.M_eventDefinition_1, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_eventDefinition_1[len(xmlElement.M_eventDefinition_1)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_eventDefinition_2 = append(xmlElement.M_eventDefinition_2, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_eventDefinition_2[len(xmlElement.M_eventDefinition_2)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_eventDefinition_3 = append(xmlElement.M_eventDefinition_3, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_eventDefinition_3[len(xmlElement.M_eventDefinition_3)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_eventDefinition_4 = append(xmlElement.M_eventDefinition_4, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_eventDefinition_4[len(xmlElement.M_eventDefinition_4)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_eventDefinition_5 = append(xmlElement.M_eventDefinition_5, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_eventDefinition_5[len(xmlElement.M_eventDefinition_5)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_eventDefinition_6 = append(xmlElement.M_eventDefinition_6, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_eventDefinition_6[len(xmlElement.M_eventDefinition_6)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_eventDefinition_7 = append(xmlElement.M_eventDefinition_7, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_eventDefinition_7[len(xmlElement.M_eventDefinition_7)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_eventDefinition_8 = append(xmlElement.M_eventDefinition_8, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_eventDefinition_8[len(xmlElement.M_eventDefinition_8)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_eventDefinition_9 = append(xmlElement.M_eventDefinition_9, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_eventDefinition_9[len(xmlElement.M_eventDefinition_9)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:TimerEventDefinition")
		}
	}

	/** Serialyze ref eventDefinitionRef **/
	xmlElement.M_eventDefinitionRef = object.M_eventDefinitionRef

}

/** serialysation of TerminateEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeTerminateEventDefinition(xmlElement *BPMN20.XsdTerminateEventDefinition, object *BPMN20.TerminateEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** TerminateEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Assignment **/
func (this *BPMSXmlFactory) SerialyzeAssignment(xmlElement *BPMN20.XsdAssignment, object *BPMN20.Assignment) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Assignment **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/
	if object.M_from != nil {
		xmlElement.M_from = new(BPMN20.XsdFrom)
	}

	/** Now I will save the value of from **/
	if object.M_from != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_from)), object.M_from)
	}

	/** Serialyze FormalExpression **/
	if object.M_to != nil {
		xmlElement.M_to = new(BPMN20.XsdTo)
	}

	/** Now I will save the value of to **/
	if object.M_to != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_to)), object.M_to)
	}
}

/** serialysation of CallChoreography **/
func (this *BPMSXmlFactory) SerialyzeCallChoreography(xmlElement *BPMN20.XsdCallChoreography, object *BPMN20.CallChoreography) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CallChoreography **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref initiatingParticipantRef **/
	xmlElement.M_initiatingParticipantRef = object.M_initiatingParticipantRef

	/** ChoreographyLoopType **/
	if object.M_loopType == BPMN20.ChoreographyLoopType_None {
		xmlElement.M_loopType = "##None"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_Standard {
		xmlElement.M_loopType = "##Standard"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_MultiInstanceSequential {
		xmlElement.M_loopType = "##MultiInstanceSequential"
	} else if object.M_loopType == BPMN20.ChoreographyLoopType_MultiInstanceParallel {
		xmlElement.M_loopType = "##MultiInstanceParallel"
	} else {
		xmlElement.M_loopType = "##None"
	}

	/** Serialyze ParticipantAssociation **/
	if len(object.M_participantAssociation) > 0 {
		xmlElement.M_participantAssociation = make([]*BPMN20.XsdParticipantAssociation, 0)
	}

	/** Now I will save the value of participantAssociation **/
	for i := 0; i < len(object.M_participantAssociation); i++ {
		xmlElement.M_participantAssociation = append(xmlElement.M_participantAssociation, new(BPMN20.XsdParticipantAssociation))
		this.SerialyzeParticipantAssociation(xmlElement.M_participantAssociation[i], object.M_participantAssociation[i])
	}

	/** Serialyze ref calledChoreographyRef **/
	xmlElement.M_calledChoreographyRef = object.M_calledChoreographyRef
}

/** serialysation of EndEvent **/
func (this *BPMSXmlFactory) SerialyzeEndEvent(xmlElement *BPMN20.XsdEndEvent, object *BPMN20.EndEvent) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** EndEvent **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInput **/
	if len(object.M_dataInput) > 0 {
		xmlElement.M_dataInput = make([]*BPMN20.XsdDataInput, 0)
	}

	/** Now I will save the value of dataInput **/
	for i := 0; i < len(object.M_dataInput); i++ {
		xmlElement.M_dataInput = append(xmlElement.M_dataInput, new(BPMN20.XsdDataInput))
		this.SerialyzeDataInput(xmlElement.M_dataInput[i], object.M_dataInput[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze InputSet **/
	if object.M_inputSet != nil {
		xmlElement.M_inputSet = new(BPMN20.XsdInputSet)
	}

	/** Now I will save the value of inputSet **/
	if object.M_inputSet != nil {
		this.SerialyzeInputSet(xmlElement.M_inputSet, object.M_inputSet)
	}

	/** Serialyze EventDefinition **/
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_0 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_1 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_2 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_3 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_4 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_5 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_6 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_7 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_8 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_9 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}

	/** Now I will save the value of eventDefinition **/
	for i := 0; i < len(object.M_eventDefinition); i++ {
		switch v := object.M_eventDefinition[i].(type) {
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_eventDefinition_0 = append(xmlElement.M_eventDefinition_0, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_eventDefinition_0[len(xmlElement.M_eventDefinition_0)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_eventDefinition_1 = append(xmlElement.M_eventDefinition_1, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_eventDefinition_1[len(xmlElement.M_eventDefinition_1)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_eventDefinition_2 = append(xmlElement.M_eventDefinition_2, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_eventDefinition_2[len(xmlElement.M_eventDefinition_2)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_eventDefinition_3 = append(xmlElement.M_eventDefinition_3, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_eventDefinition_3[len(xmlElement.M_eventDefinition_3)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_eventDefinition_4 = append(xmlElement.M_eventDefinition_4, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_eventDefinition_4[len(xmlElement.M_eventDefinition_4)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_eventDefinition_5 = append(xmlElement.M_eventDefinition_5, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_eventDefinition_5[len(xmlElement.M_eventDefinition_5)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_eventDefinition_6 = append(xmlElement.M_eventDefinition_6, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_eventDefinition_6[len(xmlElement.M_eventDefinition_6)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_eventDefinition_7 = append(xmlElement.M_eventDefinition_7, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_eventDefinition_7[len(xmlElement.M_eventDefinition_7)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_eventDefinition_8 = append(xmlElement.M_eventDefinition_8, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_eventDefinition_8[len(xmlElement.M_eventDefinition_8)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_eventDefinition_9 = append(xmlElement.M_eventDefinition_9, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_eventDefinition_9[len(xmlElement.M_eventDefinition_9)-1], v)
			log.Println("Serialyze ThrowEvent:eventDefinition:TimerEventDefinition")
		}
	}

	/** Serialyze ref eventDefinitionRef **/
	xmlElement.M_eventDefinitionRef = object.M_eventDefinitionRef

}

/** serialysation of Definitions **/
func (this *BPMSXmlFactory) SerialyzeDefinitions(xmlElement *BPMN20.XsdDefinitions, object *BPMN20.Definitions) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Import **/
	if len(object.M_import) > 0 {
		xmlElement.M_import = make([]*BPMN20.XsdImport, 0)
	}

	/** Now I will save the value of import **/
	for i := 0; i < len(object.M_import); i++ {
		xmlElement.M_import = append(xmlElement.M_import, new(BPMN20.XsdImport))
		this.SerialyzeImport(xmlElement.M_import[i], object.M_import[i])
	}

	/** Serialyze Extension **/
	if len(object.M_extension) > 0 {
		xmlElement.M_extension = make([]*BPMN20.XsdExtension, 0)
	}

	/** Now I will save the value of extension **/
	for i := 0; i < len(object.M_extension); i++ {
		xmlElement.M_extension = append(xmlElement.M_extension, new(BPMN20.XsdExtension))
		this.SerialyzeExtension(xmlElement.M_extension[i], object.M_extension[i])
	}

	/** Serialyze RootElement **/
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_0 = make([]*BPMN20.XsdCategory, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_1 = make([]*BPMN20.XsdGlobalChoreographyTask, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_2 = make([]*BPMN20.XsdChoreography, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_3 = make([]*BPMN20.XsdGlobalConversation, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_4 = make([]*BPMN20.XsdCollaboration, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_5 = make([]*BPMN20.XsdCorrelationProperty, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_6 = make([]*BPMN20.XsdDataStore, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_7 = make([]*BPMN20.XsdEndPoint, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_8 = make([]*BPMN20.XsdError, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_9 = make([]*BPMN20.XsdEscalation, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_10 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_11 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_12 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_13 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_14 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_15 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_16 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_17 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_18 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_19 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_20 = make([]*BPMN20.XsdGlobalBusinessRuleTask, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_21 = make([]*BPMN20.XsdGlobalManualTask, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_22 = make([]*BPMN20.XsdGlobalScriptTask, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_23 = make([]*BPMN20.XsdGlobalTask, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_24 = make([]*BPMN20.XsdGlobalUserTask, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_25 = make([]*BPMN20.XsdInterface, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_26 = make([]*BPMN20.XsdItemDefinition, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_27 = make([]*BPMN20.XsdMessage, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_28 = make([]*BPMN20.XsdPartnerEntity, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_29 = make([]*BPMN20.XsdPartnerRole, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_30 = make([]*BPMN20.XsdProcess, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_31 = make([]*BPMN20.XsdResource, 0)
	}
	if len(object.M_rootElement) > 0 {
		xmlElement.M_rootElement_32 = make([]*BPMN20.XsdSignal, 0)
	}

	/** Now I will save the value of rootElement **/
	for i := 0; i < len(object.M_rootElement); i++ {
		switch v := object.M_rootElement[i].(type) {
		case *BPMN20.Category:
			xmlElement.M_rootElement_0 = append(xmlElement.M_rootElement_0, new(BPMN20.XsdCategory))
			this.SerialyzeCategory(xmlElement.M_rootElement_0[len(xmlElement.M_rootElement_0)-1], v)
			log.Println("Serialyze Definitions:rootElement:Category")
		case *BPMN20.GlobalChoreographyTask:
			xmlElement.M_rootElement_1 = append(xmlElement.M_rootElement_1, new(BPMN20.XsdGlobalChoreographyTask))
			this.SerialyzeGlobalChoreographyTask(xmlElement.M_rootElement_1[len(xmlElement.M_rootElement_1)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalChoreographyTask")
		case *BPMN20.Choreography_impl:
			xmlElement.M_rootElement_2 = append(xmlElement.M_rootElement_2, new(BPMN20.XsdChoreography))
			this.SerialyzeChoreography(xmlElement.M_rootElement_2[len(xmlElement.M_rootElement_2)-1], v)
			log.Println("Serialyze Definitions:rootElement:Choreography")
		case *BPMN20.GlobalConversation:
			xmlElement.M_rootElement_3 = append(xmlElement.M_rootElement_3, new(BPMN20.XsdGlobalConversation))
			this.SerialyzeGlobalConversation(xmlElement.M_rootElement_3[len(xmlElement.M_rootElement_3)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalConversation")
		case *BPMN20.Collaboration_impl:
			xmlElement.M_rootElement_4 = append(xmlElement.M_rootElement_4, new(BPMN20.XsdCollaboration))
			this.SerialyzeCollaboration(xmlElement.M_rootElement_4[len(xmlElement.M_rootElement_4)-1], v)
			log.Println("Serialyze Definitions:rootElement:Collaboration")
		case *BPMN20.CorrelationProperty:
			xmlElement.M_rootElement_5 = append(xmlElement.M_rootElement_5, new(BPMN20.XsdCorrelationProperty))
			this.SerialyzeCorrelationProperty(xmlElement.M_rootElement_5[len(xmlElement.M_rootElement_5)-1], v)
			log.Println("Serialyze Definitions:rootElement:CorrelationProperty")
		case *BPMN20.DataStore:
			xmlElement.M_rootElement_6 = append(xmlElement.M_rootElement_6, new(BPMN20.XsdDataStore))
			this.SerialyzeDataStore(xmlElement.M_rootElement_6[len(xmlElement.M_rootElement_6)-1], v)
			log.Println("Serialyze Definitions:rootElement:DataStore")
		case *BPMN20.EndPoint:
			xmlElement.M_rootElement_7 = append(xmlElement.M_rootElement_7, new(BPMN20.XsdEndPoint))
			this.SerialyzeEndPoint(xmlElement.M_rootElement_7[len(xmlElement.M_rootElement_7)-1], v)
			log.Println("Serialyze Definitions:rootElement:EndPoint")
		case *BPMN20.Error:
			xmlElement.M_rootElement_8 = append(xmlElement.M_rootElement_8, new(BPMN20.XsdError))
			this.SerialyzeError(xmlElement.M_rootElement_8[len(xmlElement.M_rootElement_8)-1], v)
			log.Println("Serialyze Definitions:rootElement:Error")
		case *BPMN20.Escalation:
			xmlElement.M_rootElement_9 = append(xmlElement.M_rootElement_9, new(BPMN20.XsdEscalation))
			this.SerialyzeEscalation(xmlElement.M_rootElement_9[len(xmlElement.M_rootElement_9)-1], v)
			log.Println("Serialyze Definitions:rootElement:Escalation")
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_rootElement_10 = append(xmlElement.M_rootElement_10, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_rootElement_10[len(xmlElement.M_rootElement_10)-1], v)
			log.Println("Serialyze Definitions:rootElement:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_rootElement_11 = append(xmlElement.M_rootElement_11, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_rootElement_11[len(xmlElement.M_rootElement_11)-1], v)
			log.Println("Serialyze Definitions:rootElement:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_rootElement_12 = append(xmlElement.M_rootElement_12, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_rootElement_12[len(xmlElement.M_rootElement_12)-1], v)
			log.Println("Serialyze Definitions:rootElement:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_rootElement_13 = append(xmlElement.M_rootElement_13, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_rootElement_13[len(xmlElement.M_rootElement_13)-1], v)
			log.Println("Serialyze Definitions:rootElement:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_rootElement_14 = append(xmlElement.M_rootElement_14, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_rootElement_14[len(xmlElement.M_rootElement_14)-1], v)
			log.Println("Serialyze Definitions:rootElement:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_rootElement_15 = append(xmlElement.M_rootElement_15, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_rootElement_15[len(xmlElement.M_rootElement_15)-1], v)
			log.Println("Serialyze Definitions:rootElement:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_rootElement_16 = append(xmlElement.M_rootElement_16, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_rootElement_16[len(xmlElement.M_rootElement_16)-1], v)
			log.Println("Serialyze Definitions:rootElement:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_rootElement_17 = append(xmlElement.M_rootElement_17, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_rootElement_17[len(xmlElement.M_rootElement_17)-1], v)
			log.Println("Serialyze Definitions:rootElement:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_rootElement_18 = append(xmlElement.M_rootElement_18, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_rootElement_18[len(xmlElement.M_rootElement_18)-1], v)
			log.Println("Serialyze Definitions:rootElement:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_rootElement_19 = append(xmlElement.M_rootElement_19, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_rootElement_19[len(xmlElement.M_rootElement_19)-1], v)
			log.Println("Serialyze Definitions:rootElement:TimerEventDefinition")
		case *BPMN20.GlobalBusinessRuleTask:
			xmlElement.M_rootElement_20 = append(xmlElement.M_rootElement_20, new(BPMN20.XsdGlobalBusinessRuleTask))
			this.SerialyzeGlobalBusinessRuleTask(xmlElement.M_rootElement_20[len(xmlElement.M_rootElement_20)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalBusinessRuleTask")
		case *BPMN20.GlobalManualTask:
			xmlElement.M_rootElement_21 = append(xmlElement.M_rootElement_21, new(BPMN20.XsdGlobalManualTask))
			this.SerialyzeGlobalManualTask(xmlElement.M_rootElement_21[len(xmlElement.M_rootElement_21)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalManualTask")
		case *BPMN20.GlobalScriptTask:
			xmlElement.M_rootElement_22 = append(xmlElement.M_rootElement_22, new(BPMN20.XsdGlobalScriptTask))
			this.SerialyzeGlobalScriptTask(xmlElement.M_rootElement_22[len(xmlElement.M_rootElement_22)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalScriptTask")
		case *BPMN20.GlobalTask_impl:
			xmlElement.M_rootElement_23 = append(xmlElement.M_rootElement_23, new(BPMN20.XsdGlobalTask))
			this.SerialyzeGlobalTask(xmlElement.M_rootElement_23[len(xmlElement.M_rootElement_23)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalTask")
		case *BPMN20.GlobalUserTask:
			xmlElement.M_rootElement_24 = append(xmlElement.M_rootElement_24, new(BPMN20.XsdGlobalUserTask))
			this.SerialyzeGlobalUserTask(xmlElement.M_rootElement_24[len(xmlElement.M_rootElement_24)-1], v)
			log.Println("Serialyze Definitions:rootElement:GlobalUserTask")
		case *BPMN20.Interface:
			xmlElement.M_rootElement_25 = append(xmlElement.M_rootElement_25, new(BPMN20.XsdInterface))
			this.SerialyzeInterface(xmlElement.M_rootElement_25[len(xmlElement.M_rootElement_25)-1], v)
			log.Println("Serialyze Definitions:rootElement:Interface")
		case *BPMN20.ItemDefinition:
			xmlElement.M_rootElement_26 = append(xmlElement.M_rootElement_26, new(BPMN20.XsdItemDefinition))
			this.SerialyzeItemDefinition(xmlElement.M_rootElement_26[len(xmlElement.M_rootElement_26)-1], v)
			log.Println("Serialyze Definitions:rootElement:ItemDefinition")
		case *BPMN20.Message:
			xmlElement.M_rootElement_27 = append(xmlElement.M_rootElement_27, new(BPMN20.XsdMessage))
			this.SerialyzeMessage(xmlElement.M_rootElement_27[len(xmlElement.M_rootElement_27)-1], v)
			log.Println("Serialyze Definitions:rootElement:Message")
		case *BPMN20.PartnerEntity:
			xmlElement.M_rootElement_28 = append(xmlElement.M_rootElement_28, new(BPMN20.XsdPartnerEntity))
			this.SerialyzePartnerEntity(xmlElement.M_rootElement_28[len(xmlElement.M_rootElement_28)-1], v)
			log.Println("Serialyze Definitions:rootElement:PartnerEntity")
		case *BPMN20.PartnerRole:
			xmlElement.M_rootElement_29 = append(xmlElement.M_rootElement_29, new(BPMN20.XsdPartnerRole))
			this.SerialyzePartnerRole(xmlElement.M_rootElement_29[len(xmlElement.M_rootElement_29)-1], v)
			log.Println("Serialyze Definitions:rootElement:PartnerRole")
		case *BPMN20.Process:
			xmlElement.M_rootElement_30 = append(xmlElement.M_rootElement_30, new(BPMN20.XsdProcess))
			this.SerialyzeProcess(xmlElement.M_rootElement_30[len(xmlElement.M_rootElement_30)-1], v)
			log.Println("Serialyze Definitions:rootElement:Process")
		case *BPMN20.Resource:
			xmlElement.M_rootElement_31 = append(xmlElement.M_rootElement_31, new(BPMN20.XsdResource))
			this.SerialyzeResource(xmlElement.M_rootElement_31[len(xmlElement.M_rootElement_31)-1], v)
			log.Println("Serialyze Definitions:rootElement:Resource")
		case *BPMN20.Signal:
			xmlElement.M_rootElement_32 = append(xmlElement.M_rootElement_32, new(BPMN20.XsdSignal))
			this.SerialyzeSignal(xmlElement.M_rootElement_32[len(xmlElement.M_rootElement_32)-1], v)
			log.Println("Serialyze Definitions:rootElement:Signal")
		}
	}

	/** Serialyze BPMNDiagram **/
	if len(object.M_BPMNDiagram) > 0 {
		xmlElement.M_BPMNDiagram = make([]*BPMNDI.XsdBPMNDiagram, 0)
	}

	/** Now I will save the value of BPMNDiagram **/
	for i := 0; i < len(object.M_BPMNDiagram); i++ {
		xmlElement.M_BPMNDiagram = append(xmlElement.M_BPMNDiagram, new(BPMNDI.XsdBPMNDiagram))
		this.SerialyzeBPMNDiagram(xmlElement.M_BPMNDiagram[i], object.M_BPMNDiagram[i])
	}

	/** Serialyze Relationship **/
	if len(object.M_relationship) > 0 {
		xmlElement.M_relationship = make([]*BPMN20.XsdRelationship, 0)
	}

	/** Now I will save the value of relationship **/
	for i := 0; i < len(object.M_relationship); i++ {
		xmlElement.M_relationship = append(xmlElement.M_relationship, new(BPMN20.XsdRelationship))
		this.SerialyzeRelationship(xmlElement.M_relationship[i], object.M_relationship[i])
	}

	/** Definitions **/
	xmlElement.M_id = object.M_id

	xmlElement.M_name = object.M_name

	/** Definitions **/
	xmlElement.M_targetNamespace = object.M_targetNamespace

	/** Definitions **/
	xmlElement.M_expressionLanguage = object.M_expressionLanguage

	/** Definitions **/
	xmlElement.M_typeLanguage = object.M_typeLanguage

	/** Definitions **/
	xmlElement.M_exporter = object.M_exporter

	/** Definitions **/
	xmlElement.M_exporterVersion = object.M_exporterVersion
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of DataState **/
func (this *BPMSXmlFactory) SerialyzeDataState(xmlElement *BPMN20.XsdDataState, object *BPMN20.DataState) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataState **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of Message **/
func (this *BPMSXmlFactory) SerialyzeMessage(xmlElement *BPMN20.XsdMessage, object *BPMN20.Message) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Message **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref itemRef **/
	xmlElement.M_itemRef = object.M_itemRef
}

/** serialysation of Category **/
func (this *BPMSXmlFactory) SerialyzeCategory(xmlElement *BPMN20.XsdCategory, object *BPMN20.Category) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Category **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze CategoryValue **/
	if len(object.M_categoryValue) > 0 {
		xmlElement.M_categoryValue = make([]*BPMN20.XsdCategoryValue, 0)
	}

	/** Now I will save the value of categoryValue **/
	for i := 0; i < len(object.M_categoryValue); i++ {
		xmlElement.M_categoryValue = append(xmlElement.M_categoryValue, new(BPMN20.XsdCategoryValue))
		this.SerialyzeCategoryValue(xmlElement.M_categoryValue[i], object.M_categoryValue[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of ErrorEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeErrorEventDefinition(xmlElement *BPMN20.XsdErrorEventDefinition, object *BPMN20.ErrorEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ErrorEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref errorRef **/
	xmlElement.M_errorRef = object.M_errorRef
}

/** serialysation of DataStoreReference **/
func (this *BPMSXmlFactory) SerialyzeDataStoreReference(xmlElement *BPMN20.XsdDataStoreReference, object *BPMN20.DataStoreReference) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataStoreReference **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef

	/** Serialyze ref dataStoreRef **/
	xmlElement.M_dataStoreRef = object.M_dataStoreRef
}

/** serialysation of DataObjectReference **/
func (this *BPMSXmlFactory) SerialyzeDataObjectReference(xmlElement *BPMN20.XsdDataObjectReference, object *BPMN20.DataObjectReference) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataObjectReference **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef

	/** Serialyze ref dataObjectRef **/
	xmlElement.M_dataObjectRef = object.M_dataObjectRef
}

/** serialysation of EventBasedGateway **/
func (this *BPMSXmlFactory) SerialyzeEventBasedGateway(xmlElement *BPMN20.XsdEventBasedGateway, object *BPMN20.EventBasedGateway) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** EventBasedGateway **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** GatewayDirection **/
	if object.M_gatewayDirection == BPMN20.GatewayDirection_Unspecified {
		xmlElement.M_gatewayDirection = "##Unspecified"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Converging {
		xmlElement.M_gatewayDirection = "##Converging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Diverging {
		xmlElement.M_gatewayDirection = "##Diverging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Mixed {
		xmlElement.M_gatewayDirection = "##Mixed"
	} else {
		xmlElement.M_gatewayDirection = "##Unspecified"
	}

	/** Gateway **/
	xmlElement.M_instantiate = object.M_instantiate

	/** EventBasedGatewayType **/
	if object.M_eventGatewayType == BPMN20.EventBasedGatewayType_Exclusive {
		xmlElement.M_eventGatewayType = "##Exclusive"
	} else if object.M_eventGatewayType == BPMN20.EventBasedGatewayType_Parallel {
		xmlElement.M_eventGatewayType = "##Parallel"
	} else {
		xmlElement.M_eventGatewayType = "##Exclusive"
	}
}

/** serialysation of BPMNDiagram **/
func (this *BPMSXmlFactory) SerialyzeBPMNDiagram(xmlElement *BPMNDI.XsdBPMNDiagram, object *BPMNDI.BPMNDiagram) {
	if xmlElement == nil {
		return
	}

	/** BPMNDiagram **/
	xmlElement.M_name = object.M_name

	/** BPMNDiagram **/
	xmlElement.M_documentation = object.M_documentation

	/** BPMNDiagram **/
	xmlElement.M_resolution = object.M_resolution

	/** BPMNDiagram **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze BPMNPlane **/
	if object.M_BPMNPlane != nil {
		xmlElement.M_BPMNPlane = new(BPMNDI.XsdBPMNPlane)
	}

	/** Now I will save the value of BPMNPlane **/
	if object.M_BPMNPlane != nil {
		this.SerialyzeBPMNPlane(xmlElement.M_BPMNPlane, object.M_BPMNPlane)
	}

	/** Serialyze BPMNLabelStyle **/
	if len(object.M_BPMNLabelStyle) > 0 {
		xmlElement.M_BPMNLabelStyle = make([]*BPMNDI.XsdBPMNLabelStyle, 0)
	}

	/** Now I will save the value of BPMNLabelStyle **/
	for i := 0; i < len(object.M_BPMNLabelStyle); i++ {
		xmlElement.M_BPMNLabelStyle = append(xmlElement.M_BPMNLabelStyle, new(BPMNDI.XsdBPMNLabelStyle))
		this.SerialyzeBPMNLabelStyle(xmlElement.M_BPMNLabelStyle[i], object.M_BPMNLabelStyle[i])
	}
}

/** serialysation of DataStore **/
func (this *BPMSXmlFactory) SerialyzeDataStore(xmlElement *BPMN20.XsdDataStore, object *BPMN20.DataStore) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataStore **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_capacity = object.M_capacity

	/** RootElement **/
	xmlElement.M_isUnlimited = object.M_isUnlimited

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef
}

/** serialysation of OutputSet **/
func (this *BPMSXmlFactory) SerialyzeOutputSet(xmlElement *BPMN20.XsdOutputSet, object *BPMN20.OutputSet) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** OutputSet **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref dataOutputRefs **/
	xmlElement.M_dataOutputRefs = object.M_dataOutputRefs

	/** Serialyze ref optionalOutputRefs **/
	xmlElement.M_optionalOutputRefs = object.M_optionalOutputRefs

	/** Serialyze ref whileExecutingOutputRefs **/
	xmlElement.M_whileExecutingOutputRefs = object.M_whileExecutingOutputRefs

	/** Serialyze ref inputSetRefs **/
	xmlElement.M_inputSetRefs = object.M_inputSetRefs

	/** BaseElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of Script **/
func (this *BPMSXmlFactory) SerialyzeScript(xmlElement *BPMN20.XsdScript, object *BPMN20.Script) {
	if xmlElement == nil {
		return
	}
	xmlElement.M_script = object.M_script
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of BusinessRuleTask **/
func (this *BPMSXmlFactory) SerialyzeBusinessRuleTask(xmlElement *BPMN20.XsdBusinessRuleTask, object *BPMN20.BusinessRuleTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** BusinessRuleTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##unspecified"
		}
	}
}

/** serialysation of FormalExpression **/
func (this *BPMSXmlFactory) SerialyzeFormalExpression(xmlElement *BPMN20.XsdFormalExpression, object *BPMN20.FormalExpression) {
	if xmlElement == nil {
		return
	}

	/** Expression **/
	xmlElement.M_language = object.M_language

	/** Serialyze ref evaluatesToTypeRef **/
	xmlElement.M_evaluatesToTypeRef = object.M_evaluatesToTypeRef
	/** other content **/
	exprStr := object.GetOther().(string)
	xmlElement.M_other = exprStr
}

/** serialysation of Rendering **/
func (this *BPMSXmlFactory) SerialyzeRendering(xmlElement *BPMN20.XsdRendering, object *BPMN20.Rendering) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Rendering **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of StandardLoopCharacteristics **/
func (this *BPMSXmlFactory) SerialyzeStandardLoopCharacteristics(xmlElement *BPMN20.XsdStandardLoopCharacteristics, object *BPMN20.StandardLoopCharacteristics) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** StandardLoopCharacteristics **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/
	if object.M_loopCondition != nil {
		xmlElement.M_loopCondition = new(BPMN20.XsdLoopCondition)
	}

	/** Now I will save the value of loopCondition **/
	if object.M_loopCondition != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_loopCondition)), object.M_loopCondition)
	}

	/** LoopCharacteristics **/
	xmlElement.M_testBefore = object.M_testBefore

	/** LoopCharacteristics **/
	xmlElement.M_loopMaximum = object.M_loopMaximum
}

/** serialysation of InclusiveGateway **/
func (this *BPMSXmlFactory) SerialyzeInclusiveGateway(xmlElement *BPMN20.XsdInclusiveGateway, object *BPMN20.InclusiveGateway) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** InclusiveGateway **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** GatewayDirection **/
	if object.M_gatewayDirection == BPMN20.GatewayDirection_Unspecified {
		xmlElement.M_gatewayDirection = "##Unspecified"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Converging {
		xmlElement.M_gatewayDirection = "##Converging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Diverging {
		xmlElement.M_gatewayDirection = "##Diverging"
	} else if object.M_gatewayDirection == BPMN20.GatewayDirection_Mixed {
		xmlElement.M_gatewayDirection = "##Mixed"
	} else {
		xmlElement.M_gatewayDirection = "##Unspecified"
	}

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
}

/** serialysation of ExtensionElements **/
func (this *BPMSXmlFactory) SerialyzeExtensionElements(xmlElement *BPMN20.XsdExtensionElements, object *BPMN20.ExtensionElements) {
	if xmlElement == nil {
		return
	}
	xmlElement.M_value = object.M_value
	this.m_references["ExtensionElements"] = object
}

/** serialysation of GlobalTask **/
func (this *BPMSXmlFactory) SerialyzeGlobalTask(xmlElement *BPMN20.XsdGlobalTask, object *BPMN20.GlobalTask_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref supportedInterfaceRef **/
	xmlElement.M_supportedInterfaceRef = object.M_supportedInterfaceRef

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze InputOutputBinding **/
	if len(object.M_ioBinding) > 0 {
		xmlElement.M_ioBinding = make([]*BPMN20.XsdInputOutputBinding, 0)
	}

	/** Now I will save the value of ioBinding **/
	for i := 0; i < len(object.M_ioBinding); i++ {
		xmlElement.M_ioBinding = append(xmlElement.M_ioBinding, new(BPMN20.XsdInputOutputBinding))
		this.SerialyzeInputOutputBinding(xmlElement.M_ioBinding[i], object.M_ioBinding[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze GlobalTask:resourceRole:ResourceRole")
		}
	}
}

/** serialysation of BPMNLabel **/
func (this *BPMSXmlFactory) SerialyzeBPMNLabel(xmlElement *BPMNDI.XsdBPMNLabel, object *BPMNDI.BPMNLabel) {
	if xmlElement == nil {
		return
	}

	/** BPMNLabel **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Bounds **/
	if object.M_Bounds != nil {
		xmlElement.M_Bounds = new(DC.XsdBounds)
	}

	/** Now I will save the value of Bounds **/
	if object.M_Bounds != nil {
		this.SerialyzeBounds(xmlElement.M_Bounds, object.M_Bounds)
	}

	/** Serialyze ref labelStyle **/
	xmlElement.M_labelStyle = object.M_labelStyle
}

/** serialysation of PotentialOwner **/
func (this *BPMSXmlFactory) SerialyzePotentialOwner(xmlElement *BPMN20.XsdPotentialOwner, object *BPMN20.PotentialOwner) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** PotentialOwner **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref resourceRef **/
	xmlElement.M_resourceRef = &object.M_resourceRef

	/** Serialyze ResourceParameterBinding **/
	if len(object.M_resourceParameterBinding) > 0 {
		xmlElement.M_resourceParameterBinding = make([]*BPMN20.XsdResourceParameterBinding, 0)
	}

	/** Now I will save the value of resourceParameterBinding **/
	for i := 0; i < len(object.M_resourceParameterBinding); i++ {
		xmlElement.M_resourceParameterBinding = append(xmlElement.M_resourceParameterBinding, new(BPMN20.XsdResourceParameterBinding))
		this.SerialyzeResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], object.M_resourceParameterBinding[i])
	}

	/** Serialyze ResourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		xmlElement.M_resourceAssignmentExpression = new(BPMN20.XsdResourceAssignmentExpression)
	}

	/** Now I will save the value of resourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		this.SerialyzeResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)
	}
}

/** serialysation of Auditing **/
func (this *BPMSXmlFactory) SerialyzeAuditing(xmlElement *BPMN20.XsdAuditing, object *BPMN20.Auditing) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Auditing **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of DataOutputAssociation **/
func (this *BPMSXmlFactory) SerialyzeDataOutputAssociation(xmlElement *BPMN20.XsdDataOutputAssociation, object *BPMN20.DataOutputAssociation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataOutputAssociation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = object.M_targetRef

	/** Serialyze FormalExpression **/
	if object.M_transformation != nil {
		xmlElement.M_transformation = new(BPMN20.XsdTransformation)
	}

	/** Now I will save the value of transformation **/
	if object.M_transformation != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_transformation)), object.M_transformation)
	}

	/** Serialyze Assignment **/
	if len(object.M_assignment) > 0 {
		xmlElement.M_assignment = make([]*BPMN20.XsdAssignment, 0)
	}

	/** Now I will save the value of assignment **/
	for i := 0; i < len(object.M_assignment); i++ {
		xmlElement.M_assignment = append(xmlElement.M_assignment, new(BPMN20.XsdAssignment))
		this.SerialyzeAssignment(xmlElement.M_assignment[i], object.M_assignment[i])
	}
}

/** serialysation of CallActivity **/
func (this *BPMSXmlFactory) SerialyzeCallActivity(xmlElement *BPMN20.XsdCallActivity, object *BPMN20.CallActivity) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CallActivity **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default

	/** Serialyze ref calledElement **/
	xmlElement.M_calledElement = object.M_calledElement
}

/** serialysation of Participant **/
func (this *BPMSXmlFactory) SerialyzeParticipant(xmlElement *BPMN20.XsdParticipant, object *BPMN20.Participant) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Participant **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref interfaceRef **/
	xmlElement.M_interfaceRef = object.M_interfaceRef

	/** Serialyze ref endPointRef **/
	xmlElement.M_endPointRef = object.M_endPointRef

	/** Serialyze ParticipantMultiplicity **/
	if object.M_participantMultiplicity != nil {
		xmlElement.M_participantMultiplicity = new(BPMN20.XsdParticipantMultiplicity)
	}

	/** Now I will save the value of participantMultiplicity **/
	if object.M_participantMultiplicity != nil {
		this.SerialyzeParticipantMultiplicity(xmlElement.M_participantMultiplicity, object.M_participantMultiplicity)
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref processRef **/
	xmlElement.M_processRef = object.M_processRef
}

/** serialysation of CorrelationProperty **/
func (this *BPMSXmlFactory) SerialyzeCorrelationProperty(xmlElement *BPMN20.XsdCorrelationProperty, object *BPMN20.CorrelationProperty) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CorrelationProperty **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze CorrelationPropertyRetrievalExpression **/
	if len(object.M_correlationPropertyRetrievalExpression) > 0 {
		xmlElement.M_correlationPropertyRetrievalExpression = make([]*BPMN20.XsdCorrelationPropertyRetrievalExpression, 0)
	}

	/** Now I will save the value of correlationPropertyRetrievalExpression **/
	for i := 0; i < len(object.M_correlationPropertyRetrievalExpression); i++ {
		xmlElement.M_correlationPropertyRetrievalExpression = append(xmlElement.M_correlationPropertyRetrievalExpression, new(BPMN20.XsdCorrelationPropertyRetrievalExpression))
		this.SerialyzeCorrelationPropertyRetrievalExpression(xmlElement.M_correlationPropertyRetrievalExpression[i], object.M_correlationPropertyRetrievalExpression[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref type **/
	xmlElement.M_type = object.M_type
}

/** serialysation of Expression **/
func (this *BPMSXmlFactory) SerialyzeExpression(xmlElement *BPMN20.XsdExpression, object *BPMN20.Expression_impl) {
	if xmlElement == nil {
		return
	}
	/** other content **/
	exprStr := object.GetOther().(string)
	xmlElement.M_other = exprStr
}

/** serialysation of InputOutputBinding **/
func (this *BPMSXmlFactory) SerialyzeInputOutputBinding(xmlElement *BPMN20.XsdInputOutputBinding, object *BPMN20.InputOutputBinding) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** InputOutputBinding **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref operationRef **/
	xmlElement.M_operationRef = object.M_operationRef

	/** Serialyze ref inputDataRef **/
	xmlElement.M_inputDataRef = object.M_inputDataRef

	/** Serialyze ref outputDataRef **/
	xmlElement.M_outputDataRef = object.M_outputDataRef
}

/** serialysation of Signal **/
func (this *BPMSXmlFactory) SerialyzeSignal(xmlElement *BPMN20.XsdSignal, object *BPMN20.Signal) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Signal **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref structureRef **/
	xmlElement.M_structureRef = object.M_structureRef
}

/** serialysation of BPMNShape **/
func (this *BPMSXmlFactory) SerialyzeBPMNShape(xmlElement *BPMNDI.XsdBPMNShape, object *BPMNDI.BPMNShape) {
	if xmlElement == nil {
		return
	}

	/** BPMNShape **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Bounds **/

	/** Now I will save the value of Bounds **/
	if object.M_Bounds != nil {
		this.SerialyzeBounds(&xmlElement.M_Bounds, object.M_Bounds)
	}

	/** Serialyze BPMNLabel **/
	if object.M_BPMNLabel != nil {
		xmlElement.M_BPMNLabel = new(BPMNDI.XsdBPMNLabel)
	}

	/** Now I will save the value of BPMNLabel **/
	if object.M_BPMNLabel != nil {
		this.SerialyzeBPMNLabel(xmlElement.M_BPMNLabel, object.M_BPMNLabel)
	}
	/** bpmnElement **/
	xmlElement.M_bpmnElement = object.M_bpmnElement

	/** LabeledShape **/
	xmlElement.M_isHorizontal = object.M_isHorizontal

	/** LabeledShape **/
	xmlElement.M_isExpanded = object.M_isExpanded

	/** LabeledShape **/
	xmlElement.M_isMarkerVisible = object.M_isMarkerVisible

	/** LabeledShape **/
	xmlElement.M_isMessageVisible = object.M_isMessageVisible

	/** ParticipantBandKind **/
	if object.M_participantBandKind == BPMNDI.ParticipantBandKind_Top_initiating {
		xmlElement.M_participantBandKind = "##top_initiating"
	} else if object.M_participantBandKind == BPMNDI.ParticipantBandKind_Middle_initiating {
		xmlElement.M_participantBandKind = "##middle_initiating"
	} else if object.M_participantBandKind == BPMNDI.ParticipantBandKind_Bottom_initiating {
		xmlElement.M_participantBandKind = "##bottom_initiating"
	} else if object.M_participantBandKind == BPMNDI.ParticipantBandKind_Top_non_initiating {
		xmlElement.M_participantBandKind = "##top_non_initiating"
	} else if object.M_participantBandKind == BPMNDI.ParticipantBandKind_Middle_non_initiating {
		xmlElement.M_participantBandKind = "##middle_non_initiating"
	} else if object.M_participantBandKind == BPMNDI.ParticipantBandKind_Bottom_non_initiating {
		xmlElement.M_participantBandKind = "##bottom_non_initiating"
	}

	/** Serialyze ref choreographyActivityShape **/
	xmlElement.M_choreographyActivityShape = object.M_choreographyActivityShape
}

/** serialysation of Conversation **/
func (this *BPMSXmlFactory) SerialyzeConversation(xmlElement *BPMN20.XsdConversation, object *BPMN20.Conversation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Conversation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** Serialyze ref messageFlowRef **/
	xmlElement.M_messageFlowRef = object.M_messageFlowRef

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of Relationship **/
func (this *BPMSXmlFactory) SerialyzeRelationship(xmlElement *BPMN20.XsdRelationship, object *BPMN20.Relationship) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Relationship **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref source **/
	xmlElement.M_source = object.M_source

	/** Serialyze ref target **/
	xmlElement.M_target = object.M_target

	/** BaseElement **/
	xmlElement.M_type = object.M_type

	/** RelationshipDirection **/
	if object.M_direction == BPMN20.RelationshipDirection_None {
		xmlElement.M_direction = "##None"
	} else if object.M_direction == BPMN20.RelationshipDirection_Forward {
		xmlElement.M_direction = "##Forward"
	} else if object.M_direction == BPMN20.RelationshipDirection_Backward {
		xmlElement.M_direction = "##Backward"
	} else if object.M_direction == BPMN20.RelationshipDirection_Both {
		xmlElement.M_direction = "##Both"
	}
}

/** serialysation of InputSet **/
func (this *BPMSXmlFactory) SerialyzeInputSet(xmlElement *BPMN20.XsdInputSet, object *BPMN20.InputSet) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** InputSet **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref dataInputRefs **/
	xmlElement.M_dataInputRefs = object.M_dataInputRefs

	/** Serialyze ref optionalInputRefs **/
	xmlElement.M_optionalInputRefs = object.M_optionalInputRefs

	/** Serialyze ref whileExecutingInputRefs **/
	xmlElement.M_whileExecutingInputRefs = object.M_whileExecutingInputRefs

	/** Serialyze ref outputSetRefs **/
	xmlElement.M_outputSetRefs = object.M_outputSetRefs

	/** BaseElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of ResourceAssignmentExpression **/
func (this *BPMSXmlFactory) SerialyzeResourceAssignmentExpression(xmlElement *BPMN20.XsdResourceAssignmentExpression, object *BPMN20.ResourceAssignmentExpression) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ResourceAssignmentExpression **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Expression **/

	/** Now I will save the value of expression **/
	if object.M_expression != nil {
		switch v := object.M_expression.(type) {
		case *BPMN20.FormalExpression:
			xmlElement.M_expression_0 = new(BPMN20.XsdFormalExpression)
			this.SerialyzeFormalExpression(xmlElement.M_expression_0, v)
			log.Println("Serialyze ResourceAssignmentExpression:expression:FormalExpression")
		case *BPMN20.Expression_impl:
			xmlElement.M_expression_1 = new(BPMN20.XsdExpression)
			this.SerialyzeExpression(xmlElement.M_expression_1, v)
			log.Println("Serialyze ResourceAssignmentExpression:expression:Expression")
		}
	}
}

/** serialysation of CorrelationPropertyBinding **/
func (this *BPMSXmlFactory) SerialyzeCorrelationPropertyBinding(xmlElement *BPMN20.XsdCorrelationPropertyBinding, object *BPMN20.CorrelationPropertyBinding) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CorrelationPropertyBinding **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/

	/** Now I will save the value of dataPath **/
	if object.M_dataPath != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(&xmlElement.M_dataPath)), object.M_dataPath)
	}

	/** Serialyze ref correlationPropertyRef **/
	xmlElement.M_correlationPropertyRef = object.M_correlationPropertyRef
}

/** serialysation of AdHocSubProcess **/
func (this *BPMSXmlFactory) SerialyzeAdHocSubProcess(xmlElement *BPMN20.XsdAdHocSubProcess, object *BPMN20.AdHocSubProcess) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** AdHocSubProcess **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default

	/** Serialyze LaneSet **/
	if len(object.M_laneSet) > 0 {
		xmlElement.M_laneSet = make([]*BPMN20.XsdLaneSet, 0)
	}

	/** Now I will save the value of laneSet **/
	for i := 0; i < len(object.M_laneSet); i++ {
		xmlElement.M_laneSet = append(xmlElement.M_laneSet, new(BPMN20.XsdLaneSet))
		this.SerialyzeLaneSet(xmlElement.M_laneSet[i], object.M_laneSet[i])
	}

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze SubProcess:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze SubProcess:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze SubProcess:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze SubProcess:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze SubProcess:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze SubProcess:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze SubProcess:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze SubProcess:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze SubProcess:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze SubProcess:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze SubProcess:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze SubProcess:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze SubProcess:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze SubProcess:flowElement:UserTask")
		}
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze SubProcess:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze SubProcess:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze SubProcess:artifact:TextAnnotation")
		}
	}

	/** Activity **/
	xmlElement.M_triggeredByEvent = object.M_triggeredByEvent

	/** Serialyze FormalExpression **/
	if object.M_completionCondition != nil {
		xmlElement.M_completionCondition = new(BPMN20.XsdCompletionCondition)
	}

	/** Now I will save the value of completionCondition **/
	if object.M_completionCondition != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_completionCondition)), object.M_completionCondition)
	}

	/** SubProcess **/
	xmlElement.M_cancelRemainingInstances = object.M_cancelRemainingInstances

	/** AdHocOrdering **/
	if object.M_ordering == BPMN20.AdHocOrdering_Parallel {
		xmlElement.M_ordering = "##Parallel"
	} else if object.M_ordering == BPMN20.AdHocOrdering_Sequential {
		xmlElement.M_ordering = "##Sequential"
	}
}

/** serialysation of SequenceFlow **/
func (this *BPMSXmlFactory) SerialyzeSequenceFlow(xmlElement *BPMN20.XsdSequenceFlow, object *BPMN20.SequenceFlow) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** SequenceFlow **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze FormalExpression **/
	if object.M_conditionExpression != nil {
		xmlElement.M_conditionExpression = new(BPMN20.XsdConditionExpression)
	}

	/** Now I will save the value of conditionExpression **/
	if object.M_conditionExpression != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_conditionExpression)), object.M_conditionExpression)
	}

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = object.M_targetRef

	/** FlowElement **/
	xmlElement.M_isImmediate = object.M_isImmediate
}

/** serialysation of Transaction **/
func (this *BPMSXmlFactory) SerialyzeTransaction(xmlElement *BPMN20.XsdTransaction, object *BPMN20.Transaction) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Transaction **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default

	/** Serialyze LaneSet **/
	if len(object.M_laneSet) > 0 {
		xmlElement.M_laneSet = make([]*BPMN20.XsdLaneSet, 0)
	}

	/** Now I will save the value of laneSet **/
	for i := 0; i < len(object.M_laneSet); i++ {
		xmlElement.M_laneSet = append(xmlElement.M_laneSet, new(BPMN20.XsdLaneSet))
		this.SerialyzeLaneSet(xmlElement.M_laneSet[i], object.M_laneSet[i])
	}

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze SubProcess:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze SubProcess:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze SubProcess:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze SubProcess:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze SubProcess:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze SubProcess:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze SubProcess:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze SubProcess:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze SubProcess:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze SubProcess:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze SubProcess:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze SubProcess:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze SubProcess:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze SubProcess:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze SubProcess:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze SubProcess:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze SubProcess:flowElement:UserTask")
		}
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze SubProcess:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze SubProcess:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze SubProcess:artifact:TextAnnotation")
		}
	}

	/** Activity **/
	xmlElement.M_triggeredByEvent = object.M_triggeredByEvent
	if len(object.M_methodStr) > 0 {
		xmlElement.M_method = object.M_methodStr
	} else {

		/** TransactionMethod **/
		if object.M_method == BPMN20.TransactionMethod_Compensate {
			xmlElement.M_method = "##Compensate"
		} else if object.M_method == BPMN20.TransactionMethod_Image {
			xmlElement.M_method = "##Image"
		} else if object.M_method == BPMN20.TransactionMethod_Store {
			xmlElement.M_method = "##Store"
		} else {
			xmlElement.M_method = "##Compensate"
		}
	}
}

/** serialysation of Bounds **/
func (this *BPMSXmlFactory) SerialyzeBounds(xmlElement *DC.XsdBounds, object *DC.Bounds) {
	if xmlElement == nil {
		return
	}

	/** Bounds **/
	xmlElement.M_x = object.M_x

	/** Bounds **/
	xmlElement.M_y = object.M_y

	/** Bounds **/
	xmlElement.M_width = object.M_width

	/** Bounds **/
	xmlElement.M_height = object.M_height
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of CategoryValue **/
func (this *BPMSXmlFactory) SerialyzeCategoryValue(xmlElement *BPMN20.XsdCategoryValue, object *BPMN20.CategoryValue) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CategoryValue **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_value = object.M_value
}

/** serialysation of CorrelationKey **/
func (this *BPMSXmlFactory) SerialyzeCorrelationKey(xmlElement *BPMN20.XsdCorrelationKey, object *BPMN20.CorrelationKey) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CorrelationKey **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref correlationPropertyRef **/
	xmlElement.M_correlationPropertyRef = object.M_correlationPropertyRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of InputOutputSpecification **/
func (this *BPMSXmlFactory) SerialyzeInputOutputSpecification(xmlElement *BPMN20.XsdInputOutputSpecification, object *BPMN20.InputOutputSpecification) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** InputOutputSpecification **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze DataInput **/
	if len(object.M_dataInput) > 0 {
		xmlElement.M_dataInput = make([]*BPMN20.XsdDataInput, 0)
	}

	/** Now I will save the value of dataInput **/
	for i := 0; i < len(object.M_dataInput); i++ {
		xmlElement.M_dataInput = append(xmlElement.M_dataInput, new(BPMN20.XsdDataInput))
		this.SerialyzeDataInput(xmlElement.M_dataInput[i], object.M_dataInput[i])
	}

	/** Serialyze DataOutput **/
	if len(object.M_dataOutput) > 0 {
		xmlElement.M_dataOutput = make([]*BPMN20.XsdDataOutput, 0)
	}

	/** Now I will save the value of dataOutput **/
	for i := 0; i < len(object.M_dataOutput); i++ {
		xmlElement.M_dataOutput = append(xmlElement.M_dataOutput, new(BPMN20.XsdDataOutput))
		this.SerialyzeDataOutput(xmlElement.M_dataOutput[i], object.M_dataOutput[i])
	}

	/** Serialyze InputSet **/
	if len(object.M_inputSet) > 0 {
		xmlElement.M_inputSet = make([]*BPMN20.XsdInputSet, 0)
	}

	/** Now I will save the value of inputSet **/
	for i := 0; i < len(object.M_inputSet); i++ {
		xmlElement.M_inputSet = append(xmlElement.M_inputSet, new(BPMN20.XsdInputSet))
		this.SerialyzeInputSet(xmlElement.M_inputSet[i], object.M_inputSet[i])
	}

	/** Serialyze OutputSet **/
	if len(object.M_outputSet) > 0 {
		xmlElement.M_outputSet = make([]*BPMN20.XsdOutputSet, 0)
	}

	/** Now I will save the value of outputSet **/
	for i := 0; i < len(object.M_outputSet); i++ {
		xmlElement.M_outputSet = append(xmlElement.M_outputSet, new(BPMN20.XsdOutputSet))
		this.SerialyzeOutputSet(xmlElement.M_outputSet[i], object.M_outputSet[i])
	}
}

/** serialysation of Error **/
func (this *BPMSXmlFactory) SerialyzeError(xmlElement *BPMN20.XsdError, object *BPMN20.Error) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Error **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_errorCode = object.M_errorCode

	/** Serialyze ref structureRef **/
	xmlElement.M_structureRef = object.M_structureRef
}

/** serialysation of UserTask **/
func (this *BPMSXmlFactory) SerialyzeUserTask(xmlElement *BPMN20.XsdUserTask, object *BPMN20.UserTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** UserTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default

	/** Serialyze Rendering **/
	if len(object.M_rendering) > 0 {
		xmlElement.M_rendering = make([]*BPMN20.XsdRendering, 0)
	}

	/** Now I will save the value of rendering **/
	for i := 0; i < len(object.M_rendering); i++ {
		xmlElement.M_rendering = append(xmlElement.M_rendering, new(BPMN20.XsdRendering))
		this.SerialyzeRendering(xmlElement.M_rendering[i], object.M_rendering[i])
	}
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##unspecified"
		}
	}
}

/** serialysation of SubConversation **/
func (this *BPMSXmlFactory) SerialyzeSubConversation(xmlElement *BPMN20.XsdSubConversation, object *BPMN20.SubConversation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** SubConversation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** Serialyze ref messageFlowRef **/
	xmlElement.M_messageFlowRef = object.M_messageFlowRef

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ConversationNode **/
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_0 = make([]*BPMN20.XsdCallConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_1 = make([]*BPMN20.XsdConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_2 = make([]*BPMN20.XsdSubConversation, 0)
	}

	/** Now I will save the value of conversationNode **/
	for i := 0; i < len(object.M_conversationNode); i++ {
		switch v := object.M_conversationNode[i].(type) {
		case *BPMN20.CallConversation:
			xmlElement.M_conversationNode_0 = append(xmlElement.M_conversationNode_0, new(BPMN20.XsdCallConversation))
			this.SerialyzeCallConversation(xmlElement.M_conversationNode_0[len(xmlElement.M_conversationNode_0)-1], v)
			log.Println("Serialyze SubConversation:conversationNode:CallConversation")
		case *BPMN20.Conversation:
			xmlElement.M_conversationNode_1 = append(xmlElement.M_conversationNode_1, new(BPMN20.XsdConversation))
			this.SerialyzeConversation(xmlElement.M_conversationNode_1[len(xmlElement.M_conversationNode_1)-1], v)
			log.Println("Serialyze SubConversation:conversationNode:Conversation")
		case *BPMN20.SubConversation:
			xmlElement.M_conversationNode_2 = append(xmlElement.M_conversationNode_2, new(BPMN20.XsdSubConversation))
			this.SerialyzeSubConversation(xmlElement.M_conversationNode_2[len(xmlElement.M_conversationNode_2)-1], v)
			log.Println("Serialyze SubConversation:conversationNode:SubConversation")
		}
	}
}

/** serialysation of ResourceParameter **/
func (this *BPMSXmlFactory) SerialyzeResourceParameter(xmlElement *BPMN20.XsdResourceParameter, object *BPMN20.ResourceParameter) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ResourceParameter **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref type **/
	xmlElement.M_type = object.M_type

	/** BaseElement **/
	xmlElement.M_isRequired = object.M_isRequired
}

/** serialysation of ConversationLink **/
func (this *BPMSXmlFactory) SerialyzeConversationLink(xmlElement *BPMN20.XsdConversationLink, object *BPMN20.ConversationLink) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ConversationLink **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = object.M_targetRef
}

/** serialysation of MessageEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeMessageEventDefinition(xmlElement *BPMN20.XsdMessageEventDefinition, object *BPMN20.MessageEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** MessageEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref operationRef **/
	xmlElement.M_operationRef = &object.M_operationRef

	/** Serialyze ref messageRef **/
	xmlElement.M_messageRef = object.M_messageRef
}

/** serialysation of DataInputAssociation **/
func (this *BPMSXmlFactory) SerialyzeDataInputAssociation(xmlElement *BPMN20.XsdDataInputAssociation, object *BPMN20.DataInputAssociation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** DataInputAssociation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = object.M_targetRef

	/** Serialyze FormalExpression **/
	if object.M_transformation != nil {
		xmlElement.M_transformation = new(BPMN20.XsdTransformation)
	}

	/** Now I will save the value of transformation **/
	if object.M_transformation != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(xmlElement.M_transformation)), object.M_transformation)
	}

	/** Serialyze Assignment **/
	if len(object.M_assignment) > 0 {
		xmlElement.M_assignment = make([]*BPMN20.XsdAssignment, 0)
	}

	/** Now I will save the value of assignment **/
	for i := 0; i < len(object.M_assignment); i++ {
		xmlElement.M_assignment = append(xmlElement.M_assignment, new(BPMN20.XsdAssignment))
		this.SerialyzeAssignment(xmlElement.M_assignment[i], object.M_assignment[i])
	}
}

/** serialysation of BPMNPlane **/
func (this *BPMSXmlFactory) SerialyzeBPMNPlane(xmlElement *BPMNDI.XsdBPMNPlane, object *BPMNDI.BPMNPlane) {
	if xmlElement == nil {
		return
	}

	/** BPMNPlane **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze DiagramElement **/
	if len(object.M_DiagramElement) > 0 {
		xmlElement.M_DiagramElement_0 = make([]*BPMNDI.XsdBPMNShape, 0)
	}
	if len(object.M_DiagramElement) > 0 {
		xmlElement.M_DiagramElement_1 = make([]*BPMNDI.XsdBPMNEdge, 0)
	}

	/** Now I will save the value of DiagramElement **/
	for i := 0; i < len(object.M_DiagramElement); i++ {
		switch v := object.M_DiagramElement[i].(type) {
		case *BPMNDI.BPMNShape:
			xmlElement.M_DiagramElement_0 = append(xmlElement.M_DiagramElement_0, new(BPMNDI.XsdBPMNShape))
			this.SerialyzeBPMNShape(xmlElement.M_DiagramElement_0[len(xmlElement.M_DiagramElement_0)-1], v)
			log.Println("Serialyze Plane:DiagramElement:BPMNShape")
		case *BPMNDI.BPMNEdge:
			xmlElement.M_DiagramElement_1 = append(xmlElement.M_DiagramElement_1, new(BPMNDI.XsdBPMNEdge))
			this.SerialyzeBPMNEdge(xmlElement.M_DiagramElement_1[len(xmlElement.M_DiagramElement_1)-1], v)
			log.Println("Serialyze Plane:DiagramElement:BPMNEdge")
		}
	}
	/** bpmnElement **/
	xmlElement.M_bpmnElement = object.M_bpmnElement
}

/** serialysation of TextAnnotation **/
func (this *BPMSXmlFactory) SerialyzeTextAnnotation(xmlElement *BPMN20.XsdTextAnnotation, object *BPMN20.TextAnnotation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** TextAnnotation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Text **/
	if object.M_text != nil {
		xmlElement.M_text = new(BPMN20.XsdText)
	}

	/** Now I will save the value of text **/
	if object.M_text != nil {
		this.SerialyzeText(xmlElement.M_text, object.M_text)
	}

	/** Artifact **/
	xmlElement.M_textFormat = object.M_textFormat
}

/** serialysation of PartnerEntity **/
func (this *BPMSXmlFactory) SerialyzePartnerEntity(xmlElement *BPMN20.XsdPartnerEntity, object *BPMN20.PartnerEntity) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** PartnerEntity **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** RootElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of Property **/
func (this *BPMSXmlFactory) SerialyzeProperty(xmlElement *BPMN20.XsdProperty, object *BPMN20.Property) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Property **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze DataState **/
	if object.M_dataState != nil {
		xmlElement.M_dataState = new(BPMN20.XsdDataState)
	}

	/** Now I will save the value of dataState **/
	if object.M_dataState != nil {
		this.SerialyzeDataState(xmlElement.M_dataState, object.M_dataState)
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref itemSubjectRef **/
	xmlElement.M_itemSubjectRef = object.M_itemSubjectRef
}

/** serialysation of StartEvent **/
func (this *BPMSXmlFactory) SerialyzeStartEvent(xmlElement *BPMN20.XsdStartEvent, object *BPMN20.StartEvent) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** StartEvent **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataOutput **/
	if len(object.M_dataOutput) > 0 {
		xmlElement.M_dataOutput = make([]*BPMN20.XsdDataOutput, 0)
	}

	/** Now I will save the value of dataOutput **/
	for i := 0; i < len(object.M_dataOutput); i++ {
		xmlElement.M_dataOutput = append(xmlElement.M_dataOutput, new(BPMN20.XsdDataOutput))
		this.SerialyzeDataOutput(xmlElement.M_dataOutput[i], object.M_dataOutput[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze OutputSet **/
	if object.M_outputSet != nil {
		xmlElement.M_outputSet = new(BPMN20.XsdOutputSet)
	}

	/** Now I will save the value of outputSet **/
	if object.M_outputSet != nil {
		this.SerialyzeOutputSet(xmlElement.M_outputSet, object.M_outputSet)
	}

	/** Serialyze EventDefinition **/
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_0 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_1 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_2 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_3 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_4 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_5 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_6 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_7 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_8 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_9 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}

	/** Now I will save the value of eventDefinition **/
	for i := 0; i < len(object.M_eventDefinition); i++ {
		switch v := object.M_eventDefinition[i].(type) {
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_eventDefinition_0 = append(xmlElement.M_eventDefinition_0, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_eventDefinition_0[len(xmlElement.M_eventDefinition_0)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_eventDefinition_1 = append(xmlElement.M_eventDefinition_1, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_eventDefinition_1[len(xmlElement.M_eventDefinition_1)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_eventDefinition_2 = append(xmlElement.M_eventDefinition_2, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_eventDefinition_2[len(xmlElement.M_eventDefinition_2)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_eventDefinition_3 = append(xmlElement.M_eventDefinition_3, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_eventDefinition_3[len(xmlElement.M_eventDefinition_3)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_eventDefinition_4 = append(xmlElement.M_eventDefinition_4, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_eventDefinition_4[len(xmlElement.M_eventDefinition_4)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_eventDefinition_5 = append(xmlElement.M_eventDefinition_5, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_eventDefinition_5[len(xmlElement.M_eventDefinition_5)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_eventDefinition_6 = append(xmlElement.M_eventDefinition_6, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_eventDefinition_6[len(xmlElement.M_eventDefinition_6)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_eventDefinition_7 = append(xmlElement.M_eventDefinition_7, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_eventDefinition_7[len(xmlElement.M_eventDefinition_7)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_eventDefinition_8 = append(xmlElement.M_eventDefinition_8, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_eventDefinition_8[len(xmlElement.M_eventDefinition_8)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_eventDefinition_9 = append(xmlElement.M_eventDefinition_9, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_eventDefinition_9[len(xmlElement.M_eventDefinition_9)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:TimerEventDefinition")
		}
	}

	/** Serialyze ref eventDefinitionRef **/
	xmlElement.M_eventDefinitionRef = object.M_eventDefinitionRef

	/** Event **/
	xmlElement.M_parallelMultiple = object.M_parallelMultiple

	/** CatchEvent **/
	xmlElement.M_isInterrupting = object.M_isInterrupting
}

/** serialysation of CompensateEventDefinition **/
func (this *BPMSXmlFactory) SerialyzeCompensateEventDefinition(xmlElement *BPMN20.XsdCompensateEventDefinition, object *BPMN20.CompensateEventDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** CompensateEventDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** EventDefinition **/
	xmlElement.M_waitForCompletion = object.M_waitForCompletion

	/** Serialyze ref activityRef **/
	xmlElement.M_activityRef = object.M_activityRef
}

/** serialysation of PartnerRole **/
func (this *BPMSXmlFactory) SerialyzePartnerRole(xmlElement *BPMN20.XsdPartnerRole, object *BPMN20.PartnerRole) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** PartnerRole **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref participantRef **/
	xmlElement.M_participantRef = object.M_participantRef

	/** RootElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of Process **/
func (this *BPMSXmlFactory) SerialyzeProcess(xmlElement *BPMN20.XsdProcess, object *BPMN20.Process) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Process **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref supportedInterfaceRef **/
	xmlElement.M_supportedInterfaceRef = object.M_supportedInterfaceRef

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze InputOutputBinding **/
	if len(object.M_ioBinding) > 0 {
		xmlElement.M_ioBinding = make([]*BPMN20.XsdInputOutputBinding, 0)
	}

	/** Now I will save the value of ioBinding **/
	for i := 0; i < len(object.M_ioBinding); i++ {
		xmlElement.M_ioBinding = append(xmlElement.M_ioBinding, new(BPMN20.XsdInputOutputBinding))
		this.SerialyzeInputOutputBinding(xmlElement.M_ioBinding[i], object.M_ioBinding[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze LaneSet **/
	if len(object.M_laneSet) > 0 {
		xmlElement.M_laneSet = make([]*BPMN20.XsdLaneSet, 0)
	}

	/** Now I will save the value of laneSet **/
	for i := 0; i < len(object.M_laneSet); i++ {
		xmlElement.M_laneSet = append(xmlElement.M_laneSet, new(BPMN20.XsdLaneSet))
		this.SerialyzeLaneSet(xmlElement.M_laneSet[i], object.M_laneSet[i])
	}

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze Process:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze Process:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze Process:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze Process:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze Process:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze Process:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze Process:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze Process:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze Process:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze Process:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze Process:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze Process:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze Process:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze Process:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze Process:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze Process:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze Process:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze Process:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze Process:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze Process:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze Process:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze Process:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze Process:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze Process:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze Process:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze Process:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze Process:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze Process:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze Process:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze Process:flowElement:UserTask")
		}
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze Process:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze Process:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze Process:artifact:TextAnnotation")
		}
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Process:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Process:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Process:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Process:resourceRole:ResourceRole")
		}
	}

	/** Serialyze CorrelationSubscription **/
	if len(object.M_correlationSubscription) > 0 {
		xmlElement.M_correlationSubscription = make([]*BPMN20.XsdCorrelationSubscription, 0)
	}

	/** Now I will save the value of correlationSubscription **/
	for i := 0; i < len(object.M_correlationSubscription); i++ {
		xmlElement.M_correlationSubscription = append(xmlElement.M_correlationSubscription, new(BPMN20.XsdCorrelationSubscription))
		this.SerialyzeCorrelationSubscription(xmlElement.M_correlationSubscription[i], object.M_correlationSubscription[i])
	}

	/** Serialyze ref supports **/
	xmlElement.M_supports = object.M_supports

	/** ProcessType **/
	if object.M_processType == BPMN20.ProcessType_None {
		xmlElement.M_processType = "##None"
	} else if object.M_processType == BPMN20.ProcessType_Public {
		xmlElement.M_processType = "##Public"
	} else if object.M_processType == BPMN20.ProcessType_Private {
		xmlElement.M_processType = "##Private"
	} else {
		xmlElement.M_processType = "##None"
	}

	/** CallableElement **/
	xmlElement.M_isClosed = object.M_isClosed

	/** CallableElement **/
	xmlElement.M_isExecutable = object.M_isExecutable

	/** Serialyze ref definitionalCollaborationRef **/
	xmlElement.M_definitionalCollaborationRef = object.M_definitionalCollaborationRef
}

/** serialysation of GlobalChoreographyTask **/
func (this *BPMSXmlFactory) SerialyzeGlobalChoreographyTask(xmlElement *BPMN20.XsdGlobalChoreographyTask, object *BPMN20.GlobalChoreographyTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalChoreographyTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Participant **/
	if len(object.M_participant) > 0 {
		xmlElement.M_participant = make([]*BPMN20.XsdParticipant, 0)
	}

	/** Now I will save the value of participant **/
	for i := 0; i < len(object.M_participant); i++ {
		xmlElement.M_participant = append(xmlElement.M_participant, new(BPMN20.XsdParticipant))
		this.SerialyzeParticipant(xmlElement.M_participant[i], object.M_participant[i])
	}

	/** Serialyze MessageFlow **/
	if len(object.M_messageFlow) > 0 {
		xmlElement.M_messageFlow = make([]*BPMN20.XsdMessageFlow, 0)
	}

	/** Now I will save the value of messageFlow **/
	for i := 0; i < len(object.M_messageFlow); i++ {
		xmlElement.M_messageFlow = append(xmlElement.M_messageFlow, new(BPMN20.XsdMessageFlow))
		this.SerialyzeMessageFlow(xmlElement.M_messageFlow[i], object.M_messageFlow[i])
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze Collaboration:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze Collaboration:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze Collaboration:artifact:TextAnnotation")
		}
	}

	/** Serialyze ConversationNode **/
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_0 = make([]*BPMN20.XsdCallConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_1 = make([]*BPMN20.XsdConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_2 = make([]*BPMN20.XsdSubConversation, 0)
	}

	/** Now I will save the value of conversationNode **/
	for i := 0; i < len(object.M_conversationNode); i++ {
		switch v := object.M_conversationNode[i].(type) {
		case *BPMN20.CallConversation:
			xmlElement.M_conversationNode_0 = append(xmlElement.M_conversationNode_0, new(BPMN20.XsdCallConversation))
			this.SerialyzeCallConversation(xmlElement.M_conversationNode_0[len(xmlElement.M_conversationNode_0)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:CallConversation")
		case *BPMN20.Conversation:
			xmlElement.M_conversationNode_1 = append(xmlElement.M_conversationNode_1, new(BPMN20.XsdConversation))
			this.SerialyzeConversation(xmlElement.M_conversationNode_1[len(xmlElement.M_conversationNode_1)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:Conversation")
		case *BPMN20.SubConversation:
			xmlElement.M_conversationNode_2 = append(xmlElement.M_conversationNode_2, new(BPMN20.XsdSubConversation))
			this.SerialyzeSubConversation(xmlElement.M_conversationNode_2[len(xmlElement.M_conversationNode_2)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:SubConversation")
		}
	}

	/** Serialyze ConversationAssociation **/
	if len(object.M_conversationAssociation) > 0 {
		xmlElement.M_conversationAssociation = make([]*BPMN20.XsdConversationAssociation, 0)
	}

	/** Now I will save the value of conversationAssociation **/
	for i := 0; i < len(object.M_conversationAssociation); i++ {
		xmlElement.M_conversationAssociation = append(xmlElement.M_conversationAssociation, new(BPMN20.XsdConversationAssociation))
		this.SerialyzeConversationAssociation(xmlElement.M_conversationAssociation[i], object.M_conversationAssociation[i])
	}

	/** Serialyze ParticipantAssociation **/
	if len(object.M_participantAssociation) > 0 {
		xmlElement.M_participantAssociation = make([]*BPMN20.XsdParticipantAssociation, 0)
	}

	/** Now I will save the value of participantAssociation **/
	for i := 0; i < len(object.M_participantAssociation); i++ {
		xmlElement.M_participantAssociation = append(xmlElement.M_participantAssociation, new(BPMN20.XsdParticipantAssociation))
		this.SerialyzeParticipantAssociation(xmlElement.M_participantAssociation[i], object.M_participantAssociation[i])
	}

	/** Serialyze MessageFlowAssociation **/
	if len(object.M_messageFlowAssociation) > 0 {
		xmlElement.M_messageFlowAssociation = make([]*BPMN20.XsdMessageFlowAssociation, 0)
	}

	/** Now I will save the value of messageFlowAssociation **/
	for i := 0; i < len(object.M_messageFlowAssociation); i++ {
		xmlElement.M_messageFlowAssociation = append(xmlElement.M_messageFlowAssociation, new(BPMN20.XsdMessageFlowAssociation))
		this.SerialyzeMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], object.M_messageFlowAssociation[i])
	}

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref choreographyRef **/
	xmlElement.M_choreographyRef = object.M_choreographyRef

	/** Serialyze ConversationLink **/
	if len(object.M_conversationLink) > 0 {
		xmlElement.M_conversationLink = make([]*BPMN20.XsdConversationLink, 0)
	}

	/** Now I will save the value of conversationLink **/
	for i := 0; i < len(object.M_conversationLink); i++ {
		xmlElement.M_conversationLink = append(xmlElement.M_conversationLink, new(BPMN20.XsdConversationLink))
		this.SerialyzeConversationLink(xmlElement.M_conversationLink[i], object.M_conversationLink[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_isClosed = object.M_isClosed

	/** Serialyze FlowElement **/
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_0 = make([]*BPMN20.XsdAdHocSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_1 = make([]*BPMN20.XsdBoundaryEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_2 = make([]*BPMN20.XsdBusinessRuleTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_3 = make([]*BPMN20.XsdCallActivity, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_4 = make([]*BPMN20.XsdCallChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_5 = make([]*BPMN20.XsdChoreographyTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_6 = make([]*BPMN20.XsdComplexGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_7 = make([]*BPMN20.XsdDataObject, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_8 = make([]*BPMN20.XsdDataObjectReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_9 = make([]*BPMN20.XsdDataStoreReference, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_10 = make([]*BPMN20.XsdEndEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_11 = make([]*BPMN20.XsdEventBasedGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_12 = make([]*BPMN20.XsdExclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_13 = make([]*BPMN20.XsdImplicitThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_14 = make([]*BPMN20.XsdInclusiveGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_15 = make([]*BPMN20.XsdIntermediateCatchEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_16 = make([]*BPMN20.XsdIntermediateThrowEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_17 = make([]*BPMN20.XsdManualTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_18 = make([]*BPMN20.XsdParallelGateway, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_19 = make([]*BPMN20.XsdReceiveTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_20 = make([]*BPMN20.XsdScriptTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_21 = make([]*BPMN20.XsdSendTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_22 = make([]*BPMN20.XsdSequenceFlow, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_23 = make([]*BPMN20.XsdServiceTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_24 = make([]*BPMN20.XsdStartEvent, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_25 = make([]*BPMN20.XsdSubChoreography, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_26 = make([]*BPMN20.XsdSubProcess, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_27 = make([]*BPMN20.XsdTask, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_28 = make([]*BPMN20.XsdTransaction, 0)
	}
	if len(object.M_flowElement) > 0 {
		xmlElement.M_flowElement_29 = make([]*BPMN20.XsdUserTask, 0)
	}

	/** Now I will save the value of flowElement **/
	for i := 0; i < len(object.M_flowElement); i++ {
		switch v := object.M_flowElement[i].(type) {
		case *BPMN20.AdHocSubProcess:
			xmlElement.M_flowElement_0 = append(xmlElement.M_flowElement_0, new(BPMN20.XsdAdHocSubProcess))
			this.SerialyzeAdHocSubProcess(xmlElement.M_flowElement_0[len(xmlElement.M_flowElement_0)-1], v)
			log.Println("Serialyze Choreography:flowElement:AdHocSubProcess")
		case *BPMN20.BoundaryEvent:
			xmlElement.M_flowElement_1 = append(xmlElement.M_flowElement_1, new(BPMN20.XsdBoundaryEvent))
			this.SerialyzeBoundaryEvent(xmlElement.M_flowElement_1[len(xmlElement.M_flowElement_1)-1], v)
			log.Println("Serialyze Choreography:flowElement:BoundaryEvent")
		case *BPMN20.BusinessRuleTask:
			xmlElement.M_flowElement_2 = append(xmlElement.M_flowElement_2, new(BPMN20.XsdBusinessRuleTask))
			this.SerialyzeBusinessRuleTask(xmlElement.M_flowElement_2[len(xmlElement.M_flowElement_2)-1], v)
			log.Println("Serialyze Choreography:flowElement:BusinessRuleTask")
		case *BPMN20.CallActivity:
			xmlElement.M_flowElement_3 = append(xmlElement.M_flowElement_3, new(BPMN20.XsdCallActivity))
			this.SerialyzeCallActivity(xmlElement.M_flowElement_3[len(xmlElement.M_flowElement_3)-1], v)
			log.Println("Serialyze Choreography:flowElement:CallActivity")
		case *BPMN20.CallChoreography:
			xmlElement.M_flowElement_4 = append(xmlElement.M_flowElement_4, new(BPMN20.XsdCallChoreography))
			this.SerialyzeCallChoreography(xmlElement.M_flowElement_4[len(xmlElement.M_flowElement_4)-1], v)
			log.Println("Serialyze Choreography:flowElement:CallChoreography")
		case *BPMN20.ChoreographyTask:
			xmlElement.M_flowElement_5 = append(xmlElement.M_flowElement_5, new(BPMN20.XsdChoreographyTask))
			this.SerialyzeChoreographyTask(xmlElement.M_flowElement_5[len(xmlElement.M_flowElement_5)-1], v)
			log.Println("Serialyze Choreography:flowElement:ChoreographyTask")
		case *BPMN20.ComplexGateway:
			xmlElement.M_flowElement_6 = append(xmlElement.M_flowElement_6, new(BPMN20.XsdComplexGateway))
			this.SerialyzeComplexGateway(xmlElement.M_flowElement_6[len(xmlElement.M_flowElement_6)-1], v)
			log.Println("Serialyze Choreography:flowElement:ComplexGateway")
		case *BPMN20.DataObject:
			xmlElement.M_flowElement_7 = append(xmlElement.M_flowElement_7, new(BPMN20.XsdDataObject))
			this.SerialyzeDataObject(xmlElement.M_flowElement_7[len(xmlElement.M_flowElement_7)-1], v)
			log.Println("Serialyze Choreography:flowElement:DataObject")
		case *BPMN20.DataObjectReference:
			xmlElement.M_flowElement_8 = append(xmlElement.M_flowElement_8, new(BPMN20.XsdDataObjectReference))
			this.SerialyzeDataObjectReference(xmlElement.M_flowElement_8[len(xmlElement.M_flowElement_8)-1], v)
			log.Println("Serialyze Choreography:flowElement:DataObjectReference")
		case *BPMN20.DataStoreReference:
			xmlElement.M_flowElement_9 = append(xmlElement.M_flowElement_9, new(BPMN20.XsdDataStoreReference))
			this.SerialyzeDataStoreReference(xmlElement.M_flowElement_9[len(xmlElement.M_flowElement_9)-1], v)
			log.Println("Serialyze Choreography:flowElement:DataStoreReference")
		case *BPMN20.EndEvent:
			xmlElement.M_flowElement_10 = append(xmlElement.M_flowElement_10, new(BPMN20.XsdEndEvent))
			this.SerialyzeEndEvent(xmlElement.M_flowElement_10[len(xmlElement.M_flowElement_10)-1], v)
			log.Println("Serialyze Choreography:flowElement:EndEvent")
		case *BPMN20.EventBasedGateway:
			xmlElement.M_flowElement_11 = append(xmlElement.M_flowElement_11, new(BPMN20.XsdEventBasedGateway))
			this.SerialyzeEventBasedGateway(xmlElement.M_flowElement_11[len(xmlElement.M_flowElement_11)-1], v)
			log.Println("Serialyze Choreography:flowElement:EventBasedGateway")
		case *BPMN20.ExclusiveGateway:
			xmlElement.M_flowElement_12 = append(xmlElement.M_flowElement_12, new(BPMN20.XsdExclusiveGateway))
			this.SerialyzeExclusiveGateway(xmlElement.M_flowElement_12[len(xmlElement.M_flowElement_12)-1], v)
			log.Println("Serialyze Choreography:flowElement:ExclusiveGateway")
		case *BPMN20.ImplicitThrowEvent:
			xmlElement.M_flowElement_13 = append(xmlElement.M_flowElement_13, new(BPMN20.XsdImplicitThrowEvent))
			this.SerialyzeImplicitThrowEvent(xmlElement.M_flowElement_13[len(xmlElement.M_flowElement_13)-1], v)
			log.Println("Serialyze Choreography:flowElement:ImplicitThrowEvent")
		case *BPMN20.InclusiveGateway:
			xmlElement.M_flowElement_14 = append(xmlElement.M_flowElement_14, new(BPMN20.XsdInclusiveGateway))
			this.SerialyzeInclusiveGateway(xmlElement.M_flowElement_14[len(xmlElement.M_flowElement_14)-1], v)
			log.Println("Serialyze Choreography:flowElement:InclusiveGateway")
		case *BPMN20.IntermediateCatchEvent:
			xmlElement.M_flowElement_15 = append(xmlElement.M_flowElement_15, new(BPMN20.XsdIntermediateCatchEvent))
			this.SerialyzeIntermediateCatchEvent(xmlElement.M_flowElement_15[len(xmlElement.M_flowElement_15)-1], v)
			log.Println("Serialyze Choreography:flowElement:IntermediateCatchEvent")
		case *BPMN20.IntermediateThrowEvent:
			xmlElement.M_flowElement_16 = append(xmlElement.M_flowElement_16, new(BPMN20.XsdIntermediateThrowEvent))
			this.SerialyzeIntermediateThrowEvent(xmlElement.M_flowElement_16[len(xmlElement.M_flowElement_16)-1], v)
			log.Println("Serialyze Choreography:flowElement:IntermediateThrowEvent")
		case *BPMN20.ManualTask:
			xmlElement.M_flowElement_17 = append(xmlElement.M_flowElement_17, new(BPMN20.XsdManualTask))
			this.SerialyzeManualTask(xmlElement.M_flowElement_17[len(xmlElement.M_flowElement_17)-1], v)
			log.Println("Serialyze Choreography:flowElement:ManualTask")
		case *BPMN20.ParallelGateway:
			xmlElement.M_flowElement_18 = append(xmlElement.M_flowElement_18, new(BPMN20.XsdParallelGateway))
			this.SerialyzeParallelGateway(xmlElement.M_flowElement_18[len(xmlElement.M_flowElement_18)-1], v)
			log.Println("Serialyze Choreography:flowElement:ParallelGateway")
		case *BPMN20.ReceiveTask:
			xmlElement.M_flowElement_19 = append(xmlElement.M_flowElement_19, new(BPMN20.XsdReceiveTask))
			this.SerialyzeReceiveTask(xmlElement.M_flowElement_19[len(xmlElement.M_flowElement_19)-1], v)
			log.Println("Serialyze Choreography:flowElement:ReceiveTask")
		case *BPMN20.ScriptTask:
			xmlElement.M_flowElement_20 = append(xmlElement.M_flowElement_20, new(BPMN20.XsdScriptTask))
			this.SerialyzeScriptTask(xmlElement.M_flowElement_20[len(xmlElement.M_flowElement_20)-1], v)
			log.Println("Serialyze Choreography:flowElement:ScriptTask")
		case *BPMN20.SendTask:
			xmlElement.M_flowElement_21 = append(xmlElement.M_flowElement_21, new(BPMN20.XsdSendTask))
			this.SerialyzeSendTask(xmlElement.M_flowElement_21[len(xmlElement.M_flowElement_21)-1], v)
			log.Println("Serialyze Choreography:flowElement:SendTask")
		case *BPMN20.SequenceFlow:
			xmlElement.M_flowElement_22 = append(xmlElement.M_flowElement_22, new(BPMN20.XsdSequenceFlow))
			this.SerialyzeSequenceFlow(xmlElement.M_flowElement_22[len(xmlElement.M_flowElement_22)-1], v)
			log.Println("Serialyze Choreography:flowElement:SequenceFlow")
		case *BPMN20.ServiceTask:
			xmlElement.M_flowElement_23 = append(xmlElement.M_flowElement_23, new(BPMN20.XsdServiceTask))
			this.SerialyzeServiceTask(xmlElement.M_flowElement_23[len(xmlElement.M_flowElement_23)-1], v)
			log.Println("Serialyze Choreography:flowElement:ServiceTask")
		case *BPMN20.StartEvent:
			xmlElement.M_flowElement_24 = append(xmlElement.M_flowElement_24, new(BPMN20.XsdStartEvent))
			this.SerialyzeStartEvent(xmlElement.M_flowElement_24[len(xmlElement.M_flowElement_24)-1], v)
			log.Println("Serialyze Choreography:flowElement:StartEvent")
		case *BPMN20.SubChoreography:
			xmlElement.M_flowElement_25 = append(xmlElement.M_flowElement_25, new(BPMN20.XsdSubChoreography))
			this.SerialyzeSubChoreography(xmlElement.M_flowElement_25[len(xmlElement.M_flowElement_25)-1], v)
			log.Println("Serialyze Choreography:flowElement:SubChoreography")
		case *BPMN20.SubProcess_impl:
			xmlElement.M_flowElement_26 = append(xmlElement.M_flowElement_26, new(BPMN20.XsdSubProcess))
			this.SerialyzeSubProcess(xmlElement.M_flowElement_26[len(xmlElement.M_flowElement_26)-1], v)
			log.Println("Serialyze Choreography:flowElement:SubProcess")
		case *BPMN20.Task_impl:
			xmlElement.M_flowElement_27 = append(xmlElement.M_flowElement_27, new(BPMN20.XsdTask))
			this.SerialyzeTask(xmlElement.M_flowElement_27[len(xmlElement.M_flowElement_27)-1], v)
			log.Println("Serialyze Choreography:flowElement:Task")
		case *BPMN20.Transaction:
			xmlElement.M_flowElement_28 = append(xmlElement.M_flowElement_28, new(BPMN20.XsdTransaction))
			this.SerialyzeTransaction(xmlElement.M_flowElement_28[len(xmlElement.M_flowElement_28)-1], v)
			log.Println("Serialyze Choreography:flowElement:Transaction")
		case *BPMN20.UserTask:
			xmlElement.M_flowElement_29 = append(xmlElement.M_flowElement_29, new(BPMN20.XsdUserTask))
			this.SerialyzeUserTask(xmlElement.M_flowElement_29[len(xmlElement.M_flowElement_29)-1], v)
			log.Println("Serialyze Choreography:flowElement:UserTask")
		}
	}

	/** Serialyze ref initiatingParticipantRef **/
	xmlElement.M_initiatingParticipantRef = object.M_initiatingParticipantRef
}

/** serialysation of IntermediateCatchEvent **/
func (this *BPMSXmlFactory) SerialyzeIntermediateCatchEvent(xmlElement *BPMN20.XsdIntermediateCatchEvent, object *BPMN20.IntermediateCatchEvent) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** IntermediateCatchEvent **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataOutput **/
	if len(object.M_dataOutput) > 0 {
		xmlElement.M_dataOutput = make([]*BPMN20.XsdDataOutput, 0)
	}

	/** Now I will save the value of dataOutput **/
	for i := 0; i < len(object.M_dataOutput); i++ {
		xmlElement.M_dataOutput = append(xmlElement.M_dataOutput, new(BPMN20.XsdDataOutput))
		this.SerialyzeDataOutput(xmlElement.M_dataOutput[i], object.M_dataOutput[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze OutputSet **/
	if object.M_outputSet != nil {
		xmlElement.M_outputSet = new(BPMN20.XsdOutputSet)
	}

	/** Now I will save the value of outputSet **/
	if object.M_outputSet != nil {
		this.SerialyzeOutputSet(xmlElement.M_outputSet, object.M_outputSet)
	}

	/** Serialyze EventDefinition **/
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_0 = make([]*BPMN20.XsdCancelEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_1 = make([]*BPMN20.XsdCompensateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_2 = make([]*BPMN20.XsdConditionalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_3 = make([]*BPMN20.XsdErrorEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_4 = make([]*BPMN20.XsdEscalationEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_5 = make([]*BPMN20.XsdLinkEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_6 = make([]*BPMN20.XsdMessageEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_7 = make([]*BPMN20.XsdSignalEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_8 = make([]*BPMN20.XsdTerminateEventDefinition, 0)
	}
	if len(object.M_eventDefinition) > 0 {
		xmlElement.M_eventDefinition_9 = make([]*BPMN20.XsdTimerEventDefinition, 0)
	}

	/** Now I will save the value of eventDefinition **/
	for i := 0; i < len(object.M_eventDefinition); i++ {
		switch v := object.M_eventDefinition[i].(type) {
		case *BPMN20.CancelEventDefinition:
			xmlElement.M_eventDefinition_0 = append(xmlElement.M_eventDefinition_0, new(BPMN20.XsdCancelEventDefinition))
			this.SerialyzeCancelEventDefinition(xmlElement.M_eventDefinition_0[len(xmlElement.M_eventDefinition_0)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:CancelEventDefinition")
		case *BPMN20.CompensateEventDefinition:
			xmlElement.M_eventDefinition_1 = append(xmlElement.M_eventDefinition_1, new(BPMN20.XsdCompensateEventDefinition))
			this.SerialyzeCompensateEventDefinition(xmlElement.M_eventDefinition_1[len(xmlElement.M_eventDefinition_1)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:CompensateEventDefinition")
		case *BPMN20.ConditionalEventDefinition:
			xmlElement.M_eventDefinition_2 = append(xmlElement.M_eventDefinition_2, new(BPMN20.XsdConditionalEventDefinition))
			this.SerialyzeConditionalEventDefinition(xmlElement.M_eventDefinition_2[len(xmlElement.M_eventDefinition_2)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:ConditionalEventDefinition")
		case *BPMN20.ErrorEventDefinition:
			xmlElement.M_eventDefinition_3 = append(xmlElement.M_eventDefinition_3, new(BPMN20.XsdErrorEventDefinition))
			this.SerialyzeErrorEventDefinition(xmlElement.M_eventDefinition_3[len(xmlElement.M_eventDefinition_3)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:ErrorEventDefinition")
		case *BPMN20.EscalationEventDefinition:
			xmlElement.M_eventDefinition_4 = append(xmlElement.M_eventDefinition_4, new(BPMN20.XsdEscalationEventDefinition))
			this.SerialyzeEscalationEventDefinition(xmlElement.M_eventDefinition_4[len(xmlElement.M_eventDefinition_4)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:EscalationEventDefinition")
		case *BPMN20.LinkEventDefinition:
			xmlElement.M_eventDefinition_5 = append(xmlElement.M_eventDefinition_5, new(BPMN20.XsdLinkEventDefinition))
			this.SerialyzeLinkEventDefinition(xmlElement.M_eventDefinition_5[len(xmlElement.M_eventDefinition_5)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:LinkEventDefinition")
		case *BPMN20.MessageEventDefinition:
			xmlElement.M_eventDefinition_6 = append(xmlElement.M_eventDefinition_6, new(BPMN20.XsdMessageEventDefinition))
			this.SerialyzeMessageEventDefinition(xmlElement.M_eventDefinition_6[len(xmlElement.M_eventDefinition_6)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:MessageEventDefinition")
		case *BPMN20.SignalEventDefinition:
			xmlElement.M_eventDefinition_7 = append(xmlElement.M_eventDefinition_7, new(BPMN20.XsdSignalEventDefinition))
			this.SerialyzeSignalEventDefinition(xmlElement.M_eventDefinition_7[len(xmlElement.M_eventDefinition_7)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:SignalEventDefinition")
		case *BPMN20.TerminateEventDefinition:
			xmlElement.M_eventDefinition_8 = append(xmlElement.M_eventDefinition_8, new(BPMN20.XsdTerminateEventDefinition))
			this.SerialyzeTerminateEventDefinition(xmlElement.M_eventDefinition_8[len(xmlElement.M_eventDefinition_8)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:TerminateEventDefinition")
		case *BPMN20.TimerEventDefinition:
			xmlElement.M_eventDefinition_9 = append(xmlElement.M_eventDefinition_9, new(BPMN20.XsdTimerEventDefinition))
			this.SerialyzeTimerEventDefinition(xmlElement.M_eventDefinition_9[len(xmlElement.M_eventDefinition_9)-1], v)
			log.Println("Serialyze CatchEvent:eventDefinition:TimerEventDefinition")
		}
	}

	/** Serialyze ref eventDefinitionRef **/
	xmlElement.M_eventDefinitionRef = object.M_eventDefinitionRef

	/** Event **/
	xmlElement.M_parallelMultiple = object.M_parallelMultiple
}

/** serialysation of Association **/
func (this *BPMSXmlFactory) SerialyzeAssociation(xmlElement *BPMN20.XsdAssociation, object *BPMN20.Association) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** Association **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = object.M_targetRef

	/** AssociationDirection **/
	if object.M_associationDirection == BPMN20.AssociationDirection_None {
		xmlElement.M_associationDirection = "##None"
	} else if object.M_associationDirection == BPMN20.AssociationDirection_One {
		xmlElement.M_associationDirection = "##One"
	} else if object.M_associationDirection == BPMN20.AssociationDirection_Both {
		xmlElement.M_associationDirection = "##Both"
	} else {
		xmlElement.M_associationDirection = "##None"
	}
}

/** serialysation of ConversationAssociation **/
func (this *BPMSXmlFactory) SerialyzeConversationAssociation(xmlElement *BPMN20.XsdConversationAssociation, object *BPMN20.ConversationAssociation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ConversationAssociation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref innerConversationNodeRef **/
	xmlElement.M_innerConversationNodeRef = object.M_innerConversationNodeRef

	/** Serialyze ref outerConversationNodeRef **/
	xmlElement.M_outerConversationNodeRef = object.M_outerConversationNodeRef
}

/** serialysation of LaneSet **/
func (this *BPMSXmlFactory) SerialyzeLaneSet(xmlElement *BPMN20.XsdLaneSet, object *BPMN20.LaneSet) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** LaneSet **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Lane **/
	if len(object.M_lane) > 0 {
		xmlElement.M_lane = make([]*BPMN20.XsdLane, 0)
	}

	/** Now I will save the value of lane **/
	for i := 0; i < len(object.M_lane); i++ {
		xmlElement.M_lane = append(xmlElement.M_lane, new(BPMN20.XsdLane))
		this.SerialyzeLane(xmlElement.M_lane[i], object.M_lane[i])
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name
}

/** serialysation of ResourceParameterBinding **/
func (this *BPMSXmlFactory) SerialyzeResourceParameterBinding(xmlElement *BPMN20.XsdResourceParameterBinding, object *BPMN20.ResourceParameterBinding) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ResourceParameterBinding **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Expression **/

	/** Now I will save the value of expression **/
	if object.M_expression != nil {
		switch v := object.M_expression.(type) {
		case *BPMN20.FormalExpression:
			xmlElement.M_expression_0 = new(BPMN20.XsdFormalExpression)
			this.SerialyzeFormalExpression(xmlElement.M_expression_0, v)
			log.Println("Serialyze ResourceParameterBinding:expression:FormalExpression")
		case *BPMN20.Expression_impl:
			xmlElement.M_expression_1 = new(BPMN20.XsdExpression)
			this.SerialyzeExpression(xmlElement.M_expression_1, v)
			log.Println("Serialyze ResourceParameterBinding:expression:Expression")
		}
	}

	/** Serialyze ref parameterRef **/
	xmlElement.M_parameterRef = object.M_parameterRef
}

/** serialysation of GlobalConversation **/
func (this *BPMSXmlFactory) SerialyzeGlobalConversation(xmlElement *BPMN20.XsdGlobalConversation, object *BPMN20.GlobalConversation) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** GlobalConversation **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Participant **/
	if len(object.M_participant) > 0 {
		xmlElement.M_participant = make([]*BPMN20.XsdParticipant, 0)
	}

	/** Now I will save the value of participant **/
	for i := 0; i < len(object.M_participant); i++ {
		xmlElement.M_participant = append(xmlElement.M_participant, new(BPMN20.XsdParticipant))
		this.SerialyzeParticipant(xmlElement.M_participant[i], object.M_participant[i])
	}

	/** Serialyze MessageFlow **/
	if len(object.M_messageFlow) > 0 {
		xmlElement.M_messageFlow = make([]*BPMN20.XsdMessageFlow, 0)
	}

	/** Now I will save the value of messageFlow **/
	for i := 0; i < len(object.M_messageFlow); i++ {
		xmlElement.M_messageFlow = append(xmlElement.M_messageFlow, new(BPMN20.XsdMessageFlow))
		this.SerialyzeMessageFlow(xmlElement.M_messageFlow[i], object.M_messageFlow[i])
	}

	/** Serialyze Artifact **/
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_0 = make([]*BPMN20.XsdAssociation, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_1 = make([]*BPMN20.XsdGroup, 0)
	}
	if len(object.M_artifact) > 0 {
		xmlElement.M_artifact_2 = make([]*BPMN20.XsdTextAnnotation, 0)
	}

	/** Now I will save the value of artifact **/
	for i := 0; i < len(object.M_artifact); i++ {
		switch v := object.M_artifact[i].(type) {
		case *BPMN20.Association:
			xmlElement.M_artifact_0 = append(xmlElement.M_artifact_0, new(BPMN20.XsdAssociation))
			this.SerialyzeAssociation(xmlElement.M_artifact_0[len(xmlElement.M_artifact_0)-1], v)
			log.Println("Serialyze Collaboration:artifact:Association")
		case *BPMN20.Group:
			xmlElement.M_artifact_1 = append(xmlElement.M_artifact_1, new(BPMN20.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_artifact_1[len(xmlElement.M_artifact_1)-1], v)
			log.Println("Serialyze Collaboration:artifact:Group")
		case *BPMN20.TextAnnotation:
			xmlElement.M_artifact_2 = append(xmlElement.M_artifact_2, new(BPMN20.XsdTextAnnotation))
			this.SerialyzeTextAnnotation(xmlElement.M_artifact_2[len(xmlElement.M_artifact_2)-1], v)
			log.Println("Serialyze Collaboration:artifact:TextAnnotation")
		}
	}

	/** Serialyze ConversationNode **/
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_0 = make([]*BPMN20.XsdCallConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_1 = make([]*BPMN20.XsdConversation, 0)
	}
	if len(object.M_conversationNode) > 0 {
		xmlElement.M_conversationNode_2 = make([]*BPMN20.XsdSubConversation, 0)
	}

	/** Now I will save the value of conversationNode **/
	for i := 0; i < len(object.M_conversationNode); i++ {
		switch v := object.M_conversationNode[i].(type) {
		case *BPMN20.CallConversation:
			xmlElement.M_conversationNode_0 = append(xmlElement.M_conversationNode_0, new(BPMN20.XsdCallConversation))
			this.SerialyzeCallConversation(xmlElement.M_conversationNode_0[len(xmlElement.M_conversationNode_0)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:CallConversation")
		case *BPMN20.Conversation:
			xmlElement.M_conversationNode_1 = append(xmlElement.M_conversationNode_1, new(BPMN20.XsdConversation))
			this.SerialyzeConversation(xmlElement.M_conversationNode_1[len(xmlElement.M_conversationNode_1)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:Conversation")
		case *BPMN20.SubConversation:
			xmlElement.M_conversationNode_2 = append(xmlElement.M_conversationNode_2, new(BPMN20.XsdSubConversation))
			this.SerialyzeSubConversation(xmlElement.M_conversationNode_2[len(xmlElement.M_conversationNode_2)-1], v)
			log.Println("Serialyze Collaboration:conversationNode:SubConversation")
		}
	}

	/** Serialyze ConversationAssociation **/
	if len(object.M_conversationAssociation) > 0 {
		xmlElement.M_conversationAssociation = make([]*BPMN20.XsdConversationAssociation, 0)
	}

	/** Now I will save the value of conversationAssociation **/
	for i := 0; i < len(object.M_conversationAssociation); i++ {
		xmlElement.M_conversationAssociation = append(xmlElement.M_conversationAssociation, new(BPMN20.XsdConversationAssociation))
		this.SerialyzeConversationAssociation(xmlElement.M_conversationAssociation[i], object.M_conversationAssociation[i])
	}

	/** Serialyze ParticipantAssociation **/
	if len(object.M_participantAssociation) > 0 {
		xmlElement.M_participantAssociation = make([]*BPMN20.XsdParticipantAssociation, 0)
	}

	/** Now I will save the value of participantAssociation **/
	for i := 0; i < len(object.M_participantAssociation); i++ {
		xmlElement.M_participantAssociation = append(xmlElement.M_participantAssociation, new(BPMN20.XsdParticipantAssociation))
		this.SerialyzeParticipantAssociation(xmlElement.M_participantAssociation[i], object.M_participantAssociation[i])
	}

	/** Serialyze MessageFlowAssociation **/
	if len(object.M_messageFlowAssociation) > 0 {
		xmlElement.M_messageFlowAssociation = make([]*BPMN20.XsdMessageFlowAssociation, 0)
	}

	/** Now I will save the value of messageFlowAssociation **/
	for i := 0; i < len(object.M_messageFlowAssociation); i++ {
		xmlElement.M_messageFlowAssociation = append(xmlElement.M_messageFlowAssociation, new(BPMN20.XsdMessageFlowAssociation))
		this.SerialyzeMessageFlowAssociation(xmlElement.M_messageFlowAssociation[i], object.M_messageFlowAssociation[i])
	}

	/** Serialyze CorrelationKey **/
	if len(object.M_correlationKey) > 0 {
		xmlElement.M_correlationKey = make([]*BPMN20.XsdCorrelationKey, 0)
	}

	/** Now I will save the value of correlationKey **/
	for i := 0; i < len(object.M_correlationKey); i++ {
		xmlElement.M_correlationKey = append(xmlElement.M_correlationKey, new(BPMN20.XsdCorrelationKey))
		this.SerialyzeCorrelationKey(xmlElement.M_correlationKey[i], object.M_correlationKey[i])
	}

	/** Serialyze ref choreographyRef **/
	xmlElement.M_choreographyRef = object.M_choreographyRef

	/** Serialyze ConversationLink **/
	if len(object.M_conversationLink) > 0 {
		xmlElement.M_conversationLink = make([]*BPMN20.XsdConversationLink, 0)
	}

	/** Now I will save the value of conversationLink **/
	for i := 0; i < len(object.M_conversationLink); i++ {
		xmlElement.M_conversationLink = append(xmlElement.M_conversationLink, new(BPMN20.XsdConversationLink))
		this.SerialyzeConversationLink(xmlElement.M_conversationLink[i], object.M_conversationLink[i])
	}

	/** RootElement **/
	xmlElement.M_name = object.M_name

	/** RootElement **/
	xmlElement.M_isClosed = object.M_isClosed
}

/** serialysation of HumanPerformer **/
func (this *BPMSXmlFactory) SerialyzeHumanPerformer(xmlElement *BPMN20.XsdHumanPerformer, object *BPMN20.HumanPerformer_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** HumanPerformer **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref resourceRef **/
	xmlElement.M_resourceRef = &object.M_resourceRef

	/** Serialyze ResourceParameterBinding **/
	if len(object.M_resourceParameterBinding) > 0 {
		xmlElement.M_resourceParameterBinding = make([]*BPMN20.XsdResourceParameterBinding, 0)
	}

	/** Now I will save the value of resourceParameterBinding **/
	for i := 0; i < len(object.M_resourceParameterBinding); i++ {
		xmlElement.M_resourceParameterBinding = append(xmlElement.M_resourceParameterBinding, new(BPMN20.XsdResourceParameterBinding))
		this.SerialyzeResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], object.M_resourceParameterBinding[i])
	}

	/** Serialyze ResourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		xmlElement.M_resourceAssignmentExpression = new(BPMN20.XsdResourceAssignmentExpression)
	}

	/** Now I will save the value of resourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		this.SerialyzeResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)
	}
}

/** serialysation of ComplexBehaviorDefinition **/
func (this *BPMSXmlFactory) SerialyzeComplexBehaviorDefinition(xmlElement *BPMN20.XsdComplexBehaviorDefinition, object *BPMN20.ComplexBehaviorDefinition) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ComplexBehaviorDefinition **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FormalExpression **/

	/** Now I will save the value of condition **/
	if object.M_condition != nil {
		this.SerialyzeFormalExpression((*BPMN20.XsdFormalExpression)(unsafe.Pointer(&xmlElement.M_condition)), object.M_condition)
	}

	/** Serialyze ImplicitThrowEvent **/
	if object.M_event != nil {
		xmlElement.M_event = new(BPMN20.XsdEvent)
	}

	/** Now I will save the value of event **/
	if object.M_event != nil {
		this.SerialyzeImplicitThrowEvent((*BPMN20.XsdImplicitThrowEvent)(unsafe.Pointer(xmlElement.M_event)), object.M_event)
	}
}

/** serialysation of ServiceTask **/
func (this *BPMSXmlFactory) SerialyzeServiceTask(xmlElement *BPMN20.XsdServiceTask, object *BPMN20.ServiceTask) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ServiceTask **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Auditing **/
	if object.M_auditing != nil {
		xmlElement.M_auditing = new(BPMN20.XsdAuditing)
	}

	/** Now I will save the value of auditing **/
	if object.M_auditing != nil {
		this.SerialyzeAuditing(xmlElement.M_auditing, object.M_auditing)
	}

	/** Serialyze Monitoring **/
	if object.M_monitoring != nil {
		xmlElement.M_monitoring = new(BPMN20.XsdMonitoring)
	}

	/** Now I will save the value of monitoring **/
	if object.M_monitoring != nil {
		this.SerialyzeMonitoring(xmlElement.M_monitoring, object.M_monitoring)
	}

	/** Serialyze ref categoryValueRef **/
	xmlElement.M_categoryValueRef = object.M_categoryValueRef

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref incoming **/
	xmlElement.M_incoming = object.M_incoming

	/** Serialyze ref outgoing **/
	xmlElement.M_outgoing = object.M_outgoing

	/** Serialyze InputOutputSpecification **/
	if object.M_ioSpecification != nil {
		xmlElement.M_ioSpecification = new(BPMN20.XsdInputOutputSpecification)
	}

	/** Now I will save the value of ioSpecification **/
	if object.M_ioSpecification != nil {
		this.SerialyzeInputOutputSpecification(xmlElement.M_ioSpecification, object.M_ioSpecification)
	}

	/** Serialyze Property **/
	if len(object.M_property) > 0 {
		xmlElement.M_property = make([]*BPMN20.XsdProperty, 0)
	}

	/** Now I will save the value of property **/
	for i := 0; i < len(object.M_property); i++ {
		xmlElement.M_property = append(xmlElement.M_property, new(BPMN20.XsdProperty))
		this.SerialyzeProperty(xmlElement.M_property[i], object.M_property[i])
	}

	/** Serialyze DataInputAssociation **/
	if len(object.M_dataInputAssociation) > 0 {
		xmlElement.M_dataInputAssociation = make([]*BPMN20.XsdDataInputAssociation, 0)
	}

	/** Now I will save the value of dataInputAssociation **/
	for i := 0; i < len(object.M_dataInputAssociation); i++ {
		xmlElement.M_dataInputAssociation = append(xmlElement.M_dataInputAssociation, new(BPMN20.XsdDataInputAssociation))
		this.SerialyzeDataInputAssociation(xmlElement.M_dataInputAssociation[i], object.M_dataInputAssociation[i])
	}

	/** Serialyze DataOutputAssociation **/
	if len(object.M_dataOutputAssociation) > 0 {
		xmlElement.M_dataOutputAssociation = make([]*BPMN20.XsdDataOutputAssociation, 0)
	}

	/** Now I will save the value of dataOutputAssociation **/
	for i := 0; i < len(object.M_dataOutputAssociation); i++ {
		xmlElement.M_dataOutputAssociation = append(xmlElement.M_dataOutputAssociation, new(BPMN20.XsdDataOutputAssociation))
		this.SerialyzeDataOutputAssociation(xmlElement.M_dataOutputAssociation[i], object.M_dataOutputAssociation[i])
	}

	/** Serialyze ResourceRole **/
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_0 = make([]*BPMN20.XsdHumanPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_1 = make([]*BPMN20.XsdPotentialOwner, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_2 = make([]*BPMN20.XsdPerformer, 0)
	}
	if len(object.M_resourceRole) > 0 {
		xmlElement.M_resourceRole_3 = make([]*BPMN20.XsdResourceRole, 0)
	}

	/** Now I will save the value of resourceRole **/
	for i := 0; i < len(object.M_resourceRole); i++ {
		switch v := object.M_resourceRole[i].(type) {
		case *BPMN20.HumanPerformer_impl:
			xmlElement.M_resourceRole_0 = append(xmlElement.M_resourceRole_0, new(BPMN20.XsdHumanPerformer))
			this.SerialyzeHumanPerformer(xmlElement.M_resourceRole_0[len(xmlElement.M_resourceRole_0)-1], v)
			log.Println("Serialyze Activity:resourceRole:HumanPerformer")
		case *BPMN20.PotentialOwner:
			xmlElement.M_resourceRole_1 = append(xmlElement.M_resourceRole_1, new(BPMN20.XsdPotentialOwner))
			this.SerialyzePotentialOwner(xmlElement.M_resourceRole_1[len(xmlElement.M_resourceRole_1)-1], v)
			log.Println("Serialyze Activity:resourceRole:PotentialOwner")
		case *BPMN20.Performer_impl:
			xmlElement.M_resourceRole_2 = append(xmlElement.M_resourceRole_2, new(BPMN20.XsdPerformer))
			this.SerialyzePerformer(xmlElement.M_resourceRole_2[len(xmlElement.M_resourceRole_2)-1], v)
			log.Println("Serialyze Activity:resourceRole:Performer")
		case *BPMN20.ResourceRole_impl:
			xmlElement.M_resourceRole_3 = append(xmlElement.M_resourceRole_3, new(BPMN20.XsdResourceRole))
			this.SerialyzeResourceRole(xmlElement.M_resourceRole_3[len(xmlElement.M_resourceRole_3)-1], v)
			log.Println("Serialyze Activity:resourceRole:ResourceRole")
		}
	}

	/** Serialyze LoopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
	}
	if object.M_loopCharacteristics != nil {
		xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
	}

	/** Now I will save the value of loopCharacteristics **/
	if object.M_loopCharacteristics != nil {
		switch v := object.M_loopCharacteristics.(type) {
		case *BPMN20.MultiInstanceLoopCharacteristics:
			xmlElement.M_loopCharacteristics_0 = new(BPMN20.XsdMultiInstanceLoopCharacteristics)
			this.SerialyzeMultiInstanceLoopCharacteristics(xmlElement.M_loopCharacteristics_0, v)
			log.Println("Serialyze Activity:loopCharacteristics:MultiInstanceLoopCharacteristics")
		case *BPMN20.StandardLoopCharacteristics:
			xmlElement.M_loopCharacteristics_1 = new(BPMN20.XsdStandardLoopCharacteristics)
			this.SerialyzeStandardLoopCharacteristics(xmlElement.M_loopCharacteristics_1, v)
			log.Println("Serialyze Activity:loopCharacteristics:StandardLoopCharacteristics")
		}
	}

	/** FlowNode **/
	xmlElement.M_isForCompensation = object.M_isForCompensation

	/** FlowNode **/
	xmlElement.M_startQuantity = object.M_startQuantity

	/** FlowNode **/
	xmlElement.M_completionQuantity = object.M_completionQuantity

	/** Serialyze ref default **/
	xmlElement.M_default = object.M_default
	if len(object.M_implementationStr) > 0 {
		xmlElement.M_implementation = object.M_implementationStr
	} else {

		/** Implementation **/
		if object.M_implementation == BPMN20.Implementation_Unspecified {
			xmlElement.M_implementation = "##unspecified"
		} else if object.M_implementation == BPMN20.Implementation_WebService {
			xmlElement.M_implementation = "##WebService"
		} else {
			xmlElement.M_implementation = "##WebService"
		}
	}

	/** Serialyze ref operationRef **/
	xmlElement.M_operationRef = object.M_operationRef
}

/** serialysation of Import **/
func (this *BPMSXmlFactory) SerialyzeImport(xmlElement *BPMN20.XsdImport, object *BPMN20.Import) {
	if xmlElement == nil {
		return
	}

	/** Import **/
	xmlElement.M_namespace = object.M_namespace

	/** Import **/
	xmlElement.M_location = object.M_location

	/** Import **/
	xmlElement.M_importType = object.M_importType
	this.m_references["Import"] = object
}

/** serialysation of MessageFlow **/
func (this *BPMSXmlFactory) SerialyzeMessageFlow(xmlElement *BPMN20.XsdMessageFlow, object *BPMN20.MessageFlow) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** MessageFlow **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = object.M_targetRef

	/** Serialyze ref messageRef **/
	xmlElement.M_messageRef = object.M_messageRef
}

/** serialysation of EndPoint **/
func (this *BPMSXmlFactory) SerialyzeEndPoint(xmlElement *BPMN20.XsdEndPoint, object *BPMN20.EndPoint) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** EndPoint **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ResourceRole **/
func (this *BPMSXmlFactory) SerialyzeResourceRole(xmlElement *BPMN20.XsdResourceRole, object *BPMN20.ResourceRole_impl) {
	if xmlElement == nil {
		return
	}

	/** Serialyze Documentation **/
	if len(object.M_documentation) > 0 {
		xmlElement.M_documentation = make([]*BPMN20.XsdDocumentation, 0)
	}

	/** Now I will save the value of documentation **/
	for i := 0; i < len(object.M_documentation); i++ {
		xmlElement.M_documentation = append(xmlElement.M_documentation, new(BPMN20.XsdDocumentation))
		this.SerialyzeDocumentation(xmlElement.M_documentation[i], object.M_documentation[i])
	}

	/** Serialyze ExtensionElements **/
	if object.M_extensionElements != nil {
		xmlElement.M_extensionElements = new(BPMN20.XsdExtensionElements)
	}

	/** Now I will save the value of extensionElements **/
	if object.M_extensionElements != nil {
		this.SerialyzeExtensionElements(xmlElement.M_extensionElements, object.M_extensionElements)
	}

	/** ResourceRole **/
	xmlElement.M_id = object.M_id
	//	xmlElement.M_other= object.M_other.(string)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** BaseElement **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref resourceRef **/
	xmlElement.M_resourceRef = &object.M_resourceRef

	/** Serialyze ResourceParameterBinding **/
	if len(object.M_resourceParameterBinding) > 0 {
		xmlElement.M_resourceParameterBinding = make([]*BPMN20.XsdResourceParameterBinding, 0)
	}

	/** Now I will save the value of resourceParameterBinding **/
	for i := 0; i < len(object.M_resourceParameterBinding); i++ {
		xmlElement.M_resourceParameterBinding = append(xmlElement.M_resourceParameterBinding, new(BPMN20.XsdResourceParameterBinding))
		this.SerialyzeResourceParameterBinding(xmlElement.M_resourceParameterBinding[i], object.M_resourceParameterBinding[i])
	}

	/** Serialyze ResourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		xmlElement.M_resourceAssignmentExpression = new(BPMN20.XsdResourceAssignmentExpression)
	}

	/** Now I will save the value of resourceAssignmentExpression **/
	if object.M_resourceAssignmentExpression != nil {
		this.SerialyzeResourceAssignmentExpression(xmlElement.M_resourceAssignmentExpression, object.M_resourceAssignmentExpression)
	}
}
