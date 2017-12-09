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
 * The JSON functions...
 */
var JSON = {}

JSON.stringify = function(obj){
    return stringify(obj)
}

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
        typeName = "Object"
    } else if (variableType == "BOOLEAN") {
        variableType = Data_BOOLEAN
        typeName = "bool"
    } else {
        console.log("undefined type: ", variableType, "for", variable)
        return undefined
    }

    // Now I will create the rpc data.
    return new RpcData({ "name": variableName, "type": variableType, "dataBytes": variable, "typeName": typeName });
}

//////////////////////////////////// XML/SQL type ////////////////////////////////////

/**
 * Dertermine if the value is a base type.
 */
function isXsBaseType(fieldType) {
    return isXsId(fieldType) || isXsRef(fieldType) || isXsInt(fieldType) || isXsString(fieldType) || isXsBinary(fieldType) || isXsNumeric(fieldType) || isXsBoolean(fieldType) || isXsDate(fieldType) || isXsTime(fieldType) || isXsMoney(fieldType)
}

/**
 * Helper function use to dertermine if a XS type must be considere integer.
 */
function isXsInt(fieldType) {
    if (endsWith(fieldType, "byte")|| endsWith(fieldType, "long") || endsWith(fieldType, "int") || endsWith(fieldType, "integer") || endsWith(fieldType, "short")  // XML
        || endsWith(fieldType, "unsignedInt") || endsWith(fieldType, "unsignedBtype") || endsWith(fieldType, "unsignedShort") || endsWith(fieldType, "unsignedLong")  // XML
        || endsWith(fieldType, "negativeInteger") || endsWith(fieldType, "nonNegativeInteger") || endsWith(fieldType, "nonPositiveInteger") || endsWith(fieldType, "positiveInteger") // XML
        || endsWith(fieldType, "tinyint") || endsWith(fieldType, "smallint") || endsWith(fieldType, "bigint"))// SQL
    {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere String.
 */
function isXsString(fieldType) {
    if (endsWith(fieldType, "string")
        || endsWith(fieldType, "Name") || endsWith(fieldType, "QName") || endsWith(fieldType, "NMTOKEN")  // XML
        || endsWith(fieldType, "gDay") || endsWith(fieldType, "gMonth") || endsWith(fieldType, "gMonthDay") || endsWith(fieldType, "gYear") // XML
        || endsWith(fieldType, "gYearMonth") || endsWith(fieldType, "token") || endsWith(fieldType, "normalizedString") || endsWith(fieldType, "hexBinary") // XML
        || endsWith(fieldType, "language") || endsWith(fieldType, "NMTOKENS") || endsWith(fieldType, "NOTATION")  // XML
        || endsWith(fieldType, "char") || endsWith(fieldType, "nchar") || endsWith(fieldType, "varchar") // SQL
        || endsWith(fieldType, "nvarchar") || endsWith(fieldType, "text") || endsWith(fieldType, "ntext") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere binary value.
 */
function isXsBinary(fieldType) {
    if (endsWith(fieldType, "base64Binary") // XML
        || endsWith(fieldType, "varbinary") || endsWith(fieldType, "binary") // SQL
        || endsWith(fieldType, "image") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere numeric value.
 */
function isXsNumeric(fieldType) {
    if (endsWith(fieldType, "double") || endsWith(fieldType, "decimal") || endsWith(fieldType, "float") // XML
        || endsWith(fieldType, "numeric") || endsWith(fieldType, "real") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere boolean value.
 */
function isXsBoolean(fieldType) {
    if (endsWith(fieldType, "boolean") // XML
        || endsWith(fieldType, "bit")  // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere date value.
 */
function isXsDate(fieldType) {
    if (endsWith(fieldType, "date") || endsWith(fieldType, "dateTime") // XML
        || endsWith(fieldType, "datetime2") || endsWith(fieldType, "smalldatetime") || endsWith(fieldType, "datetimeoffset") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere time value.
 */
function isXsTime(fieldType) {
    if (endsWith(fieldType, "time") // XML
        || endsWith(fieldType, "timestampNumeric") || endsWith(fieldType, "timestamp") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere money value.
 */
function isXsMoney(fieldType) {
    if (
        endsWith(fieldType, "money") || endsWith(fieldType, "smallmoney") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere id value.
 */
function isXsId(fieldType) {
    if (
        endsWith(fieldType, "ID") || endsWith(fieldType, "NCName") // XML
        || endsWith(fieldType, "uniqueidentifier") // SQL
    ) {
        return true
    }
    return false
}

/**
 * Helper function use to dertermine if a XS type must be considere id value.
 */
function isXsRef(fieldType) {
    if (
        endsWith(fieldType, "anyURI") || endsWith(fieldType, "IDREF") // XML
    ) {
        return true
    }
    return false
}

exports.JSON = JSON
exports.isXsBaseType = isXsBaseType
exports.isXsInt = isXsInt
exports.isXsString = isXsString
exports.isXsBinary = isXsBinary
exports.isXsNumeric = isXsNumeric
exports.isXsBoolean = isXsBoolean
exports.isXsDate = isXsDate
exports.isXsTime = isXsTime
exports.isXsMoney = isXsMoney
exports.isXsId = isXsId
exports.isXsRef = isXsRef


