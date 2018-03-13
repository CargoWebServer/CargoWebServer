/*
 * The list of instances.
 */
var InstanceListView = function (parent, svgDiagram) {
	// The parent div.
	this.parent = parent

	/* This is the link to the related diagram view **/
	this.svgDiagram = svgDiagram

	/** So here I will get the list of instance for the process */
	//this.svgDiagram.bpmnDiagram.M_BPMNPlane.M_bpmnElement
	var definitions = this.svgDiagram.bpmnDiagram.getParentDefinitions()
	var bpmnProcess = null
	for(var i=0; i < definitions.M_rootElement.length; i++){
		if(definitions.M_rootElement[i].M_id == this.svgDiagram.bpmnDiagram.M_BPMNPlane.M_bpmnElement){
			bpmnProcess = definitions.M_rootElement[i]
			break
		}
	}

	if(bpmnProcess != null){
		server.workflowProcessor.getActiveProcessInstances(bpmnProcess.UUID, 
		// Success callback...
		function(instances, instanceListView){
			console.log(instances)
		}, 
		function(){

		}, this)
	}
	return this
}


/////////////////////////////////////////////////////////////////////////////////////////
// The data input wizard...
/////////////////////////////////////////////////////////////////////////////////////////
var ProcessWizard = function (parent, startEvent) {
	this.parent = parent
	this.id = randomUUID()

	// The wizard dialog...
	this.dialog = new Dialog(this.id, this.parent, false, "New Process")

	// Set the dialog position...
	var diagramElement = startEvent.getDiagramElement()
	var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	this.dialog.setPosition(x, y)

	this.content = this.dialog.content.appendElement({ "tag": "div", "class": "process_wizard_content" }).down()

	// That will contain the values ask by the user...
	this.dataView = new BpmnDataView(this.content, startEvent)

	this.dialog.ok.element.onclick = function (dataView, process, dialog) {
		return function () {
			// I will close the dialogue first...
			dialog.close()

			// I will save the data view...
			dataView.save(function (process) {
				return function (itemAwareInstances) {
					// Here I will create the new process...
					server.workflowManager.startProcess(process.UUID, itemAwareInstances, [],
						// Success Callback
						function (result, caller) {
							/* Nothing here */
						},
						// Error Callback
						function () {/* Nothing here */ },
						{})
				}
			} (process))
			
		}
	} (this.dataView, startEvent.getParent(), this.dialog)

	return this
}