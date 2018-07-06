/**
 * The bpmn view.
 */
var BpmnExplorer = function (parent) {
    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "bpmn_explorer" }).down()

    // Contain the list of defintion.
    this.definitions = {}

    this.header = this.panel.appendElement({ "tag": "div", "class": "bpmn_explorer_header", "style":"text-align: left;" }).down()

    // Now the definitions upload button...
    this.uploadDefintionsBtn = this.header.appendElement({ "tag": "div", "class": "entity_panel_header_button", "style":"width: 20px;" }).down()
    this.uploadDefintionsBtn.appendElement({ "tag": "i", "class": "fa fa-folder" })

    var fileExplorer = this.uploadDefintionsBtn.appendElement({ "tag": "input", "type": "file", "accept": ".bpmn, .BPMN", "multiple": "", "style": "display: none;" }).down()

    fileExplorer.element.onchange = function (bpmnExplorer) {
        return function (evt) {
            var files = evt.target.files; // FileList object
            for (var i = 0; i < files.length; i++) {
                bpmnExplorer.importDefinitions(files[i])
            }
        }
    }(this)

    // The upload bpmn file action...
    this.uploadDefintionsBtn.element.onclick = function (fileExplorer) {
        return function () {
            // Display the file explorer...
            fileExplorer.element.click()
        }
    }(fileExplorer)

    this.definitionsDiv = this.panel.appendElement({ "tag": "div", "class": "definitionsLst" }).down()

    /** Now the list of defintions. */
    server.fileManager.getFileByPath("/Cargo/svg/filters.svg",
        function (file, caller) {
            var reader = new FileReader();
            // Create a blob from the data received.
            var blob = base64toBlob(file.M_data, file.M_mime)
            reader.onload = function (bpmnExplorer) {
                return function (e) {
                    var body = document.getElementsByTagName("body")[0]
                    var bodyElement = new Element(body, { "tag": "div", "style": "height: 100%; width: 100%;" });
                    var svgSrc = e.target.result;
                    var svgfiltersDiv = bodyElement.appendElement({ "tag": "div", "id": "SvgfiltersDiv", "innerHtml": svgSrc }).down

                    server.entityManager.getEntities("BPMN20.Definitions", "BPMN20", "", 0, -1, [], true, false,
                    // index, total
                    function(index, total, caller){
                        
                    },
                    // success callback
                    function(results, caller){
                        for (var i = 0; i < results.length; i++) {
                            // append the defintions to the bpmnExplorer.
                            initDefinitions(results[i])
                            bpmnExplorer.appendDefinitions(results[i])
                        }
                    },
                    // error callback
                    function(errObj, caller){
                        
                    }, bpmnExplorer)
                }
            }(caller)
            // Read the blob as a text file...
            reader.readAsText(blob);
        },
        // The error callback.
        function () { },
        this)

    // Now the append definition event handeler...
    server.workflowManager.attach(this, NewBpmnDefinitionsEvent, function (bpmnExplorer) {
        return function (evt) {
            if (evt.dataMap["definitionsInfo"] !== undefined) {
                var definition = entities[evt.dataMap["definitionsInfo"].UUID]
                bpmnExplorer.appendDefinitions(definition)
            }
        }
    }(this))
    return this
}

/**
 * Append a new definitions...
 */
BpmnExplorer.prototype.appendDefinitions = function (definitions) {

    /** Here I will append the definition into the list. */
    if (document.getElementById(definitions.M_id + "diagramDiv") == undefined) {
        var diagramsDiv = this.definitionsDiv.appendElement({ "tag": "div", "style": "display: table-row; padding: 2px;" }).down()
            .appendElement({ "tag": "div", "style": "display:table-cell; padding: 2px 2px 2px 10px;" }).down()
            .appendElement({ "tag": "i", "class": "fa fa-caret-right definitionsPanelBtn", "id": "showBtn" + definitions.M_id })
            .appendElement({ "tag": "i", "class": "fa fa-caret-down definitionsPanelBtn", "id": "hideBtn" + definitions.M_id, "style": "display: none;" })
            .appendElement({ "tag": "div", "class": "definitionsDiv", "innerHtml": definitions.M_name })
            .appendElement({ "tag": "div", "id": definitions.M_id + "diagramDiv", "class": "diagramDiv", "style": "display:none;" }).down()

        // Show the diagrams
        var showBtn = this.definitionsDiv.getChildById("showBtn" + definitions.M_id)

        // Hide the diagrams
        var hideBtn = this.definitionsDiv.getChildById("hideBtn" + definitions.M_id)

        showBtn.element.onclick = function (hideBtn, diagramsDiv) {
            return function () {
                this.style.display = "none"
                hideBtn.element.style.display = ""
                diagramsDiv.element.style.display = ""
            }
        }(hideBtn, diagramsDiv)

        hideBtn.element.onclick = function (showBtn, diagramsDiv) {
            return function () {
                this.style.display = "none"
                showBtn.element.style.display = ""
                diagramsDiv.element.style.display = "none"
            }
        }(showBtn, diagramsDiv)

        // Now I will append the diagrams...
        for (var i = 0; i < definitions.M_BPMNDiagram.length; i++) {
            var diagramLnk = diagramsDiv.appendElement({ "tag": "div", "style": "display: table-row; padding: 2px; height: 100%;" }).down()
                .appendElement({
                    "tag": "span", "class": "diagramLnk",
                    "draggable": "false",
                    "innerHtml": definitions.M_BPMNDiagram[i].M_name
                }).down()

            diagramLnk.element.onclick = function (bpmnDiagram) {
                return function () {
                    // Here I will create an open file event...
                    var evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "bpmnDiagramInfo": bpmnDiagram } }
                    server.eventHandler.broadcastLocalEvent(evt)
                }
            }(definitions.M_BPMNDiagram[i])
        }
    }
}

/**
 * Import a workflow on the server.
 */
BpmnExplorer.prototype.importDefinitions = function (file, bpmnExplorer) {
    var reader = new FileReader();

    /** I will read the file content... */
    reader.onload = (function (theFile, bpmnExplorer) {
        return function (e) {
            var text = e.target.result
            // Now I will load the content of the file.
            server.workflowManager.importXmlBpmnDefinitions(text,
                function (result, caller) {
                    /** Nothing todo the the action will be in the event listener. */
                },
                function (errMsg, caller) {

                }, bpmnExplorer)
        };
    })(file, bpmnExplorer);

    reader.readAsText(file);
}