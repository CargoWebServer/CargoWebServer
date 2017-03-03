/**
 * The language mapping... control id/text
 * TODO add language as needed!  
 */
var languageInfo = {
    "en": {
        "M_hostName": "host",
        "M_applicationsPath": "applications",
        "M_dataPath": "data",
    },
    "fr": {
        "M_hostName": "hôte",
        "M_applicationsPath": "applications",
        "M_dataPath": "donnée",
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
        if (evt.dataMap["entity"] != undefined) {
            if (evt.dataMap["entity"].TYPENAME == configurationPanel.typeName) {
                // so here i will remove the entity from the panel...
                for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                    if (configurationPanel.contentViews[i].entity.UUID == evt.dataMap["entity"].UUID) {
                        // append the the view to the content.
                        var entity = configurationPanel.contentViews[i].entity
                        var view = configurationPanel.contentViews.splice(i, 1)[0]
                        view.panel.element.style.display = "none"
                        if (entity.TYPENAME == "Config.DataStoreConfiguration") {
                            homepage.dataExplorer.removeDataSchema(entity.M_id)
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
        if (evt.dataMap["entity"] != undefined) {
            if (evt.dataMap["entity"].TYPENAME == configurationPanel.typeName) {
                // Hide the currently displayed view.
                for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                    configurationPanel.contentViews[i].panel.element.style.display = "none"
                }

                // Hide all data panel.
                homepage.dataExplorer.hidePanels()

                var entity = server.entityManager.entities[evt.dataMap["entity"].UUID]
                var configurationContent = configurationPanel.panel.getChildById("configurationContent")

                // Set the new configuration.
                var contentView = configurationPanel.setConfiguration(configurationContent, entity)
                var idField = contentView.getFieldControl("M_id")

                // Set focus to the id field.
                idField.element.focus()
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
        if (this.contentViews[i].entity.UUID == content.UUID) {
            this.contentViews[i].setEntity(content)
            if (this.contentViews[i].connectBtn != undefined) {
                this.contentViews[i].connectBtn.status = "disconnected"
                this.contentViews[i].connectBtn.element.click()
            }
            return this.contentViews[i]
        }
    }

    // In that case I will create a new entity panel.
    var contentView = new EntityPanel(configurationContent, content.TYPENAME,
        function (content, title) {
            return function (contentView) {
                // Always set the value after the panel was initialysed.
                contentView.setEntity(content)
                contentView.setTitle(title)
                contentView.hideNavigationButtons()

                contentView.saveCallback = function (contentView) {
                    return function (entity) {
                        // So here I will create the new dataStore.
                        if (entity.TYPENAME == "Config.DataStoreConfiguration") {
                            server.dataManager.createDataStore(entity.M_id, entity.M_dataStoreType, entity.M_dataStoreVendor,
                                // Success callback
                                function (success, caller) {
                                    // Init the schema informations.
                                    homepage.dataExplorer.initDataSchema(caller.entity, function (contentView) {
                                        return function () {
                                            // display the schemas information here...
                                            contentView.connectBtn.element.status = "disconnected"
                                            contentView.connectBtn.element.click()
                                        }
                                    } (caller))
                                },
                                // Error callback
                                function (errObj, caller) {
                                }, contentView)
                        }
                    }
                } (contentView)

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
                    homepage.dataExplorer.initDataSchema(content)

                    // Here I will append the connection button...
                    contentView.connectBtn = contentView.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled", "style": "display: table-cell; color: lightgrey;" }).down()
                    contentView.connectBtn.appendElement({ "tag": "i", "class": "fa fa-plug" })

                    // Now If the connection is activated...
                    contentView.connectBtn.element.onclick = function (UUID) {
                        return function () {
                            var entity = server.entityManager.entities[UUID]

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
                                        homepage.dataExplorer.hidePanel(caller.entity.M_id)
                                    },
                                    function (errMsg, caller) {
                                        // Fail to disconnect
                                        caller.connectBtn.style.color = "#8B0000"
                                        caller.connectBtn.status = "error"
                                        homepage.dataExplorer.hidePanel(caller.entity.M_id)
                                    }, { "connectBtn": this, "entity": entity })
                            } else if (this.status == "disconnected") {
                                server.dataManager.connect(entity.M_id,
                                    function (result, caller) {
                                        // Here the data store can be reach so I will try to connect.
                                        caller.connectBtn.style.color = "#4CAF50"
                                        caller.connectBtn.status = "connected"
                                        homepage.dataExplorer.showPanel(caller.entity.M_id)
                                    },
                                    function (errMsg, caller) {
                                        // fail to connect...
                                        caller.connectBtn.style.color = "#8B0000"
                                        caller.connectBtn.status = "error"
                                        homepage.dataExplorer.hidePanel(caller.entity.M_id)
                                    }, { "connectBtn": this, "entity": entity })
                            }

                        }
                    } (contentView.entity.UUID)

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
                    configurationContent.appendElement({ "tag": "div", "style": "display:table; border-top: 1px solid grey; padding: 5px 0px 5px 2px; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display:table-row; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell; padding: 2px;color: white; background-color: #bbbbbb;", "innerHtml": "Change admin password" }).up()
                        .appendElement({ "tag": "div", "style": "display:table-row; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "id": "adminPasswordChange" }).down()
                        .appendElement({ "tag": "div", "style": "display:table-row; width:100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "span", "innerHtml": "current password:" }).up()
                        .appendElement({ "tag": "div", "style": "display: table-cell; position: relative;" }).down()
                        .appendElement({ "tag": "input", "type": "password", "style": "width: 100%;", "id": "currentPwd" }).up().up()
                        .appendElement({ "tag": "div", "style": "display:table-row; width:100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "span", "type": "password", "innerHtml": "new password:" }).up()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "input", "type": "password", "style": "width: 100%;", "id": "newPwd" }).up().up()
                        .appendElement({ "tag": "div", "style": "display:table-row; width:100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "span", "innerHtml": "confirm password:" }).up()
                        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
                        .appendElement({ "tag": "input", "type": "password", "style": "width: 100%;", "id": "confirmPwd" }).up().up().up().up()
                        .appendElement({ "tag": "div", "style": "display:table; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display:table-row; width: 100%;" }).down()
                        .appendElement({ "tag": "div", "style": "display: table-cell; width:100%;" })
                        .appendElement({ "tag": "div", "id": "changeAdminPwdBtn", "style": "display: table-cell; with: 50px", "innerHtml": "ok" })

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
                    } (currentPwd), 3000)

                    setValidator("", newPwd, function (newPwd) {
                        return function (msgDiv) {
                            if (currentPwd.element.value.length == 0) {
                                msgDiv.element.innerHTML = "The password must contain a value!"
                                newPwd.element.focus()
                                return false
                            }
                        }
                    } (newPwd), 3000)

                    setValidator("", confirmPwd, function (newPwd) {
                        return function (msgDiv) {
                            if (confirmPwd.element.value.length == 0) {
                                msgDiv.element.innerHTML = "The password must contain a value!"
                                confirmPwd.element.focus()
                                return false
                            }
                        }
                    } (confirmPwd), 3000)

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
                                } (newPwd, confirmPwd)
                                return false
                            }
                        }
                    } (newPwd, confirmPwd), 3000)

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
                    } (currentPwd, newPwd, confirmPwd)
                }

            }

        } (content, this.title))

    // keep the index.
    this.contentViews.push(contentView)
    return contentView
}

/**
 * Set the configuration panel itself.
 */
ConfigurationPanel.prototype.setConfigurations = function (configurations) {

    // So here I will create the configuration selector...
    this.header = this.panel.appendElement({ "tag": "div", "style": "display: table; margin-top: 2px; margin-bottom: 4px; width: 100%;" }).down()
    this.configurationSelect = this.header.appendElement({ "tag": "div", "style": "display: table-cell; vertical-align: middle;", "innerHtml": "Configurations" })
        .appendElement({ "tag": "div", "style": "display: table-cell; vertical-align: middle;" }).down()
        .appendElement({ "tag": "select", "style": "margin-left: 5px;" }).down()

    for (var i = 0; i < configurations.length; i++) {
        var configuration = configurations[i]
        this.configurationSelect.appendElement({ "tag": "option", "value": configuration.M_id, "innerHtml": configuration.M_name })

        var configurationContent = this.panel.appendElement({ "tag": "div", "id": "configurationContent", "style": "display: table;" }).down()
        var content = configuration["M_" + this.propertyName]
        var prototype = server.entityManager.entityPrototypes[configuration.TYPENAME]
        var fieldType = prototype.FieldsType[prototype.getFieldIndex("M_" + this.propertyName)]

        // In case of multiple configurations element..
        if (fieldType.startsWith("[]")) {

            // Set an empty array if none exist.
            if (content == undefined) {
                content = []
            }

            // The new configuration button.
            this.header.appendElement({ "tag": "div", "style": "display: table-cell; width: 100%;" })

            this.newConfigElementBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled", "style": "display: table-cell;color: #657383;" }).down()
            this.newConfigElementBtn.appendElement({ "tag": "i", "class": "fa fa-plus", "style": "" })

            // I will append the navigation button i that case...
            this.previousConfigBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn", "style": "display: table-cell; color:lightgrey;" }).down()
            this.previousConfigBtn.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-left" })

            this.nextConfigBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn", "style": "display: table-cell; color:lightgrey;" }).down()
            this.nextConfigBtn.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-right" })

            if (content.length > 1) {
                this.nextConfigBtn.element.className += " enabled"
                this.nextConfigBtn.element.style.color = "#657383"
            }

            // Here the configuration panel contain more than one panel...
            for (var j = 0; j < content.length; j++) {
                if (content[j] != undefined) {
                    this.setConfiguration(configurationContent, content[j])
                    if (j != 0) {
                        this.contentViews[j].panel.element.style.display = "none"
                    }
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
                            homepage.dataExplorer.setDataSchema(configurationPanel.contentViews[configurationPanel.currentIndex].entity.M_id)
                        }

                        if (configurationPanel.currentIndex == configurationPanel.contentViews.length - 1) {
                            configurationPanel.nextConfigBtn.element.className = "entities_header_btn"
                            configurationPanel.nextConfigBtn.element.style.color = "lightgrey"
                        } else {
                            configurationPanel.nextConfigBtn.element.className = "entities_header_btn enabled"
                            configurationPanel.nextConfigBtn.element.style.color = "#657383"
                        }

                        configurationPanel.previousConfigBtn.element.className = "entities_header_btn enabled"
                        configurationPanel.previousConfigBtn.element.style.color = "#657383"
                    }
                    if (configurationPanel.contentViews.length <= 1) {
                        // disable the next button
                        configurationPanel.nextConfigBtn.element.className = "entities_header_btn"
                        configurationPanel.nextConfigBtn.element.style.color = "lightgrey"
                        // disable the previous button.
                        configurationPanel.previousConfigBtn.element.className = "entities_header_btn"
                        configurationPanel.previousConfigBtn.element.style.color = "lightgrey"
                    }

                }
            } (this)

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
                            homepage.dataExplorer.setDataSchema(configurationPanel.contentViews[configurationPanel.currentIndex].entity.M_id)
                        }

                        if (configurationPanel.currentIndex == 0) {
                            configurationPanel.previousConfigBtn.element.className = "entities_header_btn"
                            configurationPanel.previousConfigBtn.element.style.color = "lightgrey"
                        } else {
                            configurationPanel.previousConfigBtn.element.className = "entities_header_btn enabled"
                            configurationPanel.previousConfigBtn.element.style.color = "#657383"
                        }
                        configurationPanel.nextConfigBtn.element.className = "entities_header_btn enabled"
                        configurationPanel.nextConfigBtn.element.style.color = "#657383"
                    }
                    if (configurationPanel.contentViews.length <= 1) {
                        // disable the next button
                        configurationPanel.nextConfigBtn.element.className = "entities_header_btn"
                        configurationPanel.nextConfigBtn.element.style.color = "lightgrey"
                        // disable the previous button.
                        configurationPanel.previousConfigBtn.element.className = "entities_header_btn"
                        configurationPanel.previousConfigBtn.element.style.color = "lightgrey"
                    }
                }
            } (this)

            // Now the append configuration button.
            this.newConfigElementBtn.element.onclick = function (configurationPanel, configurationContent, configuration) {
                return function () {
                    // Here I will create a new entity...
                    var entity = eval("new " + configurationPanel.typeName + "()")
                    entity.UUID = configurationPanel.typeName + "%" + randomUUID()
                    entity.M_id = "New " + configurationPanel.typeName.split(".")[1]

                    server.entityManager.setEntity(entity)

                    server.entityManager.createEntity(configuration.UUID, "M_" + configurationPanel.propertyName, configurationPanel.typeName, entity.M_id, entity,
                        function (result, caller) {
                            /** The interface update is made by the new entity event handler... */
                        },
                        function (errMsg) {

                        }, configurationPanel)
                }
            } (this, configurationContent, configuration)

        } else {
            if (content != undefined) {
                this.setConfiguration(configurationContent, content)
            }
        }
    }
}