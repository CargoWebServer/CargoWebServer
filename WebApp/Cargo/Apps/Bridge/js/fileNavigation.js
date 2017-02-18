/**
 * The file navigator is use to display open project files.
 */
var FileNavigator = function (parent) {
    this.parent = parent
    this.panel = new Element(parent, { "tag": "div", "class": "fileNavigator" })

    // Keep the list of open files.
    this.files = {}

    // The tabs that contain the file names...
    this.tabs = {}

    this.activeFile = null

    // Show the list of file that dosent fit in the file explorer...
    this.showHiddenFilesBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-down fileNavigationBtn" })

    // The show the previous file.
    this.setPreviousFileBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-left fileNavigationBtn" })

    // Show the next file.
    this.setNextFileBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-right fileNavigationBtn" })


    // Here I will attach the file navigator to file event.
    // Open 
    server.fileManager.attach(this, OpenEntityEvent, function (evt, fileNavigator) {
        var file
        if (evt.dataMap["fileInfo"] != undefined) {
            file = server.entityManager.entities[evt.dataMap["fileInfo"].UUID]
        } else if (evt.dataMap["bpmnDiagramInfo"] != undefined) {
            file = server.entityManager.entities[evt.dataMap["bpmnDiagramInfo"].UUID]
        }

        if (file != undefined) {
            if (file.M_id != undefined) {
                // Here thats mean the file was open
                fileNavigator.appendFile(file)
            }
        }
    })

    // Attach the file close event.
    server.fileManager.attach(this, CloseEntityEvent, function (evt, fileNavigator) {
        if (evt.dataMap["fileId"] != undefined) {
            // Remove the file.
            fileNavigator.removeFile(evt.dataMap["fileId"])
        }
    })


    return this;
}

/**
 * That function will append a new file in the file navigator.
 */
FileNavigator.prototype.appendFile = function (file) {

    if (this.activeFile != undefined) {
        if (this.activeFile.M_id == file.M_id) {
            // Nothing todo here...
            return
        }
    }

    // Keep the ref of the file.
    if (this.files[file.M_id] != undefined) {
        // Set the tab active...
        this.setActiveTab(file.M_id)
        return
    }

    // Here the new file tab must be created.
    this.files[file.M_id] = file

    this.tabs[file.M_id] = this.panel.appendElement({ "tag": "div", "class": "file_tab" }).down()
    this.tabs[file.M_id].appendElement({ "tag": "div", "style": "display: table; height: 100%; width:100%;" }).down()
        .appendElement({ "tag": "div", "style": "display: table-cell;", "innerHtml": file.M_name })
        .appendElement({ "tag": "i", "id": file.id + "_file_tab_close_btn", "class": "fa fa-times file_tab_close_btn" })

    var fileCloseBtn = this.tabs[file.M_id].getChildById(file.id + "_file_tab_close_btn")
    if (fileCloseBtn != undefined) {
        fileCloseBtn.element.onclick = function (file) {
            return function () {
                // Send event localy...
                var evt = { "code": CloseEntityEvent, "name": FileEvent, "dataMap": { "fileId": file.M_id } }
                server.eventHandler.BroadcastEvent(evt)
            }
        } (file)
    }

    this.setActiveTab(file.M_id)

    // now the onclick event...
    this.tabs[file.M_id].element.onclick = function (file, fileNavigator) {
        return function () {
            if (fileNavigator.files[file.M_id] != undefined) {
                fileNavigator.setActiveTab(file.M_id)
            }
        }
    } (file, this)
}

FileNavigator.prototype.setActiveTab = function (fileId) {
    for (var tabId in this.tabs) {
        this.tabs[tabId].element.className = this.tabs[tabId].element.className.replace(" active", "")
    }

    // Keep the ref of the file.
    if (this.tabs[fileId] != undefined) {
        // Set the tab active...
        this.tabs[fileId].element.className += " active"
        this.activeFile = this.files[fileId]
    }

    // I will generate the event so other panel will set the current file... 
    if (this.files[fileId] != undefined) {
        // local event.
        var evt
        if (this.activeFile.TYPENAME == "BPMNDI.BPMNDiagram") {
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "bpmnDiagramInfo": this.activeFile } }
        } else {
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": this.activeFile } }
        }
        server.eventHandler.BroadcastEvent(evt)
    }
}

/**
 * That function will remove existing file from the file navigator.
 */
FileNavigator.prototype.removeFile = function (fileId) {
    var index
    for (index = 0; index < Object.keys(this.files).length; index++) {
        if (Object.keys(this.files)[index] == fileId) {
            break
        }
    }
    var file = this.files[fileId]
    delete this.files[fileId]
    var tab = this.tabs[fileId]
    delete this.tabs[fileId]
    this.panel.removeElement(tab)
    if (this.activeFile != null) {
        if (this.activeFile.M_id == fileId) {
            //this.activeFile = null
            // Now I will set the active file
            if (Object.keys(this.files).length == 0) {
                this.activeFile = null
            } else if (index >= Object.keys(this.files).length) {
                this.setActiveTab(Object.keys(this.files)[Object.keys(this.files).length - 1])
            } else {
                this.setActiveTab(Object.keys(this.files)[index])
            }
        }
    }


    // TODO ask the user to save the file if there is change...
}