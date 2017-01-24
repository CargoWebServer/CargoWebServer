//+build BPMN
package BPMS_Runtime

import (
	"encoding/xml"
)

type Runtimes struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Runtimes **/
	M_id               string
	M_name             string
	M_version          string
	M_definitions      []*DefinitionsInstance
	M_exceptions       []*Exception
	M_triggers         []*Trigger
	M_correlationInfos []*CorrelationInfo
	M_logInfos         []*LogInfo
}

/** Xml parser for Runtimes **/
type XsdRuntimes struct {
	XMLName            xml.Name                  `xml:"runtimes"`
	M_definitions      []*XsdDefinitionsInstance `xml:"definitions,omitempty"`
	M_exceptions       []*XsdException           `xml:"exceptions,omitempty"`
	M_triggers         []*XsdTrigger             `xml:"triggers,omitempty"`
	M_correlationInfos []*XsdCorrelationInfo     `xml:"correlationInfos,omitempty"`
	M_id               string                    `xml:"id,attr"`
	M_name             string                    `xml:"name,attr"`
	M_version          string                    `xml:"version,attr"`
}

/** UUID **/
func (this *Runtimes) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Runtimes) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Runtimes) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Runtimes) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Runtimes) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Version **/
func (this *Runtimes) GetVersion() string {
	return this.M_version
}

/** Init reference Version **/
func (this *Runtimes) SetVersion(ref interface{}) {
	this.NeedSave = true
	this.M_version = ref.(string)
}

/** Remove reference Version **/

/** Definitions **/
func (this *Runtimes) GetDefinitions() []*DefinitionsInstance {
	return this.M_definitions
}

/** Init reference Definitions **/
func (this *Runtimes) SetDefinitions(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var definitionss []*DefinitionsInstance
	for i := 0; i < len(this.M_definitions); i++ {
		if this.M_definitions[i].GetUUID() != ref.(Instance).GetUUID() {
			definitionss = append(definitionss, this.M_definitions[i])
		} else {
			isExist = true
			definitionss = append(definitionss, ref.(*DefinitionsInstance))
		}
	}
	if !isExist {
		definitionss = append(definitionss, ref.(*DefinitionsInstance))
	}
	this.M_definitions = definitionss
}

/** Remove reference Definitions **/
func (this *Runtimes) RemoveDefinitions(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(Instance)
	definitions_ := make([]*DefinitionsInstance, 0)
	for i := 0; i < len(this.M_definitions); i++ {
		if toDelete.GetUUID() != this.M_definitions[i].GetUUID() {
			definitions_ = append(definitions_, this.M_definitions[i])
		}
	}
	this.M_definitions = definitions_
}

/** Exceptions **/
func (this *Runtimes) GetExceptions() []*Exception {
	return this.M_exceptions
}

/** Init reference Exceptions **/
func (this *Runtimes) SetExceptions(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var exceptionss []*Exception
	for i := 0; i < len(this.M_exceptions); i++ {
		if this.M_exceptions[i].GetUUID() != ref.(*Exception).GetUUID() {
			exceptionss = append(exceptionss, this.M_exceptions[i])
		} else {
			isExist = true
			exceptionss = append(exceptionss, ref.(*Exception))
		}
	}
	if !isExist {
		exceptionss = append(exceptionss, ref.(*Exception))
	}
	this.M_exceptions = exceptionss
}

/** Remove reference Exceptions **/
func (this *Runtimes) RemoveExceptions(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(Exception)
	exceptions_ := make([]*Exception, 0)
	for i := 0; i < len(this.M_exceptions); i++ {
		if toDelete.GetUUID() != this.M_exceptions[i].GetUUID() {
			exceptions_ = append(exceptions_, this.M_exceptions[i])
		}
	}
	this.M_exceptions = exceptions_
}

/** Triggers **/
func (this *Runtimes) GetTriggers() []*Trigger {
	return this.M_triggers
}

/** Init reference Triggers **/
func (this *Runtimes) SetTriggers(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var triggerss []*Trigger
	for i := 0; i < len(this.M_triggers); i++ {
		if this.M_triggers[i].GetUUID() != ref.(*Trigger).GetUUID() {
			triggerss = append(triggerss, this.M_triggers[i])
		} else {
			isExist = true
			triggerss = append(triggerss, ref.(*Trigger))
		}
	}
	if !isExist {
		triggerss = append(triggerss, ref.(*Trigger))
	}
	this.M_triggers = triggerss
}

/** Remove reference Triggers **/
func (this *Runtimes) RemoveTriggers(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(Trigger)
	triggers_ := make([]*Trigger, 0)
	for i := 0; i < len(this.M_triggers); i++ {
		if toDelete.GetUUID() != this.M_triggers[i].GetUUID() {
			triggers_ = append(triggers_, this.M_triggers[i])
		}
	}
	this.M_triggers = triggers_
}

/** CorrelationInfos **/
func (this *Runtimes) GetCorrelationInfos() []*CorrelationInfo {
	return this.M_correlationInfos
}

/** Init reference CorrelationInfos **/
func (this *Runtimes) SetCorrelationInfos(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var correlationInfoss []*CorrelationInfo
	for i := 0; i < len(this.M_correlationInfos); i++ {
		if this.M_correlationInfos[i].GetUUID() != ref.(*CorrelationInfo).GetUUID() {
			correlationInfoss = append(correlationInfoss, this.M_correlationInfos[i])
		} else {
			isExist = true
			correlationInfoss = append(correlationInfoss, ref.(*CorrelationInfo))
		}
	}
	if !isExist {
		correlationInfoss = append(correlationInfoss, ref.(*CorrelationInfo))
	}
	this.M_correlationInfos = correlationInfoss
}

/** Remove reference CorrelationInfos **/
func (this *Runtimes) RemoveCorrelationInfos(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(CorrelationInfo)
	correlationInfos_ := make([]*CorrelationInfo, 0)
	for i := 0; i < len(this.M_correlationInfos); i++ {
		if toDelete.GetUUID() != this.M_correlationInfos[i].GetUUID() {
			correlationInfos_ = append(correlationInfos_, this.M_correlationInfos[i])
		}
	}
	this.M_correlationInfos = correlationInfos_
}

/** LogInfos **/
func (this *Runtimes) GetLogInfos() []*LogInfo {
	return this.M_logInfos
}

/** Init reference LogInfos **/
func (this *Runtimes) SetLogInfos(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var logInfoss []*LogInfo
	for i := 0; i < len(this.M_logInfos); i++ {
		if this.M_logInfos[i].GetUUID() != ref.(*LogInfo).GetUUID() {
			logInfoss = append(logInfoss, this.M_logInfos[i])
		} else {
			isExist = true
			logInfoss = append(logInfoss, ref.(*LogInfo))
		}
	}
	if !isExist {
		logInfoss = append(logInfoss, ref.(*LogInfo))
	}
	this.M_logInfos = logInfoss
}

/** Remove reference LogInfos **/
func (this *Runtimes) RemoveLogInfos(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(LogInfo)
	logInfos_ := make([]*LogInfo, 0)
	for i := 0; i < len(this.M_logInfos); i++ {
		if toDelete.GetUUID() != this.M_logInfos[i].GetUUID() {
			logInfos_ = append(logInfos_, this.M_logInfos[i])
		}
	}
	this.M_logInfos = logInfos_
}
