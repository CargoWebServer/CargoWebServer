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
				var entities = []
				this.values[data.UUID] = entities

				this.appendItemDefinition(table, data, itemDefinition, data.M_isCollection,
					// Append item callback 
					function (entities) {
						return function (value) {
							// append the entity
							entities.push(value)
						}
					} (entities),
					// Remove item callback
					function (entities) {
						return function (value) {
							// remove the entity.
							entities.pop(value)
						}
					} (entities))
			}
		}
	}

	this.dialog.ok.element.onclick = function (values, process, dialog, wizard) {
		return function () {

			// Here I will the itemAwareInstance...
			var index = 0
			var itemAwareInstances = []
			for (var dataId in values) {
				var data = []
				for (var i = 0; i < values[dataId].length; i++) {
					// serialyse the object...
					if (values[dataId][i].stringify != undefined) {
						var objStr = values[dataId][i].stringify()
						console.log(objStr)
						data.push(objStr)
					} else {
						data.push(values[dataId][i])
					}
				}

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

			// release the values
			wizard.values = {}

		}
	} (this.values, startEvent.getParent(), this.dialog, this)

	return this
}

/*
 * Create the interface to enter the data about an item definition.
 */
ProcessWizard.prototype.appendItemDefinition = function (parent, data, itemDefinition, isCollection, onSelect, onRemove) {

	// It must be at lest one row...
	var table = parent.appendElement({ "tag": "div", "style": "position: relative; width:100%;" }).down()
	
	// Here I will create a table.
	parent.appendElement({ "tag": "div", "style": "display: table;" }).down()
		.appendElement({ "tag": "div", "id": "labelDiv", "style": "display: table-cell; padding: 2px;" })
		.appendElement({ "tag": "i", "id": "editEntityBtn", "class": "new_item_definition_button fa fa-pencil-square-o", "style": "display: none;" })
		.appendElement({ "tag": "i", "id": "appendEntityBtn", "class": "new_item_definition_button fa fa-plus", "style": "display: none;" })
		.appendElement({ "tag": "div", "id": "valueDiv", "style": "display: table-cell;padding: 2px; width: 100%;" })

	var labelDiv = parent.getChildById("labelDiv")
	var valueDiv = parent.getChildById("valueDiv")

	labelDiv.element.innerHTML = data.M_id

	var appendEntityBtn = null
	if (isCollection) {
		appendEntityBtn = parent.getChildById("appendEntityBtn")
		appendEntityBtn.element.style.display = ""
	}

	var itemDefintionRef = ""
	if (isString(itemDefinition)) {
		itemDefintionRef = itemDefinition
	} else {
		itemDefintionRef = itemDefinition.UUID
	}

	var control = null

	// Here the data is a primitive xsd type...
	if (itemDefinition == "xsd:string") {
		if (!isCollection) {
			control = valueDiv.appendElement({ "tag": "textarea" }).down()
			labelDiv.element.style.verticalAlign = "top"
			control.element.onblur = function (onSelect, id) {
				return function () {
					// Set the text element.
					this.value.UUID = id
					onSelect(this.value)
				}
			} (onSelect, data.M_id)
		} else {
			// TODO write code for multiple input...
		}
	} else if (itemDefinition == "xsd:boolean") {

	} else if (itemDefinition == "xsd:int" || itemDefinition == "xsd:integer") {

	} else if (itemDefinition == "xsd:byte") {

	} else if (itemDefinition == "xsd:long") {

	} else if (itemDefinition == "xsd:date") {

	} else if (itemDefinition == "xsd:double" || itemDefinition == "xsd:float") {

	} else {

		// Object reference here.
		if (!isCollection) {
			appendEntityBtn = parent.getChildById("editEntityBtn")
			appendEntityBtn.element.style.display = ""
		}

		data["set_M_itemSubjectRef_" + itemDefintionRef + "_ref"](function (parent, isCollection, onSelect, onRemove) {
			return function (itemDefinition) {
				if (itemDefinition.M_structureRef != undefined) {
					// If the itemfinition is a structure.
					if (itemDefinition.M_structureRef.indexOf(".") != -1) {
						server.entityManager.getEntityPrototypes(itemDefinition.M_structureRef.split(".")[0],
							function (result, caller) {

								// Get the ui's object.
								var itemPrototype = server.entityManager.entityPrototypes[caller.itemDefinition.M_structureRef]
								var isCollection = caller.isCollection
								var parent = caller.parent
								var typeName = itemPrototype.TypeName

								appendEntityBtn.element.onclick = function (valueDiv, itemPrototype, isCollection, onSelect, onRemove) {
									return function (evt) {
										// get the child by it's id.
										var id = itemPrototype.TypeName
										var control = valueDiv.getChildById(id + "_new")
										if (control == undefined) {
											control = valueDiv.appendElement({ "tag": "input", "id": id + "_new" }).down()
											if (!isCollection) {
												var parentElement
												if (valueDiv.element.firstChild != control.element) {
													valueDiv.removeAllChilds()
												}
											}

											control.element.readOnly = true
											control.element.style.cursor = "progress"
											control.element.style.width = "auto"

											server.entityManager.getObjectsByType(itemPrototype.TypeName, itemPrototype.TypeName.split(".")[0], "",
												// Progress...
												function () { },
												// Sucess...
												function (results, caller) {
													var control = caller.control
													var onSelect = caller.onSelect
													var onRemove = caller.onRemove

													control.element.readOnly = false
													control.element.style.cursor = "default"

													if (results.length > 0) {
														// get title display a readable name for the end user
														// or the first entity id.
														var lst = []
														var objMap = {}

														for (var i = 0; i < results.length; i++) {
															var result = results[i]
															var prototype = server.entityManager.entityPrototypes[result.TYPENAME]
															var titles = result.getTitles()
															for (var j = 0; j < titles.length; j++) {
																lst.push(titles[j])
																// Link the title with the object...
																objMap[titles[j]] = result
															}
														}

														// Now i will set it autocompletion list...
														attachAutoComplete(control, lst, true)
														control.element.onblur = control.element.onchange = function (objMap, onSelect, onRemove, control, valueDiv) {
															return function (evt) {
																var value = objMap[this.value]
																if (value != undefined) {
																	onSelect(value)

																	// Compose the ref name...
																	var titles = value.getTitles()
																	var refName = ""
																	for (var j = 0; j < titles.length; j++) {
																		refName += titles[j]
																		if (j < titles.length - 1) {
																			refName += " "
																		}
																	}

																	// Remove the parent element.
																	valueDiv.removeElement(control)

																	// So now I will create a new lnk and display it...
																	var ln = valueDiv.appendElement({ "tag": "div", "class": "entities_btn_container" }).down()
																	var ref = ln.appendElement({ "tag": "div" }).down().appendElement({ "tag": "a", "href": "#", "title": value.TYPENAME, "innerHtml": refName }).down()
																	var deleteLnkButton = ln.appendElement({ "tag": "div", "class": "entities_btn" }).down().appendElement({ "tag": "i", "class": "fa fa-trash" }).down()

																	// todo display it in the propertie panel...
																	deleteLnkButton.element.onclick = function (valueDiv, ln, value, onRemove) {
																		return function () {
																			// I will call on remove with the value...
																			onRemove(value)
																			valueDiv.removeElement(ln)
																		}
																	} (valueDiv, ln, value, onRemove)
																}

															}
														} (objMap, onSelect, onRemove, control, valueDiv)

														control.element.focus()
														control.element.select();
													}
												},
												// error
												function () { }, { "control": control, "onSelect": onSelect, "onRemove": onRemove })
										}

										// Display it
										control.element.style.display = ""
									}
								} (valueDiv, itemPrototype, isCollection, caller.onSelect, caller.onRemove)

							},
							function (errMsg, caller) {

							},
							{ "parent": parent, "isCollection": isCollection, "itemDefinition": itemDefinition, "onSelect": onSelect, "onRemove": onRemove })
					}
				}
			}
		} (table, isCollection, onSelect, onRemove))
	}

	return table
}