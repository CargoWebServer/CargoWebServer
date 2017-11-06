/**
 * Here i will made use of the dynamic caracter of the object store to create a new object...
 */
var EntityPrototypeEditor = function (parent, imports, baseType, initCallback) {
    // Here I will create the content of the panel...
    this.space = new Element(parent, { "tag": "div", "class": "entity admin_table", "style": "display:block;top:0px; bottom:0px; left:0px; right:0px; position: absolute; display: none;" })
    this.panel = this.space.appendElement({ "tag": "div", "class": "admin_table", "style": "display:table;max-height: 250px; overflow-y:auto; width:auto; position: relative; border-collapse:separate;border-spacing:5px; text-align: left;" }).down()
    this.proto = null

    // Call when the initialisation is done.
    this.initCallback = initCallback

    // The list dependency
    this.imports = imports

    // the list of all typename contain 
    // in imports.
    this.typeNameLst = []

    // If there is a type from where all other type must be derived from.
    this.baseType = baseType

    // The map of field to be updated.
    this.fieldsToUpdate = {}

    // The typeName input...
    this.panel.appendElement({ "tag": "div", "style": "display: table; width: 100%;" }).down()
        .appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
        .appendElement({ "tag": "div", "id": "save_entity_prototype", "class": "entities_btn", "style": "display: none; margin-left: 8px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-floppy-o" }).up()
        .appendElement({ "tag": "div", "style": "display: table-cell; width: 100%;" }).down()
        .appendElement({ "tag": "div", "style": "display: inline-block; padding-right: 5px;", "innerHtml": "Type Name" })
        .appendElement({ "tag": "div", "style": "display: inline-block" }).down()
        .appendElement({ "tag": "input", "id": "dynamicItemName", "style": "display: inline-block; width: 250px;" }).up().up()
        .appendElement({ "tag": "div", "id": "delete_entity_prototype", "class": "entities_btn", "style": "display: none; margin-left: 8px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-trash-o" })

    this.saveBtn = this.panel.getChildById("save_entity_prototype")
    this.deleteBtn = this.panel.getChildById("delete_entity_prototype")

    // I will retreive all derived item type...
    this.typeNameInput = this.panel.getChildById("dynamicItemName")

    // I will get the list of type derived from item...
    // Here i will display the base type.
    this.panel.appendElement({ "tag": "div", "style": "display: table; width: 100%;", "innerHtml": "Base Type" })
        .appendElement({ "tag": "div", "style": "display: table; width: 100%;" }).down()
        .appendElement({ "tag": "div", "id": "allSuperTypes", "style": "display: table-cell; width: 50%; min-width: 200px; height: 150px; overflow-y: auto; border: 1px solid grey;" })
        .appendElement({ "tag": "div", "id": "superTypes", "style": "display: table-cell; width: 50%; min-width: 200px; height: 150px; overflow-y: auto; border: 1px solid grey;" })

    this.allSuperTypes = this.panel.getChildById("allSuperTypes")
    this.superTypes = this.panel.getChildById("superTypes")

    this.restrictions = this.panel
        .appendElement({ "tag": "div", "style": "display: none;", "id": "append_restriction_panel" }).down()
        .appendElement({ "tag": "div", "style": "display: table-cell; vertical-align: middle;" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Restrictions" }).up()
        .appendElement({ "tag": "div", "class": "entities_btn", "style": "diplay: table-cell;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-plus", "id": "append_restiction_btn" }).up()
        .appendElement({ "tag": "div", "style": "display: table;", "id": "restriction_edit_panel" }).up()
        .appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
        .appendElement({ "tag": "div", "style": "display: none; width: 100%;height: 150px;min-width: 200px; overflow-y: auto; border: 1px solid grey;" }).down()
        .appendElement({ "tag": "div", "style": "display: table; width: 100%; " }).down()

    this.restrictionEditPanel = this.panel.getChildById("restriction_edit_panel")

    this.appendRestrictionPanel = this.panel.getChildById("append_restriction_panel")
    this.appendRestrictionBtn = this.panel.getChildById("append_restiction_btn")

    this.appendRestrictionBtn.element.onclick = function (entityPrototypeEditor) {
        return function () {
            this.style.display = "none" // one restriction at time.
            entityPrototypeEditor.appendRestriction()
        }
    }(this)

    // The propetie panel.
    this.properties = this.panel.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;", "innerHtml": "Properties" })
        .appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
        .appendElement({ "tag": "div", "style": "display: table" })

    // So here for each base type I will get the base type propertie...
    function setAutocompleteDynamicType(panel) {

        // Now I will create the list of derived class name...
        var objMap = {}

        for (var i = 0; i < panel.typeNameLst.length; i++) {
            var typeName = panel.typeNameLst[i]
            objMap[typeName] = getEntityPrototype(typeName)
            objMap[typeName.split(".")[1]] = getEntityPrototype(typeName)
        }

        // Now I will create the auto-complete text-box.
        var input = panel.typeNameInput
        attachAutoComplete(input, panel.typeNameLst, false)

        input.element.addEventListener("keyup", function (panel) {
            return function (e) {
                // If the key is escape...
                if (this.value.length == 0) {
                    panel.clear()
                } else {
                    if (e.keyCode == 13 && panel.getCurrentEntityPrototype() == null) {
                        // In that case I will create a new entity prototype.
                        var prototype = new EntityPrototype()
                        prototype.init({
                            "TypeName": this.value,
                            "Fields": [],
                            "FieldsType": [],
                            "FieldsVisibility": [],
                            "FieldsOrder": [],
                            "Ids": [],
                            "Index": [],
                            "SuperTypeNames": [],
                            "Restrictions": [],
                            "IsAbstract": false,
                            "SubstitutionGroup": [],
                            "FieldsNillable": [],
                            "FieldsDocumentation": [],
                            "FieldsDefaultValue": [],
                            "ListOf": ""
                        })

                        prototype.notExist = true
                        if (panel.baseType != undefined) {
                            if (panel.baseType.length > 0) {
                                prototype.SuperTypeNames.push(panel.baseType)
                            }
                        }

                        panel.setCurrentPrototype(prototype)
                        // Keep reference to the prototype.
                        setEntityPrototype(prototype)
                    }
                }
            }
        }(panel))

        input.element.onblur = input.element.onchange = function (objMap, panel) {
            return function (evt) {
                var prototype = objMap[this.value]
                if (prototype != undefined) {
                    panel.setCurrentPrototype(prototype)
                }
            }
        }(objMap, panel)

        input.element.addEventListener("keyup", function (panel, input) {
            return function (e) {
                // If the key is escape...
                if (e.keyCode === 8) {
                    panel.clear()
                }
                // Only index selection will erase the panel if the input is empty.
                if (this.value.indexOf(".") != -1) {
                    var storeId = this.value.split(".")[0]
                    if (storeId.length > 0) {
                        server.entityManager.getEntityPrototype(this.value, storeId,
                            // Success Callback
                            function (prototype, caller) {
                                var panel = caller.panel
                                panel.setCurrentPrototype(prototype)
                                caller.input.autocompleteDiv.element.style.display = "none"
                            },
                            // Error Callback
                            function (errObj, caller) {
                                if (caller.panel.getCurrentEntityPrototype() != undefined) {
                                    if (caller.panel.getCurrentEntityPrototype().notExist == undefined) {
                                        caller.panel.clear()
                                    }
                                }
                            },
                            { "panel": panel, "input": input })

                    } else {
                        panel.clear()
                    }
                } else {
                    panel.clear()
                }
            }
        }(panel, input))

        if (panel.initCallback != undefined) {
            panel.initCallback(panel)
        }
    }

    function setTypeNameInputAutocomplete(panel) {

        var callback = function (panel, lst, index, total, callback) {
            if (total > 0) {
                server.entityManager.getEntityPrototypes(panel.imports[index],
                    function (results, caller) {
                        // Now I will create the list of possible type...
                        for (var i = 0; i < results.length - 1; i++) {
                            caller.lst.push(results[i].TypeName)
                        }

                        // Now I will set the autocomplete box.
                        if (caller.index == caller.total - 1) {
                            caller.lst.sort()
                            setAutocompleteDynamicType(caller.panel)
                        } else {
                            caller.callback(caller.panel, caller.lst, caller.index + 1, caller.total, caller.callback)
                        }
                    },
                    function (errObj, caller) {

                    },
                    { "panel": panel, "lst": lst, "index": index, "total": total, "callback": callback })
            } else {
                // Call the callback.
                callback(panel)
            }
        }

        // recusively append data type until all is done.
        callback(panel, panel.typeNameLst, 0, panel.imports.length, callback)
    }

    setTypeNameInputAutocomplete(this)

    this.saveBtn.element.onclick = function (panel) {
        return function () {
            panel.saveBtn.element.style.display = "none"
            // dynamic editor.
            panel.saveprototype()
        }
    }(this)

    this.deleteBtn.element.onclick = function (panel) {
        return function () {
            // Now I will ask for confirmation.
            var confirmDialog = new Dialog(randomUUID(), undefined, true)
            confirmDialog.div.element.style.maxWidth = "450px"
            confirmDialog.setCentered()
            server.languageManager.setElementText(confirmDialog.title, "delete_dialog_entity_title")
            confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete prototype " + panel.getCurrentEntityPrototype().TypeName + "?" })

            confirmDialog.ok.element.onclick = function (dialog, panel) {
                return function () {
                    var typeName = panel.getCurrentEntityPrototype().TypeName
                    server.entityManager.deleteEntityPrototype(typeName, typeName.split(".")[0],
                        function (result, caller) {
                            caller.panel.clear()
                        },
                        function () {

                        }, { "panel": panel })

                    dialog.close()
                }
            }(confirmDialog, panel)
        }
    }(this)

    /////////////////////////////////////////////////////////////////////////////////////////
    // Event handling here.
    /////////////////////////////////////////////////////////////////////////////////////////

    // New prototype event
    server.prototypeManager.attach(this, NewPrototypeEvent, function (evt, EntityPrototypeEditor) {
        // Refresh the content of autocomplte dynamic type.
        setAutocompleteDynamicType(EntityPrototypeEditor)
    })

    // Update prototype event
    server.prototypeManager.attach(this, UpdatePrototypeEvent, function (evt, EntityPrototypeEditor) {
        // if the current item is the one with change I will reset it content.
        if (evt.dataMap.prototype.TypeName == EntityPrototypeEditor.getCurrentEntityPrototype().TypeName) {
            EntityPrototypeEditor.setCurrentPrototype(evt.dataMap.prototype)
        }
    })

    // Delete prototype event.
    server.prototypeManager.attach(this, DeletePrototypeEvent, function (evt, EntityPrototypeEditor) {
        // If the current display item is deleted I will clear the panel.
        if (evt.dataMap.prototype.TypeName == EntityPrototypeEditor.getCurrentEntityPrototype().TypeName) {
            EntityPrototypeEditor.clear()
            if (EntityPrototypeEditor.typeName != null) {
                if (EntityPrototypeEditor.typeName.element.value == evt.dataMap.prototype.TypeName) {
                    EntityPrototypeEditor.typeName.element.value = ""
                }
            }
        }

        // Refresh the content of autocomplte dynamic type.
        setAutocompleteDynamicType(EntityPrototypeEditor)
    })

    return this
}

EntityPrototypeEditor.prototype.getCurrentEntityPrototype = function () {
    return this.proto
}

/**
 * Clear the content of the editor.
 */
EntityPrototypeEditor.prototype.clear = function () {
    this.superTypes.removeAllChilds()
    this.allSuperTypes.removeAllChilds()
    this.properties.removeAllChilds()
    this.proto = null
    this.saveBtn.element.style.display = "none"
    this.deleteBtn.element.style.display = "none"
}

/**
 * Set the current edited item.
 */
EntityPrototypeEditor.prototype.setCurrentPrototype = function (prototype) {

    // First of all I will display the item supertypes informations.
    this.fieldsToUpdate = {}
    this.displaySupertypes(prototype, function (EntityPrototypeEditor) {
        return function (prototype) {
            // Now I will display the item properties.
            EntityPrototypeEditor.displayPrototypeProperties(prototype, true)
            EntityPrototypeEditor.proto = prototype
            EntityPrototypeEditor.deleteBtn.element.style.display = "table-cell"

            // Now I will display the item restrictions.
            var baseType = getBaseTypeExtension(prototype.TypeName)
            // Only base type derived prototype can have restrictions.
            if (isXsBaseType(baseType)) {
                EntityPrototypeEditor.appendRestrictionPanel.element.style.display = "table"
                EntityPrototypeEditor.restrictions.removeAllChilds()
                EntityPrototypeEditor.displayPrototypeRestrictions(prototype)
            }
        }
    }(this))
}

/**
 * Dipsplay the sypertype information.
 */
EntityPrototypeEditor.prototype.displaySupertypes = function (prototype, callback) {
    // Clear the list of supertypes
    this.clear()
    var successCallback = function (results, caller) {

        // Test only...
        var editor = caller.editor
        var prototype = caller.prototype
        var superTypes = editor.superTypes.appendElement({ "tag": "div", "style": "display: table; width: 100%;" }).down()

        // set supertypename as empty array is it's null
        if (prototype.SuperTypeNames == undefined) {
            prototype.SuperTypeNames = []
        }
        for (var i = 0; i < prototype.SuperTypeNames.length; i++) {
            if (document.getElementById("superType_" + prototype.SuperTypeNames[i] + "_" + prototype.TypeName + "_row") == undefined) {
                var superType = superTypes.appendElement({ "tag": "div", "style": "display: table-row;width: 100%;", "id": "superType_" + prototype.SuperTypeNames[i] + "_" + prototype.TypeName + "_row" }).down()
                var removeSupertypeBtn = superType.appendElement({ "tag": "div", "style": "display: table-cell;width: 100%;", "innerHtml": prototype.SuperTypeNames[i] })
                    .appendElement({ "tag": "div", "class": "entities_btn" }).down()
                    .appendElement({ "tag": "i", "class": "fa fa-close" }).down()

                if (prototype.SuperTypeNames[i] == editor.baseType) {
                    removeSupertypeBtn.element.style.display = "none"
                } else {
                    removeSupertypeBtn.element.onclick = function (editor, prototype, superTypename) {
                        return function () {
                            prototype.SuperTypeNames.splice(prototype.SuperTypeNames.indexOf(superTypename), 1);
                            editor.setCurrentPrototype(prototype)
                        }
                    }(editor, prototype, prototype.SuperTypeNames[i])
                }

                // Here I will display the supertypes properties...
                editor.displayPrototypeSupertypeProperties(prototype.SuperTypeNames[i], function (callback, prototype, isDone) {
                    return function () {
                        if (isDone) {
                            // continue the execution here.
                            callback(prototype)
                        }
                    }
                }(caller.callback, prototype, i == prototype.SuperTypeNames.length - 1))
            }
        }

        // Now the list of other super type not used by this one.
        if (document.getElementById("allSuperTypes_" + prototype.TypeName) == undefined) {
            var allSuperTypes = editor.allSuperTypes.appendElement({ "tag": "div", "style": "width: 100%; height: 200px;", "id": "allSuperTypes_" + prototype.TypeName }).down()
            for (var i = 0; i < results.length; i++) {
                if (results[i] != undefined) {
                    if (document.getElementById("allSuperTypes_" + results[i].TypeName + "_" + prototype.TypeName + "_row") == undefined) {
                        // display only if the supertype is not already present in the list of supertype 
                        if (prototype.SuperTypeNames.indexOf(results[i].TypeName) == -1) {
                            // The type must not have actual type as supertype.
                            if (results[i].SuperTypeNames.indexOf(prototype.TypeName) == -1 && results[i].TypeName != prototype.TypeName && prototype.SuperTypeNames.indexOf(results[i].TypeName) == -1) {
                                var superType = allSuperTypes.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;", "id": "allSuperTypes_" + results[i].TypeName + "_" + prototype.TypeName + "_row" }).down()
                                var appendSupertypeBtn = superType.appendElement({ "tag": "div", "style": "display: table-cell; width: 100%;", "innerHtml": results[i].TypeName })
                                    .appendElement({ "tag": "div", "class": "entities_btn" }).down()
                                    .appendElement({ "tag": "i", "class": "fa fa-plus" }).down()

                                // Here I will append a new supertype to the prototype.
                                appendSupertypeBtn.element.onclick = function (editor, prototype, superPrototype) {
                                    return function () {
                                        // Recursively append all surpertype to a given prototype.
                                        function appendSuperPrototype(prototype, superPrototype, callback) {
                                            var callback_ = function (prototype, superPrototype, index, total, callback, callback_) {
                                                if (superPrototype.SuperTypeNames.length > 0) {
                                                    server.entityManager.getEntityPrototype(superPrototype.SuperTypeNames[index], superPrototype.SuperTypeNames[index].split(".")[0],
                                                        function (result, caller) {
                                                            // recursively append the prototype.
                                                            if (caller.prototype.SuperTypeNames.indexOf(result.TypeName) == -1 && caller.prototype.TypeName != result.TypeName) {
                                                                appendSuperPrototype(caller.prototype, result, caller.callback, caller.done)
                                                            }
                                                            if (caller.index == caller.total - 1) {
                                                                if (caller.prototype.SuperTypeNames.indexOf(caller.superPrototype.TypeName) == -1) {
                                                                    caller.prototype.SuperTypeNames.push(caller.superPrototype.TypeName)
                                                                }
                                                                caller.callback(caller.prototype)
                                                            } else {
                                                                caller.callback_(caller.prototype, caller.superPrototype, caller.index + 1, caller.total, caller.callback_, caller.callback)
                                                            }
                                                        },
                                                        function (errObj, caller) {
                                                        }, { "prototype": prototype, "superPrototype": superPrototype, "index": index, "total": superPrototype.SuperTypeNames.length, "callback": callback, "callback_": callback_ })
                                                } else {
                                                    if (prototype.SuperTypeNames.indexOf(superPrototype.TypeName) == -1) {
                                                        prototype.SuperTypeNames.push(superPrototype.TypeName)
                                                    }
                                                    callback(prototype)
                                                }
                                            }
                                            // call for each super
                                            callback_(prototype, superPrototype, 0, callback_, callback)
                                        }

                                        appendSuperPrototype(prototype, superPrototype, function (prototype, editor) {
                                            return function (prototype) {
                                                editor.setCurrentPrototype(prototype)
                                            }
                                        }(prototype, editor))


                                    }
                                }(editor, prototype, results[i])
                            }
                        }
                    }
                }
            }
        }

        // In the case of no surpertype exist...
        if (prototype.SuperTypeNames.length == 0) {
            callback(prototype)
        }

    }

    // In case of a given a base type name is given...
    if (this.baseType != undefined) {
        server.entityManager.getDerivedEntityPrototypes(this.baseType,
            /** Success callback */
            successCallback,
            /** Error callback */
            function (errObj, caller) {

            }, { "callback": callback, "prototype": prototype, "editor": this })
    } else {
        // In the other case i will use all prototypes from imports.
        var results = []
        for (var i = 0; i < this.typeNameLst.length; i++) {
            results.push(getEntityPrototype(this.typeNameLst[i]))
        }
        var caller = { "callback": callback, "prototype": prototype, "editor": this }
        successCallback(results, caller)
    }
}

/**
 * Set the supertype item properties.
 * The callback function is call when all propreties are display.
 */
EntityPrototypeEditor.prototype.displayPrototypeSupertypeProperties = function (superTypeName, callback) {
    // Here I will get All element derived from 
    server.entityManager.getEntityPrototype(superTypeName, superTypeName.split(".")[0],
        // success callback
        function (result, caller) {
            caller.editor.displayPrototypeProperties(result, false)
            if (caller.callback != undefined) {
                caller.callback(result.Fields)
            }
        },
        function (errObj, caller) {

        },
        { "editor": this, "callback": callback })
}

/**
 * Display the item propertie edition panel.
 */
EntityPrototypeEditor.prototype.displayPrototypeProperties = function (prototype, isEditable) {

    // The title.
    this.properties.appendElement({ "tag": "div", "style": "display: table; width: 100%;" }).down()
        .appendElement({ "tag": "div" }).down()
        .appendElement({ "tag": "div", "class": "entities_btn", "style": "display: table-cell" }).down()
        .appendElement({ "tag": "i", "id": "new_prototype_propertie_" + prototype.TypeName + "_btn", "class": "fa fa-plus" }).up()
        .appendElement({ "tag": "div", "style": "display: table-cell;", "innerHtml": prototype.TypeName })

    var properties = this.properties.appendElement({ "tag": "div", "style": "display: table; width: 100%;  padding-left: 18px;" }).down()

    // So here I will get the list of all supertype field.
    function getSuperTypesFields(superTypeNames, fields, callback) {
        var callback_ = function (superTypeNames, fields, index, total, callback, callback_) {
            if (total > 0) {
                server.entityManager.getEntityPrototype(superTypeNames[index], superTypeNames[index].split(".")[0],
                    // success callback
                    function (result, caller) {
                        caller.fields = caller.fields.concat(result.Fields)
                        if (caller.index == caller.total - 1) {
                            caller.callback(caller.fields)
                        } else {
                            caller.callback_(caller.superTypeNames, caller.fields, caller.index + 1, caller.total, caller.callback, caller.callback_)
                        }
                    },
                    // error callback
                    function () {

                    },
                    // nothing to do here.
                    { "superTypeNames": superTypeNames, "fields": fields, "index": index, "total": total, "callback": callback, "callback_": callback_ })
            } else {
                fields = fields.concat(result.Fields)
                callback(fields)
            }
        }

        callback_(superTypeNames, fields, 0, superTypeNames.length, callback, callback_)
    }


    if (prototype.SuperTypeNames == undefined) {
        prototype.SuperTypeNames = []
    }
    var superTypesFields = []
    if (prototype.SuperTypeNames.length > 0) {
        getSuperTypesFields(prototype.SuperTypeNames, superTypesFields,
            function (EntityPrototypeEditor, prototype, properties, isEditable) {
                return function (superTypesFields) {
                    for (var i = 3; i < prototype.Fields.length - 2; i++) {
                        // display attributes
                        if (superTypesFields.indexOf(prototype.Fields[i]) == -1) {
                            // Here only if the propertie is part of the entity itself and not of one of it parent.
                            EntityPrototypeEditor.displayPrototypePropertie(prototype, prototype.Fields[i], prototype.FieldsType[i], properties, isEditable, i)
                        }
                    }
                }
            }(this, prototype, properties, isEditable))
    } else {
        for (var i = 3; i < prototype.Fields.length - 2; i++) {
            // display attributes
            if (superTypesFields.indexOf(prototype.Fields[i]) == -1) {
                // Here only if the propertie is part of the entity itself and not of one of it parent.
                this.displayPrototypePropertie(prototype, prototype.Fields[i], prototype.FieldsType[i], properties, isEditable, i)
            }
        }
    }

    var newPropertieBtn = this.properties.getChildById("new_prototype_propertie_" + prototype.TypeName + "_btn")
    newPropertieBtn.element.onclick = function (EntityPrototypeEditor, prototype, properties) {
        return function () {
            var index = prototype.Fields.length - 2
            prototype.Fields.splice(index, 0, "")
            prototype.FieldsType.splice(index, 0, "")
            prototype.FieldsVisibility.splice(index, 0, true)
            prototype.FieldsNillable.splice(index, 0, true)
            prototype.FieldsDocumentation.splice(index, 0, "")
            prototype.FieldsDefaultValue.splice(index, 0, "")

            // Set the fields order.
            for (var i = 0; i < prototype.FieldsOrder.length; i++) {
                if (prototype.FieldsOrder[i] >= index) {
                    prototype.FieldsOrder[i] = prototype.FieldsOrder[i] + 1
                }
            }
            prototype.FieldsOrder.splice(index, 0, index)

            var buttons = EntityPrototypeEditor.displayPrototypePropertie(prototype, "", "", properties, true, index)
            // go in edit mode.
            buttons.editBtn.element.click()
            // Set the focus to the name input.
            buttons.nameInput.element.focus()
        }
    }(this, prototype, properties)

    if (!isEditable) {
        newPropertieBtn.element.style.display = "none"
    }

}

/**
 * Display single propertie.
 */
EntityPrototypeEditor.prototype.displayPrototypePropertie = function (prototype, propertieName, propertieTypeName, parent, isEditable, index) {

    var propertieRow = parent.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;", "id": prototype.TypeName + "_" + propertieName + "_row" }).down()
    var saveBtn = propertieRow.appendElement({ "tag": "div", "class": "entities_btn", "style": "display: none;" }).down().appendElement({ "tag": "i", "class": "fa fa-floppy-o" }).down()
    var editBtn = propertieRow.appendElement({ "tag": "div", "class": "entities_btn", "style": "vertical-align: text-top;" }).down().appendElement({ "tag": "i", "class": "fa fa-pencil-square-o" }).down()

    // The propertie name
    var nameDiv = propertieRow.appendElement({ "tag": "div", "style": "display: table-cell" }).down()
    var nameSpan = nameDiv.appendElement({ "tag": "span", "innerHtml": propertieName.replace("M_", "") }).down()
    var nameEditDiv = nameDiv.appendElement({ "tag": "div", "style": "display: none;" }).down()
    var typeNameLabel = nameEditDiv.appendElement({ "tag": "div", "id": "type_name_label", "innerHtml": "name" }).down()
    var nameInput = nameDiv.appendElement({ "tag": "input", "style": "display: none", "value": propertieName.replace("M_", "") }).down()

    // Now I will append 
    if (prototype.Ids.indexOf(propertieName) != -1) {
        // Here the propertie is an id...
        nameDiv.element.style.color = "#428bca"
        nameInput.element.style.color = "#428bca"
    } else if (prototype.Indexs.indexOf(propertieName) != -1) {
        // Here the propertie is an id...
        nameDiv.element.style.color = "green"
        nameInput.element.style.color = "green"
    }

    // The propertie type name.
    var typeNameDiv = propertieRow.appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
    var typeNameSpan = typeNameDiv.appendElement({ "tag": "span", "innerHtml": propertieTypeName }).down()
    var typeNameEditDiv = typeNameDiv.appendElement({ "tag": "div", "style": "display: none;" }).down()
    var typeNameLabel = typeNameEditDiv.appendElement({ "tag": "div", "id": "type_name_label", "innerHtml": "type" }).down()
    var typeNameInput = typeNameEditDiv.appendElement({ "tag": "input", "style": "display: none", "value": propertieTypeName }).down()

    var defaultValueDiv = propertieRow.appendElement({ "tag": "div", "id": propertieName + "_" + index, "style": "display: none;" }).down()

    // The identity is id or not
    var identityEditDiv = propertieRow.appendElement({ "tag": "div", "style": "display: none;text-align: center;" }).down()
    identityEditDiv.appendElement({ "tag": "div", "innerHtml": "Id" }).down()
    var indentityInput = identityEditDiv.appendElement({ "tag": "input", "type": "checkbox" }).down()

    if (prototype.Ids.indexOf(propertieName) != -1) {
        indentityInput.element.checked = true
    }

    // if the property is use as index.
    var indexEditDiv = propertieRow.appendElement({ "tag": "div", "style": "display: none;text-align: center;" }).down()
    indexEditDiv.appendElement({ "tag": "div", "innerHtml": "index" }).down()
    var indexInput = indexEditDiv.appendElement({ "tag": "input", "type": "checkbox" }).down()
    if (prototype.Indexs.indexOf(propertieName) != -1) {
        indexInput.element.checked = true
    }

    // if the property can be null or not.
    var nullEditDiv = propertieRow.appendElement({ "tag": "div", "style": "display: none;text-align: center;" }).down()
    nullEditDiv.appendElement({ "tag": "div", "innerHtml": "null" }).down()
    var nullInput = nullEditDiv.appendElement({ "tag": "input", "type": "checkbox" }).down()

    if (prototype.FieldsNillable[index] == true) {
        nullInput.element.checked = true
    }

    var visibleEditDiv = propertieRow.appendElement({ "tag": "div", "style": "display: none; text-align: center;" }).down()
    visibleEditDiv.appendElement({ "tag": "div", "innerHtml": "visible" }).down()
    var visibleInput = visibleEditDiv.appendElement({ "tag": "input", "type": "checkbox" }).down()
    if (prototype.FieldsVisibility[index] == true) {
        visibleInput.element.checked = true
    }

    attachAutoComplete(typeNameInput, this.typeNameLst.sort(), false, function (input) {
        return function (value) {
            if (input.element.value.startsWith("[]")) {
                input.element.value = "[]" + value
            } else {
                input.element.value = value
            }
        }
    }(typeNameInput))

    // Append the propertie as identity.
    indentityInput.element.onclick = function (editor, index, nullInput) {
        return function () {
            if (this.checked) {
                var propertieName = editor.getCurrentEntityPrototype().Fields[index]
                if (editor.getCurrentEntityPrototype().Ids.indexOf(propertieName) == -1) {
                    editor.getCurrentEntityPrototype().Ids.push(propertieName)
                    editor.getCurrentEntityPrototype().FieldsToUpdate.push(editor.getCurrentEntityPrototype().Fields[index])
                }
                nullInput.element.checked = false
            } else {
                editor.getCurrentEntityPrototype().Ids.splice(index, 1)
                editor.getCurrentEntityPrototype().FieldsToUpdate.push(editor.getCurrentEntityPrototype().Fields[index])
            }
            editor.saveBtn.element.style.display = "table-cell"
        }
    }(this, index, nullInput)

    // Append the propertie as indexation value.
    indexInput.element.onclick = function (editor, index) {
        return function () {
            if (this.checked) {
                var propertieName = editor.getCurrentEntityPrototype().Fields[index]
                if (editor.getCurrentEntityPrototype().Indexs.indexOf(propertieName) == -1) {
                    editor.getCurrentEntityPrototype().Indexs.push(propertieName)
                    editor.getCurrentEntityPrototype().FieldsToUpdate.push(editor.getCurrentEntityPrototype().Fields[index])
                }
            } else {
                editor.getCurrentEntityPrototype().Indexs.splice(index, 1)
                editor.getCurrentEntityPrototype().FieldsToUpdate.push(editor.getCurrentEntityPrototype().Fields[index])
            }
            editor.saveBtn.element.style.display = "table-cell"
        }
    }(this, index)

    nullInput.element.onclick = function (editor, index) {
        return function () {
            editor.getCurrentEntityPrototype().FieldsNillable[index] = this.checked
            editor.saveBtn.element.style.display = "table-cell"
        }
    }(this, index)

    visibleInput.element.onclick = function (editor, index) {
        return function () {
            editor.getCurrentEntityPrototype().FieldsVisibility[index] = this.checked
            editor.saveBtn.element.style.display = "table-cell"
        }
    }(this, index)

    // Edit attributes.
    var deleteBtn = propertieRow.appendElement({ "tag": "div", "class": "entities_btn" }).down().appendElement({ "tag": "i", "class": "fa fa-trash-o" }).down()

    if (isEditable == false) {
        editBtn.element.style.display = "none"
        deleteBtn.element.style.display = "none"
    } else {
        // Edit button action...
        editBtn.element.onclick = function (nameEditDiv, typeNameEditDiv, identityEditDiv, indexEditDiv, nullEditDiv, visibleEditDiv, nameInput, typeNameInput, typeNameSpan, nameSpan, defaultValueDiv) {
            return function () {
                if (nameInput.element.style.display == "none") {
                    nullEditDiv.element.style.display = "table-cell"
                    visibleEditDiv.element.style.display = "table-cell"
                    indexEditDiv.element.style.display = "table-cell"
                    identityEditDiv.element.style.display = "table-cell"
                    nameEditDiv.element.style.display = "inline"
                    nameInput.element.style.display = "inline"
                    typeNameEditDiv.element.style.display = "inline"
                    typeNameInput.element.style.display = "inline"
                    nameSpan.element.style.display = "none"
                    typeNameSpan.element.style.display = "none"
                    defaultValueDiv.element.style.display = "table-cell"
                    this.style.color = "green"
                } else {
                    nullEditDiv.element.style.display = "none"
                    visibleEditDiv.element.style.display = "none"
                    indexEditDiv.element.style.display = "none"
                    identityEditDiv.element.style.display = "none"
                    nameEditDiv.element.style.display = "none"
                    nameInput.element.style.display = "none"
                    typeNameEditDiv.element.style.display = "none"
                    typeNameInput.element.style.display = "none"
                    nameSpan.element.style.display = "inline"
                    typeNameSpan.element.style.display = "inline"
                    defaultValueDiv.element.style.display = "none"

                    // Set the values
                    nameSpan.element.innerHTML = nameInput.element.value
                    typeNameSpan.element.innerHTML = typeNameInput.element.value
                    this.style.color = ""
                }
            }
        }(nameEditDiv, typeNameEditDiv, identityEditDiv, indexEditDiv, nullEditDiv, visibleEditDiv, nameInput, typeNameInput, typeNameSpan, nameSpan, defaultValueDiv)

        // delete propertie button.
        deleteBtn.element.onclick = function (panel, propertieRow, propertieName, index) {
            return function () {
                propertieRow.removeAllChilds()
                propertieRow.element.parentNode.removeChild(propertieRow.element)
                panel.saveBtn.element.style.display = "table-cell"
                var prototype = panel.getCurrentEntityPrototype()
                if (propertieName.length > 0) {
                    prototype.FieldsToDelete.push(index)
                }

                // Remove the propertie from the prototype.
                prototype.Fields.splice(index, 1)
                prototype.FieldsType.splice(index, 1)
                prototype.FieldsVisibility.splice(index, 1)
                prototype.FieldsNillable.splice(index, 1)
                prototype.FieldsDocumentation.splice(index, 1)
                prototype.FieldsDefaultValue.splice(index, 1)

                // Now I will set the field order...
                for (var i = 0; i < prototype.FieldsOrder.length; i++) {
                    if (prototype.FieldsOrder[i] >= index) {
                        prototype.FieldsOrder[i] = prototype.FieldsOrder[i] - 1
                    }
                }

                // Remove from the field order to.
                prototype.FieldsOrder.splice(index, 1)
            }
        }(this, propertieRow, propertieName, index)

        // Display the save button when propertie change.
        nameInput.element.onchange = typeNameInput.element.onchange = function (panel, nameSpan, nameInput, typeNameSpan, typeNameInput, index, propertieRow, defaultValueDiv) {
            return function () {
                propertieRow.element.id = nameInput.value + "_row"
                panel.saveBtn.element.style.display = "table-cell"
                nameSpan.element.innerHTML = nameInput.value
                typeNameSpan.element.innerHTML = typeNameInput.value

                if (panel.fieldsToUpdate[index] == undefined) {
                    panel.fieldsToUpdate[index] = panel.getCurrentEntityPrototype().Fields[index]
                }

                // So here I will set the value of the propertie.
                var fieldType = typeNameInput.element.value

                // Set in indexs...
                if (panel.getCurrentEntityPrototype().Indexs.indexOf(panel.getCurrentEntityPrototype().Fields[index]) != -1) {
                    panel.getCurrentEntityPrototype().Indexs[panel.getCurrentEntityPrototype().Indexs.indexOf(panel.getCurrentEntityPrototype().Fields[index])] = "M_" + nameInput.element.value
                }

                // Set in ids...
                if (panel.getCurrentEntityPrototype().Ids.indexOf(panel.getCurrentEntityPrototype().Fields[index]) != -1) {
                    panel.getCurrentEntityPrototype().Ids[panel.getCurrentEntityPrototype().Ids.indexOf(panel.getCurrentEntityPrototype().Fields[index])] = "M_" + nameInput.element.value
                }

                panel.getCurrentEntityPrototype().Fields[index] = "M_" + nameInput.element.value
                panel.getCurrentEntityPrototype().FieldsType[index] = fieldType

                // Set the default value editor for the selected field type.
                defaultValueDiv.removeAllChilds()

                // because of the closure I need to get the element from the dom to effectively erease it content.
                document.getElementById(defaultValueDiv.id).innerHTML = ""
                panel.setDefaultValueEditor(defaultValueDiv, panel.getCurrentEntityPrototype(), index, 0)
                panel.setDefaultFieldValue(panel.getCurrentEntityPrototype(), index)
            }
        }(this, nameSpan, nameInput, typeNameSpan, typeNameInput, index, propertieRow, defaultValueDiv)

        nameInput.element.onblur = typeNameInput.element.onblur = function (panel, nameInput, typeNameInput, index) {
            return function () {
                // Push the name of the field with it's old name and it new name.
                if (panel.fieldsToUpdate[index] != undefined) {
                    var toUpdate = panel.fieldsToUpdate[index] + ":M_" + nameInput.element.value
                    if (panel.getCurrentEntityPrototype().FieldsToUpdate.indexOf(toUpdate) == -1) {
                        panel.getCurrentEntityPrototype().FieldsToUpdate.push(toUpdate)
                    }
                }
            }
        }(this, nameInput, typeNameInput, index)
        this.setDefaultValueEditor(defaultValueDiv, prototype, index, 0)
        this.setDefaultFieldValue(prototype, index)
    }
    return { "editBtn": editBtn, "nameInput": nameInput }
}

/**
 * Display the prototype restrictions if there's one...
 */
EntityPrototypeEditor.prototype.displayPrototypeRestrictions = function (prototype) {
    for (var i = 0; i < prototype.SuperTypeNames.length; i++) {
        // Also display parent restrictions.
        var superType = getEntityPrototype(prototype.SuperTypeNames[i])
        this.displayPrototypeRestrictions(superType)
    }
    // Here I will display the list of restriction.
    if (prototype.Restrictions != null) {
        for (var i = 0; i < prototype.Restrictions.length; i++) {
            this.restrictions.element.parentNode.style.display = "block"
            var restriction = prototype.Restrictions[i]
            // Enumeration restriction.
            var restrictionType = ""
            if (restriction.Type == 1) {
                restrictionType = "enumeration"
            } else if (restriction.Type == 2) {
                restrictionType = "fraction digits"
            } else if (restriction.Type == 3) {
                restrictionType = "length"
            } else if (restriction.Type == 4) {
                restrictionType = "max exclusive"
            } else if (restriction.Type == 5) {
                restrictionType = "max inclusive"
            } else if (restriction.Type == 6) {
                restrictionType = "max length"
            } else if (restriction.Type == 7) {
                restrictionType = "min exclusive"
            } else if (restriction.Type == 8) {
                restrictionType = "min inclusive"
            } else if (restriction.Type == 9) {
                restrictionType = "min length"
            } else if (restriction.Type == 10) {
                restrictionType = "pattern"
            } else if (restriction.Type == 11) {
                restrictionType = "total digits"
            } else if (restriction.Type == 12) {
                restrictionType = "white space"
            }

            // The restriction row.
            var restrictionPanel = this.restrictions.getChildById(prototype.TypeName + "_restriction_" + restriction.Value)
            if (this.restrictions.getChildById(prototype.TypeName + "_restriction_" + restriction.Value) == undefined) {
                restrictionPanel = this.restrictions.appendElement({ "tag": "div", "style": "display: table-row;", "id": prototype.TypeName + "_restriction_" + restriction.Value }).down()

                var removeRestrictionBtn = restrictionPanel
                    .appendElement({ "tag": "div", "style": "display: table-cell;", "innerHtml": restriction.Value })
                    .appendElement({ "tag": "div", "style": "display: table-cell;", "innerHtml": restrictionType })
                    .appendElement({ "tag": "div", "class": "entities_btn", "style": "text-align: right;" }).down()
                    .appendElement({ "tag": "i", "class": "fa fa-close" }).down()

                // The remove restriction button.
                removeRestrictionBtn.element.onclick = function (prototype, restriction, restrictionPanel, saveBtn) {
                    return function () {
                        var restrictions = []
                        var needSave = false
                        for (var i = 0; i < prototype.Restrictions.length; i++) {
                            if (prototype.Restrictions[i].Value != restriction.Value) {
                                restrictions.push(prototype.Restrictions[i])
                            } else {
                                needSave = true
                            }
                        }

                        if (needSave) {
                            saveBtn.element.style.display = "table-cell"
                            prototype.Restrictions = restrictions
                            restrictionPanel.element.parentNode.removeChild(restrictionPanel.element)
                        }
                        console.log("Remove " + restriction.Value + " from " + prototype.TypeName)
                    }
                }(prototype, restriction, restrictionPanel, this.saveBtn)
            }

        }
    } else {
        this.restrictions.element.parentNode.style.display = "none"
    }
}

/**
 * Append a new restriction in the list of restriction.
 */
EntityPrototypeEditor.prototype.appendRestriction = function () {
    var prototype = this.getCurrentEntityPrototype()
    this.displayPrototypeRestrictions(prototype)
    if (prototype.Restrictions != null) {
        this.restrictions.element.parentNode.style.display = "block"

        // Now I will append the restriction value input.
        var restrictionsValueInput = this.restrictionEditPanel.appendElement({ "tag": "div", "style": "display: table-cell; vertical-align: middle;" }).down()
            .appendElement({ "tag": "input" }).down()

        // Now the restriction type selector.
        var restrictionsTypeSelect = this.restrictionEditPanel.appendElement({ "tag": "div", "style": "display: table-cell; vertical-align: middle;" }).down()
            .appendElement({ "tag": "select" }).down()

        // Now the buttons.
        var restrictionBtns = this.restrictionEditPanel.appendElement({ "tag": "div", "style": "display: table-cell;" }).down()

        var appendRestrictionBtn = restrictionBtns.appendElement({ "tag": "div", "class": "entities_btn" }).down()
            .appendElement({ "tag": "i", "class": "fa fa-check" }).down()

        var cancelRestrictionBtn = restrictionBtns.appendElement({ "tag": "div", "class": "entities_btn" }).down()
            .appendElement({ "tag": "i", "class": "fa fa-close" }).down()


        // Now I will appdend the restriction type options.
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 1, "innerHtml": "enumeration" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 2, "innerHtml": "fraction digits" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 3, "innerHtml": "length" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 4, "innerHtml": "max exclusive" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 5, "innerHtml": "max inclusive" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 6, "innerHtml": "max length" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 7, "innerHtml": "min exclusive" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 8, "innerHtml": "min inclusive" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 9, "innerHtml": "min length" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 10, "innerHtml": "pattern" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 11, "innerHtml": "total digits" })
        restrictionsTypeSelect.appendElement({ "tag": "option", "value": 12, "innerHtml": "white space" })

        // Set the focus to the input.
        restrictionsValueInput.element.focus()

        // Now the action...
        cancelRestrictionBtn.element.onclick = function (restrictionEditPanel, appendRestrictionBtn) {
            return function () {
                restrictionEditPanel.removeAllChilds()
                appendRestrictionBtn.element.style.display = ""
            }
        }(this.restrictionEditPanel, this.appendRestrictionBtn)

        appendRestrictionBtn.element.onclick = function (prototype, value, type, restrictionEditPanel, appendRestrictionBtn, entityPrototypeEditor) {
            return function () {
                // create the new resctriction.
                var restriction = { "TYPENAME": "Server.Restriction", "Type": type.element.value, "Value": value.element.value }
                if (prototype.Restrictions == undefined) {
                    prototype.Restrictions = []
                }
                prototype.Restrictions.push(restriction)

                // Refresh the restriction.
                entityPrototypeEditor.displayPrototypeRestrictions(prototype)
                entityPrototypeEditor.saveBtn.element.style.display = "block"

                // So here I will append the new restriction.
                restrictionEditPanel.removeAllChilds()
                appendRestrictionBtn.element.style.display = ""
            }
        }(prototype, restrictionsValueInput, restrictionsTypeSelect, this.restrictionEditPanel, this.appendRestrictionBtn, this)
    }
}

/**
 * That function is use to save or create an entity prototype.
 */
EntityPrototypeEditor.prototype.saveprototype = function () {
    // Set the default fields values before saving it.
    this.setDefaultFieldsValue()
    var storeId = this.getCurrentEntityPrototype().TypeName.split(".")[0]
    if (this.getCurrentEntityPrototype().notExist != undefined) {
        // create entity prototype.
        server.entityManager.createEntityPrototype(storeId, this.getCurrentEntityPrototype(),
            // The success callback
            function (result, caller) {
                caller.panel.setCurrentPrototype(result)
            },
            function () {

            }, { "panel": this })
    } else {
        // Save entity prototype here.
        server.entityManager.saveEntityPrototype(storeId, this.getCurrentEntityPrototype(),
            // The success callback
            function (result, caller) {
                caller.panel.setCurrentPrototype(result)
            },
            function () {

            }, { "panel": this })
    }
}

/**
 * Display the default value editor for a given field.
 */
EntityPrototypeEditor.prototype.setDefaultValueEditor = function (defaultValueDiv, prototype, index, level, defaultValue) {
    if (prototype == undefined) {
        return // nothing to do in that case.
    }

    // Set the title at level 0 only.
    if (level == 0) {
        defaultValueDiv.appendElement({ "tag": "div", "innerHtml": "default value" })
    }

    // Cleanup the default value div at first to be sure nothing was present.
    var restrictions = []
    var enumerations = []

    if (defaultValue == undefined) {
        defaultValue = prototype.FieldsDefaultValue[index]
    }

    // If the prototype contain restriction...
    if (prototype.Restrictions != null) {
        restrictions = prototype.Restrictions.slice()
    }

    // The type name used to create the editor.
    var fieldType = prototype.FieldsType[index]

    if (fieldType.startsWith("[]") || fieldType.endsWith(":Ref")) {
        return
    }

    var defaultFieldValueEditor
    if (restrictions.length > 0) {
        for (var i = 0; i < restrictions.length; i++) {
            if (restrictions[i].Type == 1) {
                enumerations.push(restrictions[i].Value)
            }
        }
    }

    var id = prototype.TypeName + "_" + prototype.Fields[index] + "_editor"

    // The field type must not be an array or a reference...
    if (enumerations.length > 0) {
        defaultFieldValueEditor = new Element(null, { "tag": "select", "id": id })
        for (var i = 0; i < enumerations.length; i++) {
            defaultFieldValueEditor.appendElement({ "tag": "option", "value": i + 1, "innerHtml": enumerations[i] })
        }
    } else {

        // Create the input...
        function setDefaultValueInput(id, type, value, step, min, max) {
            var input = new Element(null, { "tag": "input", "id": id })
            input.element.type = type
            input.element.value = value
            input.element.step = step
            input.element.max = max
            input.element.min = min
            return input
        }

        if (isXsId(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "", "undefined")
        } else if (isXsRef(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "", "undefined")
        } else if (isXsInt(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "number", 0, 1)
        } else if (isXsDate(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "date")
        } else if (isXsTime(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "datetime")
        } else if (isXsString(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "", "")
        } else if (isXsBinary(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "", "")
        } else if (isXsBoolean(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "checkbox", false)
        } else if (isXsNumeric(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "number", 0.0, .01)
        } else if (isXsMoney(fieldType)) {
            defaultFieldValueEditor = setDefaultValueInput(id, "number", 0.0, .01)
        } else {
            // Here I will try to see if the fieldType is a base type.
            if (!fieldType.endsWith(":Ref") && !fieldType.startsWith("[]")) {
                // Here the field type is not xs basic type...
                var fieldPrototype = getEntityPrototype(fieldType)
                // Display field name for level superior to 0 ...
                if (level > 0) {
                    defaultValueDiv = defaultValueDiv.appendElement({ "tag": "div", "style": "display: table; width: 100%; padding-left: 4px;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell", "innerHtml": prototype.Fields[index].replace("M_", "") }).up()
                        .appendElement({ "tag": "div", "style": "display: table-row" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell" }).down()
                }
                var subDefaultValueDiv = defaultValueDiv.appendElement({ "tag": "div", "id": id }).down()
                if (fieldPrototype != undefined) {
                    level++
                    for (var i = 3; i < fieldPrototype.Fields.length - 2; i++) {
                        // Recursively create field editor.
                        this.setDefaultValueEditor(subDefaultValueDiv, fieldPrototype, i, level, prototype.FieldsDefaultValue[index])
                    }
                }
            }
        }
    }

    /**
     * Set the default value
     */
    if (defaultFieldValueEditor != undefined) {
        // Display field name for level superior to 0 ...
        if (prototype.Fields[index] != "M_valueOf" && level > 0) {
            defaultValueDiv = defaultValueDiv.appendElement({ "tag": "div", "style": "display: table;  width: 100%;" }).down()
                .appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
                .appendElement({ "tag": "div", "style": "display: table-cell", "innerHtml": prototype.Fields[index].replace("M_", "") })
                .appendElement({ "tag": "div", "style": "display: table-cell" }).down()
        }
        defaultValueDiv.element.appendChild(defaultFieldValueEditor.element)
        defaultFieldValueEditor.element.onchange = function (panel) {
            return function () {
                panel.saveBtn.element.style.display = "table-cell"
            }
        }(this)
    }
}

// Set the content of the FieldsDefaultValue from the value store in the editor.
EntityPrototypeEditor.prototype.setDefaultFieldsValue = function () {
    for (var i = 3; i < this.getCurrentEntityPrototype().Fields.length - 2; i++) {
        var fieldType = this.getCurrentEntityPrototype().FieldsType[i]
        if (fieldType.startsWith("xs.")) {
            var element = document.getElementById(this.getCurrentEntityPrototype().TypeName + "_" + this.getCurrentEntityPrototype().Fields[i] + "_editor")
            if (element != null) {
                this.getCurrentEntityPrototype().FieldsDefaultValue[i] = element.value
            }
        } else {
            // Here the value will be a strcture.
            if (!fieldType.endsWith(":Ref") && !fieldType.startsWith("[]")) {
                this.getCurrentEntityPrototype().FieldsDefaultValue[i] = JSON.stringify(this.getDefaultFieldValue(getEntityPrototype(fieldType)))
            }
        }
    }
}

// That function will return the default value contain in editor.
EntityPrototypeEditor.prototype.getDefaultFieldValue = function (prototype) {
    var values = {}
    values["TYPENAME"] = prototype.TypeName

    for (var i = 3; i < prototype.Fields.length - 2; i++) {
        var fieldType = prototype.FieldsType[i]
        if (fieldType.startsWith("xs.")) {
            var element = document.getElementById(prototype.TypeName + "_" + prototype.Fields[i] + "_editor")
            if (element != null) {
                values[prototype.Fields[i]] = element.value
            }
        } else {
            if (!fieldType.endsWith(":Ref") && !fieldType.startsWith("[]")) {
                values[prototype.Fields[i]] = this.getDefaultFieldValue(getEntityPrototype(fieldType))
            }
        }
    }
    return values
}

// Set the content of the default value in editor.
EntityPrototypeEditor.prototype.setDefaultFieldValue = function (prototype, index) {
    // I will use the default value if there is some.
    var fieldType = prototype.FieldsType[index]
    var field = prototype.Fields[index]
    var defaultValue = prototype.FieldsDefaultValue[index]
    if (defaultValue == "undefined" || defaultValue == undefined || defaultValue.length == 0) {
        return
    }

    var id = prototype.TypeName + "_" + prototype.Fields[index] + "_editor"
    var editor = document.getElementById(id)

    // Set the value...

    // simple type.
    if (fieldType.startsWith("xs.")) {
        editor.value = defaultValue
    } else {
        // Here The value must be an object.
        var obj = JSON.parse(defaultValue)
        function setValue(obj) {
            // So here I will iterate over the properties and set it given editor
            for (var field in obj) {
                if (field.startsWith("M_")) {
                    var value = obj[field]
                    if (isObject(value)) {
                        setValue(value)
                    } else {
                        var id = obj.TYPENAME + "_" + field + "_editor"
                        var editor = document.getElementById(id)
                        if (editor != null) {
                            editor.value = value
                        }
                    }
                }
            }
        }
        setValue(obj)
    }

}