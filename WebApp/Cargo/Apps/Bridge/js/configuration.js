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
                        var view = configurationPanel.contentViews.splice(i, 1)[0]
                        view.panel.element.style.display = "none"
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
                var entity = server.entityManager.entities[evt.dataMap["entity"].UUID]
                var configurationContent = configurationPanel.panel.getChildById("configurationContent")
                var contentView = new EntityPanel(configurationContent, entity.TYPENAME,
                    function (entity, configurationPanel) {
                        return function (contentView) {
                            // Always set the value after the panel was initialysed.
                            contentView.setEntity(entity)
                            contentView.setTitle(configurationPanel.title)
                            contentView.deleteCallback = function (entity) {
                                // Here I will remove the folder if the entity is 
                                // a database...
                                if (entity.TYPENAME == "CargoConfig.DataStoreConfiguration") {
                                    // also remove the data store.
                                    server.dataManager.deleteDataStore(entity.M_id)
                                }
                            }

                            configurationPanel.contentViews.push(contentView)

                            for (var i = 0; i < configurationPanel.contentViews.length; i++) {
                                configurationPanel.contentViews[i].panel.element.style.display = "none"
                            }

                            configurationPanel.currentIndex = configurationPanel.contentViews.length - 1
                            contentView.panel.element.style.display = ""

                            configurationPanel.nextConfigBtn.element.className = "entities_header_btn"
                            configurationPanel.nextConfigBtn.element.style.color = "lightgrey"

                            if (configurationPanel.currentIndex == 1) {
                                configurationPanel.previousConfigBtn.element.className = "entities_header_btn"
                                configurationPanel.previousConfigBtn.element.style.color = "lightgrey"
                            } else {
                                configurationPanel.previousConfigBtn.element.className = "entities_header_btn enabled"
                                configurationPanel.previousConfigBtn.element.style.color = ""
                            }

                            var idField = contentView.getFieldControl("M_id")
                            idField.element.focus()
                        }
                    } (entity, configurationPanel))
            }
        }
    })

    return this
}

ConfigurationPanel.prototype.setConfigurations = function (configurations) {

    // So here I will create the configuration selector...
    this.header = this.panel.appendElement({ "tag": "div", "style": "display: table; margin-top: 2px; margin-bottom: 4px;" }).down()
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

                    var contentView = new EntityPanel(configurationContent, content[j].TYPENAME,
                        function (content, title) {
                            return function (contentView) {
                                // Always set the value after the panel was initialysed.
                                contentView.setEntity(content)
                                contentView.setTitle(title)
                                contentView.hideNavigationButtons()
                                contentView.deleteCallback = function (entity) {
                                    // Here I will remove the folder if the entity is 
                                    // a database...
                                    if (entity.TYPENAME == "CargoConfig.DataStoreConfiguration") {
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
                        } (content[j], this.title))

                    // Here I will append the connection button...
                    contentView.connectBtn = contentView.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled", "style": "display: table-cell; color: lightgrey;" }).down()
                    contentView.connectBtn.appendElement({ "tag": "i", "class": "fa fa-plug" })

                    // Now If the connection is activated...
                    contentView.connectBtn.element.onclick = function(contentView){
                        return function(){
                            // Here I will try to open or close the connection...
                            server.dataManager.ping(contentView.entity.M_id,
                            function(result, caller){
                                // Here the data store can be reach so I will try to connect.

                            }, 
                            function(errMsg, caller){
                                // Here There is an error...
                                
                            }, contentView)
                        }
                    }(contentView)

                    if (j != 0) {
                        contentView.panel.element.style.display = "none"
                    }

                    // keep the index.
                    this.contentViews.push(contentView)
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
                var contentView = new EntityPanel(configurationContent, content.TYPENAME,
                    function (content, title) {
                        return function (contentView) {
                            // Always set the value after the panel was initialysed.
                            contentView.setEntity(content)
                            contentView.setTitle(title)
                            contentView.hideNavigationButtons()
                        }
                    } (content, this.title))
                contentView.deleteBtn.element.style.display = "none"

            }
        }

    }
}