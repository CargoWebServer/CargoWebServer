/**
 * The code editor.
 * TODO create the split functionnality
 * TODO create the multiuser access for a single file...
 */
var CodeEditor = function (parent) {

    // The panel...
    this.panel = parent.appendElement({ "tag": "div", "class": "codeEditor" }).down()

    // The open files...
    this.files = {}

    // The current file.
    this.activeFile = null

    // The map of file panels.
    this.filesPanel = {}

    // TODO create the new file event and the delete file event here...

    // Here I will attach the file navigator to file event.
    // Open 
    server.fileManager.attach(this, OpenEntityEvent, function (evt, codeEditor) {
        if (evt.dataMap["fileInfo"] != undefined) {
            var file = server.entityManager.entities[evt.dataMap["fileInfo"].UUID]
            if (file != undefined) {
                if (file.M_data != undefined && file.M_data != "") {
                    // Here thats mean the file was open
                    codeEditor.appendFile(file)
                }
            }
        }else if (evt.dataMap["bpmnDiagramInfo"] != undefined) {
            var diagram = server.entityManager.entities[evt.dataMap["bpmnDiagramInfo"].UUID]
            if (diagram != undefined) {
                codeEditor.appendBpmnDiagram(diagram)
            }
        }
    })

    // Attach the file close event.
    server.fileManager.attach(this, CloseEntityEvent, function (evt, codeEditor) {
        if (evt.dataMap["fileId"] != undefined) {
            codeEditor.removeFile(evt.dataMap["fileId"])
        }
    })

    return this
}

CodeEditor.prototype.appendBpmnDiagram = function (diagram) {
    // Here I will set the file
    if (this.files[diagram.M_id] != undefined) {
        // Set the tab active...
        this.setActiveFile(diagram.M_id)
        this.diagram.canvas.initWorkspace()
        return
    }

    var filePanel = this.panel.appendElement({ "tag": "div", "class": "filePanel", "id": diagram.M_id + "_editor" }).down()

    // Here I will create the new diagram...
    this.diagram = new SvgDiagram(filePanel, diagram)
    this.diagram.init()
    this.diagram.drawDiagramElements()

    this.files[diagram.M_id] = diagram

    this.filesPanel[diagram.M_id] = filePanel

    this.setActiveFile(diagram.M_id)

    // Now the resize element...
    this.diagram.canvas.initWorkspace = function (workspace) {
        return function () {
            if (workspace.lastChild == undefined) {
                return
            }

            if (workspace.lastChild.lastChild != undefined) {
                for (var childId in workspace.childs) {
                    var child = workspace.childs[childId]
                    if (child.element.viewBox != null) {
                        child.resize(workspace.element.offsetWidth, workspace.element.offsetHeight)
                    }
                }
            }
        }
    } (filePanel)

    window.addEventListener("resize", function (canvas) {
        return function () {
            canvas.initWorkspace()
        }
    } (this.diagram.canvas))


    this.diagram.canvas.initWorkspace()
}

CodeEditor.prototype.appendFile = function (file) {

    var fileMode = ""
    if (file.M_mime == "application/javascript") {
        fileMode = "ace/mode/javascript"
    } else if (file.M_mime == "text/css") {
        fileMode = "ace/mode/css"
    } else if (file.M_mime == "text/html") {
        fileMode = "ace/mode/html"
    } else if (file.M_mime == "text/json") {
        fileMode = "ace/mode/json"
    }

    if (fileMode.length == 0) {
        return
    }

    // Here I will set the file
    if (this.files[file.M_id] != undefined) {
        // Set the tab active...
        this.setActiveFile(file.M_id)
        return
    }

    // Here the new file tab must be created.
    this.files[file.M_id] = file

    // Now I will create the file editor.
    var filePanel = this.panel.appendElement({ "tag": "xmp", "class": "filePanel", "id": file.M_id + "_editor", "innerHtml": decode64(file.M_data) }).down()
    var editor = ace.edit(file.M_id + "_editor");
    // TODO connect the ace configuration panel with the Brigde configuration panel.
    // editor.setTheme("ace/theme/monokai");

    editor.getSession().setMode(fileMode);

    this.filesPanel[file.M_id] = filePanel

    this.setActiveFile(file.M_id)
}

CodeEditor.prototype.removeFile = function (fileId) {
    if (this.filesPanel[fileId] != undefined) {
        this.panel.removeElement(this.filesPanel[fileId])
        delete this.filesPanel[fileId]
        delete this.files[fileId]
        if (this.activeFile != undefined) {
            if (this.activeFile.M_id == fileId) {
                this.activeFile = null
            }
        }
    }
}

/**
 * Set the current file panel.
 */
CodeEditor.prototype.setActiveFile = function (fileId) {
    for (var id in this.filesPanel) {
        this.filesPanel[id].element.style.display = "none"
    }
    this.filesPanel[fileId].element.style.display = ""
    this.activeFile = this.files[fileId]
}