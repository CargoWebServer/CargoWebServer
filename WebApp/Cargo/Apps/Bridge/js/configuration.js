/**
 * The language mapping... control id/text
 * TODO add language as needed!  
 */
var languageInfo = {
    "en": {
        "M_hostName": "host",
        "M_applicationsPath": "applications",
        "M_dataPath": "data",
        "FrequencyType_MINUTE": "MINUTE",
    },
    "fr": {
        "M_hostName": "hôte",
        "M_applicationsPath": "applications",
        "M_dataPath": "donnée",
        "FrequencyType_MINUTE": "MINUTE",
    }
}

server.languageManager.appendLanguageInfo(languageInfo)

/**
 * The sever configurations.
 * @param parent the parent panel
 * @param type Can be server
 */
var ConfigurationPanel = function (parent, title, typeName, propertyName) {
    /** The parent div */
    this.parent = parent

    /** The panel */
    this.panel = new Element(parent, { "tag": "div", "class": "severConfiguration" })

    /** The type */
    this.typeName = typeName

    /** The name in its parent. */
    this.propertyName = propertyName

    /** The title */
    this.title = title

    /** The configuration selector */
    this.configurationSelect = null
    this.activeConfiguration = null

    /** The header that contain the configuration selector... */
    this.header = null

    /** The new configuration element button */
    this.newConfigElementBtn = null

    /** the current index */
    this.currentIndex = 0
    this.contentViews = []

    /** The navigation bettons... */
    this.nextConfigBtn = null
    this.previousConfigBtn = null

    server.entityManager.attach(this, DeleteEntityEvent, function (evt, configurationPanel) {
        if (evt.dataMap["entity"] !== undefined) {
            if (evt.dataMap["entity"].TYPENAME == configurationPanel.typeName) {
                // so here i will remove the entity from the panel...
                for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                    if (configurationPanel.contentViews[i].entity.UUID == evt.dataMap["entity"].UUID) {
                        // append the the view to the content.
                        var entity = configurationPanel.contentViews[i].entity
                        var view = configurationPanel.contentViews.splice(i, 1)[0]
                        view.panel.element.style.display = "none"
                        if (entity.TYPENAME == "Config.DataStoreConfiguration") {
                            homePage.dataExplorer.removeDataSchema(entity.M_id)
                        }
                        if (view.deleteCallback != undefined) {
                            view.deleteCallback(entity)
                        }
                    }
                }

                // Now I will set the current view if it was active...
                configurationPanel.currentIndex = -1

                // Set to the first item
                configurationPanel.nextConfigBtn.element.click()
            }
        }
    })

    server.entityManager.attach(this, NewEntityEvent, function (evt, configurationPanel) {
        if (evt.dataMap["entity"] !== undefined) {
            if (evt.dataMap["entity"].TYPENAME == configurationPanel.typeName) {
                // Hide all data panel.
                homePage.dataExplorer.hidePanels()

                var entity = entities[evt.dataMap["entity"].UUID]
                var configurationContent = configurationPanel.panel.getChildById("configurationContent")

                // Set the new configuration.
                var contentView = configurationPanel.setConfiguration(configurationContent, entity)
                contentView.header.display()
            }
        }
    })

    server.entityManager.attach(this, UpdateEntityEvent, function (evt, configurationPanel) {
        if (evt.dataMap["entity"] !== undefined) {
            if (evt.dataMap["entity"].TYPENAME == configurationPanel.typeName) {
                // Hide all data panel.
                for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                    if (configurationPanel.contentViews[i].entity.UUID == evt.dataMap["entity"].UUID) {
                        // display the header.
                        configurationPanel.contentViews[i].header.display()
                    }
                }
            }
        }
    })

    return this
}

/**
 * Append a new configuration to the configuration panel.
 */
ConfigurationPanel.prototype.setConfiguration = function (configurationContent, content) {

    // If the view already exist
    for (var i = 0; i < this.contentViews.length; i++) {
        if (this.contentViews[i].entity !== undefined) {
            if (this.contentViews[i].entity.M_id == content.M_id) {
                this.contentViews[i].setEntity(content)
                if (this.contentViews[i].connectBtn !== undefined) {
                    this.contentViews[i].connectBtn.status = "disconnected"
                    this.contentViews[i].connectBtn.element.click()
                }
                this.contentViews[i].panel.element.style.display = ""
                return this.contentViews[i]
            } else {
                this.contentViews[i].panel.element.style.display = "none"
            }
        }
    }

    // In that case I will create a new entity panel.
    var contentView = new EntityPanel(configurationContent, content.TYPENAME,
        function (content, title) {
            return function (contentView) {
                if (content.TYPENAME == "Config.DataStoreConfiguration") {
                    // The delete callback function.
                    contentView.deleteCallback = function (entity) {
                        // Here I will remove the folder if the entity is 
                        // a database...
                        if (entity.TYPENAME == "Config.DataStoreConfiguration") {
                            // also remove the data store.
                            server.dataManager.deleteDataStore(entity.M_id,
                                // success callback
                                function () {

                                },
                                // error callback.
                                function () {

                                }, this)

                        }
                    }
                }

                // Always set the value after the panel was initialysed.
                contentView.setEntity(content)
                contentView.header.display()

                // The data store configuration.
                if (content.TYPENAME == "Config.DataStoreConfiguration") {
                    // So here I will set the schema view for the releated store.
                    // Here I have a service configuration.
                    if (content.UUID != undefined) {
                        if (content.UUID.length != 0) {
                            // Set only if is not a new.
                            homePage.dataExplorer.initDataSchema(content)
                        }
                    }

                    // Here I will broadcast a local event, export data menue need that information.
                    var evt = {
                        "code": NewDataStoreEvent, "name": DataEvent,
                        "dataMap": { "storeConfig": content }
                    }
                    server.eventHandler.broadcastLocalEvent(evt)

                    // Here I will append the connection button...
                    contentView.connectBtn = contentView.header.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
                    contentView.connectBtn.appendElement({ "tag": "i", "class": "fa fa-plug" })

                    // Now If the connection is activated...
                    contentView.connectBtn.element.onclick = function (contentView) {
                        return function () {
                            var entity = entities[contentView.entity.UUID]
                            if (this.status == "error") {
                                this.status = "disconnected"
                                this.style.color = "#8B0000"
                            }

                            // Here I will try to open or close the connection...
                            if (this.status == "connected") {
                                server.dataManager.close(entity.M_id,
                                    function (result, caller) {
                                        // Here the data store can be reach so I will try to connect.
                                        caller.connectBtn.style.color = "lightgrey"
                                        caller.connectBtn.status = "disconnected"
                                        homePage.dataExplorer.hidePanel(caller.entity.M_id)
                                    },
                                    function (errMsg, caller) {
                                        // Fail to disconnect
                                        caller.connectBtn.style.color = "#8B0000"
                                        caller.connectBtn.status = "error"
                                        homePage.dataExplorer.hidePanel(caller.entity.M_id)
                                    }, { "connectBtn": this, "entity": entity })

                            } else if (this.status == "disconnected") {
                                server.dataManager.connect(entity.M_id,
                                    function (result, caller) {
                                        // Here the data store can be reach so I will try to connect.
                                        caller.connectBtn.style.color = "#4CAF50"
                                        caller.connectBtn.status = "connected"
                                        homePage.dataExplorer.showPanel(caller.entity.M_id)
                                    },
                                    function (errMsg, caller) {
                                        // fail to connect...
                                        caller.connectBtn.style.color = "#8B0000"
                                        caller.connectBtn.status = "error"
                                        homePage.dataExplorer.hidePanel(caller.entity.M_id)
                                        if (errMsg.dataMap.errorObj.M_id == "DATASTORE_DOESNT_EXIST_ERROR") {
                                            var entity = caller.entity
                                            var contentView = caller.contentView
                                            // In that case I will try to create the data store...
                                            server.dataManager.createDataStore(entity.M_id, entity.M_storeName, entity.M_hostName, entity.M_ipv4, entity.M_port, entity.M_dataStoreType, entity.M_dataStoreVendor,
                                                // Success callback
                                                function (success, caller) {
                                                    // Init the schema informations.
                                                    server.dataManager.connect(entity.M_id,
                                                        // success callback.
                                                        function (result, caller) {
                                                            // Here the data store can be reach so I will try to connect.
                                                            caller.connectBtn.element.style.color = "#4CAF50"
                                                            caller.connectBtn.element.status = "connected"
                                                            homePage.dataExplorer.initDataSchema(caller.entity, function (contentView) {
                                                                return function () {
                                                                    // display the imports information here...
                                                                    contentView.connectBtn.element.status = "disconnected"
                                                                    contentView.connectBtn.element.click()
                                                                }
                                                            }(caller))
                                                        },
                                                        // error callback
                                                        function (errObj, caller) {
                                                            caller.connectBtn.element.status = "error"
                                                            caller.connectBtn.element.style.color = "#8B0000"
                                                        }, caller)
                                                },
                                                // Error callback
                                                function (errObj, caller) {
                                                }, contentView)
                                        }
                                    }, { "connectBtn": this, "entity": entity, "contentView": contentView })
                            }

                        }
                    }(contentView)

                    // Set the connection status
                    server.dataManager.ping(contentView.entity.M_id,
                        function (result, caller) {
                            // Here the data store can be reach so I will try to connect.
                            caller.style.color = "#4CAF50"
                            caller.status = "connected"
                        },
                        function (errMsg, caller) {
                            // Here There is an error...
                            caller.style.color = "#8B0000"
                            caller.status = "error"
                        }, contentView.connectBtn.element)

                } else if (content.TYPENAME == "Config.ServerConfiguration") {
                    // If the content is server configuration I will also append the change admin password option.
                    // That pannel will be use to change admin password.
                    configurationContent.appendElement({ "tag": "div", "class": "panel" }).down()
                        .appendElement({ "tag": "div", "style": "display:table-row; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "id": "adminPasswordChange", "class": "entity_panel" }).down()
                        .appendElement({ "tag": "div", "class": "entity" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "span", "innerHtml": "current password:" }).up()
                        .appendElement({ "tag": "div", "style": "display: table-cell; position: relative;" }).down()
                        .appendElement({ "tag": "input", "type": "password", "style": "width: 100%;", "id": "currentPwd" }).up().up()
                        .appendElement({ "tag": "div", "class": "entity" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "span", "type": "password", "innerHtml": "new password:" }).up()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "input", "type": "password", "style": "width: 100%;", "id": "newPwd" }).up().up()
                        .appendElement({ "tag": "div", "class": "entity" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "span", "innerHtml": "confirm password:" }).up()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "input", "type": "password", "style": "width: 100%;", "id": "confirmPwd" }).up().up().up().up()
                        .appendElement({ "tag": "div", "class": "entity", "style": "text-align: right;" }).down()
                        .appendElement({ "tag": "div", "id": "changeAdminPwdBtn", "innerHtml": "ok" })

                    var currentPwd = configurationContent.getChildById("currentPwd")
                    var newPwd = configurationContent.getChildById("newPwd")
                    var confirmPwd = configurationContent.getChildById("confirmPwd")
                    var changeAdminPwdBtn = configurationContent.getChildById("changeAdminPwdBtn")

                    // Validation here.

                    // So here I will use a decorator to display message to the input...
                    // The first parameter is the message to use by default...
                    setValidator("", currentPwd, function (currentPwd) {
                        return function (msgDiv) {
                            if (currentPwd.element.value.length == 0) {
                                msgDiv.element.innerHTML = "The password must contain a value!"
                                currentPwd.element.focus()
                                return false
                            }
                        }
                    }(currentPwd), 3000)

                    setValidator("", newPwd, function (newPwd) {
                        return function (msgDiv) {
                            if (currentPwd.element.value.length == 0) {
                                msgDiv.element.innerHTML = "The password must contain a value!"
                                newPwd.element.focus()
                                return false
                            }
                        }
                    }(newPwd), 3000)

                    setValidator("", confirmPwd, function (newPwd) {
                        return function (msgDiv) {
                            if (confirmPwd.element.value.length == 0) {
                                msgDiv.element.innerHTML = "The password must contain a value!"
                                confirmPwd.element.focus()
                                return false
                            }
                        }
                    }(confirmPwd), 3000)

                    setValidator("", confirmPwd, function (newPwd, confirmPwd) {
                        return function (msgDiv) {
                            if (confirmPwd.element.value != newPwd.element.value) {
                                msgDiv.element.innerHTML = "The tow values enter for password does not match!"
                                newPwd.element.focus()
                                newPwd.element.style.border = "1px solid red"
                                newPwd.element.value = ""
                                confirmPwd.element.value = ""
                                newPwd.element.onkeydown = function (newPwd, confirmPwd) {
                                    return function () {
                                        newPwd.element.style.border = ""
                                        confirmPwd.element.style.border = ""
                                    }
                                }(newPwd, confirmPwd)
                                return false
                            }
                        }
                    }(newPwd, confirmPwd), 3000)

                    changeAdminPwdBtn.element.onclick = function (currentPwd, newPwd, confirmPwd) {
                        return function () {
                            if (confirmPwd.element.value == newPwd.element.value) {
                                // Now I will set the new admin password.
                                server.securityManager.changeAdminPassword(currentPwd.element.value, newPwd.element.value,
                                    // success callback
                                    function (results, caller) {

                                    },
                                    // error callback
                                    function (errMsg, caller) {

                                    }, {})
                            }
                        }
                    }(currentPwd, newPwd, confirmPwd)
                } else if (content.TYPENAME == "Config.OAuth2Configuration") {
                    // So here I will append other element in the view here.

                } else if (content.TYPENAME == "Config.ServiceConfiguration") {
                    // Here I have a service configuration.
                    homePage.serviceExplorer.initService(content)

                } else if (content.TYPENAME == "Config.ScheduledTask") {
                    // Here I will personalise input a little.
                    content.getPanel().fields["M_frequency"].panel.element.title = "The task must be execute n time per frequency type (once, daily, weekely, or mouthly). *Is ignore if frenquencyType is ONCE."

                    // The script button must be hidden...
                    content.getPanel().fields["M_script"].panel.element.style.display = "none"

                    var editBtn = content.getPanel().header.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
                    editBtn.appendElement({ "tag": "i", "title": "Edit task script.", "class": "fa fa-edit", "style": "padding-top: 2px;" })

                    // The save bnt...
                    content.getPanel().saveCallback = function (panel) {
                        return function (task) {
                            // Now I will schedule the task.
                            server.configurationManager.scheduleTask(task,
                                // Success Callback
                                function (entity, caller) {
                                    // set the title and expand the panel.
                                    if (caller != undefined) {
                                        caller.header.title.element.innerText = caller.entity.M_id
                                        caller.header.expandBtn.element.click()
                                    }
                                },
                                // Error Callback
                                function (errObj, caller) {
                                    console.log(errObj)
                                }, panel)
                        }
                    }(content.getPanel())

                    content.getPanel().deleteCallback = function (task) {
                        // If the task is delete I will delete it file.
                        server.entityManager.getEntityById("CargoEntities.File", "CargoEntities", [task.M_id], false,
                            function (fileEntity, caller) {
                                server.entityManager.removeEntity(fileEntity.UUID,
                                    function (result, fileUUID) {
                                        // send close event to editor.
                                        var evt = { "code": CloseEntityEvent, "name": FileEvent, "dataMap": { "fileId": fileUUID } }
                                        server.eventHandler.broadcastLocalEvent(evt)
                                    },
                                    function () {

                                    }, fileEntity)
                            },
                            function () {

                            }, {})
                    }

                    editBtn.element.onclick = function (ScheduledTask_M_script, entityPanel) {
                        return function () {
                            var entity = entityPanel.entity
                            if (entity.UUID != undefined) {
                                server.entityManager.getEntityById("CargoEntities.File", "CargoEntities", [entity.M_id], false,
                                    function (file, caller) {
                                        caller.entity.M_script = file.M_id
                                        server.entityManager.saveEntity(caller.entity) // Save the entity...
                                        evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
                                        server.eventHandler.broadcastLocalEvent(evt)
                                    },
                                    function (errObj, caller) {
                                        var file = new CargoEntities.File()
                                        file.M_id = caller.entity.M_id
                                        file.M_name = caller.entity.M_id + ".js"
                                        file.M_isDir = false
                                        file.M_fileType = 1
                                        file.M_mime = "application/javascript"
                                        file.M_modeTime = Date.now()

                                        server.entityManager.saveEntity(file,
                                            function (file, caller) {
                                                caller.entity.M_script = file.M_id
                                                server.entityManager.saveEntity(caller.entity) // Save the entity...
                                                evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
                                                server.eventHandler.broadcastLocalEvent(evt)
                                            },
                                            function () {

                                            }, caller)
                                    }, { "entity": entity })
                            } else {
                                // In that case the script must be save...
                                server.entityManager.createEntity(entityPanel.parentEntity.UUID, entityPanel.ParentLnk, entity,
                                    function (entity, caller) {
                                        // Set the entity.
                                        entityPanel.setEntity(entity)
                                        caller.editBtn.click()
                                    },
                                    function () {

                                    }, { "editBtn": this, "entityPanel": entityPanel })
                            }
                        }
                    }(content.getPanel().fields["M_script"].panel, content.getPanel())



                } else if (content.TYPENAME == "Config.LdapConfiguration") {
                    // Here I will append in case of sql datasotre the synchornize button.
                    contentView.refreshBtn = contentView.header.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
                    contentView.refreshBtn.appendElement({ "tag": "i", "class": "fa fa-refresh" })

                    // The refresh action.
                    contentView.refreshBtn.element.onclick = function (contentView) {
                        return function () {
                            var entity = entities[contentView.entity.UUID]
                            this.style.color = "#428bca"
                            server.ldapManager.synchronize(entity.M_id,
                                // success callback
                                function (results, caller) {
                                    console.log("synchronization success!")
                                    caller.refreshBtn.element.style.color = "#4CAF50"
                                },
                                // error callback
                                function (errObj, caller) {
                                    console.log("synchronization fail!", error)
                                    caller.refreshBtn.element.style.color = "#8B0000"
                                }, { "refreshBtn": contentView.refreshBtn })
                        }
                    }(contentView)

                    // TODO display users, groups and computers here...
                }
            }

        }(content, this.title))

    // Set parent entity informations.
    contentView.parentEntity = this.activeConfiguration
    if (contentView.entity != null) {
        if (contentView.entity.TYPENAME == "Config.DataStoreConfiguration") {
            contentView.ParentLnk = "M_dataStoreConfigs"
        } else if (contentView.entity.TYPENAME == "Config.ServerConfiguration") {
            contentView.ParentLnk = "M_serverConfig"
        } else if (contentView.entity.TYPENAME == "Config.ServiceConfiguration") {
            contentView.ParentLnk = "M_serviceConfigs"
        } else if (contentView.entity.TYPENAME == "Config.SmtpConfiguration") {
            contentView.ParentLnk = "M_smtpConfigs"
        } else if (contentView.entity.TYPENAME == "Config.LdapConfiguration") {
            contentView.ParentLnk = "M_ldapConfigs"
        } else if (contentView.entity.TYPENAME == "Config.ApplicationConfiguration") {
            contentView.ParentLnk = "M_applicationConfigs"
        } else if (contentView.entity.TYPENAME == "Config.OAuth2Configuration") {
            contentView.ParentLnk = "M_oauth2Configuration"
        } else if (contentView.entity.TYPENAME == "Config.ScheduledTask") {
            contentView.ParentLnk = "M_scheduledTasks"
        }
    }

    // keep the index.
    this.contentViews.push(contentView)

    return contentView
}

/**
 * Set the configuration panel itself.
 */
ConfigurationPanel.prototype.setConfigurations = function (configurations) {

    // So here I will create the configuration selector...
    this.header = this.panel.appendElement({ "tag": "div", "class": "panel entity" }).down()
    this.configurationSelect = this.header
        .appendElement({ "tag": "div", "innerHtml": "Configurations" })
        .appendElement({ "tag": "select", "style": "margin-left: 5px; flex-grow: 1;" }).down()

    for (var i = 0; i < configurations.length; i++) {
        if (i == 0) {
            // The first configuration as the default one.
            this.activeConfiguration = configurations[i]
        }
        var configuration = configurations[i]
        this.configurationSelect.appendElement({ "tag": "option", "value": configuration.M_id, "innerHtml": configuration.M_name })

        var configurationContent = this.panel.appendElement({ "tag": "div", "id": "configurationContent" }).down()
        var content = configuration["M_" + this.propertyName]
        var prototype = configuration.getPrototype()
        var fieldName = "M_" + this.propertyName
        var fieldType = prototype.FieldsType[prototype.getFieldIndex(fieldName)]

        var newConfiguration = function (configurationPanel, configurationContent, configuration, fieldType, fieldName) {
            return function () {
                // Here I will create a new entity...
                var entity = eval("new " + configurationPanel.typeName + "()")
                entity.M_id = "New " + configurationPanel.typeName.split(".")[1]
                entity.ParentUuid = configuration.UUID
                entity.ParentLnk = fieldName

                // Return the parent link...
                entity.getParent = function (parent) {
                    return function () {
                        return parent
                    }
                }(configuration)

                // Set the entity content.
                var configurationContent = configurationPanel.panel.getChildById("configurationContent")

                // Set the new configuration.
                var contentView = configurationPanel.setConfiguration(configurationContent, entity)
                contentView.header.expandBtn.element.click()

                var idField = contentView.fields["M_id"]

                // Hide the data explorer panel.
                homePage.dataExplorer.hidePanels()

                // Set focus to the id field.
                idField.renderer.renderer.element.click()
                idField.renderer.renderer.element.focus()
                idField.editor.editor.element.setSelectionRange(0, idField.editor.editor.element.value.length)

            }
        }(this, configurationContent, configuration, fieldType, fieldName)

        // In case of multiple configurations element..
        if (fieldType.startsWith("[]")) {

            // Set an empty array if none exist.
            if (content == undefined) {
                content = []
            }

            // The new configuration button.
            this.header.appendElement({ "tag": "div", "class": "entity_panel_header" })

            this.newConfigElementBtn = this.header.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
            this.newConfigElementBtn.appendElement({ "tag": "i", "class": "fa fa-plus", "style": "" })

            // I will append the navigation button i that case...
            this.previousConfigBtn = this.header.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
            this.previousConfigBtn.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-left" })

            this.nextConfigBtn = this.header.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
            this.nextConfigBtn.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-right" })

            if (content.length > 1) {
                this.nextConfigBtn.element.className += " enabled"
            }

            // Here the configuration panel contain more than one panel...
            for (var j = 0; j < content.length; j++) {
                if (content[j] != undefined) {
                    this.setConfiguration(configurationContent, content[j])
                    this.contentViews[j].panel.element.style.display = "none"
                }
            }

            // The next configuration button.
            this.nextConfigBtn.element.onclick = function (configurationPanel) {
                return function () {
                    // Here I will display the next element
                    if (configurationPanel.currentIndex < configurationPanel.contentViews.length - 1) {
                        for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                            configurationPanel.contentViews[i].panel.element.style.display = "none"
                        }
                        configurationPanel.currentIndex++
                        configurationPanel.contentViews[configurationPanel.currentIndex].panel.element.style.display = ""

                        if (configurationPanel.contentViews[configurationPanel.currentIndex].entity.TYPENAME == "Config.DataStoreConfiguration") {
                            homePage.dataExplorer.setDataSchema(configurationPanel.contentViews[configurationPanel.currentIndex].entity.M_id)
                        }

                        if (configurationPanel.contentViews[configurationPanel.currentIndex].entity.TYPENAME == "Config.ServiceConfiguration") {
                            homePage.serviceExplorer.setService(configurationPanel.contentViews[configurationPanel.currentIndex].entity.M_id)
                        }

                        if (configurationPanel.currentIndex == configurationPanel.contentViews.length - 1) {
                            configurationPanel.nextConfigBtn.element.className = "entity_panel_header_button"
                        } else {
                            configurationPanel.nextConfigBtn.element.className = "entity_panel_header_button"
                        }

                        configurationPanel.previousConfigBtn.element.className = "entity_panel_header_button"
                    }
                    if (configurationPanel.contentViews.length <= 1) {
                        // disable the next button
                        configurationPanel.nextConfigBtn.element.className = "entity_panel_header_button"

                        // disable the previous button.
                        configurationPanel.previousConfigBtn.element.className = "entity_panel_header_button"
                    }
                }
            }(this)

            // The previous configuration button.
            this.previousConfigBtn.element.onclick = function (configurationPanel) {
                return function () {
                    if (configurationPanel.currentIndex > 0) {
                        for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                            configurationPanel.contentViews[i].panel.element.style.display = "none"
                        }
                        configurationPanel.currentIndex--
                        configurationPanel.contentViews[configurationPanel.currentIndex].panel.element.style.display = ""

                        if (configurationPanel.contentViews[configurationPanel.currentIndex].entity.TYPENAME == "Config.DataStoreConfiguration") {
                            homePage.dataExplorer.setDataSchema(configurationPanel.contentViews[configurationPanel.currentIndex].entity.M_id)
                        }

                        if (configurationPanel.contentViews[configurationPanel.currentIndex].entity.TYPENAME == "Config.ServiceConfiguration") {
                            homePage.serviceExplorer.setService(configurationPanel.contentViews[configurationPanel.currentIndex].entity.M_id)
                        }

                        if (configurationPanel.currentIndex == 0) {
                            configurationPanel.previousConfigBtn.element.className = "entity_panel_header_button"
                        } else {
                            configurationPanel.previousConfigBtn.element.className = "entity_panel_header_button"
                        }
                        configurationPanel.nextConfigBtn.element.className = "entity_panel_header_button"
                    }
                    if (configurationPanel.contentViews.length <= 1) {
                        // disable the next button
                        configurationPanel.nextConfigBtn.element.className = "entity_panel_header_button"
                        // disable the previous button.
                        configurationPanel.previousConfigBtn.element.className = "entity_panel_header_button"
                    }
                }
            }(this)

            // Set the new configuration click handler.
            this.newConfigElementBtn.element.onclick = newConfiguration

        } else {
            if (content != undefined) {
                if (isObject(content)) {
                    if (this.contentViews[0] == null) {
                        this.newConfigElementBtn = this.header.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
                        this.newConfigElementBtn.appendElement({ "tag": "i", "class": "fa fa-plus", "style": "" })

                        // Set the new configuration click handler.
                        this.newConfigElementBtn.element.onclick = newConfiguration
                    }
                    this.setConfiguration(configurationContent, content)
                }
            }
        }
    }

    // Show the first panel if there is one.
    if (this.contentViews[0] != undefined) {
        if (this.contentViews[0].panel != null) {
            this.contentViews[0].panel.element.style.display = ""
        }
    }
}