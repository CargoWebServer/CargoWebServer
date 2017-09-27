/*
 * (C) Copyright 2016 Mycelius SA (http://mycelius.com/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * @fileOverview Entities related functionalities.
 * @author Dave Courtois
 * @version 1.0
 */

/**
  Restriction are expression defining limitation on the
  range of value that a variable can take. Type of restriction
  are :
    <ul>
        <li>Enumeration</li>
        <li>FractionDigits</li>
        <li>Length</li>
        <li>MaxExclusive</li>
        <li> MaxInclusive</li>
        <li>MaxLength</li>
        <li>MinExclusive</li>
        <li> MinInclusive</li>
        <li>MinLength</li>
        <li>Pattern</li>
        <li>TotalDigits</li>
        <li>WhiteSpace</li>
    </ul>
    */
var Restriction = function () {
    // The the of the restriction (Facet)
    this.Type

    // The value.
    this.Value
}

/**
 * Append a new object value into an entity.
 */
function appendObjectValue(object, field, value) {
    var prototype = entityPrototypes[object.TYPENAME]
    var fieldIndex = prototype.getFieldIndex(field)
    var prototype = entityPrototypes[object.TYPENAME]
    var fieldType = prototype.FieldsType[fieldIndex]
    var isArray = fieldType.startsWith("[]")
    var isRef = fieldType.endsWith(":Ref")
    if (fieldIndex > -1) {
        if (isArray) {
            var index = 0
            var isExist = false
            // Create an array if is not exist.
            if (object[field] == undefined) {
                object[field] = []
            } else {
                for (var i = 0; i < object[field].length; i++) {
                    index = i
                    if (object[field][i].UUID == value.UUID) {
                        if (object[field][i].IsInit && value.IsInit) {
                            isExist = true
                        }
                        break
                    }
                }
            }

            // Set the reference in case of reference
            if (isRef) {
                setRef(object, field, value.UUID, true)
            } else {
                // Append or replace the value in case of an array.
                if (!isExist) {
                    index = object[field].length
                    object[field].push(value)
                } else {
                    object[field][index] = value
                }
            }

        } else {
            object[field] = value
            if (isRef) {
                setRef(object, field, value.UUID, false)
            }
        }
    }

    // Set need save to the object.
    object["NeedSave"] = true

    // Can be usefull to intercept change event...
    if (object.onChange != undefined) {
        object.onChange(object)
    }
}

/**
 * Remove an object from a given object.
 */
function removeObjectValue(object, field, value) {
    var prototype = entityPrototypes[object.TYPENAME]
    var index = prototype.getFieldIndex(field)
    if (index > -1) {
        var fieldType = prototype.FieldsType[index]
        var isArray = fieldType.startsWith("[]")
        var isRef = fieldType.endsWith(":Ref")

        if (isArray) {
            // Here the entity is an array.
            for (var i = 0; i < object[field].length; i++) {
                var uuid = ""
                if (isObject(object[field][i])) {
                    uuid = object[field][i].UUID
                } else {
                    uuid = object[field][i]
                }

                if (uuid == value.UUID) {
                    object[field].splice(i, 1)
                    // remove the set and reset function.
                    if (isRef) {
                        delete object["reset_" + field + "_" + uuid + "_ref"]
                        delete object["set_" + field + "_" + uuid + "_ref"]
                    }
                }
            }
        } else {
            if (isRef) {
                var uuid = ""
                if (isObject(object[field])) {
                    uuid = object[field].UUID
                } else {
                    uuid = object[field]
                }
                delete object["reset_" + field + "_" + uuid + "_ref"]
                delete object["set_" + field + "_" + uuid + "_ref"]
            }
            delete object[field]
        }
    }
    object["NeedSave"] = true
}

/**
 * That function is use to reset the entity of it's original value.
 */
function resetObjectValues(object) {
    // Remove the object panel...
    delete object["panel"]
    var prototype = entityPrototypes[object.TYPENAME]

    for (var propertyId in object) {
        var propretyType = prototype.FieldsType[prototype.getFieldIndex(propertyId)]
        if (propretyType != undefined && object[propertyId] != null) {
            var isRef = propretyType.endsWith(":Ref")
            var isArray = propretyType.startsWith("[]")
            var isBaseType = propretyType.startsWith("[]xs.") || propretyType.startsWith("xs.")
            if (isArray) {
                for (var i = 0; i < object[propertyId].length; i++) {
                    if (isObject(object[propertyId][i])) {
                        if (object[propertyId][i]["UUID"] != undefined) {
                            // Reset it's sub-objects.
                            if (!isRef && !isBaseType) {
                                resetObjectValues(object[propertyId][i])
                            }
                        }
                    }
                }
            } else if (isObject(object[propertyId])) {
                if (object[propertyId]["UUID"] != undefined) {
                    if (!isRef && !isBaseType) {
                        resetObjectValues(object[propertyId])
                    }
                }
            }
        }

        /** Only reference must be reset here. */
        if (propertyId.startsWith("reset_") && propertyId.endsWith("_ref")) {
            // Call the reset function.
            object[propertyId]()
        }
    }
}

/**
 * Return a property field type for a given field for a given type name.
 */
function getPropertyType(typeName, property) {
    var prototype = entityPrototypes[typeName]
    var propertyType = null
    for (var i = 0; i < prototype.Fields.length; i++) {
        if (prototype.Fields[i] == property) {
            propertyType = prototype.FieldsType[i]
            break
        }
    }
    return propertyType
}

/////////////////////////////////////////////////////////////////////////
// Entity referenced information.

/**
 * That stucture is use to keep track of object 
 * referenced by another object.
 */
var EntityRef = function (name, owner, value) {
    this.TypeName = "Server.EntityRef"
    this.Name = name
    this.OwnerUuid = owner.UUID
    this.Value = value

    return this
}

/**
 * Append a reference to a target if it dosen't already exist.
 */
function appendReferenced(refName, target, owner) {

    if (target.referenced == undefined) {
        target.referenced = []
    }

    // Look if the value is not already there...
    for (var i = 0; i < target.referenced.length; i++) {
        var ref = target.referenced[i]
        if (ref.Name == refName && ref.OwnerUuid == owner.UUID) {
            // Nothing to do here.
            return
        }
    }

    // Append the referenced.
    target.referenced.push(new EntityRef(refName, owner, ""))
}

/////////////////////////////////////////////////////////////////////////
// Initialisation code here.

/**
 * Set object reference.
 */
function setRef(owner, property, refValue, isArray) {
    if (refValue.length == 0) {
        return owner
    }

    // Keep track of the references targets.
    if (owner.references.indexOf(refValue) == -1) {
        owner.references.push(refValue)
    }

    if (!isObjectReference(refValue)) {
        owner[property] = refValue
        return owner
    }


    if (isArray) {
        var index = owner[property].length
        if (owner[property].indexOf(refValue) == -1) {
            owner[property].push(refValue)
            /* The reset reference fucntion **/
            owner["reset_" + property + "_" + refValue + "_ref"] = function (entity, propertyName) {
                return function () {
                    for (var i = 0; i < entity[propertyName].length; i++) {
                        if (isObject(entity[propertyName][i])) {
                            entity[propertyName][i] = entity[propertyName][i].UUID
                        }
                    }
                }
            }(owner, property)

            /* The set reference fucntion **/
            owner["set_" + property + "_" + refValue + "_ref"] = function (entityUuid, propertyName, index, refValue) {
                return function (initCallback) {
                    var isExist = entities[refValue] != undefined
                    var isInit = false

                    if (isExist) {
                        isInit = entities[refValue].IsInit
                    }

                    if (isExist && isInit) {
                        // Here the reference exist on the server.
                        var entity = entities[entityUuid]
                        var ref = entities[refValue]
                        entity[propertyName][index] = ref
                        appendReferenced(propertyName, ref, entity)
                        if (initCallback != undefined) {
                            initCallback(ref)
                            initCallback = undefined
                        }

                    } else {
                        server.entityManager.getEntityByUuid(refValue,
                            function (result, caller) {
                                var propertyName = caller.propertyName
                                var index = caller.index
                                var entity = entities[caller.entityUuid]
                                var ref = entities[caller.refValue]
                                entity[propertyName][index] = ref
                                appendReferenced(propertyName, ref, entity)
                                if (caller.initCallback != undefined) {
                                    caller.initCallback(ref)
                                    caller.initCallback = undefined
                                }
                            },
                            function (errorMsg, caller) {
                            },
                            { "entityUuid": entityUuid, "propertyName": propertyName, "index": index, "refValue": refValue, "initCallback": initCallback }
                        )
                    }
                }
            }(owner.UUID, property, index, refValue)
        }
    } else {
        owner[property] = refValue

        /* The reset fucntion **/
        owner["reset_" + property + "_" + refValue + "_ref"] = function (entityUuid, propertyName) {
            return function () {
                var entity = entities[entityUuid]
                // Set back the id of the reference
                if (entity != null) {
                    if (isObject(entity[propertyName])) {
                        entity[propertyName] = entity[propertyName].UUID
                    }
                }
            }
        }(owner.UUID, property)

        /* The set fucntion **/
        owner["set_" + property + "_" + refValue + "_ref"] = function (entityUuid, propertyName, refValue) {
            return function (initCallback) {

                var isExist = entities[refValue] != undefined
                var isInit = false

                if (isExist) {
                    isInit = entities[refValue].IsInit
                }

                // If the entity is already on the client side...
                if (isExist && isInit) {
                    // Here the reference exist on the server.
                    var entity = entities[entityUuid]
                    var ref = entities[refValue]
                    entity[propertyName] = ref
                    appendReferenced(propertyName, ref, entity)
                    if (initCallback != undefined) {
                        initCallback(ref)
                        initCallback = undefined
                    }
                } else {
                    server.entityManager.getEntityByUuid(refValue,
                        function (result, caller) {
                            var propertyName = caller.propertyName
                            var entity = entities[caller.entityUuid]
                            var ref = entities[caller.refValue]
                            entity[propertyName] = ref
                            appendReferenced(propertyName, ref, entity)
                            if (caller.initCallback != undefined) {
                                caller.initCallback(ref)
                                caller.initCallback = undefined
                            }
                        },
                        function () { },
                        { "entityUuid": entityUuid, "propertyName": propertyName, "refValue": refValue, "initCallback": initCallback }
                    )
                }
            }
        }(owner.UUID, property, refValue)
    }
    return owner
}

/**
 * Set object, that function call setObjectValues in this path so it's recursive.
 */
function setSubObject(parent, property, values, isArray) {

    if (values.TYPENAME == undefined || values.UUID.length == 0) {
        return parent
    }

    server.entityManager.getEntityPrototype(values.TYPENAME, values.TYPENAME.split(".")[0],
        function (result, caller) {
            var parent = caller.parent
            var property = caller.property
            var values = caller.values
            var isArray = caller.isArray
            if (values.TYPENAME == "BPMN20.StartEvent") {
                i = 0;
            }

            var object = entities[values.UUID]
            if (object == undefined) {
                object = eval("new " + values.TYPENAME + "()")
                // Keep track of the parent uuid in the child.
                object.UUID = values.UUID
            }

            // Keep track of the child uuid inside the parent.
            if (parent.childsUuid == undefined) {
                parent.childsUuid = []
            }

            if (parent.childsUuid.indexOf(object.UUID)) {
                parent.childsUuid.push(object.UUID)
            }

            if (isArray) {
                if (parent[property] == undefined) {
                    parent[property] = []
                }
                object.init(values)
                parent[property].push(object)

            } else {
                object.init(values)
                parent[property] = object
            }

            object.ParentUuid = parent.UUID
            object.ParentLnk = property
            server.entityManager.setEntity(object)

        },
        function () {

        }, { "parent": parent, "property": property, "values": values, "isArray": isArray })


    return parent
}

/**
 * That function initialyse an object created from a given prototype constructor with the values from a plain JSON object.
 * @param {object} object The object to initialyse.
 * @param {object} values The plain JSON object that contain values.
 */
function setObjectValues(object, values) {

    // Get the entity prototype.
    var prototype = entityPrototypes[object["TYPENAME"]]
    if (prototype == undefined) {
        return
    }

    ////////////////////////////////////////////////////////////////////////
    // Set back the reference...
    if (values == undefined) {
        // call set_property on each object if there is a function defined.
        for (var property in object) {
            var propertyType = getPropertyType(object["TYPENAME"], property)
            if (propertyType != null) {
                var isRef = propertyType.endsWith(":Ref")
                if (propertyType.startsWith("[]")) {
                    // The property is an array.
                    if (object[property] != null) {
                        if (object[property].length > 0) {
                            for (var i = 0; i < object[property].length; i++) {
                                if (isString(object[property][i])) {
                                    if (isRef) {
                                        if (object["set_" + property + "_" + object[property] + "_ref"] != undefined) {
                                            // Call it.
                                            object["set_" + property + "_" + object[property] + "_ref"]()
                                        }
                                    }
                                }
                            }
                        }
                    }
                } else {
                    if (object[property] != undefined) {
                        if (isString(object[property])) {
                            if (isRef) {
                                if (object["set_" + property + "_" + object[property] + "_ref"] != undefined) {
                                    // Call it.
                                    object["set_" + property + "_" + object[property] + "_ref"]()
                                }
                            }
                        }
                    }
                }
            }
        }
        return
    }

    ////////////////////////////////////////////////////////////////////
    // Reset actual object fields and cound number of sub-objects...

    // The list of properties to set.
    // reference values must be put at end of the list.
    var properties = []
    for (var i = 0; i < prototype.FieldsType.length; i++) {
        var fieldType = prototype.FieldsType[i]
        var field = prototype.Fields[i]
        if (field.startsWith("M_") || field.startsWith("[]M_")) {
            // Reset the objet fields.
            if (!fieldType.startsWith("sqltypes.") && !fieldType.startsWith("[]sqltypes.") && !fieldType.startsWith("xs.") && !fieldType.startsWith("[]xs.")) {
                if (fieldType.startsWith("[]")) {
                    object[field] = []
                } else {
                    object[field] = ""
                }
            } else {
                // Reset the base type fields.
                if (fieldType.startsWith("[]xs.") || fieldType.startsWith("[]sqltypes.")) {
                    object[field] = []
                } else if (fieldType.startsWith("xs.") || fieldType.startsWith("sqltypes.")) {
                    object[field] = ""
                }
            }
        }
    }

    ////////////////////////////////////////////////////////////////////////////////
    // Generate sub-object and reference set and reset function
    for (var property in values) {

        var propertyType = prototype.FieldsType[prototype.getFieldIndex(property)]

        if (propertyType != null) {
            // Condition...
            var isRef = propertyType.endsWith(":Ref")
            var baseType = getBaseTypeExtension(propertyType)

            // M_listOf, M_valueOf field or enumeration type contain plain value.
            var isBaseType = isXsBaseType(baseType) || propertyType.startsWith("sqltypes.") && !propertyType.startsWith("[]sqltypes.") || propertyType.startsWith("[]xs.") || propertyType.startsWith("xs.") || property == "M_listOf" || property == "M_valueOf" || propertyType.startsWith("enum:")

            if (values[property] != null) {
                if (isBaseType) {
                    // String, int, double...
                    if (propertyType == "xs.base64Binary") {
                        // In case of binairy object I will try to create object if information is given for it.
                        // TODO see why to decode is necessary here...
                        var strVal = decode64(decode64(values[property]))
                        if (strVal.indexOf("TYPENAME") != -1 && strVal.indexOf("__class__") != -1) {
                            var jsonObj = JSON.parse(strVal)
                            if (!isArray(jsonObj)) {
                                // In case of the object is not an array...
                                var obj = eval("new " + jsonObj.TYPENAME + "()")
                                obj.initCallback = function (uuid, property) {
                                    return function (val) {
                                        entities[uuid][property] = val
                                    }
                                }(object.UUID, property)

                                obj.init(jsonObj)
                            } else {
                                for (var i = 0; i < jsonObj.length; i++) {
                                    var jsonObj_ = JSON.parse(jsonObj[i])
                                    var obj = eval("new " + jsonObj_.TYPENAME + "()")
                                    obj.initCallback = function (uuid, property) {
                                        return function (val) {
                                            if (entities[uuid][property] == "") {
                                                entities[uuid][property] = []
                                            }
                                            entities[uuid][property].push(val)
                                        }
                                    }(object.UUID, property)
                                    obj.init(jsonObj_)
                                }
                            }
                        } else {
                            // No information available to create an object, so the string will be use....
                            object[property] = strVal
                        }
                    } else {
                        object[property] = values[property]
                        // If the object is a simple derived type.
                        if (object[property].UUID != undefined) {
                            server.entityManager.setEntity(object[property])
                        }
                    }
                } else {
                    // Set object ref or values... only property begenin with M_ will be set here...
                    var isArray_ = propertyType.startsWith("[]")
                    if (property.startsWith("[]M_") || property.startsWith("M_")) {
                        if (isArray_) {
                            object[property] = []
                            for (var i = 0; i < values[property].length; i++) {
                                if (isRef) {
                                    object = setRef(object, property, values[property][i], isArray_)
                                } else {
                                    object = setSubObject(object, property, values[property][i], isArray_)
                                }
                            }
                        } else {
                            if (isRef) {
                                object = setRef(object, property, values[property], isArray_)
                            } else {
                                object = setSubObject(object, property, values[property], isArray_)
                            }
                        }
                    }
                }
            }
        }
    }

    //////////////////////////////////////////////////////
    // Set common values...
    object.UUID = values.UUID
    object.NeedSave = false
    object.exist = true
    object.IsInit = true // The object part only and not the refs...
    object.ParentUuid = values.ParentUuid // set the parent uuid.
    object.ParentLnk = values.ParentLnk
    
    // Set the initialyse object.
    server.entityManager.setEntity(object)

    // Call the init callback.
    if (object.initCallback != undefined) {
        object.initCallback(object)
        object.initCallback == undefined
    }
}

/**
 * Look if the given type is a list of other type.
 * @param {string} typeName The extension type name, ex. xs.string, xs.int, xs.date etc.
 */
function isListOf(typeName) {
    typeName = typeName.replace("[]", "").replace(":Ref", "")
    var prototype = entityPrototypes[typeName]

    if (prototype != null) {
        if (prototype.ListOf != null) {
            if (prototype.ListOf.length > 0) {
                return true
            }
        }
        if (prototype.SuperTypeNames != null) {
            for (var i = 0; i < prototype.SuperTypeNames.length; i++) {
                if (prototype.SuperTypeNames[i] != prototype.TypeName) {
                    if (isListOf(prototype.SuperTypeNames[i])) {
                        return true
                    }
                }
            }
        }
    }
    return false
}

/**
 * Return the name of the base type if the type is an extension of such a type.
 * @param {string} typeName The extension type name, ex. xs.string, xs.int, xs.date etc.
 */
function getBaseTypeExtension(typeName, isArray) {
    if (!isArray) {
        isArray = typeName.startsWith("[]")
    }
    typeName = typeName.replace("[]", "").replace(":Ref", "")
    var prototype = entityPrototypes[typeName]

    if (prototype != null) {
        if (prototype.SuperTypeNames != null) {
            for (var i = 0; i < prototype.SuperTypeNames.length; i++) {
                if (prototype.SuperTypeNames[i].startsWith("xs.")) {
                    if (isArray) {
                        return "[]" + prototype.SuperTypeNames[i]
                    }
                    return prototype.SuperTypeNames[i]
                } else if (getBaseTypeExtension(prototype.SuperTypeNames[i]).length > 0) {
                    return getBaseTypeExtension(prototype.SuperTypeNames[i])
                }
            }
        } else if (prototype.ListOf != null) {
            if (prototype.ListOf.length > 0) {
                if (prototype.ListOf.startsWith("xs.")) {
                    return "[]" + prototype.ListOf
                }
                return getBaseTypeExtension(prototype.ListOf, true)
            }
        }
    }
    return ""
}
