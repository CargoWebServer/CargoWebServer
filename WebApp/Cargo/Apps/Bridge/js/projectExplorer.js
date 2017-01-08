/**
 * That class contain the code to display the projects to the end
 * user.
 */
var ProjectExplorer = function (parent) {

    // Keep reference to the parent.
    this.parent = parent

    this.panel = parent.appendElement({ "tag": "div", "class": "project_explorer" }).down()

    // Here I will init the projects...
    server.projectManager.GetAllProjects(
        // success callback
        function (results, caller) {
            caller.initProjects(results)
            //caller.panel.element.style.display = ""
        },
        // error callback
        function (errorMsg, caller) {
            console.log(errMsg)
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
                        panel.initFilesView(panel.panel, project.M_filesRef[0], 0)
                    }
                } (this, project)
            )
        }
    }
    return this
}

/**
 * Initialisation of the files/folders view.
 */
ProjectView.prototype.initFilesView = function (parent, dir, level) {

    var folderDiv = parent.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
        .appendElement({ "tag": "div", "class": "project_folder" }).down()

    /** The expend button */
    folderDiv.expandBtn = folderDiv.appendElement({ "tag": "i", "class": "fa fa-caret-right", "style": "display:inline;" }).down()

    /** The shrink button */
    folderDiv.shrinkBtn = folderDiv.appendElement({ "tag": "i", "class": "fa fa-caret-down", "style": "display:none;" }).down()

    /** The project title. */
    var proejectTitle = folderDiv.appendElement({ "tag": "div", "innerHtml": dir.M_name, "style": "display:inline; padding-left: 2px;" })

    if (level != 0) {
        folderDiv.element.style.display = "none"
        folderDiv.expandBtn.element.style.display = "inline"
    }

    folderDiv.childsDiv = []
    if (dir.M_files != null) {
        for (var i = 0; i < dir.M_files.length; i++) {
            var file = dir.M_files[i]
            // display the file...
            if (file.M_isDir == true) {
                level = level + 1
                folderDiv.childsDiv.push(this.initFilesView(folderDiv, file, level))
            } else {
                // Here I will append the file...
                var fileDiv = parent.appendElement({ "tag": "div", "style": "display: table-row; width: 100%" }).down()
                    .appendElement({ "tag": "div", "class": "project_file" }).down()

                // Set the file title
                fileDiv.appendElement({ "tag": "div", "innerHtml": file.M_name, "style": "display:inline;" })
                if (level != 0 || file.M_name.startsWith(".")) {
                    fileDiv.element.style.display = "none"
                }

                // Here I will set the action on the file.
                fileDiv.element.onclick = function (file) {
                    return function (evt) {
                        evt.stopPropagation()
                        // - Manage the file in order that all user have the same file view.
                        server.fileManager.openFile(file.M_id,
                            // Success callback
                            function (result, caller) {

                            },
                            // Error callback
                            function (errMsg, caller) {

                            }, this)
                    }
                } (file)
                folderDiv.childsDiv.push(fileDiv)
            }
        }
    }

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
    } (folderDiv, folderDiv.expandBtn, folderDiv.shrinkBtn)

    return folderDiv
}