////////////////////////////////// Server //////////////////////////

// SetInterval function like the one on the widows object in the browser.
// exported by default. resetInterval are also define.
function setInterval(callback, interval) {
    setInterval_(callback, interval, sessionId)
}

// SetTimeout function like the one on the widows object in the browser.
// exported by default. resetTimeout are also define.
function setTimeout(callback, timeout) {
    setTimeout_(callback, timeout, sessionId)
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

// Exported function...
exports.ExcelToCsv = ExcelToCsv
