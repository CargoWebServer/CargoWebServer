/////////////////////////////////////////  BPMN function calls ////////////////////////////////////////////
// BPMN20
function WorkflowManagerGetDefinitionsIds() {
    var ids = null
    ids = server.GetWorkflowManager().GetDefinitionsIds(messageId, sessionId)
    return ids
}
function WorkflowManagerGetDefinitionsById(id) {
    var definition = null
    definition = server.GetWorkflowManager().GetDefinitionsById(id, messageId, sessionId)
    return definition
}

function WorkflowManagerGetAllDefinitions() {
    var allDefintions = null
    allDefintions = server.GetWorkflowManager().GetAllDefinitions(messageId, sessionId)
    return allDefintions
}

function WorkflowManagerImportXmlBpmnDefinitions(content) {
    err = server.GetWorkflowManager().ImportXmlBpmnDefinitions(content, messageId, sessionId)
    return err
}
// BPMN Runtime
function WorkflowManagerGetDefinitionInstances(definitionsId) {
    var instances = null
    instances = server.GetWorkflowManager().GetDefinitionInstances(definitionsId, messageId, sessionId)
    return instances
}

function WorkflowManagerNewItemAwareElementInstance(bpmnElementId, data) {
    // Now create the new process...
    var itemDefinitionEntity = null
    itemDefinitionEntity = server.GetWorkflowProcessor().NewItemAwareElementInstance(bpmnElementId, data, messageId, sessionId)
    return itemDefinitionEntity
}

function WorkflowManagerStartProcess(processUUID, eventData, eventDefinitionData) {
    // Now create the new process...
    server.GetWorkflowManager().StartProcess(processUUID, eventData, eventDefinitionData, messageId, sessionId)

}

function WorkflowProcessorGetActiveProcessInstances(bpmnElementId) {
    server.GetWorkflowProcessor().GetActiveProcessInstances(bpmnElementId, messageId, sessionId)
}

function WorkflowProcessorActivateActivityInstance(instanceUuid) {
    server.GetWorkflowProcessor().ActivateActivityInstance(instanceUuid, messageId, sessionId)
}

/////////////////////////////////////////  Account Manager calls ////////////////////////////////////////////
function AccountManagerRegister(name, password, email) {
    var newAccount = null
    newAccount = server.GetAccountManager().Register(name, password, email, messageId, sessionId)
    return newAccount
}

function AccountManagerGetAccountById(id) {
    var account
    account = server.GetAccountManager().GetAccountById(id, messageId, sessionId)
    return account
}

function AccountManagerGetUserById(id) {
    var objects
    objects = server.GetAccountManager().GetUserById(id, messageId, sessionId)
    return objects
}


/////////////////////////////////////////  Entity Manager calls /////////////////////////////////////////////
function EntityManagerGetObjectsByType(typeName, queryStr, storeId) {
    var objects
    objects = server.GetEntityManager().GetObjectsByType(typeName, queryStr, storeId, messageId, sessionId)
    return objects
}

function EntityManagerGetEntityLnks(uuid) {
    var lnkLst = null
    lnkLst = server.GetEntityManager().GetEntityLnks(uuid, messageId, sessionId)
    return lnkLst
}

function EntityManagerGetEntityByUuid(uuid) {
    var entity = null
    entity = server.GetEntityManager().GetObjectByUuid(uuid, messageId, sessionId)
    return entity
}

function EntityManagerGetEntityById(storeId, typeName, ids) {
    var entity = null
    entity = server.GetEntityManager().GetObjectById(storeId, typeName, ids, messageId, sessionId)
    return entity
}

function EntityManagerCreateEntity(parentUuid, attributeName, typeName, id, values) {
    var entity = null
    entity = server.GetEntityManager().CreateEntity(parentUuid, attributeName, typeName, id, values, messageId, sessionId)
    return entity
}

function EntityManagerRemoveEntity(uuid) {
    entity = server.GetEntityManager().RemoveEntity(uuid, messageId, sessionId)
}

function EntityManagerSaveEntity(entity, typeName) {
    // save the entity value.
    entity = server.GetEntityManager().SaveEntity(entity, typeName, messageId, sessionId)
    return entity
}

function EntityManagerCreateEntityPrototype(storeId, prototype) {
    var proto = null
    proto = server.GetEntityManager().CreateEntityPrototype(storeId, prototype, messageId, sessionId)
    return proto
}

function EntityManagerGetEntityPrototype(typeName, storeId) {
    var proto = null
    proto = server.GetEntityManager().GetEntityPrototype(typeName, storeId, messageId, sessionId)
    return proto
}

function EntityManagerGetEntityPrototypes(storeId) {
    var protos = null
    protos = server.GetEntityManager().GetEntityPrototypes(storeId, messageId, sessionId)
    return protos
}

function EntityManagerGetDerivedEntityPrototypes(typeName, storeId) {
    var proto = null
    proto = server.GetEntityManager().GetDerivedEntityPrototypes(typeName, messageId, sessionId)
    return proto
}

/////////////////////////////////////////  Data Manager calls /////////////////////////////////////////////
function DataManagerRead(connectionId, query, fields, params) {
    var values = null
    values = server.GetDataManager().Read(connectionId, query, fields, params, messageId, sessionId)
    return values
}

function DataManagerCreate(connectionId, query, values) {
    var id = null
    id = server.GetDataManager().Create(connectionId, query, values, messageId, sessionId)
    return id
}

function DataManagerUpdate(connectionId, query, fields, params) {
    // No value are return.
    server.GetDataManager().Update(connectionId, query, fields, params, messageId, sessionId)
}

function DataManagerDelete(connectionId, query, params) {
    // No value are return.
    server.GetDataManager().Delete(connectionId, query, params, messageId, sessionId)
}

function DataManagerCreateDataStore(storeId, storeType, storeVendor) {
    server.GetDataManager().CreateDataStore(storeId, storeType, storeVendor, messageId, sessionId)
}

function DataManagerDeleteDataStore(storeId) {
    server.GetDataManager().DeleteDataStore(storeId, messageId, sessionId)
}

function DataManagerPing(connectionId) {
    // No value are return.
    server.GetDataManager().Ping(connectionId, messageId, sessionId)
}

function DataManagerConnect(connectionId) {
    // No value are return.
    server.GetDataManager().Connect(connectionId, messageId, sessionId)
}

function DataManagerClose(connectionId) {
    // No value are return.
    server.GetDataManager().Close(connectionId, messageId, sessionId)
}

function DataManagerImportXsdSchema(fileName, fileContent) {
    err = server.GetDataManager().ImportXsdSchema(fileName, fileContent, messageId, sessionId)
    return err
}

function DataManagerImportXmlData(content) {
    err = server.GetDataManager().ImportXmlData(content, messageId, sessionId)
    return err
}

function DataManagerSynchronize(storeId) {
    err = server.GetDataManager().Synchronize(storeId, messageId, sessionId)
    return err
}

/////////////////////////////////////////  Event Manager calls /////////////////////////////////////////////
function EventManagerAppendEventFilter(filter, channelId) {
    server.GetEventManager().AppendEventFilter(filter, channelId, messageId, sessionId)
}

function EventManagerBroadcastNetworkEvent(evtNumber, eventName, eventDatas) {
    // Call the method.
    server.GetEventManager().BroadcastEventData(evtNumber, eventName, eventDatas, messageId, sessionId)
}

/////////////////////////////////////////  File Manager calls /////////////////////////////////////////////
function FileManagerCreateDir(dirName, dirPath) {
    var dir = null
    dir = server.GetFileManager().CreateDir(dirName, dirPath, messageId, sessionId)
    return dir
}

function FileManagerCreateFile(filename, filepath, thumbnailMaxHeight, thumbnailMaxWidth, dbfile) {
    var file = null
    file = server.GetFileManager().CreateFile(filename, filepath, thumbnailMaxHeight, thumbnailMaxWidth, dbfile, messageId, sessionId)
    return file
}

function FileManagerGetFileByPath(path) {
    var file = null
    file = server.GetFileManager().GetFileByPath(path, messageId, sessionId)
    return file
}

function FileManagerOpenFile(fileId) {
    var file = null
    file = server.GetFileManager().OpenFile(fileId, messageId, sessionId)
    return file
}

function FileManagerGetMimeTypeByExtension(fileExtension) {
    var file = null
    file = server.GetFileManager().GetMimeTypeByExtension(fileExtension, messageId, sessionId)
    return file
}

function FileManagerIsFileExist(fileName, filePath) {
    var isFileExist
    isFileExist = server.GetFileManager().IsFileExist(fileName, filePath)
    return isFileExist
}

function FileManagerDeleteFile(uuid) {
    var err
    err = server.GetFileManager().DeleteFile(uuid, messageId, sessionId)
    return err
}

function FileManagerRemoveFile(filePath) {
    var err
    err = server.GetFileManager().RemoveFile(filePath, messageId, sessionId)
    return err
}

////////////////////////////////// Email Manager //////////////////////////
function EmailManagerSendEmail(serverId, from_, to, cc, title, msg, attachs, bodyType) {
    var err = null
    err = server.GetEmailManager().SendEmail(serverId, from_, to, cc, title, msg, attachs, bodyType, messageId, sessionId)
    return err
}

////////////////////////////////// Security Manager //////////////////////////
function SecurityManagerCreateRole(id) {
    var newRole = null
    newRole = server.GetSecurityManager().CreateRole(id, messageId, sessionId)
    return newRole
}

function SecurityManagerGetRole(id) {
    var role = null
    role = server.GetSecurityManager().GetRole(id, messageId, sessionId)
    return role
}

function SecurityManagerDeleteRole(id) {
    server.GetSecurityManager().DeleteRole(id, messageId, sessionId)
}

function SecurityManagerHasAccount(roleId, accountId) {
    var roleHasAccount = server.GetSecurityManager().HasAccount(roleId, accountId, messageId, sessionId)
    return roleHasAccount
}

function SecurityManagerAppendAccount(roleId, accountId) {
    server.GetSecurityManager().AppendAccount(roleId, accountId, messageId, sessionId)
}

function SecurityManagerRemoveAccount(roleId, accountId) {
    server.GetSecurityManager().RemoveAccount(roleId, accountId, messageId, sessionId)
}

function SecurityManagerHasAction(roleId, actionName) {
    var roleHasAction = server.GetSecurityManager().HasAction(roleId, actionName, messageId, sessionId)
    return roleHasAction
}

function SecurityManagerAppendAction(roleId, actionName) {
    server.GetSecurityManager().AppendAction(roleId, actionName, messageId, sessionId)
}

function SecurityManagerRemoveAction(roleId, actionName) {
    server.GetSecurityManager().RemoveAction(roleId, actionName, messageId, sessionId)
}

function SecurityManagerAppendPermission(accountId, permissionType, pattern) {
    server.GetSecurityManager().AppendPermission(accountId, permissionType, pattern, messageId, sessionId)
}

function SecurityManagerRemovePermission(accountId, permissionPattern) {
    server.GetSecurityManager().RemovePermission(accountId, permissionPattern, messageId, sessionId)
}

function SecurityManagerChangeAdminPassword(pwd, newPwd) {
    server.GetSecurityManager().ChangeAdminPassword(pwd, newPwd, messageId, sessionId)
}

function SecurityManagerGetResource(clientId, scope, query, idTokenUuid){
    var accessUuid = "" // not know by the client side.
    var result = server.GetOAuth2Manager().GetResource(clientId,scope, query, idTokenUuid, accessUuid, messageId, sessionId)
    return result
}

////////////////////////////////// Service Manager ///////////////////////
function ServiceManagerGetServiceActions(serviceName){
    var result = server.GetServiceManager().GetServiceActions(serviceName, messageId, sessionId)
    return result
}


////////////////////////////////// Server //////////////////////////
function SetRootPath(path) {
    // Call set root path on the server...
    server.SetRootPath(path)
}

function Connect(host, port) {
    server.Connect(host, port)
}

function Disconnect(host, port) {
    server.Disconnect(host, port)
}

function Stop() {
    server.Stop()
}

///////////////////////////////// Session Manager //////////////////////

function SessionManagerLogin(name, password, serverId) {
    var newSession = null
    newSession = server.GetSessionManager().Login(name, password, serverId, messageId, sessionId)
    return newSession
}

function SessionManagerGetActiveSessions() {
    var sessions = null
    sessions = server.GetSessionManager().GetActiveSessions(messageId, sessionId)
    return sessions
}

function SessionManagerGetActiveSessionByAccountId(accountId) {
    var sessions = null
    sessions = server.GetSessionManager().GetActiveSessionByAccountId(accountId)
    return sessions
}

function SessionManagerUpdateSessionState(state) {
    var err = server.GetSessionManager().UpdateSessionState(state, messageId, sessionId)
    return err
}

function SessionManagerLogout(toCloseId) {
    var err = server.GetSessionManager().Logout(toCloseId, messageId, sessionId)
    return err
}

/////////////////////////////////////// Other script /////////////////////////////////////////
/**
 * Read excel file and get it content as comma separated values.
 */
function ExcelToCsv(filePath) {
    var val = filePath.split("\\")
    var fileName = val[val.length - 1].split(".")[0]
    var outputFile = server.GetConfigurationManager().GetTmpPath() + "/" + fileName + ".csv"
    // C:\\ instead of C:/
    outputFile = outputFile.replace("/", "\\", -1)
    server.ExecuteVbScript("xlsx2csv.vbs", [filePath, outputFile], messageId, sessionId)

    // I will read the file content.
    var results = {"data" : server.GetFileManager().ReadCsvFile(outputFile, messageId, sessionId), "outputFile":outputFile}

    return results
}
