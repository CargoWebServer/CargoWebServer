/**
 * That class contain the code to display the projects to the end
 * user.
 */
var ProjectExplorer = function (parent) {

    // Keep reference to the parent.
    this.parent = parent

    this.panel = parent.appendElement({ "tag": "div", "class": "project_explorer" }).down()

    // Here I will init the projects...
    server.projectManager.getAllProjects(
        // success callback
        function (results, caller) {
            caller.initProjects(results)
            //caller.panel.element.style.display = ""
        },
        // error callback
        function (errorMsg, caller) {

        }, this)

    return this
}

ProjectExplorer.prototype.initProjects = function (projectsInfo) {
    // Now I will set each project...
    for (var i = 0; i < projectsInfo.length; i++) {
        var projectInfo = projectsInfo[i]
        var projectView = new ProjectView(this.panel, projectInfo)
    }
}

/**
 * The project view object display the content of a project.
 */
var ProjectView = function (parent, project) {
    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "project_view" }).down()

    if (project.M_filesRef != undefined) {
        if (project.M_filesRef.length == 1) {
            project["set_M_filesRef_" + project.M_filesRef + "_ref"](
                function (panel, project) {
                    return function (ref) {
                        var projectView = panel.initFilesView(panel.panel, project.M_filesRef[0], 0)
                        projectView.shrinkBtn.element.click()
                    }
                }(this, project)
            )
        }
    }

    // Now I will connect the events...
    server.entityManager.attach(this, UpdateEntityEvent, function (evt, projectView) {
        // generatePrototypesView = function (storeId, prototypes)
        if (evt.dataMap.entity.TYPENAME == "CargoEntities.File") {
            var parentFolderDiv = projectView.panel.getChildById(evt.dataMap.entity.ParentUuid)
            if (parentFolderDiv != null) {
                // Here I will append the new create folder/file...
                projectView.initFilesView(parentFolderDiv, evt.dataMap.entity, 0)
            }
        }
    })

    server.entityManager.attach(this, NewEntityEvent, function (evt, projectView) {
        // generatePrototypesView = function (storeId, prototypes)
        if (evt.dataMap.entity.TYPENAME == "CargoEntities.File") {
            if (evt.dataMap.entity.M_isDir == false) {
                var parentFolderDiv = projectView.panel.getChildById(evt.dataMap.entity.ParentUuid)
                if (parentFolderDiv != null) {
                    // Here I will append the new create folder/file...
                    projectView.initFilesView(parentFolderDiv, evt.dataMap.entity, 0)
                }
            }
        }
    })

    server.entityManager.attach(this, DeleteEntityEvent, function (evt, projectView) {
        // generatePrototypesView = function (storeId, prototypes)
        if (evt.dataMap.entity.TYPENAME == "CargoEntities.File") {
            var folderDiv = projectView.panel.getChildById(evt.dataMap.entity.UUID)
            if (folderDiv != undefined) {
                folderDiv.element.parentNode.removeChild(folderDiv.element)
            }
        }
    })

    return this
}

/**
 * Create a directory view.
 */
ProjectView.prototype.createDirView = function (parent, dir, level) {

    var folderDiv = parent.getChildById(dir.UUID)
    if (folderDiv == undefined) {

        folderDiv = parent.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
            .appendElement({ "tag": "div", "class": "project_folder", "id": dir.UUID }).down()

        /** The expend button */
        folderDiv.expandBtn = folderDiv.appendElement({ "tag": "i", "class": "fa fa-caret-right", "style": "display:none;" }).down()

        /** The shrink button */
        folderDiv.shrinkBtn = folderDiv.appendElement({ "tag": "i", "class": "fa fa-caret-down", "style": "display:inline;" }).down()

        /** The project title. */
        folderDiv.appendElement({ "tag": "div", "id": dir.UUID + "_folder_name", "innerHtml": dir.M_name, "style": "display:inline; padding-left: 2px;" })

        folderDiv.childsDiv = []

        if (parent.childsDiv == undefined) {
            parent.childsDiv = []
        }

        parent.childsDiv.push(folderDiv)
    } else {
        folderDiv.getChildById(dir.UUID + "_folder_name").element.innerHTML = dir.M_name
        folderDiv.getChildById(dir.UUID + "_folder_name").element.style.display = "inline"
    }

    if (dir.M_files != null) {
        for (var i = 0; i < dir.M_files.length; i++) {
            var file = dir.M_files[i]
            level = level + 1
            folderDiv.childsDiv.push(this.initFilesView(folderDiv, file, level))
        }
    }

    folderDiv.element.addEventListener('contextmenu',
        function (folderDiv, dir) {
            return function (evt) {
                evt.preventDefault();
                evt.stopPropagation();

                // Create a new folder in the project.
                var newFolderMenuItem = new MenuItem("new_folder_menu", "New folder", {}, 0,
                    function (folderDiv, dir) {
                        return function () {
                            // So here I will append an input where to enter the name of the new 
                            // folder.
                            if (folderDiv.shrinkBtn.element.style.display == "none") {
                                // expand the div view.
                                folderDiv.expandBtn.element.click()
                            }
                            var editPanel = folderDiv.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
                            var newFolderNameInput = editPanel.appendElement({ "tag": "div", "class": "project_folder" }).down()
                                .appendElement({ "tag": "i", "class": "fa fa-caret-right", "style": "display: inline;" })
                                .appendElement({ "tag": "input", "style": "display: inline;" }).down()

                            // The user type the name, esc to cancel it action or enter to confirm it...
                            newFolderNameInput.element.focus()
                            newFolderNameInput.element.onclick = function (evt) {
                                evt.stopPropagation()
                            }

                            newFolderNameInput.element.onkeyup =
                                function (editPanel) {
                                    return function (evt) {
                                        evt.stopPropagation()
                                        if (evt.keyCode == 27) {
                                            // escape key
                                            editPanel.element.parentNode.removeChild(editPanel.element)
                                        } else if (evt.keyCode == 13) {
                                            // enter key
                                            // So here I will create a new directory on the sever with the given name.
                                            var path = dir.M_path + "/" + dir.M_name
                                            server.fileManager.createDir(this.value, path,
                                                // success callback
                                                function (result, editPanel) {
                                                    // synchronize the projects...
                                                    server.projectManager.synchronize(function () {

                                                    }, function () {

                                                    }, {})
                                                    editPanel.element.parentNode.removeChild(editPanel.element)
                                                },
                                                // error callback
                                                function (errObj, editPanel) {

                                                }, editPanel)
                                        }
                                    }
                                }(editPanel)
                        }
                    }(folderDiv, dir), "fa fa-folder-o")

                // Append a new file in the project.
                var newFileMenuItem = new MenuItem("new_file_menu", "New file", {}, 0, function (dir) {
                    return function () {
                        if (folderDiv.shrinkBtn.element.style.display == "none") {
                            // expand the div view.
                            folderDiv.expandBtn.element.click()
                        }

                        var editPanel = folderDiv.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
                        var newFileNameInput = editPanel.appendElement({ "tag": "div", "class": "project_folder" }).down()
                            .appendElement({ "tag": "input", "style": "display: inline;" }).down()

                        // The user type the name, esc to cancel it action or enter to confirm it...
                        newFileNameInput.element.focus()
                        newFileNameInput.element.onclick = function (evt) {
                            evt.stopPropagation()
                        }

                        newFileNameInput.element.onkeyup =
                            function (editPanel) {
                                return function (evt) {
                                    evt.stopPropagation()
                                    if (evt.keyCode == 27) {
                                        // escape key
                                        editPanel.element.parentNode.removeChild(editPanel.element)
                                    } else if (evt.keyCode == 13) {
                                        // enter key
                                        // So here I will create a new directory on the sever with the given name.
                                        var path = dir.M_path + "/" + dir.M_name

                                        // Here I will create an empty text file.
                                        var f = null
                                        var mimeType = "text/plain"

                                        if (this.value.toLowerCase().endsWith(".js")) {
                                            mimeType = "application/javascript"
                                        } else if (this.value.toLowerCase().endsWith(".json")) {
                                            mimeType = "application/json"
                                        } else if (this.value.toLowerCase().endsWith(".html")) {
                                            mimeType = "text/html"
                                        } else if (this.value.toLowerCase().endsWith(".xml")) {
                                            mimeType = "text/xml"
                                        } else if (this.value.toLowerCase().endsWith(".css")) {
                                            mimeType = "text/css"
                                        } else if (this.value.toLowerCase().endsWith(".csv")) {
                                            mimeType = "text/csv"
                                        }

                                        var f = new File([""], this.value, { type: mimeType, lastModified: new Date(0) })

                                        // Now I will create the new file...
                                        server.fileManager.createFile(this.value, path, f, 256, 256, false,
                                            // Success callback.
                                            function (result, editPanel) {
                                                // Here is the new file...
                                                editPanel.element.parentNode.removeChild(editPanel.element)
                                            },
                                            // Progress callback.
                                            function () {

                                            },
                                            // Error callback.
                                            function () {

                                            }, editPanel)
                                    }
                                }
                            }(editPanel)

                    }
                }(dir), "fa fa-file-o")

                // Rename existing folder in the project.
                var renameMenuItem = new MenuItem("rename_menu", "Rename", {}, 0, function (folderDiv, dir) {
                    return function () {
                        // Now I will hide the fileDiv...
                        var text = folderDiv.element.childNodes[2].innerText
                        folderDiv.element.childNodes[2].style.display = "none"
                        var renameDirInput = new Element(null, { "tag": "input", "value": text })

                        // Insert at position 3
                        folderDiv.element.insertBefore(renameDirInput.element, folderDiv.element.childNodes[3])

                        renameDirInput.element.onclick = function (evt) {
                            evt.stopPropagation()
                            // nothing todo here.
                        }

                        renameDirInput.element.onkeyup = function (renameDirInput, folderDiv, text) {
                            return function (evt) {
                                evt.stopPropagation()
                                if (evt.keyCode == 27) {
                                    // escape key
                                    renameDirInput.element.parentNode.removeChild(renameDirInput.element)
                                    folderDiv.element.childNodes[2].style.display = "inline"
                                } else if (evt.keyCode == 13) {
                                    // Rename the file here.
                                    server.fileManager.renameFile(dir.UUID, this.value,
                                        // Success callback
                                        function (result, caller) {
                                            caller.renameDirInput.element.parentNode.removeChild(caller.renameDirInput.element)
                                            caller.folderDiv.element.childNodes[2].style.display = "inline"
                                        },
                                        // Error callback
                                        function (errObj, caller) {
                                            caller.renameDirInput.element.parentNode.removeChild(caller.renameDirInput.element)
                                            caller.folderDiv.element.childNodes[2].style.display = "inline"
                                        }, {"renameDirInput":renameDirInput, "folderDiv":folderDiv})
                                        
                                }
                            }
                        }(renameDirInput, folderDiv, text)
                        renameDirInput.element.setSelectionRange(0, text.length)
                        renameDirInput.element.focus()
                    }
                }(folderDiv, dir), "fa fa-edit")

                // Delete a file from a project.
                var deleteMenuItem = new MenuItem("delete_menu", "Delete", {}, 0, function (dir) {
                    return function () {
                        var confirmDialog = new Dialog(randomUUID(), undefined, true)
                        confirmDialog.div.element.style.maxWidth = "450px"
                        confirmDialog.setCentered()
                        server.languageManager.setElementText(confirmDialog.title, "Delete file")
                        confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete file " + dir.M_path + "/" + dir.M_name + "?" })
                        confirmDialog.ok.element.onclick = function (dialog, dir) {
                            return function () {
                                // Remove the folder.
                                server.fileManager.deleteFile(dir.UUID,
                                    // Success callback
                                    function (result, caller) {

                                    },
                                    // Error callback
                                    function () {

                                    }, {})
                                dialog.close()
                            }
                        }(confirmDialog, dir)
                    }
                }(dir), "fa fa-trash-o")

                // The main menu will be display in the body element, so nothing will be over it.
                var contextMenu = new PopUpMenu(folderDiv, [newFolderMenuItem, newFileMenuItem, "|", renameMenuItem, deleteMenuItem], evt)

                return false;
            }
        }(folderDiv, dir), false);

    folderDiv.element.onclick = function (folderDiv, expandBtn, shrinkBtn) {
        return function (evt) {
            evt.stopPropagation()

            // That function is use to shrink the folder recursively.
            function shrinkFolder(childsDiv) {
                for (var i = 0; i < childsDiv.length; i++) {
                    childsDiv[i].element.style.display = "none"
                    if (childsDiv[i].childsDiv != undefined) {
                        shrinkFolder(childsDiv[i].childsDiv)
                        childsDiv[i].expandBtn.element.style.display = ""
                        childsDiv[i].shrinkBtn.element.style.display = "none"
                    }
                }
            }

            var childsDiv = folderDiv.childsDiv
            if (expandBtn.element.style.display == "none") {
                shrinkFolder(childsDiv)
                expandBtn.element.style.display = ""
                shrinkBtn.element.style.display = "none"
            } else {
                for (var i = 0; i < childsDiv.length; i++) {
                    childsDiv[i].element.style.display = ""
                }
                expandBtn.element.style.display = "none"
                shrinkBtn.element.style.display = ""
            }

        }
    }(folderDiv, folderDiv.expandBtn, folderDiv.shrinkBtn)

    return folderDiv
}

/**
 * Create a file view.
 */
ProjectView.prototype.createFileView = function (parent, file) {
    // Here I will append the file...
    var fileDiv = parent.getChildById(file.UUID)
    if (fileDiv == undefined) {
        fileDiv = parent.appendElement({ "tag": "div", "style": "display: table-row; width: 100%" }).down()
            .appendElement({ "tag": "div", "class": "project_file", "style": "display: inline;", "id": file.UUID }).down()
        // Set the file title
        fileDiv.appendElement({ "tag": "div", "id": file.UUID + "_file_name_div", "innerHtml": file.M_name, "style": "display:inline;" })
    } else {
        fileDiv.getChildById(file.UUID + "_file_name_div").element.innerHTML = file.M_name
    }
    fileDiv.element.firstChild.style.display = "inline"
    // Here I will set the action on the file.
    fileDiv.element.onclick = function (file) {
        return function (evt) {
            evt.stopPropagation()
            // - Manage the file in order that all user have the same file view.
            server.fileManager.openFile(file.M_id,
                // Progress callback.
                function (index, totatl, caller) {

                },
                // Success callback
                function (result, caller) {

                },
                // Error callback
                function (errMsg, caller) {

                }, this)
        }
    }(file)

    // Now the context menu of file.
    fileDiv.element.addEventListener('contextmenu',
        function (fileDiv, file) {
            return function (evt) {
                evt.preventDefault();
                evt.stopPropagation();

                // So here I will create the popup menu for file.
                // Rename existing file in the project.
                var renameMenuItem = new MenuItem("rename_menu", "Rename", {}, 0, function (fileDiv, file) {
                    return function () {
                        // Now I will hide the fileDiv...
                        var text = fileDiv.element.firstChild.innerText
                        fileDiv.element.firstChild.style.display = "none"
                        var renameFileInput = new Element(fileDiv, { "tag": "input", "value": text })
                        renameFileInput.element.onclick = function (evt) {
                            evt.stopPropagation()
                            // nothing todo here.
                        }

                        renameFileInput.element.onkeyup = function (renameFileInput, fileDiv, text) {
                            return function (evt) {
                                evt.stopPropagation()
                                if (evt.keyCode == 27) {
                                    // escape key
                                    renameFileInput.element.parentNode.removeChild(renameFileInput.element)
                                    fileDiv.element.firstChild.style.display = "inline"
                                } else if (evt.keyCode == 13) {
                                    // Rename the file here.
                                    server.fileManager.renameFile(file.UUID, this.value,
                                        // Success callback
                                        function (result, caller) {
                                            caller.renameFileInput.element.parentNode.removeChild(caller.renameFileInput.element)
                                            caller.fileDiv.element.firstChild.style.display = "inline"
                                        },
                                        // Error callback
                                        function (errMsg, caller) {
                                            caller.renameFileInput.element.parentNode.removeChild(caller.renameFileInput.element)
                                            caller.fileDiv.element.firstChild.style.display = "inline"
                                        }, {"renameFileInput":renameFileInput, "fileDiv":fileDiv})
                                        
                                }
                            }
                        }(renameFileInput, fileDiv, text)

                        renameFileInput.element.setSelectionRange(0, text.indexOf("."))
                        renameFileInput.element.focus()
                    }
                }(fileDiv, file), "fa fa-edit")

                // Delete a file from a project.
                var deleteMenuItem = new MenuItem("delete_menu", "Delete", {}, 0, function (file) {
                    return function () {
                        var confirmDialog = new Dialog(randomUUID(), undefined, true)
                        confirmDialog.div.element.style.maxWidth = "450px"
                        confirmDialog.setCentered()
                        server.languageManager.setElementText(confirmDialog.title, "Delete file")
                        confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete file " + file.M_name + "?" })
                        confirmDialog.ok.element.onclick = function (dialog, dir) {
                            return function () {
                                // Remove the folder.
                                server.fileManager.deleteFile(file.UUID,
                                    // Success callback
                                    function (result, caller) {

                                    },
                                    // Error callback
                                    function () {

                                    }, {})
                                dialog.close()
                            }
                        }(confirmDialog, file)
                    }
                }(file), "fa fa-trash-o")
                // The main menu will be display in the body element, so nothing will be over it.
                var contextMenu = new PopUpMenu(fileDiv, [renameMenuItem, deleteMenuItem], evt)
                return false
            }
        }(fileDiv, file), false)

    return fileDiv
}

/**
 * Initialisation of the files/folders view.
 */
ProjectView.prototype.initFilesView = function (parent, file, level) {
    this.parent = parent
    var view = null
    parent.id = file.ParentUuid
    if (file.M_isDir) {
        view = this.createDirView(parent, file, level)
    } else {
        view = this.createFileView(parent, file)
    }
    return view
}