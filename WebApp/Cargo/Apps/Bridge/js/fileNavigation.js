// global variable use to parse JavaScript file.
var antlr4;
var ECMAScriptLexer;
var ECMAScriptParser;

/**
 * That component is use to reach code symbols.
 */
 var CodeNavigator = function(parent){
    // So here I will create the file parser.
    antlr4 = require('antlr4/index');
    ECMAScriptLexer = require('generated-parser/ECMAScriptLexer');
    ECMAScriptParser = require('generated-parser/ECMAScriptParser');

    return this;
 }
 
/**
 * Set active file.
 */
CodeNavigator.prototype.setFile = function(file){
    
}

/**
 * Append a file.
 */
CodeNavigator.prototype.appendFile = function(file){
    if(file.M_mime == "application/javascript"){
        // The source code.
        var input = Base64.decode(file.M_data);
        var chars = new antlr4.InputStream(input);
        var lexer = new ECMAScriptLexer.ECMAScriptLexer(chars);
        var tokens  = new antlr4.CommonTokenStream(lexer);
        var parser = new ECMAScriptParser.ECMAScriptParser(tokens);
        parser.buildParseTrees = true;
        var tree = parser.program(); // 'program' is the start rule.
        console.log(tree);
    }
}

/**
 * Remove a file.
 */
CodeNavigator.prototype.removeFile = function(fileId){
    
}

/**
 * The file selector.
 */
var FileSelector = function(parent){
    this.files = {}
    
    // Keep ref to it parent.
    this.parent = parent
    
    this.selector = parent.panel.appendElement({"tag":"select"}).down();
    
    // Set the current file.
    this.selector.element.onchange = function(fileSelector){
        return function(){
            fileSelector.parent.setActiveFile(this.value)
        }
    }(this)
    return this;
}

/**
 * Set active file.
 */
FileSelector.prototype.setFile = function(file){
    // Set the selected value.
    this.selector.element.value = file.UUID
}

/**
 * Append a file.
 */
FileSelector.prototype.appendFile = function(file){
    this.files[file.UUID] = file
    this.selector.appendElement({"tag":"option", "id":file.UUID + "_file_selector_option", "value":file.UUID, "innerHtml":file.M_name})
}

/**
 * Remove a file.
 */
FileSelector.prototype.removeFile = function(fileId){
    var option = document.getElementById(fileId + "_file_selector_option");
    option.parentNode.removeChild(option)
    delete this.files[fileId]
}

/**
 * Remove asterisk to file name.
 */
FileSelector.prototype.resetToSave = function(file){
    var option = document.getElementById(file.UUID + "_file_selector_option");
    option.innerHTML = file.M_name
}

/**
 * Append the asterisk
 */
FileSelector.prototype.setToSave = function(file){
    var option = document.getElementById(file.UUID + "_file_selector_option");
    option.innerHTML = "* " + file.M_name
}

/**
 * The file navigator is use to display open project files..
 */
var FileNavigator = function (parent) {
    this.parent = parent
    this.panel = new Element(parent, { "tag": "div", "class": "toolbar file_navigator", "style":"display:none;" })

    // Keep the list of open files.
    this.files = {}

    // Keep list of files that need save.
    this.toSaves = {}

    // Keep a link to active file.
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
            if (fileNavigator.toSaves[fileNavigator.activeFile.UUID]) {
                // The file need to be save here...
                fileNavigator.saveFile(fileNavigator.activeFile.UUID)
            }
        }
    }(this)


    // The show the previous file.
    this.setPreviousFileBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-angle-left fileNavigationBtn" }).down()

    // Show the next file.
    this.setNextFileBtn = this.panel.appendElement({ "tag": "i", "class": "fa fa-angle-right fileNavigationBtn" }).down()

    // Create the file selector.
    this.fileSelector = new FileSelector(this)
    
    // The code selector.
    //this.codeNavigator = new CodeNavigator();
    
    // Now the close button.
    this.closeButton = this.panel.appendElement({"tag":"i", "class":"fa fa-times fileNavigationBtn"}).down()
    
    this.closeButton.element.onclick = function(fileNavigator){
        return function(){
            var file = fileNavigator.activeFile
            // Send event localy...
            var evt = { "code": CloseEntityEvent, "name": FileEvent, "dataMap": { "fileId": file.UUID } }
            server.eventHandler.broadcastLocalEvent(evt)
        }
    }(this)

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
            if (evt.dataMap["fileInfo"] != undefined) {
                file = evt.dataMap["fileInfo"]
            } else if (evt.dataMap["prototypeInfo"] != undefined) {
                file = {}
                file.M_id = evt.dataMap["prototypeInfo"].TypeName
                file.M_name = evt.dataMap["prototypeInfo"].TypeName
                file.UUID = evt.dataMap["prototypeInfo"].TypeName
            } else if (evt.dataMap["bpmnDiagramInfo"] != undefined) {
                file = evt.dataMap["bpmnDiagramInfo"]
            } else if (evt.dataMap["searchInfo"] != undefined) {
                file = evt.dataMap["searchInfo"]
                if (file.M_id == -1) {
                    for (var key in fileNavigator.files) {
                        if (fileNavigator.files[key].constructor.name == "SearchInfo" && fileNavigator.files[key].UUID != file.UUID) {
                            if (file.M_id <= fileNavigator.files[key].M_id) {
                                file.M_id = fileNavigator.files[key].M_id
                            }
                        }
                    }
                    file.M_id += 1
                }
                // Set the name of the search here...
                file.M_name = "search " + (file.M_id + 1)
            }
        }

        if (file !== undefined) {
            if (file.UUID !== undefined) {
                // Here thats mean the file was open
                fileNavigator.appendFile(file)
            }
        }
    })

    server.fileManager.attach(this, UpdateFileEvent, function (evt, fileNavigator) {
        if (evt.dataMap.fileInfo !== undefined) {
            var fileId = evt.dataMap["fileInfo"].UUID
            fileNavigator.saveBtn.element.title = ""
            fileNavigator.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn"
            
            if (fileNavigator.toSaves[fileId] != undefined) {
                // Remove from the save map
                delete fileNavigator.toSaves[fileId]
                fileNavigator.fileSelector.resetToSave(fileNavigator.files[fileId])
            }
    
            // Now the file save button...
            if (Object.keys(fileNavigator.toSaves).length <= 1) {
                fileNavigator.saveAllBtn.element.style.display = "none"
            }

            if (Object.keys(fileNavigator.toSaves).length === 0) {
                fileNavigator.saveBtn.element.style.display = "none"
                fileNavigator.saveBtn.element.title = ""
            }
        }
    })

    server.entityManager.attach(this, UpdateEntityEvent, function (evt, fileNavigator) {
        if (evt.dataMap.entity !== undefined) {
            var fileId = evt.dataMap.entity.UUID
            fileNavigator.saveBtn.element.title = ""
            fileNavigator.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn"

            if (fileNavigator.toSaves[fileId] != undefined) {
                // Remove from the save map
                delete fileNavigator.toSaves[fileId]
                fileNavigator.fileSelector.resetToSave(fileNavigator.files[fileId])
            }
            
            // Now the file save button...
            if (Object.keys(fileNavigator.toSaves).length <= 1) {
                fileNavigator.saveAllBtn.element.style.display = "none"
            }

            if (Object.keys(fileNavigator.toSaves).length === 0) {
                fileNavigator.saveBtn.element.style.display = "none"
                fileNavigator.saveBtn.element.title = ""
            }

        }
    })

    // Attach the file close event.
    server.fileManager.attach(this, CloseEntityEvent, function (evt, fileNavigator) {
        if (evt.dataMap.fileId !== undefined) {
            // Remove the file.
            fileNavigator.fileSelector.removeFile(evt.dataMap["fileId"])
            fileNavigator.removeFile(evt.dataMap["fileId"])
        }

        // Now the network event.
        var entityInfo = {
            "TYPENAME": "Server.MessageData",
            "Name": "FileInfo",
            "Value": {
                "sessionId": server.sessionId,
                "fileId": evt.dataMap.fileId
            }
        }

        // Also broadcast the event over the network...
        server.eventManager.broadcastNetworkEvent(
            CloseFileEvent,
            evt.dataMap.fileId + "_editor",
            [entityInfo],
            function () { },
            function () { },
            undefined
        )
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
            fileNavigator.fileSelector.setToSave(fileNavigator.files[fileId])
            
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

FileNavigator.prototype.setActiveFile = function (fileId) {

    // Keep the ref of the file.
    this.activeFile = this.files[fileId]
    this.panel.element.style.display = ""
    
    this.fileSelector.setFile(this.activeFile)
    // I will generate the event so other panel will set the current file... 
    if (this.files[fileId] != undefined) {
        // local event.
        var evt
        if (this.activeFile.TYPENAME == "BPMNDI.BPMNDiagram") {
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "bpmnDiagramInfo": this.activeFile } }
        } else if (this.activeFile.TYPENAME == "CargoEntities.File") {
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": this.activeFile } }
        } else if (this.activeFile.TYPENAME == "SearchInfo") {
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "searchInfo": this.activeFile } }
        } else {
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "prototypeInfo": getEntityPrototype(this.activeFile.M_name) } }
        }

        server.eventHandler.broadcastLocalEvent(evt)

        if (this.toSaves[fileId] != undefined) {
            this.saveBtn.element.title = "save " + this.files[fileId].M_name
            this.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn active"
        } else {
            this.saveBtn.element.title = ""
            this.saveBtn.element.className = "fa fa-floppy-o fileNavigationBtn"
        }
        this.panel.element.style.display = "";
    }
}

/**
 * That function will append a new file in the file navigator.
 */
FileNavigator.prototype.appendFile = function (file) {
    // diplay the panel.
    if (this.activeFile != undefined) {
        if (this.activeFile.UUID == file.UUID) {
            // Nothing todo here...
            return
        }
    }

    // Keep the ref of the file.
    if (this.files[file.UUID] != undefined) {
        // Set the tab active...
        this.setActiveFile(file.UUID)
        return
    }

    // Here the new file tab must be created.
    this.files[file.UUID] = file

    // Append the file to the file navigator.
    this.fileSelector.appendFile(file)
    
    // Append the file to the code navigator.
   // this.codeNavigator.appendFile(file)

    // Set the active file.
    this.setActiveFile(file.UUID)

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
    delete this.toSaves[fileId]

    // Now the file save button...
    if (Object.keys(this.toSaves).length <= 1) {
        this.saveAllBtn.element.style.display = "none"
    }
    if (Object.keys(this.toSaves).length == 0) {
        this.saveBtn.element.style.display = "none"
        this.saveBtn.element.title = ""
    }
    
    if (this.activeFile != null) {
        if (this.activeFile.UUID == fileId) {
            //this.activeFile = null
            // Now I will set the active file
            if (Object.keys(this.files).length == 0) {
                this.activeFile = null
                this.panel.element.style.display = "none"
            } else if (index >= Object.keys(this.files).length) {
                this.setActiveFile(Object.keys(this.files)[Object.keys(this.files).length - 1])
            } else {
                this.setActiveFile(Object.keys(this.files)[index])
            }
        }
    }

}

/**
 * Save a file with a given id.
 */
FileNavigator.prototype.saveFile = function (fileId) {

    // Now I will save the file.
    var file = entities[fileId]

    if (file.M_fileType == 2) {
        // In case of disk file.
        var data = [decode64(file.M_data)]
        var f = null
        try {
            f = new File(data, file.M_name, { type: "text/plain", lastModified: new Date(0) })
        } catch (error) {
            f = new Blob(data, { type: f.M_mime });
            f.name = f.M_name
            f.File_ModifiedDateTime = new Date(0);
        }
        if (file.UUID.length == 0) {
            file.M_data = ""
            server.fileManager.createFile(file.M_name, file.M_path, f, 256, 256, false,
                // Success callback
                function (result, caller) {
                    //server.entityManager.saveEntity(caller)
                },
                // Progress callback
                function (index, total, caller) {
                },
                // Error callback
                function (errMsg, caller) {
                },
                /*file*/ {})
        } else {
            // File data are send via http post and not the websocket in that case.
            file.M_data = ""
            server.fileManager.saveFile(file, f, 256, 256,
                // Success callback
                function (result, caller) {
                },
                // Progress callback
                function (index, total, caller) {
                },
                // Error callback
                function (errMsg, caller) {
                },
                null)
        }

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
