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
	this.values = {}

	// Now initialyse the dataouput...
	if (startEvent.M_dataOutput != undefined) {
		for (var i = 0; i < startEvent.M_dataOutput.length; i++) {
			// The data to input.
			var data = startEvent.M_dataOutput[i]

			// The item definition will be use to know what kind of
			// data must be enter here by the user.
			var itemDefinition = data.M_itemSubjectRef

			//console.log(itemDefinition)
			if (itemDefinition != undefined && itemDefinition != "") {
				// append item...
				var table = this.content.appendElement({ "tag": "div", "style": "display: table; position: relative; width:100%;" }).down()
				this.appendItemDefinition(table, data, itemDefinition, data.M_isCollection)
			}
		}
	}

	this.dialog.ok.element.onclick = function (values, process, dialog) {
		return function () {

			// Here I will the itemAwareInstance...
			var index = 0
			var itemAwareInstances = []
			for (var dataId in values) {
				var data = values[dataId]
				// The array of item aware instances.
				if (data != undefined) {
					server.workflowManager.newItemAwareElementInstance(dataId, data,
						function (result, caller) {
							caller.itemAwareInstances.push(result)

							if (caller.createProcess == true) {
								// Todo set the event definiton data and event properties...
								server.workflowManager.startProcess(caller.process.UUID, caller.itemAwareInstances, [],
									// Success Callback
									function (result, dialog) {
										dialog.close()
									},
									// Error Callback
									function () {/* Nothing here */ },
									caller.dialog)
							}

						},
						function () { /* Nothing here */ },
						{ "dialog": dialog, "process": process, "itemAwareInstances": itemAwareInstances, "createProcess": index == Object.keys(values).length - 1 })
				}
				index++

			}

		}
	} (this.values, startEvent.getParent(), this.dialog)

	return this
}

/*
 * Create the interface to enter the data about an item definition.
 */
ProcessWizard.prototype.appendItemDefinition = function (parent, data, itemDefinition, isCollection) {

	// It must be at lest one row...
	var table = parent.appendElement({ "tag": "div", "style": "display: table; position: relative; width:100%;" }).down()
	data["set_M_itemSubjectRef_" + itemDefinition + "_ref"](function (parent, isCollection) {
		return function (itemDefinition) {
			if (itemDefinition.M_structureRef != undefined) {
				if (itemDefinition.M_structureRef.indexOf(".") != -1) {
					server.entityManager.getEntityPrototype(itemDefinition.M_structureRef, itemDefinition.M_structureRef.split(".")[0],
						function (result, caller) {
							// Here I will display the panel.
							

						},
						function (errMsg, caller) {

						},
						{ "parent": parent, "isCollection": isCollection })
				}
			} else {
				// Here the data is a primitive xsd type...
				if (itemDefinition == "xsd:string") {

				} else if (itemDefinition == "xsd:boolean") {

				} else if (itemDefinition == "xsd:int" || itemDefinition == "xsd:integer") {

				} else if (itemDefinition == "xsd:byte") {

				} else if (itemDefinition == "xsd:long") {

				} else if (itemDefinition == "xsd:date") {

				} else if (itemDefinition == "xsd:double" || itemDefinition == "xsd:float") {

				}
			}
		}
	} (table, isCollection))

	return table
}