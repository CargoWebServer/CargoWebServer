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
                var idField = contentView.getFieldControl("M_id")
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
                // Always set the value after the panel was initialysed.
                contentView.setEntity(content)
                contentView.header.display()

                contentView.saveCallback = function (contentView) {
                    return function (entity) {
                        // So here I will create the new dataStore.
                        if (entity.TYPENAME == "Config.DataStoreConfiguration") {
                            server.dataManager.createDataStore(entity.M_id, entity.M_storeName, entity.M_hostName, entity.M_ipv4, entity.M_port, entity.M_dataStoreType, entity.M_dataStoreVendor,
                                // Success callback
                                function (success, caller) {
                                    // Init the schema informations.
                                    homePage.dataExplorer.initDataSchema(caller.entity, function (contentView) {
                                        return function () {
                                            // display the imports information here...
                                            contentView.connectBtn.element.status = "disconnected"
                                            contentView.connectBtn.element.click()
                                        }
                                    }(caller))
                                },
                                // Error callback
                                function (errObj, caller) {
                                }, contentView)
                        } else if (entity.TYPENAME == "Config.ServiceConfiguration") {
                            console.log("------> service configuration! ", entity)
                        }
                    }
                }(contentView)

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
                    contentView.connectBtn = contentView.header.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button"}).down()
                    contentView.connectBtn.appendElement({ "tag": "i", "class": "fa fa-plug" })

                    // Here I will append in case of sql datasotre the synchornize button.
                    contentView.refreshBtn = contentView.header.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
                    contentView.refreshBtn.appendElement({ "tag": "i", "class": "fa fa-refresh" })

                    // Now If the connection is activated...
                    contentView.connectBtn.element.onclick = function (contentView) {
                        return function () {
                            var entity = entities[contentView.entity.UUID]

                            if (this.status == "error") {
                                this.status = "disconnected"
                            }
                            // Here I will try to open or close the connection...
                            if (this.status == "connected") {
                                server.dataManager.close(entity.M_id,
                                    function (result, caller) {
                                        // Here the data store can be reach so I will try to connect.
                                        caller.connectBtn.style.color = "lightgrey"
                                        caller.connectBtn.status = "disconnected"
                                        homePage.dataExplorer.hidePanel(caller.entity.M_id)
                                        caller.refreshBtn.element.style.display = "none"
                                    },
                                    function (errMsg, caller) {
                                        // Fail to disconnect
                                        caller.connectBtn.style.color = "#8B0000"
                                        caller.connectBtn.status = "error"
                                        caller.refreshBtn.element.style.display = "none"
                                        homePage.dataExplorer.hidePanel(caller.entity.M_id)
                                    }, { "connectBtn": this, "entity": entity, "refreshBtn": contentView.refreshBtn })
                            } else if (this.status == "disconnected") {
                                server.dataManager.connect(entity.M_id,
                                    function (result, caller) {
                                        // Here the data store can be reach so I will try to connect.
                                        caller.connectBtn.style.color = "#4CAF50"
                                        caller.connectBtn.status = "connected"
                                        //caller.refreshBtn.element.style.display = "table-cell"
                                        homePage.dataExplorer.showPanel(caller.entity.M_id)

                                    },
                                    function (errMsg, caller) {
                                        // fail to connect...
                                        caller.connectBtn.style.color = "#8B0000"
                                        caller.connectBtn.status = "error"
                                        caller.refreshBtn.element.style.display = "none"
                                        homePage.dataExplorer.hidePanel(caller.entity.M_id)
                                    }, { "connectBtn": this, "entity": entity, "refreshBtn": contentView.refreshBtn })
                            }

                        }
                    }(contentView)

                    // The refresh action.
                    contentView.refreshBtn.element.onclick = function (contentView) {
                        return function () {
                            var entity = entities[contentView.entity.UUID]
                            this.style.color = "#428bca"
                            server.dataManager.synchronize(entity.M_id,
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
                    var parent = content.getPanel().panel //.parentElement.parentElement

                    // Keep the reference in the content.
                    var actionsDiv = parent
                        .appendElement({ "tag": "div", "id": content.UUID + "_actions_div", "style": "position:absolute; left: 0px; bottom: 0px; overflow-y: auto; overflow-x: hidden;" }).down()

                    // set the scrolling shadow...
                    actionsDiv.element.onscroll = function (header) {
                        return function () {
                            var position = this.scrollTop;
                            if (this.scrollTop > 0) {
                                if (header.className.indexOf(" scrolling") == -1) {
                                    header.className += " scrolling"
                                }
                            } else {
                                header.className = header.className.replaceAll(" scrolling", "")
                            }
                        }
                    }(actionsDiv.parentElement.parentElement.parentElement.element)

                    window.addEventListener('resize',
                        function (actionsDiv) {
                            return function () {
                                var parent = actionsDiv.parentElement.parentElement.parentElement.parentElement
                                var top = parent.element.firstChild.offsetHeight
                                var right = parent.element.clientWidth;
                                if (top > 0 && right > 0) {
                                    actionsDiv.element.style.width = right + "px"
                                    actionsDiv.element.style.top = top + "px"
                                }
                            }
                        }(actionsDiv), true);

                    // Now I will get the list of action for a given services.
                    server.serviceManager.getServiceActions(content.M_id,
                        // success callback
                        function (results, parent) {
                            // Now I will display the list of action in panel.
                            for (var i = 0; i < results.length; i++) {
                                var result = results[i]
                                new EntityPanel(parent, result.TYPENAME, function (entity) {
                                    return function (panel) {
                                        panel.setEntity(entity)
                                        // Now I will display the action documentation correctly...
                                        var documentationInput = CargoEntities.Action_M_doc
                                        //panel.fields["M_doc"].panel.element.style.display = "none"
                                        var doc = panel.fields["M_doc"].value.element.innerText
                                        if (doc.indexOf("@src") != -1) {
                                            doc = doc.split("@src")[0]
                                        }

                                        var values = doc.split("@")
                                        doc = ""
                                        for (var i = 0; i < values.length; i++) {
                                            doc += "<div>"
                                            if (values[i].startsWith("api")) {
                                                doc += values[i].replaceAll("api 1.0", "<span style='color: green;'>api 1.0</span>")
                                            } else if (values[i].startsWith("param") && values[i].indexOf("{callback}") == -1) {
                                                var values_ = values[i].split("param")[1].split(" ")
                                                doc += "<span class='doc_tag' style='vertical-align: top;'>param</span><span>"
                                                var description = ""
                                                for (var j = 1; j < values_.length; j++) {
                                                    if (j == 1) {
                                                        // The type:
                                                        doc += "<span style='color: darkgreen; vertical-align: text-top'>" + values_[j] + "</span>"
                                                    } else if (j == 2) {
                                                        // The name
                                                        doc += "<span style='color: color: #657383; font-weight:bold; vertical-align: text-top'>" + values_[j] + "</span>"
                                                    } else {
                                                        description = description + " " + values_[j]
                                                    }
                                                }
                                                doc += "<span style='vertical-align: text-top'>" + description + "</span>"
                                                doc += "</span>"
                                            }
                                            doc += "</div>"
                                        }
                                        panel.fields["M_doc"].value.element.innerHTML = doc
                                    }
                                }(result), undefined, false, result, "")
                            }
                        },
                        // error callback
                        function (errObj, caller) {

                        }, actionsDiv)
                } else if (content.TYPENAME == "Config.ScheduledTask") {
                    // Here I will personalise input a little.
                    content.getPanel().fields["M_frequency"].panel.element.title = "The task must be execute n time per frequency type (once, daily, weekely, or mouthly). *Is ignore if frenquencyType is ONCE."

                    // The script button must be hidden...
                    content.getPanel().fields["M_script"].panel.element.style.display = "none"
                             
                    var editBtn = content.getPanel().header.panel.appendElement({"tag":"div", "class":"entity_panel_header_button"}).down()
                    editBtn.appendElement({ "tag": "i", "title": "Edit task script.", "class": "fa fa-edit", "style":"padding-top: 2px;" })

                    // The save bnt...
                    content.getPanel().saveCallback = function () {
                        return function (task) {
                            // Now I will schedule the task.
                            server.configurationManager.scheduleTask(task,
                                // Success Callback
                                function (results, caller) {
                                    /** Nothing to do here */
                                },
                                // Error Callback
                                function (errObj, caller) {
                                    console.log(errObj)
                                }, {})
                        }
                    }()

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
                                    },  { "entity": entity })
                            } else {
                                // In that case the script must be save...
                                server.entityManager.createEntity(entityPanel.parentEntity.UUID, entityPanel.ParentLnk, entity,
                                    function (entity, caller) {
                                        // Set the entity.
                                        console.log("Task entity:", entity)
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
        var fieldType = prototype.FieldsType[prototype.getFieldIndex("M_" + this.propertyName)]

        var newConfiguration = function (configurationPanel, configurationContent, configuration) {
            return function () {
                // Here I will create a new entity...
                var entity = eval("new " + configurationPanel.typeName + "()")
                entity.M_id = "New " + configurationPanel.typeName.split(".")[1]

                // Set the entity content.
                var configurationContent = configurationPanel.panel.getChildById("configurationContent")

                // Set the new configuration.
                var contentView = configurationPanel.setConfiguration(configurationContent, entity)

                var idField = contentView.getFieldControl("M_id")

                // Hide the data explorer panel.
                homePage.dataExplorer.hidePanels()

                // Set focus to the id field.
                idField.element.focus()
                idField.element.setSelectionRange(0, idField.element.value.length)

            }
        }(this, configurationContent, configuration)

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

    // Show the first panel if there is one.
    if (this.contentViews[0] != undefined) {
        if (this.contentViews[0].panel != null) {
            this.contentViews[0].panel.element.style.display = ""
        }
    }
}