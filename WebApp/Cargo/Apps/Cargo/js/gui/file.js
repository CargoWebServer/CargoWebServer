var languageInfo = {
    "en": {
        "delete_file_question": "Do you really want to remove this file?",
        "delete_dialog_title": "Remove file"
    },
    "fr": {
        "delete_file_question": "Voulez-vous vraiment enlever ce fichier?",
        "delete_dialog_title": "Enlever le fichier"
    }
}

// Set the text...
server.languageManager.appendLanguageInfo(languageInfo)

/*
 * The file panel is use manipulate file. Files are store in the server side in the db
 * and the server file system depending if the file is a db file or disk file.
 * file information are use to keep information of a file.
 */
var FilePanel = function (parent, file, filePath, fileInfo, loadCallback, readCallback) {

    // The parent panel.
    this.parent = parent

    // Keep the callback refs...
    this.loadCallback = loadCallback
    this.readCallback = readCallback

    // The file panel.
    this.panel = new Element(parent.panel, { "tag": "div", "class": "file_panel" })

    // The progress bar use when the file is read, download or upload.
    this.progressBar = this.panel.appendElement({ "tag": "progress" }).down()

    // The local file...
    this.file = file

    // The file data
    this.fileData = null

    // The file name 
    this.fileName = ""

    // The file type.
    this.fileMime = ""

    // The filepath on the server is ther is one...
    this.filePath = ""

    if (file != undefined) {
        // Get the information from the local file.
        this.fileName = file.name
        this.fileMime = file.type
    }

    this.filePath = ""
    if (filePath != undefined) {
        // Set the file path, (path on the server)
        this.filePath = filePath
    }

    // The file info...
    this.fileInfo = null
    if (fileInfo != null) {
        // Set the file info...
        this.fileInfo = fileInfo
        this.fileName = fileInfo.M_name
        this.fileMime = fileInfo.M_mime
        this.filePath = fileInfo.M_path
    }

    // Html5 file reader...
    this.reader = new FileReader();
    this.img = null

    // Closure to capture the file information.
    this.reader.onload = function (filePanel, fileName, loadCallback) {
        return function (e) {
            filePanel.fileName = fileName
            if (filePanel.fileMime.startsWith("image/")) {
                filePanel.fileData = e.target.result
                filePanel.setImg()
            } else {
                // Here i will get the mime type image...
                var fileExtension = fileName.split(".")
                fileExtension = fileExtension[fileExtension.length - 1]
                server.fileManager.getMimeTypeByExtension("." + fileExtension,
                    // Success callback
                    function (result, filePanel) {
                        filePanel.fileData = result.Thumbnail
                        filePanel.setImg()
                    },
                    // Error callback
                    function () { }
                    , filePanel
                )
            }
            if (loadCallback != undefined) {
                loadCallback(filePanel)
                filePanel.parent.panel.element.style.borderColor = "lightgrey"
            }
        }
    } (this, this.fileName, loadCallback);

    /*
     * Set the local file loading progress bar... usefull for large files... 
     **/
    this.reader.onprogress = function (progressBar) {
        return function (evt) {
            if (evt.lengthComputable) {
                var percentLoaded = Math.round((evt.loaded / evt.total) * 100);
                // Increase the progress bar length.
                if (percentLoaded < 100) {
                    progressBar.element.value = percentLoaded;
                }
            }
        }
    } (this.progressBar)

    this.reader.onloadstart = function (progressBar) {
        return function () {
            progressBar.element.max = 100
            progressBar.element.value = 0
        }
    } (this.progressBar)

    // Read in the image file as a data URL.
    if (this.file != undefined) {
        this.reader.readAsDataURL(this.file);
    } else if (this.fileInfo != undefined) {
        if (this.fileMime.startsWith("image/")) {
            this.fileData = this.fileInfo.M_thumbnail
            this.setImg()
        } else {
            // Here i will get the mime type image...
            var fileExtension = this.fileName.split(".")
            fileExtension = fileExtension[fileExtension.length - 1]
            server.fileManager.getMimeTypeByExtension("." + fileExtension.toLowerCase(),
                // Success callback
                function (result, filePanel) {
                    filePanel.fileData = result.Thumbnail
                    filePanel.setImg()
                },
                // Error callback
                function () { }
                , this
            )
        }
        this.fileName = this.fileInfo.M_name
        this.filePath = this.fileInfo.M_path
        this.fileMime = this.fileInfo.M_mime
    }

    // Now the delete button...
    this.deleteBtn = this.panel.appendElement({ "tag": "div", "class": "file_panel_delete_button" }).down()
    this.deleteBtn.appendElement({ "tag": "i", "class": "fa fa-close" }).down()

    // Mousse in and out button...
    this.panel.element.addEventListener("mouseout", function (deleteBtn) {
        return function () {
            deleteBtn.element.style.display = "none"
        }
    } (this.deleteBtn), false);

    this.panel.element.addEventListener("mouseover", function (deleteBtn) {
        return function () {
            deleteBtn.element.style.display = "block"
        }
    } (this.deleteBtn), false);

    // Now the remove file...
    this.deleteBtn.element.onclick = function (filePanel) {
        return function (e) {
            e.stopPropagation()
            var confirmDialog = new Dialog(randomUUID(), undefined, true)
            server.languageManager.setElementText(confirmDialog.title, "delete_dialog_title")
            confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Voulez-vous enlever le fichier " + filePanel.fileName + "?" })
            confirmDialog.div.element.style.maxWidth = "450px"
            confirmDialog.setCentered()
            confirmDialog.ok.element.onclick = function (dialog, filePanel) {
                return function () {
                    // I will call delete file
                    delete filePanel.parent.filePanels[filePanel.fileName]
                    filePanel.parent.panel.removeElement(filePanel.panel)
                    if (filePanel.fileInfo != undefined) {
                        server.fileManager.deleteFile(filePanel.fileInfo.UUID,
                            // Success callback
                            function (file, filePanel) {
                                // Nothing todo here...
                            },
                            // Error callback
                            function () { },
                            filePanel)
                    }
                    dialog.close()
                }
            } (confirmDialog, filePanel)
        }
    } (this)

    return this
}

FilePanel.prototype.setImg = function () {
    this.img = this.panel.appendElement({ "tag": "img", "src": this.fileData, "title": escape(this.fileName) }).down()

    // If the file is not an image I will also display the file name over it...
    if (this.fileMime.startsWith("image/") == false) {
        this.panel.appendElement({ "tag": "div", "class": "file_title", "style": "font-size:8pt;", "innerHtml": this.fileName })
        this.img.element.style.width = "64px"
        this.img.element.style.height = "64px"
    }

    var paddingTop = ((this.panel.element.offsetHeight - this.img.element.offsetHeight) / 2) - 2
    this.img.element.style.paddingTop = paddingTop + "px"
    this.progressBar.element.value = 100
    this.progressBar.element.style.display = "none"
    this.progressBar.element.style.zIndex = 10

    // Here I will set the methode to open the file...
    this.img.element.onclick = function (filePanel) {
        return function () {
            // Here I will test if the file data is not undefined or not...
            if (filePanel.file != undefined) {
                filePanel.readCallback(filePanel.file, filePanel.fileName, filePanel.filePath, filePanel.file.type, filePanel)
            } else if (filePanel.fileInfo.M_data != undefined && filePanel.fileInfo.M_data != "") {
                var file = base64toBlob(filePanel.fileInfo.M_data, filePanel.fileInfo.M_mime)
                filePanel.readCallback(file, filePanel.fileName, filePanel.filePath, filePanel.fileInfo.M_mime, filePanel)
            } else {
                filePanel.readCallback(null, filePanel.fileName, filePanel.filePath, filePanel.fileMime, filePanel)
            }
        }
    } (this)
}

/*
 * Upload the file 
 */
FilePanel.prototype.uploadFile = function (path, functor) {
    if (this.file == undefined) {
        if (functor != undefined) {
            // call the functor!!!
            functor()
        }
        return
    }

    server.fileManager.isFileExist(this.file.name, path,
        // Success Callback. 
        function (result, caller) {
            if (result == false) {
                server.fileManager.createFile(caller.name, caller.path, caller.filePanel.file, 150, 150, false,
                    // The success callback
                    function (result, caller) {
                        var filePanel = caller.filePanel
                        var functor = caller.functor

                        filePanel.progressBar.element.style.display = "none"
                        filePanel.fileInfo = result
                        if (functor != undefined) {
                            // call the functor!!! return the file entity.
                            functor(result)
                        }

                        console.log("file upload sucessfully!")
                    },
                    // The progress callback
                    function (index, total, caller) {
                        var filePanel = caller.filePanel
                        if (index == 0) {
                            // Here I will display the loading bar...
                            filePanel.progressBar.element.style.display = "block"
                        }
                        // Set the progress value.
                        filePanel.progressBar.element.value = (index / total) * 100

                        if (index == total - 1) {
                            filePanel.progressBar.element.value = 100
                            filePanel.progressBar.element.style.display = "none"
                        }
                    },
                    // The error callback.
                    function (errMsg, caller) {
                        console.log(errMsg)
                    },
                    // Caller content...
                    { "filePanel": caller.filePanel, "functor": caller.functor }
                )

            } else {
                // Nothing todo here...
                console.log("file already exist!!!!")
                // nothing to upload here...
                functor()
            }
        },
        // Error callback.
        function () { },
        { "name": this.file.name, "path": path, "filePanel": this, "functor": functor })
}


/*
 * That panel contain a group of files (file panel)
 */
var FilesPanel = function (parent, path, filesLoadCallback, filesReadCallback) {

    this.path = path // set latter...

    // Here I will keep the file to upload...
    this.filePanels = {}

    // The files load callback
    this.filesLoadCallback = filesLoadCallback

    // The file read callback...
    this.filesReadCallback = filesReadCallback

    /* The file upload panel... **/
    this.panel = new Element(parent, { "tag": "div", "class": "file_upload" })

    this.fileInput = parent.appendElement({ "tag": "input", "type": "file", "multiple": "", "style": "display: none;" }).down()

    this.fileInput.element.onclick = function () {
        this.value = null;
    };

    // The upload file botton
    this.uploadFileBtn = parent.appendElement({ "tag": "div", "class": "upload_file_btn" }).down()
    this.uploadFileBtn.appendElement({ "tag": "img", "src": "/Cargo/svg/paper_clip.svg" })

    this.uploadFileBtn.element.onclick = function (fileInput) {
        return function () {
            fileInput.element.click()
        }
    } (this.fileInput)

    this.fileInput.element.onchange = function (fileUpload, filesLoadCallback, filesReadCallback) {
        return function (evt) {
            /* Here I will get the list of files... **/
            var files = evt.target.files // FileList object.
            for (var i = 0, f; f = files[i]; i++) {
                // Append the new file panel...
                if (fileUpload.filePanels[f.name] == undefined) {
                    if (i == files.length - 1) {
                        fileUpload.filePanels[f.name] = new FilePanel(fileUpload, f, "", undefined, filesLoadCallback, filesReadCallback)
                    } else {
                        fileUpload.filePanels[f.name] = new FilePanel(fileUpload, f, "", undefined, filesLoadCallback, filesReadCallback)
                    }
                }
            }
        }
    } (this, filesLoadCallback, filesReadCallback)

    this.panel.element.addEventListener('dragover', function (fileUpload) {
        return function (evt) {
            evt.stopPropagation();
            evt.preventDefault();
            evt.dataTransfer.dropEffect = 'copy'; // Explicitly show this is a copy.
            fileUpload.panel.element.style.borderColor = "darkgrey"
        }
    } (this), false)

    this.panel.element.addEventListener('dragleave', function (fileUpload) {
        return function (evt) {
            fileUpload.panel.element.style.borderColor = "lightgrey"
        }
    } (this), false)

    this.panel.element.addEventListener('drop', function (fileUpload, filesLoadCallback, filesReadCallback) {
        return function (evt) {
            /* Here I will get the list of files... **/
            evt.stopPropagation();
            evt.preventDefault();
            var files = evt.dataTransfer.files // FileList object.
            for (var i = 0, f; f = files[i]; i++) {
                // Append the new file panel...
                if (fileUpload.filePanels[f.name] == undefined) {
                    if (i == files.length - 1) {
                        fileUpload.filePanels[f.name] = new FilePanel(fileUpload, f, "", undefined, filesLoadCallback, filesReadCallback)
                    } else {
                        fileUpload.filePanels[f.name] = new FilePanel(fileUpload, f, "", undefined, filesLoadCallback, filesReadCallback)
                    }
                }
            }
        }
    } (this, filesLoadCallback, filesReadCallback), false)

    ///////////////////////////////////////////////////////////////////////////////////////
    // The listeners...
    ///////////////////////////////////////////////////////////////////////////////////////
    server.fileManager.attach(this, DeleteFileEvent, function (evt,filesPanel) {
            var toDelete = evt.dataMap.fileInfo.M_name
            if (filesPanel.filePanels[toDelete] != undefined && filesPanel.filePanels[toDelete].filePath == evt.dataMap.fileInfo.M_path) {
                // I will remove the file panel...
                filesPanel.filePanels[toDelete].panel.element.parentNode.removeChild(filesPanel.filePanels[toDelete].panel.element)
                delete filesPanel.filePanels[toDelete]
            }
        })

    server.fileManager.attach(this, NewFileEvent, function (evt, filesPanel) {
            /*if(filesPanel.path == ""){
                return
            }*/
            var toAppend = evt.dataMap.fileInfo.M_name
            if (filesPanel.filePanels[toAppend] == undefined && evt.dataMap.fileInfo.M_path == filesPanel.path) {
                filesPanel.filePanels[evt.dataMap.fileInfo.M_name] = new FilePanel(filesPanel, undefined, "", evt.dataMap.fileInfo, filesPanel.filesLoadCallback, filesPanel.filesReadCallback)
            }
        })

    // Read the list of file for the given path and display
    // files in the panel.
    if(path.length > 0){
        var name = path.substring(path.lastIndexOf("/") + 1)
        path = path.substring(0, path.lastIndexOf("/"))
        this.readFiles(name, path, filesReadCallback) 
    }

    return this
}

/*
 * Clear actual files display in the panel...
 */
FilesPanel.prototype.clearFiles = function (name, path) {
    for (var id in this.filePanels) {
        delete this.filePanels[id]
    }

    // Remove all the files...
    this.panel.removeAllChilds()
}

/*
 * That function will be call to get the content of a directory...
 */
FilesPanel.prototype.readFiles = function (name, path, callback) {
    // First of all I will clear the files...
    this.clearFiles()
    server.fileManager.getFileByPath(path + "/" + name,
        function (dir, caller) {
            var callback = caller.callback
            var caller = caller.caller
            if (dir.M_files != null) {
                for (var i = 0; i < dir.M_files.length; i++) {
                    if (caller.filePanels[dir.M_files[i].M_name] == undefined) {
                        caller.filePanels[dir.M_files[i].M_name] = new FilePanel(caller, undefined, "", dir.M_files[i], caller.filesLoadCallback, caller.filesReadCallback)
                    }
                }
            }
            if (callback != undefined) {
                callback()
            }
        },
        function () { },
        { "caller": this, "callback": callback })
}

/*
 * Upload the file from a given path...
 */
FilesPanel.prototype.uploadFiles = function (name, path, successCallback) {
    // I will create the dir...
    server.fileManager.createDir(name, path,
        // Progress callback.
        function (result, caller) {
            var filesPanel = caller.caller
            var successCallback = caller.successCallback

            for (var i = 0; i < Object.keys(filesPanel.filePanels).length; i++) {
                var id = Object.keys(filesPanel.filePanels)[i]

                filesPanel.filePanels[id].uploadFile(result.M_path + "/" + result.M_name, function (successCallback, done) {
                    return function () {
                        if (done) {
                            successCallback()
                        }
                    }
                } (successCallback, i == Object.keys(filesPanel.filePanels).length - 1))

            }
        },
        // Error callback
        function () {
            // Error management here...
        }, { "caller": this, "successCallback": successCallback })
}