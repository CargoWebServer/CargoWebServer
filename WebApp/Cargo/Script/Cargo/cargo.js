////////////////////////////////// Server //////////////////////////
function SetRootPath(path) {
    // Call set root path on the server...
    GetServer().SetRootPath(path)
}

function Connect(address) {
    GetServer().Connect(address)
}

function Stop() {
    GetServer().Stop()
}

/////////////////////////////////////// Other script /////////////////////////////////////////
/**
 * Read excel file and get it content as comma separated values.
 */
function ExcelToCsv(filePath) {
    var val = filePath.split("\\")
    var fileName = val[val.length - 1].split(".")[0]
    var outputFile = GetServer().GetConfigurationManager().GetTmpPath() + "/" + fileName + ".csv"
    // C:\\ instead of C:/
    outputFile = outputFile.replace("/", "\\", -1)
    GetServer().ExecuteVbScript("xlsx2csv.vbs", [filePath, outputFile], messageId, sessionId)

    // I will read the file content.
    var results = {"data" : GetServer().GetFileManager().ReadCsvFile(outputFile, messageId, sessionId), "outputFile":outputFile}

    return results
}
