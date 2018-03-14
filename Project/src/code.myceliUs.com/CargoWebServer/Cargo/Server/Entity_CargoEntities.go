// +build CargoEntities

package Server

import (
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_EntityEntityPrototype() {

	var entityEntityProto EntityPrototype
	entityEntityProto.TypeName = "CargoEntities.Entity"
	entityEntityProto.IsAbstract = true
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.User")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Group")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Error")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.LogEntry")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.File")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Notification")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.TextMessage")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Account")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Computer")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Log")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Project")
	entityEntityProto.Ids = append(entityEntityProto.Ids, "UUID")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "UUID")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 0)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")
	entityEntityProto.Indexs = append(entityEntityProto.Indexs, "ParentUuid")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "ParentUuid")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 1)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "ParentLnk")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 2)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	entityEntityProto.Ids = append(entityEntityProto.Ids, "M_id")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 3)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, true)
	entityEntityProto.Fields = append(entityEntityProto.Fields, "M_id")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.ID")
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")

	/** associations of Entity **/
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 4)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.Fields = append(entityEntityProto.Fields, "M_entitiesPtr")
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "undefined")
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "undefined")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&entityEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ParameterEntityPrototype() {

	var parameterEntityProto EntityPrototype
	parameterEntityProto.TypeName = "CargoEntities.Parameter"
	parameterEntityProto.Ids = append(parameterEntityProto.Ids, "UUID")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "UUID")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 0)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.Indexs = append(parameterEntityProto.Indexs, "ParentUuid")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "ParentUuid")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 1)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "ParentLnk")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 2)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")

	/** members of Parameter **/
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 3)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_name")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 4)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_type")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 5)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_isArray")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.boolean")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "false")

	/** associations of Parameter **/
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 6)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_parametersPtr")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "undefined")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "undefined")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "CargoEntities.Parameter:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&parameterEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ActionEntityPrototype() {

	var actionEntityProto EntityPrototype
	actionEntityProto.TypeName = "CargoEntities.Action"
	actionEntityProto.Ids = append(actionEntityProto.Ids, "UUID")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "UUID")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 0)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.Indexs = append(actionEntityProto.Indexs, "ParentUuid")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "ParentUuid")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 1)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "ParentLnk")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 2)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")

	/** members of Action **/
	actionEntityProto.Ids = append(actionEntityProto.Ids, "M_name")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 3)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_name")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.ID")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 4)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_doc")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 5)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_parameters")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "[]")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]CargoEntities.Parameter")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 6)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_results")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "[]")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]CargoEntities.Parameter")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 7)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_accessType")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "1")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "enum:AccessType_Hidden:AccessType_Public:AccessType_Restricted")

	/** associations of Action **/
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 8)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_entitiesPtr")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "undefined")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "undefined")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&actionEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ErrorEntityPrototype() {

	var errorEntityProto EntityPrototype
	errorEntityProto.TypeName = "CargoEntities.Error"
	errorEntityProto.SuperTypeNames = append(errorEntityProto.SuperTypeNames, "CargoEntities.Entity")
	errorEntityProto.SuperTypeNames = append(errorEntityProto.SuperTypeNames, "CargoEntities.Message")
	errorEntityProto.Ids = append(errorEntityProto.Ids, "UUID")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "UUID")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 0)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")
	errorEntityProto.Indexs = append(errorEntityProto.Indexs, "ParentUuid")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "ParentUuid")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 1)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "ParentLnk")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 2)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	errorEntityProto.Ids = append(errorEntityProto.Ids, "M_id")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 3)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_id")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.ID")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 4)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_body")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")

	/** members of Error **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 5)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_errorPath")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 6)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_code")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.int")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "0")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 7)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_accountRef")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "CargoEntities.Account:Ref")

	/** associations of Error **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 8)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_entitiesPtr")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&errorEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_LogEntryEntityPrototype() {

	var logEntryEntityProto EntityPrototype
	logEntryEntityProto.TypeName = "CargoEntities.LogEntry"
	logEntryEntityProto.SuperTypeNames = append(logEntryEntityProto.SuperTypeNames, "CargoEntities.Entity")
	logEntryEntityProto.Ids = append(logEntryEntityProto.Ids, "UUID")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "UUID")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 0)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")
	logEntryEntityProto.Indexs = append(logEntryEntityProto.Indexs, "ParentUuid")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "ParentUuid")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 1)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "ParentLnk")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 2)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	logEntryEntityProto.Ids = append(logEntryEntityProto.Ids, "M_id")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 3)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_id")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.ID")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")

	/** members of LogEntry **/
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 4)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_creationTime")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.date")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "new Date()")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 5)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_entityRef")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Entity:Ref")

	/** associations of LogEntry **/
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 6)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_loggerPtr")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Log:Ref")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 7)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_entitiesPtr")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&logEntryEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_LogEntityPrototype() {

	var logEntityProto EntityPrototype
	logEntityProto.TypeName = "CargoEntities.Log"
	logEntityProto.SuperTypeNames = append(logEntityProto.SuperTypeNames, "CargoEntities.Entity")
	logEntityProto.Ids = append(logEntityProto.Ids, "UUID")
	logEntityProto.Fields = append(logEntityProto.Fields, "UUID")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 0)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")
	logEntityProto.Indexs = append(logEntityProto.Indexs, "ParentUuid")
	logEntityProto.Fields = append(logEntityProto.Fields, "ParentUuid")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 1)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")
	logEntityProto.Fields = append(logEntityProto.Fields, "ParentLnk")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 2)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	logEntityProto.Ids = append(logEntityProto.Ids, "M_id")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 3)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, true)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_id")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.ID")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")

	/** members of Log **/
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 4)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, true)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_entries")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "[]")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "[]CargoEntities.LogEntry")

	/** associations of Log **/
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 5)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_entitiesPtr")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "undefined")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "undefined")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&logEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ProjectEntityPrototype() {

	var projectEntityProto EntityPrototype
	projectEntityProto.TypeName = "CargoEntities.Project"
	projectEntityProto.SuperTypeNames = append(projectEntityProto.SuperTypeNames, "CargoEntities.Entity")
	projectEntityProto.Ids = append(projectEntityProto.Ids, "UUID")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "UUID")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 0)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")
	projectEntityProto.Indexs = append(projectEntityProto.Indexs, "ParentUuid")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "ParentUuid")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 1)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "ParentLnk")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 2)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	projectEntityProto.Ids = append(projectEntityProto.Ids, "M_id")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 3)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_id")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.ID")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")

	/** members of Project **/
	projectEntityProto.Indexs = append(projectEntityProto.Indexs, "M_name")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 4)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_name")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 5)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_filesRef")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "undefined")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "[]")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "[]CargoEntities.File:Ref")

	/** associations of Project **/
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 6)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_entitiesPtr")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "undefined")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "undefined")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&projectEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_MessageEntityPrototype() {

	var messageEntityProto EntityPrototype
	messageEntityProto.TypeName = "CargoEntities.Message"
	messageEntityProto.IsAbstract = true
	messageEntityProto.SuperTypeNames = append(messageEntityProto.SuperTypeNames, "CargoEntities.Entity")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.Error")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.Notification")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.TextMessage")
	messageEntityProto.Ids = append(messageEntityProto.Ids, "UUID")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "UUID")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 0)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")
	messageEntityProto.Indexs = append(messageEntityProto.Indexs, "ParentUuid")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "ParentUuid")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 1)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "ParentLnk")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 2)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	messageEntityProto.Ids = append(messageEntityProto.Ids, "M_id")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 3)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, true)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_id")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.ID")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 4)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, true)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_body")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")

	/** associations of Message **/
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 5)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_entitiesPtr")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "undefined")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "undefined")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&messageEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_NotificationEntityPrototype() {

	var notificationEntityProto EntityPrototype
	notificationEntityProto.TypeName = "CargoEntities.Notification"
	notificationEntityProto.SuperTypeNames = append(notificationEntityProto.SuperTypeNames, "CargoEntities.Entity")
	notificationEntityProto.SuperTypeNames = append(notificationEntityProto.SuperTypeNames, "CargoEntities.Message")
	notificationEntityProto.Ids = append(notificationEntityProto.Ids, "UUID")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "UUID")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 0)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")
	notificationEntityProto.Indexs = append(notificationEntityProto.Indexs, "ParentUuid")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "ParentUuid")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 1)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "ParentLnk")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 2)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	notificationEntityProto.Ids = append(notificationEntityProto.Ids, "M_id")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 3)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_id")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.ID")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 4)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_body")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")

	/** members of Notification **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 5)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_fromRef")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Account:Ref")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 6)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_toRef")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Account:Ref")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 7)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_type")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 8)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_code")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.int")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "0")

	/** associations of Notification **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 9)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_entitiesPtr")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&notificationEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_TextMessageEntityPrototype() {

	var textMessageEntityProto EntityPrototype
	textMessageEntityProto.TypeName = "CargoEntities.TextMessage"
	textMessageEntityProto.SuperTypeNames = append(textMessageEntityProto.SuperTypeNames, "CargoEntities.Entity")
	textMessageEntityProto.SuperTypeNames = append(textMessageEntityProto.SuperTypeNames, "CargoEntities.Message")
	textMessageEntityProto.Ids = append(textMessageEntityProto.Ids, "UUID")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "UUID")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 0)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")
	textMessageEntityProto.Indexs = append(textMessageEntityProto.Indexs, "ParentUuid")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "ParentUuid")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 1)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "ParentLnk")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 2)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	textMessageEntityProto.Ids = append(textMessageEntityProto.Ids, "M_id")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 3)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_id")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.ID")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 4)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_body")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** members of TextMessage **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 5)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_creationTime")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.date")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "new Date()")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 6)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_fromRef")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Account:Ref")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 7)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_toRef")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Account:Ref")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 8)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_title")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** associations of TextMessage **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 9)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_entitiesPtr")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&textMessageEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_SessionEntityPrototype() {

	var sessionEntityProto EntityPrototype
	sessionEntityProto.TypeName = "CargoEntities.Session"
	sessionEntityProto.Ids = append(sessionEntityProto.Ids, "UUID")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "UUID")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 0)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")
	sessionEntityProto.Indexs = append(sessionEntityProto.Indexs, "ParentUuid")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "ParentUuid")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 1)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "ParentLnk")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 2)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")

	/** members of Session **/
	sessionEntityProto.Ids = append(sessionEntityProto.Ids, "M_id")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 3)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_id")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.ID")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 4)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_startTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.date")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "new Date()")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 5)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_endTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.date")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "new Date()")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 6)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_statusTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.date")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "new Date()")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 7)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_sessionState")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "1")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "enum:SessionState_Online:SessionState_Away:SessionState_Offline")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 8)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_computerRef")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "CargoEntities.Computer:Ref")

	/** associations of Session **/
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 9)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_accountPtr")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "CargoEntities.Account:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&sessionEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_RoleEntityPrototype() {

	var roleEntityProto EntityPrototype
	roleEntityProto.TypeName = "CargoEntities.Role"
	roleEntityProto.Ids = append(roleEntityProto.Ids, "UUID")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "UUID")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 0)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")
	roleEntityProto.Indexs = append(roleEntityProto.Indexs, "ParentUuid")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "ParentUuid")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 1)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "ParentLnk")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 2)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")

	/** members of Role **/
	roleEntityProto.Ids = append(roleEntityProto.Ids, "M_id")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 3)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_id")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.ID")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 4)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_accountsRef")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "[]")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]CargoEntities.Account:Ref")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 5)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_actionsRef")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "[]")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]CargoEntities.Action:Ref")

	/** associations of Role **/
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 6)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_entitiesPtr")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&roleEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_AccountEntityPrototype() {

	var accountEntityProto EntityPrototype
	accountEntityProto.TypeName = "CargoEntities.Account"
	accountEntityProto.SuperTypeNames = append(accountEntityProto.SuperTypeNames, "CargoEntities.Entity")
	accountEntityProto.Ids = append(accountEntityProto.Ids, "UUID")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "UUID")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 0)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "ParentUuid")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "ParentUuid")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 1)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "ParentLnk")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 2)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	accountEntityProto.Ids = append(accountEntityProto.Ids, "M_id")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 3)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_id")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.ID")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")

	/** members of Account **/
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "M_name")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 4)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_name")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 5)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_password")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "M_email")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 6)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_email")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 7)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_sessions")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Session")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 8)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_messages")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Message")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 9)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_userRef")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "CargoEntities.User:Ref")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 10)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_rolesRef")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Role:Ref")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 11)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_permissionsRef")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Permission:Ref")

	/** associations of Account **/
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 12)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_entitiesPtr")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&accountEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ComputerEntityPrototype() {

	var computerEntityProto EntityPrototype
	computerEntityProto.TypeName = "CargoEntities.Computer"
	computerEntityProto.SuperTypeNames = append(computerEntityProto.SuperTypeNames, "CargoEntities.Entity")
	computerEntityProto.Ids = append(computerEntityProto.Ids, "UUID")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "UUID")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 0)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.Indexs = append(computerEntityProto.Indexs, "ParentUuid")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "ParentUuid")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 1)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "ParentLnk")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 2)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	computerEntityProto.Ids = append(computerEntityProto.Ids, "M_id")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 3)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_id")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.ID")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")

	/** members of Computer **/
	computerEntityProto.Indexs = append(computerEntityProto.Indexs, "M_name")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 4)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_name")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 5)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_ipv4")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 6)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_osType")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "1")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "enum:OsType_Unknown:OsType_Linux:OsType_Windows7:OsType_Windows8:OsType_Windows10:OsType_OSX:OsType_IOS")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 7)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_platformType")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "1")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "enum:PlatformType_Unknown:PlatformType_Tablet:PlatformType_Phone:PlatformType_Desktop:PlatformType_Laptop")

	/** associations of Computer **/
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 8)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_entitiesPtr")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "undefined")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "undefined")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&computerEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_PermissionEntityPrototype() {

	var permissionEntityProto EntityPrototype
	permissionEntityProto.TypeName = "CargoEntities.Permission"
	permissionEntityProto.Ids = append(permissionEntityProto.Ids, "UUID")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "UUID")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 0)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")
	permissionEntityProto.Indexs = append(permissionEntityProto.Indexs, "ParentUuid")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "ParentUuid")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 1)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "ParentLnk")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 2)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")

	/** members of Permission **/
	permissionEntityProto.Ids = append(permissionEntityProto.Ids, "M_id")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 3)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_id")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.ID")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 4)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_types")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.int")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "0")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 5)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_accountsRef")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "undefined")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "[]")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "[]CargoEntities.Account:Ref")

	/** associations of Permission **/
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 6)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_entitiesPtr")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "undefined")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "undefined")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&permissionEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_FileEntityPrototype() {

	var fileEntityProto EntityPrototype
	fileEntityProto.TypeName = "CargoEntities.File"
	fileEntityProto.SuperTypeNames = append(fileEntityProto.SuperTypeNames, "CargoEntities.Entity")
	fileEntityProto.Ids = append(fileEntityProto.Ids, "UUID")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "UUID")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 0)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "ParentUuid")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "ParentUuid")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 1)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "ParentLnk")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 2)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	fileEntityProto.Ids = append(fileEntityProto.Ids, "M_id")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 3)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_id")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.ID")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")

	/** members of File **/
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "M_name")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 4)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_name")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "M_path")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 5)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_path")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 6)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_size")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.int")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "0")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 7)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_modeTime")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.date")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "new Date()")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 8)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_isDir")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.boolean")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "false")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 9)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_checksum")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 10)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_data")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 11)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_thumbnail")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 12)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_mime")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 13)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_files")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "[]")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "[]CargoEntities.File")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 14)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_fileType")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "1")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "enum:FileType_DbFile:FileType_DiskFile")

	/** associations of File **/
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 15)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_parentDirPtr")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "CargoEntities.File:Ref")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 16)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_entitiesPtr")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&fileEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_UserEntityPrototype() {

	var userEntityProto EntityPrototype
	userEntityProto.TypeName = "CargoEntities.User"
	userEntityProto.SuperTypeNames = append(userEntityProto.SuperTypeNames, "CargoEntities.Entity")
	userEntityProto.Ids = append(userEntityProto.Ids, "UUID")
	userEntityProto.Fields = append(userEntityProto.Fields, "UUID")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 0)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.Indexs = append(userEntityProto.Indexs, "ParentUuid")
	userEntityProto.Fields = append(userEntityProto.Fields, "ParentUuid")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 1)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.Fields = append(userEntityProto.Fields, "ParentLnk")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 2)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	userEntityProto.Ids = append(userEntityProto.Ids, "M_id")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 3)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_id")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.ID")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")

	/** members of User **/
	userEntityProto.Indexs = append(userEntityProto.Indexs, "M_firstName")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 4)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_firstName")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.Indexs = append(userEntityProto.Indexs, "M_lastName")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 5)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_lastName")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 6)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_middle")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 7)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_phone")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.Indexs = append(userEntityProto.Indexs, "M_email")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 8)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_email")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 9)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_memberOfRef")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "[]")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]CargoEntities.Group:Ref")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 10)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_accounts")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "[]")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]CargoEntities.Account:Ref")

	/** associations of User **/
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 11)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_entitiesPtr")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&userEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_GroupEntityPrototype() {

	var groupEntityProto EntityPrototype
	groupEntityProto.TypeName = "CargoEntities.Group"
	groupEntityProto.SuperTypeNames = append(groupEntityProto.SuperTypeNames, "CargoEntities.Entity")
	groupEntityProto.Ids = append(groupEntityProto.Ids, "UUID")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "UUID")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 0)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")
	groupEntityProto.Indexs = append(groupEntityProto.Indexs, "ParentUuid")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "ParentUuid")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 1)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "ParentLnk")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 2)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	groupEntityProto.Ids = append(groupEntityProto.Ids, "M_id")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 3)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_id")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.ID")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")

	/** members of Group **/
	groupEntityProto.Indexs = append(groupEntityProto.Indexs, "M_name")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 4)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_name")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 5)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_membersRef")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "undefined")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "[]")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "[]CargoEntities.User:Ref")

	/** associations of Group **/
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 6)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_entitiesPtr")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "undefined")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "undefined")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&groupEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_EntitiesEntityPrototype() {

	var entitiesEntityProto EntityPrototype
	entitiesEntityProto.TypeName = "CargoEntities.Entities"
	entitiesEntityProto.Ids = append(entitiesEntityProto.Ids, "UUID")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "UUID")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 0)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.Indexs = append(entitiesEntityProto.Indexs, "ParentUuid")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "ParentUuid")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 1)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "ParentLnk")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 2)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")

	/** members of Entities **/
	entitiesEntityProto.Ids = append(entitiesEntityProto.Ids, "M_id")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 3)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_id")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.ID")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 4)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_name")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 5)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_version")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 6)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_entities")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Entity")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 7)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_roles")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Role")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 8)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_permissions")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Permission")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 9)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_actions")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Action")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*GraphStore)
	store.CreateEntityPrototype(&entitiesEntityProto)

}

/** Register the entity to the dynamic typing system. **/
func (this *EntityManager) registerCargoEntitiesObjects() {
	Utility.RegisterType((*CargoEntities.Parameter)(nil))
	Utility.RegisterType((*CargoEntities.Action)(nil))
	Utility.RegisterType((*CargoEntities.Error)(nil))
	Utility.RegisterType((*CargoEntities.LogEntry)(nil))
	Utility.RegisterType((*CargoEntities.Log)(nil))
	Utility.RegisterType((*CargoEntities.Project)(nil))
	Utility.RegisterType((*CargoEntities.Notification)(nil))
	Utility.RegisterType((*CargoEntities.TextMessage)(nil))
	Utility.RegisterType((*CargoEntities.Session)(nil))
	Utility.RegisterType((*CargoEntities.Role)(nil))
	Utility.RegisterType((*CargoEntities.Account)(nil))
	Utility.RegisterType((*CargoEntities.Computer)(nil))
	Utility.RegisterType((*CargoEntities.Permission)(nil))
	Utility.RegisterType((*CargoEntities.File)(nil))
	Utility.RegisterType((*CargoEntities.User)(nil))
	Utility.RegisterType((*CargoEntities.Group)(nil))
	Utility.RegisterType((*CargoEntities.Entities)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createCargoEntitiesPrototypes() {
	this.create_CargoEntities_EntityEntityPrototype()
	this.create_CargoEntities_ParameterEntityPrototype()
	this.create_CargoEntities_ActionEntityPrototype()
	this.create_CargoEntities_ErrorEntityPrototype()
	this.create_CargoEntities_LogEntryEntityPrototype()
	this.create_CargoEntities_LogEntityPrototype()
	this.create_CargoEntities_ProjectEntityPrototype()
	this.create_CargoEntities_MessageEntityPrototype()
	this.create_CargoEntities_NotificationEntityPrototype()
	this.create_CargoEntities_TextMessageEntityPrototype()
	this.create_CargoEntities_SessionEntityPrototype()
	this.create_CargoEntities_RoleEntityPrototype()
	this.create_CargoEntities_AccountEntityPrototype()
	this.create_CargoEntities_ComputerEntityPrototype()
	this.create_CargoEntities_PermissionEntityPrototype()
	this.create_CargoEntities_FileEntityPrototype()
	this.create_CargoEntities_UserEntityPrototype()
	this.create_CargoEntities_GroupEntityPrototype()
	this.create_CargoEntities_EntitiesEntityPrototype()
}
