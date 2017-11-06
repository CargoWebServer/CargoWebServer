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

    /** Display the list of task instance since the server started */
    this.taskInstancesExplorer = null

    /** The file navigation  */
    this.fileNavigator = null

    /** Use to configure roles and permissions */
    this.rolePermissionManager = null
    this.rolePermissionDiv = null

    /** The task scheduler */
    this.scheduledTaskDiv = null

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
    this.oauth2Configuration = null
    this.scheduledTasksConfiguration = null

    /** The propertie div a the right */
    this.propertiesDiv = null
    this.propertiesView = null

    // The main menu.
    this.mainMenu = null

    // The toolbar div
    this.toolbarDiv = null

    homepage = this

    // That function is use to change the theme in the project explorer.
    server.fileManager.attach(this, ChangeThemeEvent, function (evt, homepage) {

        // Change the propertie in the class iteself.
        function changePropertyByClassName(propertie, className, themeClass, propertie_) {
            var rule = getCSSRule(className)
            var newValue = propertyFromStylesheet(themeClass, propertie)
            if (newValue != undefined) {
                if (propertie_ == undefined) {
                    propertie_ = propertie
                }
                if(rule == undefined){
                    console.log(className, "not defined!!!")
                }
                rule.style[propertie_] = propertyFromStylesheet(themeClass, propertie)
            }

            // Here I will keep the style information in the local storage.
            cargoThemeInfos = JSON.parse(localStorage.getItem("bridge_theme_infos"))
            if (cargoThemeInfos == undefined) {
                cargoThemeInfos = {}
            }

            if (cargoThemeInfos[className] == undefined) {
                cargoThemeInfos[className] = {}
            }

            // Here I will save the value.
            cargoThemeInfos[className][propertie_] = rule.style[propertie_]
            localStorage.setItem("bridge_theme_infos", JSON.stringify(cargoThemeInfos))
        }

        // I will set class values with theme class value
        changePropertyByClassName("background-color", ".navigation_div", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".navigation_div", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".header", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".header", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".vertical_submenu", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".vertical_submenu", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".menu_row", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".menu_row", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".splitter", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".splitter", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".menu_separator", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".menu_separator", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".home_page", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".home_page", "." + evt.dataMap.themeClass)
        
        changePropertyByClassName("color", ".autoCompleteDiv", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".autoCompleteDiv", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".autoCompleteDiv", "." + evt.dataMap.themeClass)

        changePropertyByClassName("color", ".session_panel", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".session_display_panel", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".session_display_panel", "." + evt.dataMap.themeClass)
        changePropertyByClassName("color", ".session_display_panel", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".session_state_menu", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".session_state_menu", "." + evt.dataMap.themeClass)
        changePropertyByClassName("color", ".session_state_menu", "." + evt.dataMap.themeClass)

        changePropertyByClassName("background-color", ".file_tab", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".file_tab", "." + evt.dataMap.themeClass)


        // I will use the gutter color for the background color of the workspace div.
        changePropertyByClassName("background-color", ".workspace_div", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".workspace_div", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("color", ".fileNavigationBtn", "." + evt.dataMap.themeClass + " .ace_gutter")

        if (evt.dataMap.isDark) {
            changePropertyByClassName("background-color", ".contextSelector", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("background", ".contextSelector", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".navigation_div", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".vertical_menu", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".vertical_submenu", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".vertical_submenu_items", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".popup_menu", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".menu_separator", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".file_tab", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".panel", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".entity_panel", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".entity", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".entity input", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".admin_table input", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".popup_div input", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".toolbar select", "." + evt.dataMap.themeClass)
            changePropertyByClassName("background-color", ".toolbar option", "." + evt.dataMap.themeClass+ " .ace_gutter")
            changePropertyByClassName("background", ".toolbar option", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".project_explorer input", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".permission_panel input", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".entity select", "." + evt.dataMap.themeClass)
            changePropertyByClassName("background-color", ".entity option", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("background", ".entity option", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".entity textarea", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".body_cell", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".data_explorer", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".security_manager_content", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".dialog_header", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".entities_header_btn.enabled", "." + evt.dataMap.themeClass + " .ace_gutter")
        } else {
            changePropertyByClassName("background-color", ".contextSelector", "." + evt.dataMap.themeClass)
            changePropertyByClassName("background", ".contextSelector", "." + evt.dataMap.themeClass)
            changePropertyByClassName("color", ".navigation_div", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".vertical_menu", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".vertical_submenu", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".vertical_submenu_items", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".popup_menu", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".menu_separator", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".file_tab", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".panel", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".entity_panel", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".entity", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".entity input", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".admin_table input", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".popup_div input", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".toolbar select", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".project_explorer input", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".permission_panel input", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".entity select", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".entity textarea", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".body_cell", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".data_explorer", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".security_manager_content", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("background-color", ".entity option", "." + evt.dataMap.themeClass+ " .ace_gutter")
            changePropertyByClassName("background", ".entity option", "." + evt.dataMap.themeClass + " .ace_gutter")
            changePropertyByClassName("color", ".dialog_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        }

        changePropertyByClassName("background-color", "#rightDiv", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background", "#rightDiv", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background-color", "#leftDiv", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background", "#leftDiv", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background-color", ".contextSelector", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background", ".contextSelector", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background-color", ".splitter", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background", ".splitter", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("color", ".severConfiguration", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".entities_header_btn", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        changePropertyByClassName("background", ".entities_header_btn", "." + evt.dataMap.themeClass + " .ace_gutter", "color")
        
        changePropertyByClassName("background-color", ".entities_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".entities_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("color", ".entities_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".result_query_panel", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".result_query_panel", "." + evt.dataMap.themeClass)

        changePropertyByClassName("color", ".role_table_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".role_table_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".role_table_header", "." + evt.dataMap.themeClass + " .ace_gutter")

        changePropertyByClassName("color", ".table_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".header_cell", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".header_cell", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("color", ".header_cell", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".header_cell", "." + evt.dataMap.themeClass, "border-color")
        changePropertyByClassName("background-color", ".header_cell", "." + evt.dataMap.themeClass, "border-color")
        changePropertyByClassName("background", ".body_cell", "." + evt.dataMap.themeClass + " .ace_gutter", "border-color")
        changePropertyByClassName("background-color", ".body_cell", "." + evt.dataMap.themeClass + " .ace_gutter", "border-color")
        changePropertyByClassName("background", "::-webkit-scrollbar-thumb", "." + evt.dataMap.themeClass + " .ace_gutter", "border-color")
        changePropertyByClassName("background-color", "::-webkit-scrollbar-thumb", "." + evt.dataMap.themeClass + " .ace_gutter", "border-color")
        
        changePropertyByClassName("color", ".permissions_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".permissions_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".permissions_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("color", ".permissions_panel_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".permissions_panel_header", "." + evt.dataMap.themeClass, "border-color")
        changePropertyByClassName("background-color", ".permissions_panel_header", "." + evt.dataMap.themeClass, "border-color")
        changePropertyByClassName("background-color", ".permission_panel", "." + evt.dataMap.themeClass, "border-color")
        
        changePropertyByClassName("background-color", ".role_permission_manager", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".role_permission_manager", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".security_manager_content", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".security_manager_content", "." + evt.dataMap.themeClass)

        changePropertyByClassName("background-color", ".dialog_content", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".dialog_content", "." + evt.dataMap.themeClass)
        changePropertyByClassName("color", ".dialog_content", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".dialog_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".dialog_header", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background-color", ".dialog_footer", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".dialog_footer", "." + evt.dataMap.themeClass )
        changePropertyByClassName("color", ".dialog_footer", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".admin_table", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".admin_table", "." + evt.dataMap.themeClass)
        changePropertyByClassName("color", ".admin_table", "." + evt.dataMap.themeClass)
        
        changePropertyByClassName("background-color", ".popup_div", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".popup_div", "." + evt.dataMap.themeClass)
        changePropertyByClassName("color", ".popup_div", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background-color", ".popup_div", "." + evt.dataMap.themeClass + " .ace_gutter", "border-color")
        changePropertyByClassName("background-color", ".dialog", "." + evt.dataMap.themeClass + " .ace_gutter", "border-color")
        changePropertyByClassName("background-color", ".login-form", "." + evt.dataMap.themeClass)
        changePropertyByClassName("background", ".login-form", "." + evt.dataMap.themeClass)
        changePropertyByClassName("color", ".login-form", "." + evt.dataMap.themeClass)
        
        changePropertyByClassName("background-color", ".main_page", "." + evt.dataMap.themeClass + " .ace_gutter")
        changePropertyByClassName("background", ".main_page", "." + evt.dataMap.themeClass + " .ace_gutter")
    })


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

    // Create a new Entity Query File.

    // Entity Query Language File.
    var newEqlQueryMenuItem = new MenuItem("new_eql_query_menu_item", "EQL Query", {}, 1, function (extension) {
        return function () {
            createQuery(extension, "/** Eql query **/\n")
        }
    }(".eql"), "fa fa-file-o")

    // Structured Query Language Query Language File.
    var newSqlQueryMenuItem = new MenuItem("new_sql_query_menu_item", "SQL Query", {}, 1, function (extension) {
        return function () { createQuery(extension, "/** Sql query **/\n") }
    }(".sql"), "fa fa-file-o")

    var newProjectMenuItem = new MenuItem("new_project_menu_item", "New Project...", {}, 1,
        function (homepage) {
            return function () {
                homepage.createNewProject()
            }
        }(this), "fa fa-files-o")

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
                                server.schemaManager.importXsdSchema(file.name, e.target.result)
                            }
                        }(f);
                        reader.readAsText(f);
                    }
                }
            }(this)
            // Display the file explorer...
            fileExplorer.element.click()
        }
    }(parent), "fa fa-file-o")

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
            }(this)
            // Display the file explorer...
            fileExplorer.element.click()
        }
    }(parent), "fa fa-file-o")

    var importDataMenuItem = new MenuItem("import_data_menu_item", "Import", { "import_xsd_menu_item": importXsdSchemaMenuItem, "import_xml_menu_item": importXmlDataMenuItem }, 1)

    // The new menu in the data Menu
    var newDataMenuItem = new MenuItem("new_data_menu_item", "New", { "new_eql_query_menu_item": newEqlQueryMenuItem, "new_sql_query_menu_item": newSqlQueryMenuItem }, 1)

    // The preference edition menu.
    var preferencesServerItem = new MenuItem("preferences_server_menu_item", "Preferences", {}, 1,
        function () {
            // 
        }, "fa fa-wrench")

    var closeServerItem = new MenuItem("close_server_menu_item", "Close server", {}, 1, function () { server.stop() }, "fa fa-power-off")

    var fileMenuItem = new MenuItem("file_menu", "File", { "new_file_menu_item": newFileMenuItem, "preferences_server_menu_item": preferencesServerItem, "close_server_menu_item": closeServerItem }, 0)

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
                    if (leftDiv.element.clientWidth == 50) {
                        var w = 431
                        /*var navigationDivs = document.getElementsByClassName("navigation_div")
                        for(var i=0; i < navigationDivs.length; i++){
                            if(navigationDivs[i].style.display != "none"){
                                w = navigationDivs[i].clientWidth
                            }
                        }*/
                        var keyframe = "100% { width:" + w + "px;}"
                        leftDiv.animate(keyframe, .5,
                            function (leftDiv) {
                                return function () {
                                    leftDiv.element.style.width = "431px"
                                }
                            }(leftDiv))
                    } else {
                        var keyframe = "100% { width:50px;}"
                        leftDiv.element.style.overflowY = "hidden"
                        leftDiv.animate(keyframe, .5,
                            function (leftDiv) {
                                return function () {
                                    leftDiv.element.style.width = "50px"
                                    leftDiv.element.style.overflowY = "hidden"
                                }
                            }(leftDiv))
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

                // Set the size of absolute panel.
                fireResize()
            }

        }(div, leftDiv)
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

    // The roles and permissions configuration.
    this.rolePermissionDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display:none;" })
    this.rolePermissionManager = new RolePermissionManager(this.rolePermissionDiv)
    this.rolePermissionContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Roles/Permissions" }).appendElement({ "tag": "i", "class": "fa fa-shield" })
    setSelectAction(this.rolePermissionContext, this.rolePermissionDiv)

    // The Oauth context...
    this.securityDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": " left:50px; display:none;" })
    this.securityContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "OAuth2" }).appendElement({ "tag": "i", "class": "fa fa-lock" })
    setSelectAction(this.securityContext, this.securityDiv)
    this.oauth2Configuration = new ConfigurationPanel(this.securityDiv, "Security configuration", "Config.OAuth2Configuration", "oauth2Configuration")

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

    this.scheduledTasksDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display:none;" })
    this.taskSchedulerContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Task scheduler" }).appendElement({ "tag": "i", "class": "fa fa-clock-o" })
    this.scheduledTasksConfiguration = new ConfigurationPanel(this.scheduledTasksDiv, "Scheduled Tasks", "Config.ScheduledTask", "scheduledTasks")
    setSelectAction(this.taskSchedulerContext, this.scheduledTasksDiv)

    this.taskInstancesExplorer = new TaskInstancesExplorer(this.scheduledTasksDiv)

    // I will set the configuration of the panel...
    server.entityManager.getEntities("Config.Configurations", "Config", "", 0, -1, [], true,
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
            caller.oauth2Configuration.setConfigurations(results)
            caller.scheduledTasksConfiguration.setConfigurations(results)
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