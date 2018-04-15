/**
 * The code editor
 * TODO create the split functionnality
 * TODO create the multiuser access for a single file
 */

var CodeEditor = function (parent) {

    // The panel...
    this.panel = parent.appendElement({ "tag": "div", "class": "codeEditor" }).down()

    // The open files...
    this.files = {}

    // The toolbars associated whit each editor.
    this.toolbars = {}

    // The current file.
    this.activeFile = null

    // The editor
    this.editors = {}

    // The map of file panels.
    this.filesPanel = {}

    // TODO create the new file event and the delete file event here...
    this.quiet = false

    // Here I will create the file toolbar...
    //this.fileToolbar = new Element(null, { "tag": "div", "class": "toolbar" })
    this.theme = localStorage.getItem("bridge_editor_theme")
    if (this.theme == undefined) {
        this.theme = "ace/theme/chrome"
    }

    this.themeClass = localStorage.getItem("bridge_editor_theme_class")
    if (this.themeClass == undefined) {
        this.themeClass = "ace-chrome"
    }

    // Here I will attach the file navigator to file event.
    // Open.
    server.fileManager.attach(this, OpenEntityEvent, function (evt, codeEditor) {
        if (evt.dataMap["fileInfo"] !== undefined) {
            var file = entities[evt.dataMap["fileInfo"].UUID]
            if (file == undefined) {
                file = evt.dataMap["fileInfo"]
            }

            if (file.M_data !== undefined) {
                // Here thats mean the file was open
                codeEditor.appendFile(file, evt.dataMap["coord"])
            }

        } else if (evt.dataMap["bpmnDiagramInfo"] !== undefined) {
            var diagram = entities[evt.dataMap["bpmnDiagramInfo"].UUID]
            if (diagram !== undefined) {
                codeEditor.appendBpmnDiagram(diagram)
            }
        } else if (evt.dataMap["prototypeInfo"] != undefined) {
            var prototype = evt.dataMap["prototypeInfo"]
            if (prototype !== undefined) {
                codeEditor.appendPrototypeEditor(prototype)
            }
        } else if (evt.dataMap["searchInfo"] != undefined) {
            codeEditor.appendSearchPage(evt.dataMap["searchInfo"])
        }
    })

    // Attach the file close event.
    server.fileManager.attach(this, CloseEntityEvent, function (evt, codeEditor) {
        var fileId = evt.dataMap["fileId"]
        if (fileId !== undefined) {
            codeEditor.removeFile(fileId)
            if (codeEditor.toolbars[fileId] !== undefined) {
                for (var i = 0; i < codeEditor.toolbars[fileId].length; i++) {
                    var toolbar = codeEditor.toolbars[fileId][i];
                    homePage.toolbarDiv.removeElement(toolbar);
                }
            }
            codeEditor.toolbars[fileId] = []
        }
    })

    // Attach the file update event.
    server.fileManager.attach(this, UpdateFileEvent, function (evt, codeEditor) {
        if (evt.dataMap.fileInfo !== undefined) {
            var file = evt.dataMap["fileInfo"]
            var editor = codeEditor.editors[file.UUID + "_editor"]
            if (editor !== undefined) {
                // Supend the change event propagation
                codeEditor.quiet = true
                editor.setValue(decode64(file.M_data), -1)
                editor.clearSelection()
                if (editor.getCursorPosition != undefined) {
                    var position = editor.getCursorPosition()
                    editor.scrollToLine(position.row + 1, true, true, function () { });
                    editor.gotoLine(position.row + 1, position.column)
                }
                // Resume the chage event propagation.
                codeEditor.quiet = false
            }
        }
    })

    server.fileManager.attach(this, ChangeThemeEvent, function (evt, codeEditor) {
        codeEditor.theme = evt.dataMap.theme
        for (var editorUuid in codeEditor.editors) {
            if (codeEditor.editors[editorUuid].setTheme !== undefined) {
                codeEditor.editors[editorUuid].setTheme(evt.dataMap.theme);
            } else if (codeEditor.editors[editorUuid].editor.setTheme !== undefined) {
                codeEditor.editors[editorUuid].editor.setTheme(evt.dataMap.theme);
            }
        }
    })

    server.entityManager.attach(this, UpdateEntityEvent, function (evt, codeEditor) {
        if (evt.dataMap.entity !== undefined) {
            var file = evt.dataMap["entity"]
            var editor = codeEditor.editors[file.UUID + "_editor"]
            if (editor !== undefined && file.TYPENAME == "CargoEntities.File") {
                // Supend the change event propagation
                codeEditor.quiet = true
                if (editor.getCursorPosition != undefined) {
                    var position = editor.getCursorPosition()
                    editor.setValue(decode64(file.M_data), -1)
                    editor.clearSelection()
                    editor.scrollToLine(position.row + 1, true, true, function () { });
                    editor.gotoLine(position.row + 1, position.column)
                }
                // Resume the chage event propagation.
                codeEditor.quiet = false
            }
        }
    })

    /** Always display the license. */
    /*server.fileManager.getFileByPath("/LICENSE",
        // Success callback 
        function (result, caller) {
            result.M_mime = "text/plain"
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": result } }
            server.eventHandler.broadcastLocalEvent(evt)
        },
        // Error callback
        function (errObj, caller) {

        }, {})*/

    return this
}

/**
 * Create a new Search page.
 */
CodeEditor.prototype.appendSearchPage = function (searchInfo) {
    if (this.files[searchInfo.UUID] !== undefined) {
        // Set the tab active...
        this.setActiveFile(searchInfo.UUID)
        return
    }

    var filePanel = this.panel.appendElement({ "tag": "div", "class": "filePanel", "id": searchInfo.UUID + "_search_div" }).down()
    // create and display the search page.
    new SearchPage(filePanel, searchInfo)
    this.files[searchInfo.UUID] = searchInfo
    this.filesPanel[searchInfo.UUID] = filePanel
    this.setActiveFile(searchInfo.UUID)
}

/**
 * Here I will display the prototype editor.
 */
CodeEditor.prototype.appendPrototypeEditor = function (prototype) {
    // Here I will set the prototype editor.
    if (this.files[prototype.TypeName] !== undefined) {
        // Set the tab active...
        this.setActiveFile(prototype.TypeName)
        return
    }

    server.configurationManager.getActiveConfigurations(
        function (results, caller) {
            var namespaces = []
            for (var i = 0; i < results.M_dataStoreConfigs.length; i++) {
                // Sql entities are not part of the heritage system.
                if (results.M_dataStoreConfigs[i].M_dataStoreType == 2) {
                    namespaces.push(results.M_dataStoreConfigs[i].M_id)
                }
            }
            var codeEditor = caller.codeEditor
            var prototype = caller.prototype
            var entityEditor = new EntityPrototypeEditor(filePanel, namespaces, undefined, function (entityEditor) {
                entityEditor.typeNameInput.element.value = prototype.TypeName
                entityEditor.setCurrentPrototype(prototype)
                entityEditor.space.element.style.display = ""
            })
        },
        function (errObj, caller) {

        },
        { "codeEditor": this, "prototype": prototype })

    var filePanel = this.panel.appendElement({ "tag": "div", "class": "filePanel", "id": prototype.TypeName + "_editor" }).down()
    this.files[prototype.TypeName] = prototype
    this.filesPanel[prototype.TypeName] = filePanel
    this.setActiveFile(prototype.TypeName)
}

CodeEditor.prototype.appendBpmnDiagram = function (diagram) {
    // Here I will set the file
    if (this.files[diagram.UUID] !== undefined) {
        // Set the tab active...
        this.setActiveFile(diagram.UUID)
        this.diagram.canvas.initWorkspace()
        return
    }

    var filePanel = this.panel.appendElement({ "tag": "div", "class": "filePanel", "id": diagram.UUID + "_editor" }).down()


    this.diagram = new SvgDiagram(filePanel, diagram)

    this.diagram.init(function (codeEditor, diagram, filePanel) {
        return function () {
            codeEditor.diagram.drawDiagramElements()

            codeEditor.files[diagram.UUID] = diagram
            codeEditor.filesPanel[diagram.UUID] = filePanel
            codeEditor.setActiveFile(diagram.UUID)

            // Now the resize element...
            codeEditor.diagram.canvas.initWorkspace = function (workspace) {
                return function () {
                    if (workspace.lastChild === undefined) {
                        return
                    }
                    if (workspace.lastChild.lastChild !== undefined) {
                        for (var childId in workspace.childs) {
                            var child = workspace.childs[childId];
                            if (child.element.viewBox !== null) {
                                if (child.resize != undefined) {
                                    child.resize(workspace.element.offsetWidth, workspace.element.offsetHeight);
                                }
                            }
                        }
                    }
                }
            }(filePanel)

            window.addEventListener("resize", function (canvas) {
                return function () {
                    canvas.initWorkspace()
                }
            }(codeEditor.diagram.canvas))

            codeEditor.diagram.canvas.initWorkspace()
        }
    }(this, diagram, filePanel))
}

CodeEditor.prototype.appendFile = function (file, coord) {

    var fileMode = ""
    if (file.M_mime == "application/javascript") {
        fileMode = "ace/mode/javascript"
    } else if (file.M_mime == "text/css") {
        fileMode = "ace/mode/css"
    } else if (file.M_mime == "text/html") {
        fileMode = "ace/mode/html"
    } else if (file.M_mime == "text/json") {
        fileMode = "ace/mode/json"
    } else if (file.M_mime == "text/plain") {
        fileMode = "ace/mode/text"
    }


    // Here I will set the file
    if (this.files[file.UUID] != undefined) {
        // Set the tab active...
        this.setActiveFile(file.UUID, coord)
        return
    }

    // Here the new file tab must be created.
    this.files[file.UUID] = file

    //var deleteBtn = fileToolbar.appendElement({"tag":"div"}).down()

    if (fileMode.length == 0) {
        if (file.M_name.endsWith(".eql") || file.M_name.endsWith(".sql")) {
            // Here I will create a query editor insted of ace editor.
            var filePanel = this.panel.appendElement({ "tag": "div", "class": "filePanel", "id": file.UUID + "_editor" }).down()

            // The query editor.
            var queryEditor = new QueryEditor(filePanel, file, function (codeEditor, fileId) {
                return function (queryEditor) {
                    // I will append the list of dataStore that can be use to do query.
                    codeEditor.toolbars[fileId] = []
                    codeEditor.toolbars[fileId].push(queryEditor.queryToolBar)
                }
            }(this, file.UUID))

            // Init the query editor.
            queryEditor.init()

            this.editors[file.UUID + "_editor"] = queryEditor.editor

            queryEditor.editor.getSession().on('change', function (fileUUID, codeEditor) {
                return function () {
                    if (!codeEditor.quiet && entities[fileUUID] !== undefined) {
                        var editor = codeEditor.editors[fileUUID + "_editor"]
                        var evt = { "code": ChangeFileEvent, "name": FileEvent, "dataMap": { "fileId": fileUUID } }
                        var file = entities[fileUUID]
                        file.M_data = encode64(editor.getSession().getValue())
                        server.eventHandler.broadcastLocalEvent(evt)
                    }
                }
            }(file.UUID, this));

            this.filesPanel[file.UUID] = filePanel
            this.setActiveFile(file.UUID)
        }
        return
    }


    // Now I will create the file editor.
    var filePanel = this.panel.appendElement({ "tag": "xmp", "class": "filePanel", "id": file.UUID + "_editor", "innerHtml": decode64(file.M_data) }).down()

    var observer = new MutationObserver(function (codeEditor) {
        return function (multiRecord) {
            var record = multiRecord.pop()
            var themeClass = record.target.classList[record.target.classList.length - 1]
            var isDark = record.target.className.indexOf("ace_dark") != -1
            if (themeClass != codeEditor.themeClass && themeClass != "ace-tm") {
                // Keep it in the local storage.
                localStorage.setItem("bridge_editor_theme_class", themeClass)
                localStorage.setItem("bridge_editor_theme", codeEditor.theme)
                codeEditor.themeClass = themeClass
                evt = { "code": ChangeThemeEvent, "name": FileEvent, "dataMap": { "theme": codeEditor.theme, "themeClass": codeEditor.themeClass, "isDark": isDark } }
                server.eventHandler.broadcastLocalEvent(evt)
            }
        }
    }(this))

    observer.observe(filePanel.element, {
        attributes: true,
        attributeFilter: ['class'],
        childList: false,
        characterData: false
    })

    ace.require("ace/ext/language_tools");
    var editor = ace.edit(file.UUID + "_editor");
    ace.require('ace/ext/settings_menu').init(editor);
    editor.setTheme(this.theme);
    editor.getSession().setMode(fileMode);

    editor.setOptions({
        enableBasicAutocompletion: true,
        enableSnippets: true,
        enableLiveAutocompletion: true
    });

    this.editors[file.UUID + "_editor"] = editor

    // Editor command here.
    editor.commands.addCommands([{
        name: "showSettingsMenu",
        bindKey: { win: "Ctrl-q", mac: "Ctrl-q" },
        exec: function (codeEditor) {
            return function (editor) {
                editor.showSettingsMenu();
                var themeSelect = document.getElementById('setTheme');
                themeSelect.addEventListener("change", function () {
                    // Here I will throw a change theme event.
                    codeEditor.theme = this.value
                });
            }
        }(this),
        readOnly: true
    }]);

    // In case of file update...
    editor.getSession().on('change', function (fileUUID, codeEditor) {
        return function () {
            if (!codeEditor.quiet && entities[fileUUID] !== undefined) {
                var editor = codeEditor.editors[fileUUID + "_editor"]
                var evt = { "code": ChangeFileEvent, "name": FileEvent, "dataMap": { "fileId": fileUUID } }
                var file = entities[fileUUID]
                file.M_data = encode64(editor.getSession().getValue())
                server.eventHandler.broadcastLocalEvent(evt)
            }
        }
    }(file.UUID, this));

    editor.session.on("changeScrollTop", function (scrollTop) {
        var header = document.getElementById("workingFilesDiv")
        if (scrollTop > 0) {
            if (header.className.indexOf(" scrolling") == -1) {
                header.className += " scrolling"
                header.parentNode.className += " scrolling"
            }
        } else {
            header.className = header.className.replaceAll(" scrolling", "")
            header.parentNode.className = header.parentNode.className.replaceAll(" scrolling", "")
        }
    })

    this.filesPanel[file.UUID] = filePanel
    this.setActiveFile(file.UUID, coord)
}

CodeEditor.prototype.removeFile = function (uuid) {
    if (this.filesPanel[uuid] != undefined) {
        // remove the element from the panel.
        this.panel.removeElement(this.filesPanel[uuid])
        delete this.filesPanel[uuid]
        delete this.files[uuid]
        delete this.editors[uuid + "_editor"]

        // If there's no more file i will reset the shadow.
        if (Object.keys(this.files).length == 0) {
            var header = document.getElementById("workingFilesDiv")
            header.className = header.className.replaceAll(" scrolling", "")
            header.parentNode.className = header.parentNode.className.replaceAll(" scrolling", "")
        }

        if (this.activeFile != undefined) {
            if (this.activeFile.UUID == uuid) {
                this.activeFile = null
            }
        }
    }
}

/**
 * Set the current file panel.
 */
CodeEditor.prototype.setActiveFile = function (uuid, coord) {
    for (var id in this.filesPanel) {
        this.filesPanel[id].element.style.display = "none"
    }
    if (this.filesPanel[uuid] !== undefined) {
        this.filesPanel[uuid].element.style.display = ""
        var header = document.getElementById("workingFilesDiv")
        var aceContent = this.filesPanel[uuid].element.getElementsByClassName("ace_content")[0]
        if (aceContent != null) {
            if (aceContent.style.marginTop != "0px" && aceContent.style.marginTop != "") {
                if (header.className.indexOf(" scrolling") == -1) {
                    header.className += " scrolling"
                    header.parentNode.className += " scrolling"
                }
            } else {
                header.className = header.className.replaceAll(" scrolling", "")
                header.parentNode.className = header.parentNode.className.replaceAll(" scrolling", "")
            }
        } else {
            header.className = header.className.replaceAll(" scrolling", "")
            header.parentNode.className = header.parentNode.className.replaceAll(" scrolling", "")
        }
    }
    this.activeFile = this.files[uuid]

    // Now the toolbar...
    var toolbars = document.getElementsByClassName("toolbar")
    for (var i = 0; i < toolbars.length; i++) {
        toolbars[i].style.display = "none" // hide toolbar.
    }

    if (document.getElementById(uuid + "_toolbar") != undefined) {
        document.getElementById(uuid + "_toolbar").style.display = ""
    }

    // in case coord are given then I will move the editor to there.
    if (coord != undefined) {
        var editor = this.editors[uuid + "_editor"]
        editor.focus();
        editor.gotoLine(coord.ln, coord.col, true);
        editor.renderer.scrollToRow(coord.ln-3); // minus 3 to see couple line before...
    }
}