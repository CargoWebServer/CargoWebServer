var languageInfo = {
    "en": {
    },
    "fr": {
    }
}

server.languageManager.appendLanguageInfo(languageInfo)

/**
 * When the user is logged in this is the page to display.
 */
var HomePage = function () {

    /** Ref to the parent div **/
    this.parent = null

    /** Create the panel but not append it to it parent yet */
    this.panel = null

    /** header, where to put menu, working file etc... */
    this.headerDiv = null

    /** A table layout where to put elements. */
    this.menuContentDiv

    /** The session panel **/
    this.sessionPanel = null

    /** this.workingFilesDiv */
    this.workingFilesDiv = null

    /** The main section */
    this.mainArea = null

    /** The workspace div **/
    this.workspaceDiv = null

    /** The project div */
    this.projectDiv = null

    /** The vertical context selector */
    this.contextSelector = null

    /** The project explorer */
    this.projectExplorer = null

    /** Use to display the content of data store. */
    this.dataExplorer = null

    /** The file navigation  */
    this.fileNavigator = null

    /** The code editor */
    this.codeEditor = null

    /** The bpmn diagram explorer. */
    this.bpmnExplorer = null

    /** The configurations... */
    this.serverConfiguration = null
    this.servicesConfiguration = null
    this.dataConfiguration = null
    this.ldapConfiguration = null
    this.smtpConfiguration = null

    /** The propertie div a the right */
    this.propertiesDiv = null
    this.propertiesView = null

    // The main menu.
    this.mainMenu = null

    // The toolbar div
    this.toolbarDiv = null

    homepage = this

    return this
}

HomePage.prototype.init = function (parent, sessionInfo) {
    this.parent = parent
    this.parent.removeAllChilds()
    this.panel = this.parent.appendElement({ "tag": "div", "class": "home_page" }).down()

    /////////////////////////////////// Header section ///////////////////////////////////

    // On the right side there will be the menu and the workspace...
    // The menu grid who will contain the menu panel...
    this.headerDiv = this.panel.appendElement({ "tag": "div", "class": "header" }).down()

    /////////////////////////////////// Menu section ///////////////////////////////////
    var menuRow = this.headerDiv.appendElement({ "tag": "div", "style": "width:100%; height: 30px; display: table-row" }).down()

    // The toolbar file grid...
    this.toolbarDiv = this.headerDiv.appendElement({ "tag": "div", "id": "toolbarDiv" }).down()

    // This is where the menu grid will be put...
    this.menuContentDiv = menuRow.appendElement({ "tag": "div", "style": "width:100%; display: table-cell;" }).down()
        .appendElement({ "tag": "div", "style": "width:100%; display: table; height: 0px;" }).down()

    // Now I will create the session panel...
    this.sessionPanel = new SessionPanel(menuRow.appendElement({ "tag": "div", "style": "width:100%; display: table-cell; height:30px" }).down(), sessionInfo)

    // That funtion create a new file with a query in it.
    function createQuery(extension) {
        // So here I will create a new query file.
        server.fileManager.getFileByPath("/queries",
            // Progress...
            function () {

            },
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
                    var f = new File([""], "q" + lastIndex + extension, { type: "text/plain", lastModified: new Date(0) })
                } catch (error) {
                    f = new Blob([""], { type: "text/plain" });
                    f.name = "test.txt"
                    f.lastModifiedDate = new Date(0);
                }

                // Now I will create the new file...
                server.fileManager.createFile("q" + lastIndex + extension, "/queries", f, 256, 256, false,
                    // Success callback.
                    function (result, caller) {
                        // Here is the new file...
                        console.log(result)
                    },
                    // Progress callback
                    function () {

                    },
                    // Error callback.
                    function () {

                    }, caller)
            },
            // Error
            function () {

            }, {"extension":extension})
    }

    // Create a new Entity Query File.


    // Entity Query Language File.
    var newEqlQueryMenuItem = new MenuItem("new_eql_query_menu_item", "EQL Query", {}, 1, function (extension) {
        return function () { createQuery(extension) }
    } (".eql"), "fa fa-file-o")

    // Structured Query Language Query Language File.
    var newSqlQueryMenuItem = new MenuItem("new_sql_query_menu_item", "SQL Query", {}, 1, function (extension) {
        return function () { createQuery(extension) }
    } (".sql"), "fa fa-file-o")

    var newProjectMenuItem = new MenuItem("new_project_menu_item", "New Project...", {}, 1,
        function (homepage) {
            return function () {
                homepage.createNewProject()
            }
        } (this), "fa fa-files-o")

    // The new menu in the file menu
    var newFileMenuItem = new MenuItem("new_file_menu_item", "New", { "new_project_menu_item": newProjectMenuItem }, 1)

    // Now the import data menu
    var importXsdSchemaMenuItem = new MenuItem("import_xsd_menu_item", "XSD schema", {}, 2, function (parent) {
        return function () {
            var fileExplorer = parent.appendElement({ "tag": "input", "type": "file", "accept": ".xsd, .XSD, .Xsd", "multiple": "", "style": "display: none;" }).down()
            fileExplorer.element.onchange = function (bpmnExplorer) {
                return function (evt) {
                    var files = evt.target.files; // FileList object
                    for (var i = 0, f; f = files[i]; i++) {
                        //server.dataManager.importXsdSchema(f)
                        var reader = new FileReader();
                        /** I will read the file content... */
                        reader.onload = function (file) {
                            return function (e) {
                                // Now I will load the content of the file.
                                server.dataManager.importXsdSchema(file.name, e.target.result)
                            }
                        } (f);
                        reader.readAsText(f);
                    }
                }
            } (this)
            // Display the file explorer...
            fileExplorer.element.click()
        }
    } (parent), "fa fa-file-o")

    var importXmlDataMenuItem = new MenuItem("import_xml_menu_item", "XML data", {}, 2, function (parent) {
        return function () {
            var fileExplorer = parent.appendElement({ "tag": "input", "type": "file", "accept": ".xml, .XML, .dae, .DAE", "multiple": "", "style": "display: none;" }).down()
            fileExplorer.element.onchange = function (bpmnExplorer) {
                return function (evt) {
                    var files = evt.target.files; // FileList object
                    for (var i = 0, f; f = files[i]; i++) {
                        //server.dataManager.importXsdSchema(f)
                        var reader = new FileReader();
                        /** I will read the file content... */
                        reader.onload = function (e) {
                            var text = e.target.result
                            // Now I will load the content of the file.
                            server.dataManager.importXmlData(text,
                                function (result, caller) {
                                    /** Nothing todo the the action will be in the event listener. */
                                },
                                function (errMsg, caller) {

                                }, {})
                        };
                        reader.readAsText(f);
                    }
                }
            } (this)
            // Display the file explorer...
            fileExplorer.element.click()
        }
    } (parent), "fa fa-file-o")

    var importDataMenuItem = new MenuItem("import_data_menu_item", "Import", { "import_xsd_menu_item": importXsdSchemaMenuItem, "import_xml_menu_item": importXmlDataMenuItem }, 1)

    // The new menu in the data Menu
    var newDataMenuItem = new MenuItem("new_data_menu_item", "New", { "new_eql_query_menu_item": newEqlQueryMenuItem, "new_sql_query_menu_item": newSqlQueryMenuItem }, 1)

    var closeServerItem = new MenuItem("close_server_menu_item", "Close server", {}, 1, function () { server.stop() }, "fa fa-power-off")

    var fileMenuItem = new MenuItem("file_menu", "File", { "new_file_menu_item": newFileMenuItem, "close_server_menu_item": closeServerItem }, 0)

    var editMenuItem = new MenuItem("edit_menu", "Edit", {}, 0)

    var dataMenuItem = new MenuItem("data_menu", "Data", { "import_data_menu_item": importDataMenuItem, "new_data_menu_item": newDataMenuItem }, 0)

    // The main menu will be display in the body element, so nothing will be over it.
    this.mainMenu = new VerticalMenu(new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; top:2px;" }), [fileMenuItem, dataMenuItem, editMenuItem])

    /////////////////////////////////// workspace section  ///////////////////////////////////
    this.mainArea = this.panel.appendElement({ "tag": "div", "style": "display: table; width:100%; height:100%" }).down()

    // Now the left and right div...
    var splitArea1 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-cell; position: relative; height:100%" }).down()
    var leftDiv = new Element(splitArea1, { "tag": "div", "id": "leftDiv", "style": "" })
    var splitter1 = this.mainArea.appendElement({ "tag": "div", "class": "splitter vertical", "id": "splitter1" }).down()

    var splitArea2 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-cell; position: relative; width: 100%; height:100%" }).down()
    var rightDiv = new Element(splitArea2, { "tag": "div", "id": "rightDiv" })

    // Init the splitter action.
    initSplitter(splitter1, leftDiv, 50)

    // The workspace area
    this.workspaceDiv = new Element(rightDiv, { "tag": "div", "class": "workspace_div" })

    // The working file grid...
    this.workingFilesDiv = this.workspaceDiv.appendElement({ "tag": "div", "id": "workingFilesDiv" }).down()
        .appendElement({ "tag": "div", "style": "width:100%; display: inline; position: relative" }).down()
    this.fileNavigator = new FileNavigator(this.workingFilesDiv)
    // The code editor...
    this.codeEditor = new CodeEditor(this.workspaceDiv)

    // The context selector is nothing more than a simple div...
    this.contextSelector = new Element(splitArea1, { "tag": "div", "class": "contextSelector" })

    // The configuration...
    function setSelectAction(button, div) {
        button.element.onclick = function (div, leftDiv) {
            return function () {
                // if the button is already active that mean the user want to expand or shring the 
                // navigation panel.
                if (this.firstChild.className.indexOf("active") > -1) {
                    if (leftDiv.element.style.width == "50px") {
                        var keyframe = "100% { width:431px;}"
                        leftDiv.animate(keyframe, .5,
                            function (leftDiv) {
                                return function () {
                                    leftDiv.element.style.width = "431px"
                                }
                            } (leftDiv))
                    } else {
                        var keyframe = "100% { width:50px;}"
                        leftDiv.element.style.overflowY = "hidden"
                        leftDiv.animate(keyframe, .5,
                            function (leftDiv) {
                                return function () {
                                    leftDiv.element.style.width = "50px"
                                    leftDiv.element.style.overflowY = "hidden"
                                }
                            } (leftDiv))
                    }
                    return
                }

                leftDiv.element.style.width = ""

                var divs = document.getElementsByClassName("navigation_div")
                for (var i = 0; i < divs.length; i++) {
                    divs[i].style.display = "none"
                }

                div.element.style.display = ""

                var buttons = document.getElementsByClassName("navigation_btn")
                for (var i = 0; i < buttons.length; i++) {
                    buttons[i].firstChild.className = buttons[i].firstChild.className.replace(" active", "")
                    if (buttons[i].firstChild.id == "workflowImg") {
                        buttons[i].firstChild.src = "img/workflow.svg"
                    }
                }

                this.firstChild.className += " active"
                if (this.firstChild.id == "workflowImg") {
                    this.firstChild.src = "img/workflow_blue.svg"
                }

                homepage.dataExplorer.resize()

            }
        } (div, leftDiv)
    }

    // Now I will append the button inside the context selector.
    // The project.
    this.projectDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px;" })
    this.projectExplorer = new ProjectExplorer(this.projectDiv)
    this.projectContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Projects" }).appendElement({ "tag": "i", "class": "fa fa-files-o active" })
    setSelectAction(this.projectContext, this.projectDiv)

    // Workflow manager service interface here.
    if (server.workflowManager != null) {
        // The bpmn explorer...
        this.bpmnDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display:none;" })
        this.bpmnContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Workflow Manager" })
            .appendElement({ "tag": "img", "id": "workflowImg", "src": "img/workflow.svg" })

        var workflowImg = this.bpmnContext.getChildById("workflowImg")
        workflowImg.element.onmouseover = function () {
            if (this.className.indexOf("active") == -1) {
                this.src = "img/workflow_hover.svg"
            }
        }

        workflowImg.element.onmouseleave = function () {
            if (this.className.indexOf("active") == -1) {
                this.src = "img/workflow.svg"
            }
        }
        setSelectAction(this.bpmnContext, this.bpmnDiv)
        this.bpmnExplorer = new BpmnExplorer(this.bpmnDiv)
    }

    // The server context...
    this.serverSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display: none;" })
    this.serverSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Server Configuration" }).appendElement({ "tag": "i", "class": "fa fa-ship" })
    setSelectAction(this.serverSettingContext, this.serverSettingDiv)
    this.serverConfiguration = new ConfigurationPanel(this.serverSettingDiv, "Server configuration", "Config.ServerConfiguration", "serverConfig")

    // The security context...
    this.securityDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": " left:50px; display:none;" })
    this.securityContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "security" }).appendElement({ "tag": "i", "class": "fa fa-shield" })
    setSelectAction(this.securityContext, this.securityDiv)

    // The services context...
    this.servicesSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display: none;" })
    this.serviceSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Services Configuration" }).appendElement({ "tag": "i", "class": "fa fa-server" })
    setSelectAction(this.serviceSettingContext, this.servicesSettingDiv)
    this.servicesConfiguration = new ConfigurationPanel(this.servicesSettingDiv, "Services configuration", "Config.ServiceConfiguration", "serviceConfigs")

    // The database context...
    this.datasourceSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display: none;" })
    this.datasourceSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Data stores configuration" }).appendElement({ "tag": "i", "class": "fa fa-database" })
    setSelectAction(this.datasourceSettingContext, this.datasourceSettingDiv)
    this.dataConfiguration = new ConfigurationPanel(this.datasourceSettingDiv, "Data configuration", "Config.DataStoreConfiguration", "dataStoreConfigs")

    // So here I will append panel to display more information about data inside the store. **/
    this.dataExplorer = new DataExplorer(this.datasourceSettingDiv)

    // The user and group setting / ldap.
    this.userGroupSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display:none;" })
    this.userGroupSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "LDAP configuration" }).appendElement({ "tag": "i", "class": "fa fa-users" })
    setSelectAction(this.userGroupSettingContext, this.userGroupSettingDiv)
    this.ldapConfiguration = new ConfigurationPanel(this.userGroupSettingDiv, "LDAP configuration", "Config.LdapConfiguration", "ldapConfigs")

    // The mail server config.
    this.mailServerSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display:none;" })
    this.mailServerSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "SMTP configuration" }).appendElement({ "tag": "i", "class": "fa fa-envelope-o" })
    setSelectAction(this.mailServerSettingContext, this.mailServerSettingDiv)
    this.smtpConfiguration = new ConfigurationPanel(this.mailServerSettingDiv, "Email server configuration", "Config.SmtpConfiguration", "smtpConfigs")

    // That area will contain different object properties.
    /* this.propertiesDiv = new Element(rightDiv, { "tag": "div", "class": "properties_div" })
     this.propertiesView = new PropertiesView(this.propertiesDiv)*/

    // I will set the configuration of the panel...
    server.entityManager.getObjectsByType("Config.Configurations", "Config", "",
        /** Progress callback */
        function () {

        },
        /** Progress callback */
        function (results, caller) {
            caller.serverConfiguration.setConfigurations(results)
            caller.servicesConfiguration.setConfigurations(results)
            caller.ldapConfiguration.setConfigurations(results)
            caller.smtpConfiguration.setConfigurations(results)
            caller.dataConfiguration.setConfigurations(results)
        },
        /** Error callback */
        function (errMsg, caller) {

        }, this)


}

/**
 * Create a new project...
 */
HomePage.prototype.createNewProject = function () {
    // So here I will create the new wizard...
    var wiz = new ProjectWizard()
}