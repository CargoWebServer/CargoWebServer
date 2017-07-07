/**
 * Create and return a "version 4" RFC-4122 UUID string.
 */
function randomUUID() {
    var s = [], itoh = '0123456789ABCDEF';

    // Make array of random hex digits. The UUID only has 32 digits in it, but we
    // allocate an extra items to make room for the '-'s we'll be inserting.
    for (var i = 0; i < 36; i++)
        s[i] = Math.floor(Math.random() * 0x10);

    // Conform to RFC-4122, section 4.4
    s[14] = 4;  // Set 4 high bits of time_high field to version
    s[19] = (s[19] & 0x3) | 0x8;  // Specify 2 high bits of clock sequence

    // Convert to hex chars
    for (var i = 0; i < 36; i++)
        s[i] = itoh[s[i]];

    // Insert '-'s
    s[8] = s[13] = s[18] = s[23] = '-';

    return s.join('');
}

/**
 * Evaluate if an array contain a given element.
 * @param arr The target array.
 * @param obj The object to find.
 * @returns {boolean} True if the array contain the object.
 */
function contains(arr, obj) {
    var i = arr.length;
    while (i--) {
        if (arr[i] === obj) {
            return true;
        }
    }
    return false;
}

/**
 *  @constant {int}  Double data type identifier.
 */
Data_DOUBLE = 0
/**
 *  @constant {int}  Interger data type identifier.
 */
Data_INTEGER = 1
/**
 *  @constant {int}  String data type identifier.
 */
Data_STRING = 2
/**
 *  @constant {int}  Bytes data type identifier.
 */
Data_BYTES = 3
/**
 *  @constant {int}  Struct (JSON) data type identifier.
 */
Data_JSON_STR = 4
/**
 *  @constant {int}  Boolean data type identifier.
 */
Data_BOOLEAN = 5


/**
 * Creates a new RpcData.
 * @param variable The variable to create as RpcData
 * @param {string} variableType The type of the variable. Can be: DOUBLE, INTEGER, STRING, BYTES, JSON_STR, BOOLEAN
 * @param {string} variableName The name of the variable to create as RpcData. This parameter is optional
 * @param {string} typeName This is the name on the server side that must be use to interprest the data.
 * @returns {RpcData} The created RpcData or undefined if variableType was invalid
 */
function createRpcData(variable, variableType, variableName, typeName) {
    if (variableName == undefined) {
        variableName = "varName"
    }
    if (variableType == "DOUBLE") {
        variableType = Data_DOUBLE
        typeName = "double"
    } else if (variableType == "INTEGER") {
        variableType = Data_INTEGER
        typeName = "int"
    } else if (variableType == "STRING") {
        variableType = Data_STRING
        typeName = "string"
    } else if (variableType == "BYTES") {
        variableType = Data_BYTES
        typeName = "[]unit8"
    } else if (variableType == "JSON_STR") {
        variableType = Data_JSON_STR
        if (variable.stringify != undefined) {
            variableType = Data_JSON_STR
            typeName = variable.TYPENAME
            variable = variable.stringify()
        } else {
            variable = JSON.stringify(variable)
        }
    } else if (variableType == "BOOLEAN") {
        variableType = Data_BOOLEAN
        typeName = "bool"
    } else {
        return undefined
    }

    // Now I will create the rpc data.
    return new RpcData({ "name": variableName, "type": variableType, "dataBytes": variable, "typeName": typeName });
}