/**
 *  The interface to show diagram view and insatance view.
 */
var BpmnView = function(editor, diagram){
    this.editor = editor;
    this.panel = editor.panel.appendElement({ "tag": "div", "class": "filePanel", "id": diagram.UUID + "_editor", "style":"overflow: hidden;" }).down()
    var panel = this.panel.appendElement({"tag":"div", "style":"display: flex; flex-direction: column;width: 100%; height: 100%;"}).down();
    
    // So the panel will be divide in tow parts...
    // The query panel.
    var splitArea1 =panel.appendElement({ "tag": "div", "class":"query_editor_splitter_area", "style":"height: 520px;" }).down()
    
    // The edition panel.
    this.diagramPanel = new Element(splitArea1, { "tag": "div", "class": "bpmn_diagram_panel", "style":"width: 100%; height: 100%;" })

    // The splitter.
    var queryEditorSplitor = splitArea1.appendElement({ "tag": "div", "class": "splitter horizontal", "id": "query_editor_splitter" }).down()
    
    queryEditorSplitor.element.style.bottom = "0px"; // set the position of the splitter...
    splitArea1.element.style.flexBasis = "50%"; // Set the initial size of the area
    splitArea1.element.style.paddingBottom = queryEditorSplitor.element.offsetHeight + "px"; // compensate for the height of the slitter (position absolute...)

    // The result panel.
    var splitArea2 = panel.appendElement({ "tag": "div",  "class":"query_editor_splitter_area"}).down()
    
    splitArea2.element.style.flexGrow = 1;
    
    // Contain the list of bpmn instances.
    this.bpmnInstancesPanel = new Element(splitArea2, { "tag": "div", "class": "bpmn_instances_panel" })

    // Init the splitter action.
    initSplitter(queryEditorSplitor, splitArea1)

    // The diagram to display.
    this.diagram = new SvgDiagram(this.diagramPanel, diagram)
    
    this.diagram.init(function (codeEditor, diagram, panel, diagramPanel, instancesPanel) {
        return function () {
            // Display the diagram elements.
            codeEditor.diagram = diagram
            diagram.drawDiagramElements()
            codeEditor.files[diagram.bpmnDiagram.UUID] = diagram
            codeEditor.filesPanel[diagram.bpmnDiagram.UUID] = panel
            codeEditor.setActiveFile(diagram.bpmnDiagram.UUID)
            new InstanceListView(instancesPanel, diagram)
            // Now the resize element...
            diagram.canvas.initWorkspace = function (workspace) {
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
            }(diagramPanel)
            
            window.addEventListener("resize", function (canvas) {
                return function () {
                    canvas.initWorkspace()
                }
            }(diagram.canvas))
            
            // call once...
            diagram.canvas.initWorkspace()
        }
    }(editor, this.diagram, this.panel, this.diagramPanel, this.bpmnInstancesPanel))
    
    
    return this;
}

var DiagramView = function(){
    
    return this;
}

