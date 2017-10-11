/**
 * The file navigator is use to display open project files..
 */
var FileNavigator = function (parent) {
    this.parent = parent
    this.panel = new Element(parent, { "tag": "div", "class": "fileNavigator" })

    // Keep the list of open files.
    this.files = {}

    // Keep list of files that need save.
    this.toSaves = {}

    // The tabs that contain the file names...
    this.tabs = {}

    this.activeFile = null

    // The save all button.
    this.saveAllBtn = this.panel.appendElement({ "tag": "i", "title": "save all", "class": "fa fa-save-all fileNavigationBtn", "style": "display:none; font-size: 12pt;" }).down()
    this.saveAllBtn.element.onclick = function (fileNavigator) {
        return function () {
            for (var fileId in fileNavigator.toSaves) {
                fileNavigator.saveFile(fileId)
            }
        }
    }(this)

    this.saveBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-floppy-o fileNavigationBtn", "style": "display:none" }).down()

    // Now I will save the file...
    this.saveBtn.element.onclick = function (fileNavigator) {
        return function () {
            if (fileNavigator.toSaves[fileNavigator.activeFile.M_id]) {
                // The file need to be save here...
                fileNavigator.saveFile(fileNavigator.activeFile.M_id)
            }
        }
    }(this)

    // Show the list of file that dosent fit in the file explorer...
    this.showHiddenFilesBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-down fileNavigationBtn" }).down()

    // The show the previous file.
    this.setPreviousFileBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-left fileNavigationBtn" }).down()

    // Show the next file.
    this.setNextFileBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-right fileNavigationBtn" }).down()

    // Here I will attach the file navigator to file event.
    // Open 
    server.fileManager.attach(this, OpenEntityEvent, function (evt, fileNavigator) {
        var uuid
        if (evt.dataMap["fileInfo"] != undefined) {
            uuid = evt.dataMap["fileInfo"].UUID
        } else if (evt.dataMap["bpmnDiagramInfo"] != undefined) {
            uuid = evt.dataMap["bpmnDiagramInfo"].UUID
        }

        var file = entities[uuid]

        if (file == undefined) {
            // local file here.
            file = evt.dataMap["fileInfo"]
            if (file == undefined) {
                file = evt.dataMap["bpmnDiagramInfo"]
            }
            if (evt.dataMap["prototypeInfo"]!=undefined) {
                file = {}
                file.M_id = evt.dataMap["prototypeInfo"].TypeName
                file.M_name = evt.dataMap["prototypeInfo"].TypeName
            }
        }

        if (file !== undefined) {
            if (file.M_id !== undefined) {
                // Here thats mean the file was open
                fileNavigator.appendFile(file)
            }
        }
    })

    server.fileManager.attach(this, UpdateFileEvent, function (evt, codeEditor) {
        if (evt.dataMap.fileInfo !== undefined) {
            var fileId = evt.dataMap["fileInfo"].M_id
            codeEditor.saveBtn.element.title = ""
            codeEditor.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn"

            var tab = codeEditor.tabs[fileId]
            if (codeEditor.toSaves[fileId] != undefined) {
                tab.getChildById("fileNameDiv").element.innerHTML = codeEditor.toSaves[fileId]

                // Remove from the save map
                delete codeEditor.toSaves[fileId]
            }
            // Now the file save button...
            if (Object.keys(codeEditor.toSaves).length <= 1) {
                codeEditor.saveAllBtn.element.style.display = "none"
            }

            if (Object.keys(codeEditor.toSaves).length === 0) {
                codeEditor.saveBtn.element.style.display = "none"
                codeEditor.saveBtn.element.title = ""
            }
        }
    })

    server.entityManager.attach(this, UpdateEntityEvent, function (evt, codeEditor) {
        if (evt.dataMap.entity !== undefined) {
            var fileId = evt.dataMap.entity.M_id
            codeEditor.saveBtn.element.title = ""
            codeEditor.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn"

            var tab = codeEditor.tabs[fileId]
            if (codeEditor.toSaves[fileId] != undefined) {
                tab.getChildById("fileNameDiv").element.innerHTML = codeEditor.toSaves[fileId]

                // Remove from the save map
                delete codeEditor.toSaves[fileId]
            }
            // Now the file save button...
            if (Object.keys(codeEditor.toSaves).length <= 1) {
                codeEditor.saveAllBtn.element.style.display = "none"
            }

            if (Object.keys(codeEditor.toSaves).length === 0) {
                codeEditor.saveBtn.element.style.display = "none"
                codeEditor.saveBtn.element.title = ""
            }
        }
    })

    // Attach the file close event.
    server.fileManager.attach(this, CloseEntityEvent, function (evt, fileNavigator) {
        if (evt.dataMap.fileId !== undefined) {
            // Remove the file.
            fileNavigator.removeFile(evt.dataMap["fileId"])
        }
    })

    // The change file event.
    server.fileManager.attach(this, ChangeFileEvent, function (evt, fileNavigator) {
        if (evt.dataMap.fileId !== undefined) {
            // The file has already change.
            var fileId = evt.dataMap["fileId"]
            if (fileNavigator.toSaves[fileId] !== undefined) {
                return
            }

            // put in the list to save.
            fileNavigator.toSaves[fileId] = fileNavigator.files[fileId].M_name

            // Display the save button.
            fileNavigator.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn active"
            fileNavigator.saveBtn.element.style.display = ""
            fileNavigator.saveBtn.element.title = "save " + fileNavigator.toSaves[fileId]

            // I will put an asterist before the name.
            var tab = fileNavigator.tabs[fileId]
            tab.getChildById("fileNameDiv").element.innerHTML = "* " + fileNavigator.toSaves[fileId]

            // Display the save all option in case of multiple file changes.
            if (Object.keys(fileNavigator.toSaves).length > 1) {
                if (fileNavigator.saveAllBtn.element.style.display != "") {
                    fileNavigator.saveAllBtn.element.className = "fa fa-save-all fileNavigationBtn active"
                    fileNavigator.saveAllBtn.element.style.display = ""
                }
            }
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
        .appendElement({ "tag": "div", "id": "fileNameDiv", "style": "display: table-cell;", "innerHtml": file.M_name })
        .appendElement({ "tag": "i", "id": file.id + "_file_tab_close_btn", "class": "fa fa-times file_tab_close_btn" })

    var fileCloseBtn = this.tabs[file.M_id].getChildById(file.id + "_file_tab_close_btn")
    if (fileCloseBtn != undefined) {
        fileCloseBtn.element.onclick = function (file) {
            return function () {
                // Send event localy...
                var evt = { "code": CloseEntityEvent, "name": FileEvent, "dataMap": { "fileId": file.M_id } }
                server.eventHandler.broadcastLocalEvent(evt)
            }
        }(file)
    }

    this.setActiveTab(file.M_id)

    // now the onclick event...
    this.tabs[file.M_id].element.onclick = function (file, fileNavigator) {
        return function () {
            if (fileNavigator.files[file.M_id] != undefined) {
                fileNavigator.setActiveTab(file.M_id)
            }
        }
    }(file, this)
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
        this.panel.element.style.display = "block"
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
        server.eventHandler.broadcastLocalEvent(evt)

        if (this.toSaves[fileId] != undefined) {
            this.saveBtn.element.title = "save " + this.files[fileId].M_name
            this.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn active"
        } else {
            this.saveBtn.element.title = ""
            this.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn"
        }
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
                this.panel.element.style.display = "none"
            } else if (index >= Object.keys(this.files).length) {
                this.setActiveTab(Object.keys(this.files)[Object.keys(this.files).length - 1])
            } else {
                this.setActiveTab(Object.keys(this.files)[index])
            }
        }
    }

    delete this.toSaves[fileId]

    // Now the file save button...
    if (Object.keys(this.toSaves).length <= 1) {
        this.saveAllBtn.element.style.display = "none"
    }
    if (Object.keys(this.toSaves).length == 0) {
        this.saveBtn.element.style.display = "none"
        this.saveBtn.element.title = ""
    }
    // TODO ask the user to save the file if there is change...
}

/**
 * Save a file with a given id.
 */
FileNavigator.prototype.saveFile = function (fileId) {

    // Now I will save the file.
    var file = entities["CargoEntities.File:" + fileId]

    if (file.M_fileType == 2) {
        // In case of disk file.
        var data = [decode64(file.M_data)]
        var f = null
        try {
            f = new File(data, file.M_name, { type: "text/plain", lastModified: new Date(0) })
        } catch (error) {
            f = new Blob(data, { type: f.M_mime });
            f.name = f.M_name
            f.lastModifiedDate = new Date(0);
        }

        server.fileManager.createFile(file.M_name, file.M_path, f, 256, 256, false,
            // Success callback
            function (result, caller) {
                server.entityManager.saveEntity(caller)
            },
            // Progress callback
            function (index, total, caller) {
            },
            // Error callback
            function (errMsg, caller) {
            },
            file)

    } else {
        // In case of db file
        server.entityManager.saveEntity(file,
            function (result, caller) {
                console.log("-------> file save successfully")
            },
            function (errObj, caller) {
                console.log("-------> fail to save ", errObj)
            },
            {})
    }

}