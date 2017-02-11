var languageInfo = {
    "en": {
    },
    "fr": {
    }
}

server.languageManager.appendLanguageInfo(languageInfo)
var homepage = null
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

    /** The file navigation  */
    this.fileNavigator = null

    /** The code editor */
    this.codeEditor = null

    /** The bpmn diagram explorer. */
    this.bpmnExplorer = null

    /** The configurations... */
    this.serverConfiguration = null
    this.dataConfiguration = null
    this.ldapConfiguration = null
    this.smtpConfiguration = null

    /** The propertie div a the right */
    this.propertiesDiv = null
    this.propertiesView = null

    // Set the global variable...
    homepage = this

    // The main menu.
    this.mainMenu = null

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

    var menuRow = this.headerDiv.appendElement({ "tag": "div", "style": "width:100%; height: 60px; display: table-row" }).down()

    // This is where the menu grid will be put...
    this.menuContentDiv = menuRow.appendElement({ "tag": "div", "style": "width:100%; display: table-cell;" }).down()
        .appendElement({ "tag": "div", "style": "width:100%; display: table; height: 0px;" }).down()

    // Now I will create the session panel...
    this.sessionPanel = new SessionPanel(menuRow.appendElement({ "tag": "div", "style": "width:100%; display: table-cell; height:60px" }).down(), sessionInfo)

    // The menu
    var newFileOrProjectMenuItem = new MenuItem("newFileOrProjectMenuItem", "New File or Project...", {}, 1,
        function (homepage) {
            return function () {
                homepage.createNewProject()
            }
        } (this), "fa fa-files-o")
    var openFileOrProjectMenuItem = new MenuItem("openFileOrProjectMenuItem", "Close server", {}, 1, function () { server.stop() }, "fa fa-folder-open-o")
    var fileMenuItem = new MenuItem("file_menu", "File", { "newFileOrProjectMenuItem": newFileOrProjectMenuItem, "openFileOrProjectMenuItem": openFileOrProjectMenuItem }, 0)

    var editMenuItem = new MenuItem("edit_menu", "Edit", {}, 0)

    // The main menu will be display in the body element, so nothing will be over it.
    this.mainMenu = new VerticalMenu(new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; top:2px;" }), [fileMenuItem, editMenuItem])

    /////////////////////////////////// workspace section  ///////////////////////////////////
    this.mainArea = this.panel.appendElement({ "tag": "div", "style": "display: table; width:100%; height:100%" }).down()

    // The working file grid...
    this.workingFilesDiv = this.headerDiv.appendElement({ "tag": "div", "id": "workingFilesDiv", "style": "width:100%; height: 30px; display: table;" }).down()
        .appendElement({ "tag": "div", "style": "width:100%; display: inline; position: relative" }).down()
    //.appendElement({ "tag": "div", "style": "position: absolute; top:0px; left:0px; bottom:0px; right: 0px;" }).down()

    this.fileNavigator = new FileNavigator(this.workingFilesDiv)

    // Now the left, middle and right div...
    var splitArea1 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-cell; position: relative;" }).down()
    var leftDiv = new Element(splitArea1, { "tag": "div", "id": "leftDiv", "style": "" })
    var splitter1 = this.mainArea.appendElement({ "tag": "div", "class": "splitter vertical", "id": "splitter1" }).down()

    var splitArea2 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-cell; width: 100%; height:100%" }).down()
    var middleDiv = new Element(splitArea2, { "tag": "div", "id": "middleDiv" })
    var splitter2 = this.mainArea.appendElement({ "tag": "div", "class": "splitter vertical", "id": "splitter2" }).down()

    var splitArea3 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-cell; height:100%" }).down()
    var rightDiv = new Element(splitArea3, { "tag": "div", "id": "rightDiv" })

    // Init the splitter action.
    initSplitter(splitter1, leftDiv)
    initSplitter(splitter2, rightDiv)

    // The workspace area
    this.workspaceDiv = new Element(middleDiv, { "tag": "div", "class": "workspace_div" })

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
                        leftDiv.element.style.overflowY = "auto"
                        var keyframe = "100% { width:431px;}"
                        leftDiv.animate(keyframe, .5,
                            function (leftDiv) {
                                return function () {
                                    leftDiv.element.style.width = "431px"
                                    leftDiv.element.style.overflowY = "auto"
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
                }

                this.firstChild.className += " active"
            }
        } (div, leftDiv)
    }

    // Now I will append the button inside the context selector.
    // The project.
    this.projectDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px;" })
    this.projectExplorer = new ProjectExplorer(this.projectDiv)
    this.projectContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Projects" }).appendElement({ "tag": "i", "class": "fa fa-files-o active" })
    setSelectAction(this.projectContext, this.projectDiv)

    // The search context...
    this.searchDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": " left:50px; display:none;" })
    this.searchContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Search" }).appendElement({ "tag": "i", "class": "fa fa-search" })
    setSelectAction(this.searchContext, this.searchDiv)

    // The server context...
    this.serverSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display: none;" })
    this.serverSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Server Configuration" }).appendElement({ "tag": "i", "class": "fa fa-server" })
    setSelectAction(this.serverSettingContext, this.serverSettingDiv)
    this.serverConfiguration = new ConfigurationPanel(this.serverSettingDiv, "Server configuration", "Config.ServerConfiguration", "serverConfig")

    // The database context...
    this.datasourceSettingDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display: none;" })
    this.datasourceSettingContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Data stores configuration" }).appendElement({ "tag": "i", "class": "fa fa-database" })
    setSelectAction(this.datasourceSettingContext, this.datasourceSettingDiv)
    this.dataConfiguration = new ConfigurationPanel(this.datasourceSettingDiv, "Data configuration", "Config.DataStoreConfiguration", "dataStoreConfigs")

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

    // The bpmn explorer...
    this.bpmnDiv = new Element(leftDiv, { "tag": "div", "class": "navigation_div", "style": "left:50px; display:none;" })
    this.bpmnContext = new Element(this.contextSelector, { "tag": "div", "class": "navigation_btn", "title": "Workflow Manager" }).appendElement({ "tag": "i", "class": "fa fa-tasks" })
    setSelectAction(this.bpmnContext, this.bpmnDiv)
    this.bpmnExplorer = new BpmnExplorer(this.bpmnDiv)

    // That area will contain different object properties.
    this.propertiesDiv = new Element(rightDiv, { "tag": "div", "class": "properties_div" })
    this.propertiesView = new PropertiesView(this.propertiesDiv)

    // I will set the configuration of the panel...
    server.entityManager.getObjectsByType("Config.Configurations", "Config", "",
        /** Progress callback */
        function () {

        },
        /** Progress callback */
        function (results, caller) {
            caller.serverConfiguration.setConfigurations(results)
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