/////////////////////////////////////////  BPMN function calls ////////////////////////////////////////////
// BPMN20
function GetDefinitionsIds() {
    var ids = null
    ids = server.GetWorkflowManager().GetDefinitionsIds(messageId, sessionId)
    return ids
}
function GetDefinitionsById(id) {
    var definition = null
    definition = server.GetWorkflowManager().GetDefinitionsById(id, messageId, sessionId)
    return definition
}

function GetAllDefinitions() {
    var allDefintions = null
    allDefintions = server.GetWorkflowManager().GetAllDefinitions(messageId, sessionId)
    return allDefintions
}

function ImportXmlBpmnDefinitions(content) {
    err = server.GetWorkflowManager().ImportXmlBpmnDefinitions(content, messageId, sessionId)
    return err
}
// BPMN Runtime
function GetDefinitionInstances(definitionsId) {
    var instances = null
    instances = server.GetWorkflowManager().GetDefinitionInstances(definitionsId, messageId, sessionId)
    return instances
}

function NewItemAwareElementInstance(bpmnElementId, data) {
    // Now create the new process...
    var itemDefinitionEntity = null
    itemDefinitionEntity = server.GetWorkflowProcessor().NewItemAwareElementInstance(bpmnElementId, data, messageId, sessionId)
    return itemDefinitionEntity
}

function StartProcess(processUUID, eventData, eventDefinitionData) {
    // Now create the new process...
    server.GetWorkflowManager().StartProcess(processUUID, eventData, eventDefinitionData, messageId, sessionId)

}

function GetActiveProcessInstances(bpmnElementId) {
    server.GetWorkflowProcessor().GetActiveProcessInstances(bpmnElementId, messageId, sessionId)
}

function ActivateActivityInstance(instanceUuid) {
    server.GetWorkflowProcessor().ActivateActivityInstance(instanceUuid, messageId, sessionId)
}

/////////////////////////////////////////  Account Manager calls ////////////////////////////////////////////
function RegisterAccount(name, password, email) {
    var newAccount = null
    newAccount = server.GetAccountManager().Register(name, password, email, messageId, sessionId)
    return newAccount
}

function GetAccountById(id) {
    var account
    account = server.GetAccountManager().GetAccountById(id, messageId, sessionId)
    return account
}

function GetUserById(id) {
    var objects
    objects = server.GetAccountManager().GetUserById(id, messageId, sessionId)
    return objects
}


/////////////////////////////////////////  Entity Manager calls /////////////////////////////////////////////
function GetObjectsByType(typeName, queryStr, storeId) {
    var objects
    objects = server.GetEntityManager().GetObjectsByType(typeName, queryStr, storeId, messageId, sessionId)
    return objects
}

function GetEntityLnks(uuid) {
    var lnkLst = null
    lnkLst = server.GetEntityManager().GetEntityLnks(uuid, messageId, sessionId)
    return lnkLst
}

function GetEntityByUuid(uuid) {
    var entity = null
    entity = server.GetEntityManager().GetObjectByUuid(uuid, messageId, sessionId)
    return entity
}

function GetEntityById(storeId, typeName, ids) {
    var entity = null
    entity = server.GetEntityManager().GetObjectById(storeId, typeName, ids, messageId, sessionId)
    return entity
}

function CreateEntity(parentUuid, attributeName, typeName, id, values) {
    var entity = null
    entity = server.GetEntityManager().CreateEntity(parentUuid, attributeName, typeName, id, values, messageId, sessionId)
    return entity
}

function RemoveEntity(uuid) {
    entity = server.GetEntityManager().RemoveEntity(uuid, messageId, sessionId)
}

function SaveEntity(entity, typeName) {
    // save the entity value.
    entity = server.GetEntityManager().SaveEntity(entity, typeName, messageId, sessionId)
    return entity
}

function CreateEntityPrototype(storeId, prototype) {
    var proto = null
    proto = server.GetEntityManager().CreateEntityPrototype(storeId, prototype, messageId, sessionId)
    return proto
}

function GetEntityPrototype(typeName, storeId) {
    var proto = null
    proto = server.GetEntityManager().GetEntityPrototype(typeName, storeId, messageId, sessionId)
    return proto
}

function GetEntityPrototypes(storeId) {
    var protos = null
    protos = server.GetEntityManager().GetEntityPrototypes(storeId, messageId, sessionId)
    return protos
}

function GetDerivedEntityPrototypes(typeName, storeId) {
    var proto = null
    proto = server.GetEntityManager().GetDerivedEntityPrototypes(typeName, messageId, sessionId)
    return proto
}

/////////////////////////////////////////  Data Manager calls /////////////////////////////////////////////
function Read(connectionId, query, fields, params) {
    var values = null
    values = server.GetDataManager().Read(connectionId, query, fields, params, messageId, sessionId)
    return values
}

function Create(connectionId, query, values) {
    var id = null
    id = server.GetDataManager().Create(connectionId, query, values, messageId, sessionId)
    return id
}

function Update(connectionId, query, fields, params) {
    // No value are return.
    server.GetDataManager().Update(connectionId, query, fields, params, messageId, sessionId)
}

function Delete(connectionId, query, params) {
    // No value are return.
    server.GetDataManager().Delete(connectionId, query, params, messageId, sessionId)
}

function CreateDataStore(storeId, storeType, storeVendor) {
    server.GetDataManager().CreateDataStore(storeId, storeType, storeVendor, messageId, sessionId)
}

function DeleteDataStore(storeId) {
    server.GetDataManager().DeleteDataStore(storeId, messageId, sessionId)
}

function Ping_(connectionId) {
    // No value are return.
    server.GetDataManager().Ping(connectionId, messageId, sessionId)
}

function Connect_(connectionId) {
    // No value are return.
    server.GetDataManager().Connect(connectionId, messageId, sessionId)
}

function Close_(connectionId) {
    // No value are return.
    server.GetDataManager().Close(connectionId, messageId, sessionId)
}

function ImportXsdSchema(fileName, fileContent) {
    err = server.GetDataManager().ImportXsdSchema(fileName, fileContent, messageId, sessionId)
    return err
}

function ImportXmlData(content) {
    err = server.GetDataManager().ImportXmlData(content, messageId, sessionId)
    return err
}

function Synchronize(storeId) {
    err = server.GetDataManager().Synchronize(storeId, messageId, sessionId)
    return err
}

/////////////////////////////////////////  Event Manager calls /////////////////////////////////////////////
function AppendEventFilter(filter, channelId) {
    server.GetEventManager().AppendEventFilter(filter, channelId, messageId, sessionId)
}

function BroadcastNetworkEvent(evtNumber, eventName, eventDatas) {
    // Call the method.
    server.GetEventManager().BroadcastEventData(evtNumber, eventName, eventDatas, messageId, sessionId)
}

/////////////////////////////////////////  File Manager calls /////////////////////////////////////////////
function CreateDir(dirName, dirPath) {
    var dir = null
    dir = server.GetFileManager().CreateDir(dirName, dirPath, messageId, sessionId)
    return dir
}

function CreateFile(filename, filepath, thumbnailMaxHeight, thumbnailMaxWidth, dbfile) {
    var file = null
    file = server.GetFileManager().CreateFile(filename, filepath, thumbnailMaxHeight, thumbnailMaxWidth, dbfile, messageId, sessionId)
    return file
}

function GetFileByPath(path) {
    var file = null
    file = server.GetFileManager().GetFileByPath(path, messageId, sessionId)
    return file
}

function OpenFile(fileId) {
    var file = null
    file = server.GetFileManager().OpenFile(fileId, messageId, sessionId)
    return file
}

function GetMimeTypeByExtension(fileExtension) {
    var file = null
    file = server.GetFileManager().GetMimeTypeByExtension(fileExtension, messageId, sessionId)
    return file
}

function IsFileExist(fileName, filePath) {
    var isFileExist
    isFileExist = server.GetFileManager().IsFileExist(fileName, filePath)
    return isFileExist
}

function DeleteFile(uuid) {
    var err
    err = server.GetFileManager().DeleteFile(uuid, messageId, sessionId)
    return err
}

function RemoveFile(filePath) {
    var err
    err = server.GetFileManager().RemoveFile(filePath, messageId, sessionId)
    return err
}

////////////////////////////////// Email Manager //////////////////////////
function SendEmail(serverId, from_, to, cc, title, msg, attachs, bodyType) {
    var err = null
    err = server.GetEmailManager().SendEmail(serverId, from_, to, cc, title, msg, attachs, bodyType, messageId, sessionId)
    return err
}

////////////////////////////////// Security Manager //////////////////////////
function CreateRole(id) {
    var newRole = null
    newRole = server.GetSecurityManager().CreateRole(id, messageId, sessionId)
    return newRole
}

function GetRole(id) {
    var role = null
    role = server.GetSecurityManager().GetRole(id, messageId, sessionId)
    return role
}

function DeleteRole(id) {
    server.GetSecurityManager().DeleteRole(id, messageId, sessionId)
}

function HasAccount(roleId, accountId) {
    var roleHasAccount = server.GetSecurityManager().HasAccount(roleId, accountId, messageId, sessionId)
    return roleHasAccount
}

function AppendAccount(roleId, accountId) {
    server.GetSecurityManager().AppendAccount(roleId, accountId, messageId, sessionId)
}

function RemoveAccount(roleId, accountId) {
    server.GetSecurityManager().RemoveAccount(roleId, accountId, messageId, sessionId)
}

function HasAction(roleId, actionName) {
    var roleHasAction = server.GetSecurityManager().HasAction(roleId, actionName, messageId, sessionId)
    return roleHasAction
}

function AppendAction(roleId, actionName) {
    server.GetSecurityManager().AppendAction(roleId, actionName, messageId, sessionId)
}

function RemoveAction(roleId, actionName) {
    server.GetSecurityManager().RemoveAction(roleId, actionName, messageId, sessionId)
}

function AppendPermission(accountId, permissionType, pattern) {
    server.GetSecurityManager().AppendPermission(accountId, permissionType, pattern, messageId, sessionId)
}

function RemovePermission(accountId, permissionPattern) {
    server.GetSecurityManager().RemovePermission(accountId, permissionPattern, messageId, sessionId)
}

function ChangeAdminPassword(pwd, newPwd) {
    server.GetSecurityManager().ChangeAdminPassword(pwd, newPwd, messageId, sessionId)
}

function GetResource(clientId, scope, query, idTokenUuid){
    var accessUuid = "" // not know by the client side.
    var result = server.GetOAuth2Manager().GetResource(clientId,scope, query, idTokenUuid, accessUuid, messageId, sessionId)
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

function Login(name, password, serverId) {
    var newSession = null
    newSession = server.GetSessionManager().Login(name, password, serverId, messageId, sessionId)
    return newSession
}

function GetActiveSessions() {
    var sessions = null
    sessions = server.GetSessionManager().GetActiveSessions()
    return sessions
}

function GetActiveSessionByAccountId(accountId) {
    var sessions = null
    sessions = server.GetSessionManager().GetActiveSessionByAccountId(accountId)
    return sessions
}

function UpdateSessionState(state) {
    var err = server.GetSessionManager().UpdateSessionState(state, messageId, sessionId)
    return err
}

function Logout(toCloseId) {
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