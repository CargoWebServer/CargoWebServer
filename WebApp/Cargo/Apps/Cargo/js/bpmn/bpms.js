/*
 * The list of instances.
 */
var InstanceListView = function (instancesLst, parent, svgDiagram) {

	this.parent = parent
	this.instances = {}
	this.instancesDiv = {}

	/* This is the link to the related diagram view **/
	this.svgDiagram = svgDiagram

	// The process instance header...
	for (var i = 0; i < instancesLst.length; i++) {
		var instance = instancesLst[i]
		this.instances[instance.M_id] = instance
		this.appendProcessInstance(instance)
	}

	return this
}


/*
 * Append a new instance to the list.
 */
InstanceListView.prototype.appendProcessInstance = function (instance) {
	var processInstanceDiv = this.parent.appendElement({ "tag": "div", "class": "process_instance_header", "innerHtml": "Process instance " }).down()

	// Now I will iterate over the flow node instance.
	for (var i = 0; i < instance.M_flowNodeInstances.length; i++) {

		// Here I will get all necessary informations...
		var flowNodeInstance = instance.M_flowNodeInstances[i]
		var bpmnElement = this.svgDiagram.bpmnElements[flowNodeInstance.M_bpmnElementId]
		var svgElement = this.svgDiagram.svgElements[flowNodeInstance.M_bpmnElementId]
		var instanceDiv = processInstanceDiv.appendElement({ "tag": "div", "id": flowNodeInstance.M_bpmnElementId, "class": "flowNode_instance", "innerHtml": bpmnElement.M_id }).down()

		if (flowNodeInstance.M_flowNodeType == 1 /* Abstract Task **/) {

		} else if (flowNodeInstance.M_flowNodeType == 2 /* Service Task **/) {

		} else if (flowNodeInstance.M_flowNodeType == 3 /* User Task **/) {

		} else if (flowNodeInstance.M_flowNodeType == 4 /* Manual Task **/) {

		} else if (flowNodeInstance.M_flowNodeType == 5 /* Business Rule Task **/) {

		} else if (flowNodeInstance.M_flowNodeType == 6 /* Script Task **/) {

		} else if (flowNodeInstance.M_flowNodeType == 7 /* Embedded Subprocess **/) {

		} else if (flowNodeInstance.M_flowNodeType == 8 /* Event Subprocess **/) {

		} else if (flowNodeInstance.M_flowNodeType == 9 /* AdHoc Subprocess **/) {

		} else if (flowNodeInstance.M_flowNodeType == 10 /* Transaction **/) {

		} else if (flowNodeInstance.M_flowNodeType == 11 /* Call Activity **/) {

		} else if (flowNodeInstance.M_flowNodeType == 12 /* Parallel Gateway **/) {

		} else if (flowNodeInstance.M_flowNodeType == 13 /* Exclusive Gateway **/) {

		} else if (flowNodeInstance.M_flowNodeType == 14 /* Inclusive Gateway **/) {

		} else if (flowNodeInstance.M_flowNodeType == 15 /* Event Based Gateway **/) {

		} else if (flowNodeInstance.M_flowNodeType == 16 /* Complex Gateway **/) {

		} else if (flowNodeInstance.M_flowNodeType == 17 /* Start Event **/) {

		} else if (flowNodeInstance.M_flowNodeType == 18 /* Intermediate CatchEvent **/) {

		} else if (flowNodeInstance.M_flowNodeType == 19 /* Boundary Event **/) {

		} else if (flowNodeInstance.M_flowNodeType == 19 /* End Event **/) {

		} else if (flowNodeInstance.M_flowNodeType == 19 /* Intermediate Throw Event **/) {

		}

		instanceDiv.element.onclick = function (svgElement) {
			return function () {
				// Here I will set the class...

			}
		} (svgElement)

		this.instancesDiv[flowNodeInstance.M_bpmnElementId] = instanceDiv
	}
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