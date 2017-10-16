
/**
 * That class is use to display the structure of information about a given data store.
 */
var DataExplorer = function (parent) {

    // Keep reference to the parent.
    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "data_explorer" }).down()

    // Set the resize event.
    window.addEventListener('resize',
        function (dataExplorer) {
            return function () {
                dataExplorer.resize()
            }
        }(this), true);

    // That contain the map of data aready loaded.
    this.shemasView = {}
    this.prototypesView = {}
    this.configs = {}
    this.newPrototypeBtn = null
    this.storeId = ""

    // I will connect the view to the events.
    // New prototype event
    server.prototypeManager.attach(this, NewPrototypeEvent, function (evt, dataExplorer) {
       // generatePrototypesView = function (storeId, prototypes)
        if (evt.dataMap.prototype.TypeName.startsWith(dataExplorer.storeId)) {
            dataExplorer.prototypesView[evt.dataMap.prototype.TypeName] = new PrototypeTreeView(dataExplorer.shemasView[dataExplorer.storeId], evt.dataMap.prototype)
        }
    })

    // Delete prototype event.
    server.prototypeManager.attach(this, DeletePrototypeEvent, function (evt, dataExplorer) {
        if (evt.dataMap.prototype.TypeName.startsWith(dataExplorer.storeId)) {
            var treeView = dataExplorer.prototypesView[evt.dataMap.prototype.TypeName]
            treeView.panel.element.parentNode.removeChild(treeView.panel.element)
            delete dataExplorer.prototypesView[evt.dataMap.prototype.TypeName]
        }
    })

    return this
}

DataExplorer.prototype.resize = function () {
    var height = this.parent.element.offsetHeight - this.parent.element.firstChild.offsetHeight;
    this.panel.element.style.height = height - 10 + "px"
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.initDataSchema = function (storeConfig, initCallback) {
    // init one time.
    if (this.configs[storeConfig.M_id] != undefined) {
        return
    }

    this.storeId = storeConfig.M_id
    this.configs[storeConfig.M_id] = storeConfig
    this.shemasView[storeConfig.M_id] = new Element(this.panel, { "tag": "div", "class": "shemas_view" })

    // Only display the first panel at first.
    if (Object.keys(this.shemasView).length > 1) {
        this.shemasView[storeConfig.M_id].element.style.display = "none"
    }

    // So here I will get the list of all prototype from a give store and
    // create it's relavite information.
    if (storeConfig.M_dataStoreType == 1) {
        // Sql data store,
        // prototype for sql dataStore are in the 
        server.entityManager.getEntityPrototypes("sql_info." + storeConfig.M_id,
            // success callback.
            function (results, caller) {
                for (var i = 0; i < results.length; i++) {
                    caller.dataExplorer.generatePrototypesView(caller.storeId, results)
                }
                if (caller.initCallback != undefined) {
                    caller.initCallback()
                }
            },
            // error callback.
            function (errMsg, caller) {

            },
            { "dataExplorer": this, "storeId": storeConfig.M_id, "initCallback": initCallback })

    } else if (storeConfig.M_dataStoreType == 2) {
        // Entity data store.
        if (storeConfig.M_id != "sql_info") {
            server.entityManager.getEntityPrototypes(storeConfig.M_id,
                // success callback.
                function (results, caller) {
                    caller.dataExplorer.generatePrototypesView(caller.storeId, results)
                },
                // error callback.
                function (errMsg, caller) {

                },
                { "dataExplorer": this, "storeId": storeConfig.M_id })
        }
    }
}

/**
 * Generate the prototypes view.
 */
DataExplorer.prototype.generatePrototypesView = function (storeId, prototypes) {
    // I will append the append prototype...
    this.newPrototypeBtn = this.shemasView[storeId].appendElement({ "tag": "div", "class": "", "style": "display:block; color: #657383; position: absolute; top:0px;  top: 5px; right: 10px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-plus" }).down()

    // Mouse enter and leave actions...
    this.newPrototypeBtn.element.onmouseenter = function () {
        this.style.color = "#428bca";
        this.style.cursor = "pointer"
    }

    this.newPrototypeBtn.element.onmouseout = function () {
        this.style.color = "#657383";
        this.style.cursor = "default"
    }

    // Now the onclick event.
    this.newPrototypeBtn.element.onclick = function (storeId) {
        return function () {
            // Here I will create the entity prototy editor
            // Here I will create a dialog to enter the new prototype name.
            if (document.getElementById("new_prototype_popup") != null) {
                document.getElementById("new_prototype_popup_input").focus()
                return
            }

            var coord = getCoords(this)
            var dialog = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "class": "popup_div", "id": "new_prototype_popup", "style": "top:" + (coord.top + 15) + "px; left:" + (coord.left + 15) + "px;" })
            var input = dialog.appendElement({ "tag": "div", "style": "display: table; border-spacing:2px 2px;" }).down()
                .appendElement({ "tag": "div", "style": "display: table-row;" }).down()
                .appendElement({ "tag": "div", "style": "display: table-cell;", "innerHtml": "Enter the new protype name" }).up()
                .appendElement({ "tag": "div", "style": "display: table-row;" }).down()
                .appendElement({ "tag": "input", "id": "new_prototype_popup_input", "style": "display: table-cell;" }).down()

            input.element.focus()
            input.element.onkeyup = function (dialog, storeId) {
                return function (evt) {
                    if (evt.keyCode == 13) {
                        // enter key...
                        var prototype = new EntityPrototype()
                        prototype.init({
                            "TypeName": storeId + "." + this.value,
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
                        setEntityPrototype(prototype)

                        evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "prototypeInfo": prototype } }
                        server.eventHandler.broadcastLocalEvent(evt)
                        dialog.element.parentNode.removeChild(dialog.element)
                    } else if (evt.keyCode == 27) {
                        dialog.element.parentNode.removeChild(dialog.element)
                    }
                }
            }(dialog, storeId)

        }
    }(storeId)

    this.panel.element.style.borderTop = "1px solid grey"
    // Here I will create the prototype views...
    for (var i = 0; i < prototypes.length; i++) {
        // Here I will append the prototype name...
        if (this.prototypesView[prototypes[i].TypeName] == undefined) {
            this.prototypesView[prototypes[i].TypeName] = new PrototypeTreeView(this.shemasView[storeId], prototypes[i])
        }
    }
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.setDataSchema = function (storeId) {

    // here I will calculate the height...
    this.resize()

    for (var id in this.shemasView) {
        this.shemasView[id].element.style.display = "none"
    }
    if (this.shemasView[storeId] != undefined) {
        this.shemasView[storeId].element.style.display = ""
    }
}

/**
 * Hide all schema panels
 */
DataExplorer.prototype.hidePanels = function (storeId) {

    // here I will calculate the height...
    for (var id in this.shemasView) {
        this.shemasView[id].element.style.display = "none"
    }
}

/**
 * Hide a given store panel
 */
DataExplorer.prototype.hidePanel = function (storeId) {
    // here I will calculate the height...
    this.shemasView[storeId].element.style.display = "none"
}

/**
 * Show a given data panel.
 */
DataExplorer.prototype.showPanel = function (storeId) {
    if (this.shemasView[storeId] != undefined) {
        this.shemasView[storeId].element.style.display = ""
    } else {
        // I will get the list of prototypes and 
        this.initDataSchema(storeId, function (dataExplorer, storeId) {
            return function () {
                dataExplorer.setDataSchema(storeId)
            }
        }(this, storeId))
    }
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.removeDataSchema = function (storeId) {

    // here I will calculate the height...
    for (var id in this.shemasView) {
        this.shemasView[id].element.style.display = "none"
        if (this.shemasView[id] == storeId) {
            this.shemasView[storeId].element.parentNode.removeChild(this.shemasView[storeId].element)
            delete this.shemasView[storeId]
        }
    }

}

/**
 * That view is use to display prototype structures
 */
var PrototypeTreeView = function (parent, prototype) {
    this.parent = parent
    this.panel = new Element(parent, { "tag": "div", "class": "data_prototype_tree_view" })
    this.fieldsView = {}
    this.prototype = prototype // Keep the prototype reference here.

    // The type name without the imports name.
    var typeName = prototype.TypeName.substring(prototype.TypeName.indexOf(".") + 1)

    // Display the type name and the expand shrink button.
    var header = this.panel.appendElement({ "tag": "div", "class": "data_prototype_tree_view" }).down().appendElement({ "tag": "div", "class": "data_prototype_tree_view_header" }).down()
    header.appendElement()

    /** The expand button */
    this.expandBtn = header.appendElement({ "tag": "i", "class": "fa fa-caret-right", "style": "display:inline;" }).down()

    /** The shrink button */
    this.shrinkBtn = header.appendElement({ "tag": "i", "class": "fa fa-caret-down", "style": "display:none;" }).down()
    this.editLnk = header.appendElement({ "tag": "span", "innerHtml": typeName }).down()

    this.editLnk.element.onmouseover = function () {
        this.style.cursor = "pointer"
        this.style.textDecoration = "underline"
    }

    this.editLnk.element.onmouseout = function () {
        this.style.cursor = "default"
        this.style.textDecoration = ""
    }

    this.editLnk.element.onclick = function (typeName) {
        return function () {
            // So here I will generate an even...
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "prototypeInfo": entityPrototypes[typeName] } }
            server.eventHandler.broadcastLocalEvent(evt)
        }
    }(prototype.TypeName)

    this.fieldsPanel = this.panel.appendElement({ "tag": "div", "class": "data_prototype_tree_view_fields" }).down()

    // The code for display field of a given type.
    this.expandBtn.element.onclick = function (view) {
        return function () {
            view.fieldsPanel.element.style.display = "table"
            this.style.display = "none"
            view.shrinkBtn.element.style.display = ""
        }
    }(this)

    this.shrinkBtn.element.onclick = function (view) {
        return function () {
            view.fieldsPanel.element.style.display = ""
            this.style.display = "none"
            view.expandBtn.element.style.display = ""
        }
    }(this)

    // Now the fields.
    for (var i = 0; i < prototype.Fields.length; i++) {
        if (prototype.Fields[i].startsWith("M_")) {
            this.fieldsView[prototype.Fields[i]] = new PrototypeTreeViewField(this.fieldsPanel, prototype, prototype.Fields[i], prototype.FieldsType[i], prototype.FieldsVisibility[i], prototype.FieldsNillable[i])
        }
    }

    // Update prototype event
    server.prototypeManager.attach(this, UpdatePrototypeEvent, function (evt, prototypeTreeView) {
        // if the current item is the one with change I will reset it content.
        if (evt.dataMap.prototype.TypeName == prototypeTreeView.prototype.TypeName) {
            prototypeTreeView.fieldsView = {}
            prototypeTreeView.fieldsPanel.removeAllChilds()
            prototypeTreeView.prototype = evt.dataMap.prototype // Set the updated prototype version.
            for (var i = 0; i < prototypeTreeView.prototype.Fields.length; i++) {
                if (prototypeTreeView.prototype.Fields[i].startsWith("M_")) {
                    prototypeTreeView.fieldsView[prototypeTreeView.prototype.Fields[i]] = new PrototypeTreeViewField(prototypeTreeView.fieldsPanel, prototypeTreeView.prototype, prototypeTreeView.prototype.Fields[i], prototypeTreeView.prototype.FieldsType[i], prototypeTreeView.prototype.FieldsVisibility[i], prototypeTreeView.prototype.FieldsNillable[i])
                }
            }
        }
    })

    return this
}

/**
 * The view of a given field...
 */
var PrototypeTreeViewField = function (parent, prototype, fieldName, fieldType, isVisible, isNillable) {
    // if is an id...
    var isKey = contains(prototype.Ids, fieldName)
    var isIndex = contains(prototype.Indexs, fieldName)

    // Not display the M_
    fieldName = fieldName.replace("M_", "")

    // The parent prototype.
    this.prototype = prototype

    // The parent panel.
    this.parent = parent

    // Create the new panel.
    this.panel = new Element(parent, { "tag": "div", "class": "data_prototype_tree_view_field" })

    var visibilityClass = ""
    if (isVisible) {
        visibilityClass = "field_visibility visible"
    }

    var className = ""
    if (isKey) {
        className = "field_id"
    } else if (isIndex) {
        className = "field_index"
    }

    // append the the field name.
    this.panel.appendElement({ "tag": "i", "title": "visibility", "class": "fa fa-lightbulb-o " + visibilityClass }).appendElement({ "tag": "span", "innerHtml": fieldName, "class": className }).down()

    // Now the typename 
    this.panel.appendElement({ "tag": "span", "innerHtml": fieldType }).down()

}