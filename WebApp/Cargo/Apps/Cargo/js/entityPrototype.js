/**
 * The entity prototype define the schema of an entity. Entity prototype
 * are usefull to make entity persistent, to export and import schema from
 * other data source like xml schema. It's also possible to create at runtime
 * new kind of entity and make query over it.
 * </br>note: Fields, FieldsType, FieldsDocumentation, FieldsNillable and FieldsOrder 
 * must have the same number of elements. * @constructor 
 */
var EntityPrototype = function () {
    /**
     * Uniquely indentify entity prototype.
     */
    this.UUID = randomUUID()

    /**
     * The typename of the prototype itself.
     */
    this.TYPENAME = "Server.EntityPrototype"

    /**
     * @property {string} TypeName The type name of the entity must be write 'packageName.className'
     */
    this.TypeName = "" // string

    /** 
     * @property {string} Documentation The entity type documentation.
     */
    this.Documentation = ""

    /** 
     * @property {boolean} IsAbstract True if the entity prototype is an abstrac class.
     */
    this.IsAbstract = false

    /**
     * @property ListOf The prototype consiste of a list of other type
     */
    this.ListOf = ""

    /**
     * @property SubstitutionGroup The list of substitution group, must be use in abstract class.
     */
    this.SubstitutionGroup = []

    /**
     * @property SuperTypeNames The list of super class name.
     */
    this.SuperTypeNames = [] // string

    /**
     *  @property Restrictions The field restrictions.
     */
    this.Restrictions = [] // Restriction

    /**
     * @property Ids The list of fields used as ids.
     */
    this.Ids = [] // string

    /**
     * @property Indexs The list of fields used as index.
     */
    this.Indexs = [] // string

    ////////////////////////////////////////////////////////////////////////////
    //  Field's properties.
    ////////////////////////////////////////////////////////////////////////////

    /** 
     * @property Fields The fields name 
     */
    this.Fields = []

    /**
     * @property FieldsDocumentation The field documentation
     */
    this.FieldsDocumentation = []

    /**
     * @property FieldsType The fields type
     */
    this.FieldsType = [] // string

    /**
     * @property FieldsVisibility The fields visibility
     */
    this.FieldsVisibility = [] // bool

    /**
     * @property FieldsNillable If the field can be nil
     */
    this.FieldsNillable = [] // bool

    /**
     * @property The fields default value. Use to initialyse a new object with a given value.
     */
    this.FieldsOrder = []

    /**
     * @property FieldsOrder The field order
    */
    this.FieldsDefaultValue = [] // int

    /**
     * @property Contain the list of fields to delete
    */
    this.FieldsToDelete = [] // int

    /**
     * @property Contain the list of fields to rename
    */
    this.FieldsToUpdate = [] // int

    return this
}

/**
 * Return the list of field to be used as title.
 */
EntityPrototype.prototype.getTitles = function () {
    // The get title default function... can be overload.
    var titles = []

    // The ids
    for (var i = 1; i < this.Ids.length; i++) {
        var fieldIndex = this.getFieldIndex(this.Ids[i])
        var field = this.Fields[fieldIndex]
        if (this.FieldsVisibility[fieldIndex] == true) {
            titles.push(this[field])
        }
    }

    // The indexs
    for (var i = 1; i < this.Indexs.length; i++) {
        var fieldIndex = this.getFieldIndex(this.Indexs[i])
        var field = this.Fields[fieldIndex]
        if (this.FieldsVisibility[fieldIndex] == true) {
            titles.push(this[field])
        }
    }
    return titles
}

/**
 * Create a new class form json object.
 * @param {object} object The object that regroup the prototype properties.
 */
EntityPrototype.prototype.init = function (object) {
    if (object == null || object.TypeName == undefined || object.Fields == null) {
        return
    }

    if (object.FieldsNillable == null) {
        object.FieldsNillable = []
        for (var i = 0; i < object.Fields.length; i++) {
            object.FieldsNillable.push(false)
        }
    }

    if (object.FieldsDocumentation == null) {
        object.FieldsDocumentation = []
        for (var i = 0; i < object.Fields.length; i++) {
            object.FieldsDocumentation.push("")
        }
    }

    if (object.FieldsDefaultValue == null) {
        object.FieldsDefaultValue = []
        for (var i = 0; i < object.Fields.length; i++) {
            object.FieldsDefaultValue.push("")
        }
    }

    // The type name.
    this.TypeName = object.TypeName

    // The package will be an object on the global scope.
    this.PackageName = object.TypeName.split(".")[0]
    if (window[this.PackageName] == undefined) {
        window[this.PackageName] = eval(this.PackageName + " = {}")
    }

    // The type name.
    this.ClassName = object.TypeName.substring(object.TypeName.indexOf(".") + 1)

    // The object ids
    this.appendIds(object.Ids)

    // The object indexs
    this.appendIndexs(object.Indexs)

    // Now the fields.
    if (object.Fields != null && object.FieldsType != null && object.FieldsVisibility != null && object.FieldsOrder != null) {

        if (!contains(object.Fields, "ParentLnk")) {
            object.Fields.unshift("ParentLnk")
            object.FieldsType.unshift("xs.string")
            object.FieldsVisibility.unshift(false)
            object.FieldsNillable.unshift(false)
            object.FieldsDocumentation.unshift("Relation with it parent.")
            object.FieldsOrder.push(object.FieldsOrder.length)
            object.FieldsDefaultValue.unshift("")
        }

        // Append parent uuid if none is define.
        if (!contains(object.Fields, "ParentUuid")) {
            object.Fields.unshift("ParentUuid")
            object.FieldsType.unshift("xs.string")
            object.FieldsVisibility.unshift(false)
            object.FieldsNillable.unshift(false)
            object.FieldsDocumentation.unshift("The parent object UUID")
            object.FieldsOrder.push(object.FieldsOrder.length)
            object.FieldsDefaultValue.unshift("")
        }

        // Append the uuid if none is define.
        if (!contains(object.Fields, "UUID")) {
            // append the uuid...
            object.Fields.unshift("UUID")
            object.FieldsType.unshift("xs.string")
            object.FieldsVisibility.unshift(false)
            object.FieldsOrder.push(object.FieldsOrder.length)
            object.FieldsNillable.unshift(false)
            object.FieldsDocumentation.unshift("The object UUID")
            object.FieldsDefaultValue.unshift("")
            object.Ids.unshift("UUID")
        }

        for (var i = 0; i < object.Fields.length; i++) {
            this.appendField(object.Fields[i], object.FieldsType[i], object.FieldsVisibility[i], object.FieldsOrder[i], object.FieldsNillable[i], object.FieldsDocumentation[i], object.FieldsDefaultValue[i])
            if (object.Fields[i] == "UUID") {
                if (!contains(this.Ids, "UUID")) {
                    this.Ids.unshift("UUID")
                }
            } else if (object.Fields[i] == "ParentUuid") {
                if (!contains(this.Indexs, "ParentUuid")) {
                    this.Indexs.unshift("ParentUuid")
                }
            } else if (object.Fields[i] == "ParentLnk") {
                if (!contains(this.Indexs, "ParentLnk")) {
                    this.Indexs.unshift("ParentLnk")
                }
            }
        }

    } else {
        console.log(object.TypeName + " has no fields!!!")
    }

    // If the object is abstract
    this.IsAbstract = object.IsAbstract

    // Now the restriction
    this.Restrictions = object.Restrictions


    // The list of substitution type.
    if (object.SubstitutionGroup != undefined) {
        this.SubstitutionGroup = object.SubstitutionGroup
    }

    // The list of Supertype 
    if (object.SuperTypeNames != undefined) {
        this.SuperTypeNames = object.SuperTypeNames
    }

    // if the type is a collection.
    if (object.ListOf != undefined) {
        this.ListOf = object.ListOf
    }

    if (object.FieldsDefaultValue != undefined) {
        this.FieldsDefaultValue = object.FieldsDefaultValue
    }

    // Generate the class code.
    this.generateConstructor()
}


/**
 * This function generate the js class base on the entity prototype.
 */
EntityPrototype.prototype.generateConstructor = function () {
    if (this.ClassName.indexOf(" ") > 0) {
        return
    }

    var constructorSrc = this.PackageName + " || {};\n"

    var packageName = this.PackageName
    var classNames = this.ClassName.split(".")

    for (var i = 0; i < classNames.length - 1; i++) {
        packageName += "." + classNames[i]
        constructorSrc += packageName + " = " + packageName + " || {};\n"
    }

    // I will create the object constructor from the information
    // of the fields.
    constructorSrc += this.PackageName + "." + this.ClassName + " = function(values){\n"

    // Common properties share by all entity.
    constructorSrc += " this.__class__ = \"" + this.PackageName + "." + this.ClassName + "\"\n"
    constructorSrc += " this.TYPENAME = \"" + this.TypeName + "\"\n"
    constructorSrc += " if(values==undefined){\n"
    constructorSrc += "     this.UUID = undefined\n"
    constructorSrc += "     this.ParentUuid = \"\"\n"
    constructorSrc += "     this.ParentLnk = \"\"\n"
    constructorSrc += " }else{\n"
    constructorSrc += "     this.UUID = values.UUID\n"
    constructorSrc += "     this.ParentUuid = values.ParentUuid\n"
    constructorSrc += "     this.ParentLnk = values.ParentLnk\n"
    constructorSrc += " }\n"

    constructorSrc += " this.childsUuid = []\n"
    constructorSrc += " this.references = []\n"
    constructorSrc += " this.NeedSave = true\n"
    constructorSrc += " this.IsInit = false\n"
    constructorSrc += " this.exist = false\n"
    constructorSrc += " this.initCallback = undefined\n"
    constructorSrc += " this.panel = null\n"

    // Remove space accent '' from the field name
    function normalizeFieldName(fieldName) {
        // TODO make distinctive..
        fieldName = fieldName.replaceAll(" ", "_")
        fieldName = fieldName.replaceAll("'", "")
        return fieldName
    }

    // Fields.
    for (var i = 3; i < this.Fields.length; i++) {
        var fieldName = normalizeFieldName(this.Fields[i])
        if (this.FieldsDefaultValue[i] != undefined) {
            // In case of default values...
            if (this.FieldsType[i].startsWith("[]")) {
                constructorSrc += " this." + fieldName + " = []\n"
            } else if (isXsString(this.FieldsType[i]) || isXsRef(this.FieldsType[i]) || isXsId(this.FieldsType[i])) {
                constructorSrc += " this." + fieldName + " = \"" + this.FieldsDefaultValue[i] + "\"\n"
            } else {
                if (this.FieldsType[i].startsWith("xs.") || this.FieldsType[i].startsWith("sqltypes.")) {
                    if (this.FieldsDefaultValue[i].length != 0) {
                        constructorSrc += " this." + fieldName + " = " + this.FieldsDefaultValue[i] + "\n"
                    } else if (isXsNumeric(this.FieldsType[i])) {
                        constructorSrc += " this." + fieldName + " = 0.0\n"
                    } else if (isXsBoolean(this.FieldsType[i])) {
                        constructorSrc += " this." + fieldName + " = false\n"
                    } else if (isXsString(this.FieldsType[i])) {
                        constructorSrc += " this." + fieldName + " = \"\"\n"
                    } else {
                        constructorSrc += " this." + fieldName + " = null\n"
                    }
                } else if (this.FieldsType[i].startsWith("enum:")) {
                    constructorSrc += " this." + fieldName + " = 1\n"
                } else {
                    if (this.FieldsType[i].endsWith(":Ref")) {
                        constructorSrc += " this." + fieldName + " = null\n"
                    } else {

                        if (this.FieldsNillable[i] == false) {
                            constructorSrc += " this." + fieldName + " = new " + this.FieldsType[i] + "()\n"
                        } else {
                            constructorSrc += " if(getBaseTypeExtension(\"" +  this.FieldsType[i] +"\").startsWith(\"xs.\")){\n"
                            constructorSrc += "     this." + fieldName + " = new " + this.FieldsType[i] + "()\n"
                            constructorSrc += "     this." + fieldName + ".getParent = function(parent){\n"
                            constructorSrc += "         return function(){return parent}\n"
                            constructorSrc += "     }(this)\n"
                            constructorSrc += " }else{\n"
                            constructorSrc += "     this." + fieldName + " = null\n"
                            constructorSrc += " }\n"
                        }
                    }
                }
            }
        } else if (this.FieldsType[i].startsWith("[]")) {
            constructorSrc += " this." + fieldName + " = []\n"
        } else {
            if (isXsString(this.FieldsType[i]) || isXsRef(this.FieldsType[i]) || isXsId(this.FieldsType[i])) {
                constructorSrc += " this." + fieldName + " = \"\"\n"
            } else if (isXsInt(this.FieldsType[i])) {
                constructorSrc += " this." + fieldName + " = 0\n"
            } else if (isXsNumeric(this.FieldsType[i])) {
                constructorSrc += " this." + fieldName + " = 0.0\n"
            } else if (isXsDate(this.FieldsType[i])) {
                constructorSrc += " this." + fieldName + " = moment().unix()\n"
            } else if (isXsBoolean(this.FieldsType[i])) {
                constructorSrc += " this." + fieldName + " = false\n"
            } else if (this.FieldsType[i].startsWith("enum:")) {
                constructorSrc += " this." + fieldName + " = 1\n"
            } else {
                // Object here.
                if (this.FieldsType[i].startsWith("enum:")) {
                    constructorSrc += " this." + fieldName + " = 0\n"
                } else {
                    if (this.isNillable[i]) {
                        constructorSrc += " this." + fieldName + " = null\n"
                    } else {
                        constructorSrc += " this." + fieldName + " = new " + this.FieldsType[i] + "()\n"
                    }
                }
            }
        }
    }

    // Now the stringify function.
    constructorSrc += " this.stringify = function(){\n"
    constructorSrc += "       resetObjectValues(this)\n"
    constructorSrc += "       var cache = [];\n"
    constructorSrc += "       var entityStr = JSON.stringify(this, function(key, value) {\n"
    constructorSrc += "           if (typeof value === 'object' && value !== null) {\n"
    constructorSrc += "               if (cache.indexOf(value) !== -1) {\n"
    constructorSrc += "                   // Circular reference found, discard key\n"
    constructorSrc += "                   return;\n"
    constructorSrc += "               }\n"
    constructorSrc += "               // Store value in our collection\n"
    constructorSrc += "               cache.push(value);\n"
    constructorSrc += "           }\n"
    constructorSrc += "           return value;\n"
    constructorSrc += "       });\n"
    constructorSrc += "       cache = null; // Enable garbage collection\n"
    constructorSrc += "       setObjectValues(this)\n"
    constructorSrc += "       return entityStr\n"
    constructorSrc += "   }\n"

    // The get parent function
    constructorSrc += " this.getParent = function(){\n"
    constructorSrc += "       return entities[this.ParentUuid]\n"
    constructorSrc += "  }\n"

    // The get parent function
    constructorSrc += " this.getTypeName = function(){\n"
    constructorSrc += "       return getEntityPrototype(this.TYPENAME).TypeName\n"
    constructorSrc += "  }\n"

    constructorSrc += " this.getPrototype = function(){\n"
    constructorSrc += "       return getEntityPrototype(this.TYPENAME)\n"
    constructorSrc += "  }\n"

    // The setter function.
    for (var i = 0; i < this.Fields.length; i++) {
        if (!this.FieldsType[i].startsWith("xs.") && !this.FieldsType[i].startsWith("[]xs.")) {
            // So its not a basic type.
            constructorSrc += " this.set" + normalizeFieldName(this.Fields[i]).replace("M_", "").capitalizeFirstLetter() + " = function(value){\n"
            constructorSrc += "     appendObjectValue(this,\"" + normalizeFieldName(this.Fields[i]) + "\", value)\n"
            constructorSrc += " }\n"
        }
    }

    // The remove function.
    for (var i = 0; i < this.Fields.length; i++) {
        if (!this.FieldsType[i].startsWith("xs.") && !this.FieldsType[i].startsWith("[]xs.")) {
            // So its not a basic type.
            constructorSrc += " this.remove" + normalizeFieldName(this.Fields[i]).replace("M_", "").capitalizeFirstLetter() + " = function(value){\n"
            constructorSrc += "     removeObjectValue(this,\"" + normalizeFieldName(this.Fields[i]) + "\", value)\n"
            constructorSrc += " }\n"
        }
    }

    // The get title default function... can be overload.
    constructorSrc += " this.getTitles = function(){\n"
    var fields = ""
    // The ids
    for (var i = 1; i < this.Ids.length; i++) {
        var fieldIndex = this.getFieldIndex(this.Ids[i])
        var field = this.Fields[fieldIndex]
        if (this.FieldsVisibility[fieldIndex] == true) {
            if (fields.length > 0) {
                fields = fields + ", this." + field
            } else {
                fields = "this." + field
            }
        }
    }
    // The indexs
    for (var i = 1; i < this.Indexs.length; i++) {
        var fieldIndex = this.getFieldIndex(this.Indexs[i])
        var field = this.Fields[fieldIndex]
        if (this.FieldsVisibility[fieldIndex] == true) {
            if (fields.length > 0) {
                fields = fields + ",this." + field
            } else {
                fields = "this." + field
            }
        }
    }
    constructorSrc += "     return [" + fields + "]\n"
    constructorSrc += " }\n"

    // Keep the reference on the entity prototype.
    // The class level.
    constructorSrc += " return this\n"
    constructorSrc += "}\n\n"

    constructorSrc += this.PackageName + "." + this.ClassName + ".prototype.init = function(object, lazy){\n"
    // First of all i will set reference in the result.
    constructorSrc += "   this.TYPENAME = object.TYPENAME\n"
    constructorSrc += "   this.UUID = object.UUID\n"
    constructorSrc += "   this.IsInit = false\n"
    constructorSrc += "   setObjectValues(this, object, lazy)\n"
    constructorSrc += "}\n\n"

    // Set the function.
    eval(constructorSrc)
}

/**
 * Create a new prototypeField.
 * @param {string} name The field name.
 * @param {string} name The field type name.
 * @param {boolean} isVisible True, if the field is visible.
 * @param {int} order The order the field will be return, usefull to display.
 */
EntityPrototype.prototype.appendField = function (name, typeName, isVisible, order, isNillable, documentation, defaultValue) {
    // Set the field name.
    if (!contains(this.Fields, name)) {
        this.Fields.push(name)

        // Set the field type name
        this.FieldsType.push(typeName)

        // Set the field visibility
        this.FieldsVisibility.push(isVisible)
        // And the order (index in the list of fields.)
        this.FieldsOrder.push(parseInt(order))

        this.FieldsNillable.push(isNillable)

        this.FieldsDocumentation.push(documentation)

        this.FieldsDefaultValue.push(defaultValue)
    }
}

/**
 * Append the list of indexs
 * @param indexs The list of indexs to append.
 */
EntityPrototype.prototype.appendIndexs = function (indexs) {
    if (indexs != null) {
        for (var i = 0; i < indexs.length; i++) {
            this.Indexs.push(indexs[i])
        }
    }
}

/**
 * Append the list of id's
 * @param ids The list of ids to append.
 */
EntityPrototype.prototype.appendIds = function (ids) {
    if (ids != null) {
        for (var i = 0; i < ids.length; i++) {
            this.Ids.push(ids[i])
        }
    }
}

/**
 * Retreive the index of a given field in the prototype.
 * @param field The field we looking for.
 * @returns The index of the field or -1 if the field is not there.
 */
EntityPrototype.prototype.getFieldIndex = function (field) {
    for (var i = 0; i < this.Fields.length; i++) {
        if (this.Fields[i] == field) {
            return i
        }
    }
    return -1
}

/**
 * That class is use to manage entity prototype event.
 */
var EntityPrototypeManager = function () {
    if (server == undefined) {
        return
    }
    EventHub.call(this, PrototypeEvent)

    return this
}

EntityPrototypeManager.prototype = new EventHub(null);
EntityPrototypeManager.prototype.constructor = EntityPrototypeManager;

/**
 * Overload the onEvent function to set the new  prototype object on the local map
 * before propagate the event to listener.
 */
EntityPrototypeManager.prototype.onEvent = function (evt) {
    if (evt.code == NewPrototypeEvent || evt.code == UpdatePrototypeEvent) {
        // Set the prototype.
        var prototype = new EntityPrototype()
        prototype.init(evt.dataMap.prototype)
        setEntityPrototype(prototype)
        evt.dataMap.prototype = prototype
    } else if (evt.code == DeletePrototypeEvent) {
        // Remove it from the map
        deleteEntityPrototype(evt.dataMap.prototype.TypeName)
    }

    // Call the regular function.
    EventHub.prototype.onEvent.call(this, evt);
}

/**
 * Use that function to set value of the global map.
 * @param {*} prototype 
 */
function setEntityPrototype(prototype) {
    entityPrototypes[prototype.TypeName] = prototype
}

/**
 * Use that function to set value of the global map.
 * @param {*} prototype 
 */
function getEntityPrototype(typeName) {
    return entityPrototypes[typeName]
}

/**
 * Use that function to set value of the global map.
 * @param {*} prototype 
 */
function deleteEntityPrototype(typeName) {
    delete entityPrototypes[typeName]
}