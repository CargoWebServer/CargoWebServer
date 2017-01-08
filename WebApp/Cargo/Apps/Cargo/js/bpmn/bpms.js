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
			if (itemDefinition != undefined) {
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
	var row = table.appendElement({ "tag": "div", "style": "display: table-row;" }).down()

	// Append the first value here...
	var labelTxt = data.M_name

	if (isCollection && this.values[data.UUID] == undefined) {
		this.values[data.UUID] = []

		// The append button for that collection...
		var newRowButton = parent.appendElement({ "tag": "div", "style": "position: absolute; top:0px; left:2px;" }).down()
		newRowButton.appendElement({ "tag": "i", "class": "fa fa-plus new_item_definition_button", "style": "font-size: .8em;" }).down()

		newRowButton.element.onclick = function (wizard, parent, data, itemDefintion) {
			return function () {
				wizard.appendItemDefinition(parent, data, itemDefinition, isCollection)
			}
		} (this, parent, data, itemDefinition, isCollection)
	}

	var index = -1
	var label = row.appendElement({ "tag": "div", "class": "process_wizard_content_label", "style": "display: table-cell; width:35%; vertical-align: top;", "innerHtml": labelTxt }).down()
	var value = row.appendElement({ "tag": "div", "class": "process_wizard_content_value", "style": "display: table-cell; width:65%;" }).down()

	// if the item is a collection...
	if (isCollection) {
		value.element.style.paddingRight = "6px"
		this.values[data.UUID].push(undefined)
		index = this.values[data.UUID].length - 1
		labelTxt += "[" + (index + 1) + "]"
		label.element.innerHTML = labelTxt
		// I will also append a delete button for the row...
		var deleteButton = row.appendElement({ "tag": "div", "class": "", "style": "display: table-cell;padding-right:6px;" }).down()
			.appendElement({ "tag": "i", "class": "fa fa-minus wizard_row_delete_btn", "style": "font-size: .8em;" }).down()

		deleteButton.element.onclick = function (table, values, uuid, index) {
			return function () {
				var values_ = []
				for (var i = 0; i < values[uuid].length; i++) {
					if (i != index) {
						values_.push(values[uuid][i])
					}
				}
				values[uuid] = values_
				table.element.parentNode.removeChild(table.element)

			}
		} (table, this.values, data.UUID, index)
	}

	if (itemDefinition.M_structureRef != undefined) {
		// The cargo entities...
		if (itemDefinition.M_structureRef == "CargoEntities.User" || itemDefinition.M_structureRef == "CargoEntities.Group") {
			// I will create the autocomplete list...
			var input = value.appendElement({ "tag": "input", "style": "" }).down()
			server.entityManager.getObjectsByType(itemDefinition.M_structureRef, "CargoEntities", null,
				// Progress...
				function () {

				},
				// success...
				function (results, caller) {
					var lst = []
					var objMap = {}
					var input = caller.input
					var output = caller.output
					var uuid = caller.uuid
					var values = caller.values
					var index = caller.index

					for (var i = 0; i < results.length; i++) {
						var result = results[i]
						var id = ""
						if (result.TYPENAME == "CargoEntities.User") {
							// User...
							id = result.M_firstName
							if (result.M_middle.length > 0) {
								id += " " + result.M_middle
							}
							id += " " + result.M_lastName

						} else {
							// Group...
							id = result.M_name
						}

						lst.push(id)
						objMap[id] = result
					}

					attachAutoComplete(input, lst)

					// Here I will implement the keyup listener. When 
					// the value is enter and the object is found then I will
					// append it in the list of result...
					input.element.onblur = input.element.onchange = function (objMap, values, uuid, index) {
						return function (evt) {
							var value = objMap[this.value]
							if (value != undefined) {
								if (index > -1) {
									values[uuid][index] = value
								} else {
									values[uuid] = value
								}
							}
						}
					} (objMap, values, uuid, index)
				},
				// Error callback
				function () { },
				{ "input": input, "output": row, "values": this.values, "uuid": data.UUID, "index": index })

		} else if (itemDefinition.M_structureRef == "CargoEntities.File") {

		} else {
			// Here it's a late binding structure...
			console.log("definition found with item: ")
			var input = value.appendElement({ "tag": "input", "style": "" }).down()
			var objMap = {}

			server.entityManager.getObjectsByType(itemDefinition.M_structureRef, itemDefinition.M_structureRef.split(".")[0], null,
				// Progress...
				function () {

				},
				// success...
				function (results, caller) {
					var input = caller.input
					var values = caller.values
					var objMap = caller.objMap
					var uuid = caller.uuid
					var index = caller.index

					var lst = []
					for (var i = 0; i < results.length; i++) {
						var result = results[i]
						// See if the object must contain M_id...
						objMap[result.M_id] = result
						lst.push(result.M_id)
					}
					attachAutoComplete(input, lst)

					input.element.onblur = input.element.onchange = function (objMap, values, uuid, index) {
						return function (evt) {
							var value = objMap[this.value]
							if (value != undefined) {
								if (index > -1) {
									values[uuid][index] = value
								} else {
									values[uuid] = value
								}
							}
						}
					} (objMap, values, uuid, index)

				}, { "input": input, "objMap": objMap, "values": this.values, "uuid": data.UUID, "index": index })


		}
	} else {

		// Here the data is a primitive xsd type...
		if (itemDefinition == "xsd:string") {
			var input = value.appendElement({ "tag": "textArea", "style": "width: 100%" }).down()
			input.element.onblur = function (values, uuid, index) {
				return function () {
					if (index > -1) {
						values[uuid][index] = this.value
					} else {
						values[uuid] = this.value
					}
				}
			} (this.values, data.UUID, index)
		} else if (itemDefinition == "xsd:boolean") {

		} else if (itemDefinition == "xsd:int" || itemDefinition == "xsd:integer") {

		} else if (itemDefinition == "xsd:byte") {

		} else if (itemDefinition == "xsd:long") {

		} else if (itemDefinition == "xsd:date") {

		} else if (itemDefinition == "xsd:double" || itemDefinition == "xsd:float") {

		}
	}

	return table
}