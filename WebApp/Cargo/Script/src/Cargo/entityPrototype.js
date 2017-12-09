/**
 * Global variable where the entities will be keep in memory.
 */
require("./eventHub")

 // Store the prototypes.
var entityPrototypes = {};

// Store the entities.
var entities = {};

/**
 * The entity prototype define the schema of an entity. Entity prototype
 * are usefull to make entity persistent, to export and import schema from
 * other data source like xml schema. It's also possible to create at runtime
 * new kind of entity and make query over it.
 * </br>note: Fields, FieldsType, FieldsDocumentation, FieldsNillable and FieldsOrder 
 * must have the same number of elements.
 * @constructor 
 */
var EntityPrototype = function () {
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

    if (object == null || object.TypeName == undefined) {
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
	

    // The type name.
    this.TypeName = object.TypeName

    // The package will be an object on the global scope.
    this.PackageName = object.TypeName.split(".")[0]

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

    // other standard fields.
    this.appendField("childsUuid", "[]xs.string", false, this.Fields.length, false, "the array of child entities.", "[]")
    this.appendField("referenced", "[]Server.EntityRef", false, this.Fields.length, false, "The field documentation.", "[]")

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
        delete entityPrototypes[evt.dataMap.prototype.TypeName]
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

// List of exports...
// functions
exports.setEntityPrototype = setEntityPrototype
// variables
exports.entityPrototypes = entityPrototypes
exports.entities = entities
// class
exports.EntityPrototype = EntityPrototype
exports.EntityPrototypeManager = EntityPrototypeManager

