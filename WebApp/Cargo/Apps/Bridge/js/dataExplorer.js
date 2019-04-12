
/**
 * That class is use to display the structure of information about a given data store.
 */
var DataExplorer = function (parent) {

    // Keep reference to the parent.
    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "data_explorer" }).down()

    this.panel.element.onscroll = function (header) {
        return function () {
            if (this.scrollTop > 0) {
                if (header.className.indexOf(" scrolling") == -1) {
                    header.className += " scrolling"
                }
            } else {
                header.className = header.className.replaceAll(" scrolling", "")
            }
        }
    }(this.parent.element.firstChild)

    // Set the resize event.
    window.addEventListener('resize',
        function (dataExplorer) {
            return function () {
                dataExplorer.resize()
            }
        }(this), true);

    // That contain the map of data aready loaded.
    this.schemasView = {}
    this.prototypesView = {}
    this.configs = {}
    this.newPrototypeBtn = null

    // I will connect the view to the events.
    // New prototype event
    server.prototypeManager.attach(this, NewPrototypeEvent, function (evt, dataExplorer) {
        var storeId = evt.dataMap.prototype.TypeName.split(".")[0]
        if (dataExplorer.schemasView[storeId] !== undefined) {
            dataExplorer.prototypesView[evt.dataMap.prototype.TypeName] = new PrototypeTreeView(dataExplorer.schemasView[storeId], evt.dataMap.prototype)
        }
    })

    // Delete prototype event.
    server.prototypeManager.attach(this, DeletePrototypeEvent, function (evt, dataExplorer) {
        var storeId = evt.dataMap.prototype.TypeName.split(".")[0]
        if (dataExplorer.schemasView[storeId] !== undefined) {
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
    /*if (this.configs[storeConfig.M_id] != undefined) {
        return
    }*/

    this.storeId = storeConfig.M_id
    this.configs[storeConfig.M_id] = storeConfig
    this.schemasView[storeConfig.M_id] = new Element(this.panel, { "tag": "div", "class": "shemas_view" })

    // Only display the first panel at first.
    if (Object.keys(this.schemasView).length > 1) {
        this.schemasView[storeConfig.M_id].element.style.display = "none"
    }

    // So here I will get the list of all prototype from a give store and
    // create it's relavite information.
    if (storeConfig.M_dataStoreType == 1) {
        // Sql data store,
        // prototype for sql dataStore are in the
        server.entityManager.getEntityPrototypes(storeConfig.M_id,
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

/**
 * Generate the prototypes view.
 */
DataExplorer.prototype.generatePrototypesView = function (storeId, prototypes) {
    // I will append the append prototype...
    this.newPrototypeBtn = this.schemasView[storeId].appendElement({ "tag": "div", "class": "", "style": "display:block; color: #657383; position: absolute; top:0px;  top: 5px; right: 10px;" }).down()
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
            this.prototypesView[prototypes[i].TypeName] = new PrototypeTreeView(this.schemasView[storeId], prototypes[i], this.configs[storeId].M_dataStoreType)
        }
    }
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.setDataSchema = function (storeId) {

    // here I will calculate the height...
    this.resize()

    for (var id in this.schemasView) {
        this.schemasView[id].element.style.display = "none"
    }
    if (this.schemasView[storeId] != undefined) {
        this.schemasView[storeId].element.style.display = ""
    }
}

/**
 * Hide all schema panels
 */
DataExplorer.prototype.hidePanels = function (storeId) {

    // here I will calculate the height...
    for (var id in this.schemasView) {
        this.schemasView[id].element.style.display = "none"
    }
}

/**
 * Hide a given store panel
 */
DataExplorer.prototype.hidePanel = function (storeId) {
    // here I will calculate the height...
    if (this.schemasView[storeId] != undefined) {
        this.schemasView[storeId].element.style.display = "none"
    }
}

/**
 * Show a given data panel.
 */
DataExplorer.prototype.showPanel = function (storeId) {
    if (this.schemasView[storeId] != undefined) {
        this.schemasView[storeId].element.style.display = ""
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
    for (var id in this.schemasView) {
        this.schemasView[id].element.style.display = "none"
        if (this.schemasView[id] == storeId) {
            this.schemasView[storeId].element.parentNode.removeChild(this.schemasView[storeId].element)
            delete this.schemasView[storeId]
        }
    }

}

/**
 * That view is use to display prototype structures
 */
var PrototypeTreeView = function (parent, prototype, storeType) {
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
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "prototypeInfo": getEntityPrototype(typeName) } }
            server.eventHandler.broadcastLocalEvent(evt)
        }
    }(prototype.TypeName)

    // Here I will append the contextual menu...
    this.editLnk.element.addEventListener('contextmenu', function (editLnk, prototype, storeType) {
        return function (evt) {
            // Rename existing folder in the project.
            var renameMenuItem = new MenuItem("rename_menu", "Rename", {}, 0, function (editLnk, prototype) {
                return function () {
                    // Now I will hide the fileDiv...
                    var text = editLnk.element.innerText
                    var parent = editLnk.element.parentNode
                    var renameDirInput = new Element(null, { "tag": "input", "value": text })
                    editLnk.element.style.display = "none"

                    // Insert at position 3
                    parent.insertBefore(renameDirInput.element, parent.childNodes[3])

                    renameDirInput.element.onclick = function (evt) {
                        evt.stopPropagation()
                        // nothing todo here.
                    }

                    renameDirInput.element.onkeyup = function (renameDirInput, editLnk, text, prototype) {
                        return function (evt) {
                            evt.stopPropagation()
                            if (evt.keyCode == 27) {
                                // escape key
                                renameDirInput.element.parentNode.removeChild(renameDirInput.element)
                                editLnk.element.style.display = "inline"
                            } else if (evt.keyCode == 13) {
                                // Rename the file here.
                                storeId = prototype.TypeName.split(".")[0]
                                server.entityManager.renameEntityPrototype(storeId + "." + this.value, prototype, storeId,
                                    // Success callback
                                    function (result, renameDirInput) {
                                        // Remove the rename dir input.
                                        renameDirInput.element.parentNode.removeChild(renameDirInput.element)
                                    },
                                    // Error callback
                                    function () {

                                    }, renameDirInput)
                            }
                        }
                    }(renameDirInput, editLnk, text)

                    renameDirInput.element.setSelectionRange(0, text.length)
                    renameDirInput.element.focus()
                }
            }(editLnk, prototype, storeType), "fa fa-edit")

            // Delete a file from a project.
            var deleteMenuItem = new MenuItem("delete_menu", "Delete", {}, 0, function (prototype) {
                return function () {
                    var confirmDialog = new Dialog(randomUUID(), undefined, true)
                    confirmDialog.div.element.style.maxWidth = "450px"
                    confirmDialog.setCentered()
                    server.languageManager.setElementText(confirmDialog.title, "Delete file")
                    confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete prototype " + prototype.TypeName + "?" })
                    confirmDialog.ok.element.onclick = function (dialog, prototype) {
                        return function () {
                            // Remove the folder.
                            var typeName = prototype.TypeName
                            server.entityManager.deleteEntityPrototype(typeName, typeName.split(".")[0],
                                function (result, caller) {
                                },
                                function () {

                                }, undefined)
                            dialog.close()
                        }
                    }(confirmDialog, prototype)
                }
            }(prototype), "fa fa-trash-o")

            var dumpDataMenuItem = new MenuItem("dump_data_menu", "Display Data", {}, 0, function (prototype, storeType) {
                return function () {
                    // In that case I will create a local query file and open it with the query...
                    if (storeType == 1) {
                        // Sql dump
                        createQuery(".sql", "/** Sql query **/\n", function () {
                            return function (file) {
                                server.fileManager.openFile(file.M_id,
                                    // Progress callback.
                                    function (index, totatl, caller) {

                                    },
                                    // Success callback
                                    function (result, caller) {

                                    },
                                    // Error callback
                                    function (errMsg, caller) {

                                    }, this)
                            }
                        }())
                    } else if (storeType == 2) {
                        // Entities dump
                        createQuery(".eql", "/** Eql query **/\n", function () {
                            return function (file) {
                                server.fileManager.openFile(file.M_id,
                                    // Progress callback.
                                    function (index, totatl, caller) {

                                    },
                                    // Success callback
                                    function (result, caller) {
                                        // Here The file is open...
                                    },
                                    // Error callback
                                    function (errMsg, caller) {

                                    }, this)
                            }
                        }())
                    }
                }
            }(prototype, storeType), "fa fa-search")

            // The main menu will be display in the body element, so nothing will be over it.
            if (storeType == 2 && !prototype.TypeName.startsWith("xs.") && !prototype.TypeName.startsWith("Config.") && !prototype.TypeName.startsWith("CargoEntities.") && !prototype.TypeName.startsWith("sqltypes.") && !prototype.TypeName.startsWith("XMI_types.")) {
                var contextMenu = new PopUpMenu(editLnk, [dumpDataMenuItem, "|", renameMenuItem, deleteMenuItem], evt)
            } else {
                var contextMenu = new PopUpMenu(editLnk, [dumpDataMenuItem], evt)
            }
        }
    }(this.editLnk, prototype, storeType), false)

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
        if (evt.dataMap.prototype.UUID == prototypeTreeView.prototype.UUID) {
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


// That funtion create a new file with a query in it.
function createQuery(extension, query, callback) {
    // So here I will create a new query file.
    server.fileManager.getFileByPath("/queries",
        // Success
        function (results, caller) {
            var extension = caller.extension

            // query file will have a name like q1, q2... qx by default...
            var lastIndex = 0
            for (var i = 0; i < results.M_files.length; i++) {
                var f = results.M_files[i]
                if (f.M_name.match(/q[0-9]+/)) {
                    if (parseInt(f.M_name.replace("q", "").replace(extension, "")) > lastIndex) {
                        lastIndex = parseInt((f.M_name).replace("q", "").replace(extension, ""))
                    }
                }
            }
            lastIndex++

            // Here I will create an empty text file.
            var f = null
            try {
                var f = new File([query], "q" + lastIndex + extension, { type: "text/plain", lastModified: new Date(0) })
            } catch (error) {
                /** Nothing todo here. */
            }

            // Now I will create the new file...
            server.fileManager.createFile("q" + lastIndex + extension, "/queries", f, 256, 256, false,
                // Success callback.
                function (result, caller) {
                    // Here is the new file...
                    if (caller.callback != undefined) {
                        // Call the callback function with the 
                        // newly create file as it first argument.
                        caller.callback(result)
                    }
                },
                function () {

                },
                // Error callback.
                function () {

                }, caller)
        },
        // Error
        function () {

        }, { "extension": extension, "callback": callback })
}