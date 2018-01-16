/**
 * That function is use to create Entity from CSV file.
 * @param {*} filePath The csv file path
 * @param {*} typeName The type name of the entity to initialyse
 * @param {*} mappings The mappings between the entity fields and the csv column
 * @param {*} parentEntity The parent entity (at root only)
 * @param {*} parentLnk The attribute name with it parent.
 */
function importEntityFromCSV(filePath, mappings, parentEntity, parentLnk) {
    // Read the csv data.
    var importInfos = new ImportInfos(filePath)
    function createMapping(typeName, mapping) {
        var prototype = getEntityPrototype(typeName)
        var mappingInfo = new MappingInfos(typeName, mapping.isValid,
            parentEntity, parentLnk)
        if (prototype != null) {
            // I will create the mapping for each field
            for (var key in mapping) {
                // Here I got the field
                if (key != "isValid" && key != "TYPENAME") {
                    if (isObject(mapping[key])) {
                        var field = Object.keys(mapping[key])[0]
                        var index = prototype.getFieldIndex(field)
                        var fieldTypeName = prototype.FieldsType[index].replace("[]", "").replace(":Ref", "")
                        var fieldPrototype = getEntityPrototype(fieldTypeName)
                        if (fieldPrototype != null) {
                            var subMappingInfo = createMapping(fieldTypeName, mapping[key][field])
                            mappingInfo.fieldsMapping[key] = new FieldMappingInfos(field, key, subMappingInfo, null, field)
                        }
                    } else {
                        mappingInfo.fieldsMapping[key] = new FieldMappingInfos(mapping[key], key)
                    }
                }
            }
        }
        return mappingInfo
    }

    for (var i = 0; i < mappings.length; i++) {
        // Create the mapping infos...
        var mappingInfo = createMapping(mappings[i].TYPENAME, mappings[i])
        importInfos.mappingInfos.push(mappingInfo)
    }

    importInfos.import()
}

/**
 * That file contain informations to import Data.
 */
var ImportInfos = function (dataFilePath) {

    /** The path on the server were the Data is. */
    this.dataFilePath = dataFilePath

    /** A csv file can contain informations about more than on types. */
    this.mappingInfos = []

    return this
}

ImportInfos.prototype.import = function () {
    var initEntityPrototype = function (index, mappingInfos, callback) {
        server.entityManager.getEntityPrototype(mappingInfos[index].typeName, mappingInfos[index].typeName.split(".")[index],
            function (result, caller) {
                if (caller.index < caller.total) {
                    // call the next value.
                    caller.index = caller.index++
                    initEntityPrototype(caller.index, caller.mappingInfos, caller.callback)
                } else {
                    if (caller.callback != undefined) {
                        caller.callback()
                    }
                }
            },
            function (errObj) {

            }, { "index": index, "mappingInfos": mappingInfos, "callback": callback })
    }

    if (this.mappingInfos.length > 0) {
        initEntityPrototype(0, this.mappingInfos,
            // That function is call when all Mapping infos prototypes are initialysed.
            function (importInfos) {
                return function () {
                    if (importInfos.dataFilePath.endsWith(".csv")) {
                        // In that case the data to import are in csv format.
                        server.fileManager.readCsvFile(importInfos.dataFilePath,
                            function (results, caller) {
                                var header = results[0][0]
                                for (var i = 0; i < importInfos.mappingInfos.length; i++) {
                                    for (var j = 1; j < results[0].length - 1; j++) {
                                        importInfos.mappingInfos[i].import(header, results[0][j])
                                        if (importInfos.mappingInfos[i].isValid(results[0][j])) {
                                            if (importInfos.mappingInfos[i].parentEntity == null) {
                                                server.entityManager.saveEntity(importInfos.mappingInfos[i].object)
                                            } else {
                                                var parentUuid = importInfos.mappingInfos[i].parentEntity.UUID
                                                var parentLnk = importInfos.mappingInfos[i].parentLnk
                                                var entity = importInfos.mappingInfos[i].object
                                                entity.ParentUuid = parentUuid
                                                entity.ParentLnk = parentLnk
                                                
                                                server.entityManager.createEntity(parentUuid, parentLnk, entity.TYPENAME, "", entity,
                                                // success callback
                                                function(entity){
                                                    console.log(entity)
                                                },
                                                // error callback
                                                function(errObj){
                                                    console.log(errObj)
                                                },
                                                {})
                                            }
                                        }
                                    }
                                }
                            },
                            function (errorObj, caller) {

                            }, {})
                    }
                }
            }(this))
    }
}

var MappingInfos = function (typeName, isValid, parentEntity, parentLnk) {

    // If there is a parent entity.
    this.parentEntity = parentEntity

    // The lnk in it's parent.
    this.parentLnk = parentLnk

    // Is valid is a function to be evaluated to dermine if
    // the data represent the type. The paremeter is the array
    // of value of the object.
    this.isValid = isValid

    // The entity prototype type name to initialyse with data.
    this.typeName = typeName

    // That contain the assciation from 
    this.fieldsMapping = {}

    // The object.
    this.object = null

    return this
}

MappingInfos.prototype.import = function (header, data) {
    // So here I will import data.
    this.object = eval("new " + this.typeName + "()")
    var prototype = getEntityPrototype(this.typeName)
    for (var i = 0; i < header.length; i++) {
        // Now I will get the field mapping for that column
        var fieldMapping = this.fieldsMapping[header[i]]
        if (fieldMapping != null) {
            var index = prototype.getFieldIndex(fieldMapping.field)
            if (index != -1) {
                var fieldType = prototype.FieldsType[index]
                if (fieldMapping.subMappingInfos != null) {
                    // Another entity here.
                    fieldMapping.subMappingInfos.import(header, data)

                    if (fieldType.startsWith("[]")) {
                        if (this.object[fieldMapping.field] == undefined) {
                            this.object[fieldMapping.field] = []
                        }
                        this.object[fieldMapping.field].push(fieldMapping.subMappingInfos.object)
                    } else {
                        this.object[fieldMapping.field] = fieldMapping.subMappingInfos.object
                    }
                    fieldMapping.subMappingInfos.object.ParentLnk = fieldMapping.field

                } else {
                    // Not an entity.
                    if (fieldType.startsWith("[]")) {
                        if (this.object[fieldMapping.field] == undefined) {
                            this.object[fieldMapping.field] = []
                        }
                        this.object[fieldMapping.field].push(data[i])
                    } else {
                        this.object[fieldMapping.field] = data[i]
                    }
                }
            }
        }
    }
}

var FieldMappingInfos = function (field, target, subMappingInfos, transform) {
    // This path to access the value in the object
    // for example toto.titi.tata = target
    this.field = field

    // If target is a reference to another objet the 
    // taranformation will be the initialisation of that reference.
    this.transform = transform

    // Is targeted field value.
    this.target = target

    // If there is another objet and not a simple field.
    this.subMappingInfos = subMappingInfos

    return this
}
