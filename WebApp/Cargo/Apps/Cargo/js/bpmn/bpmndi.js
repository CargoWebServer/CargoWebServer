////////////////////////////////////////////////////////////////////////////////////////////////////////
// Graphical representation of Diagram.
////////////////////////////////////////////////////////////////////////////////////////////////////////
var SvgDiagram = function (parent, bpmnDiagram) {
	this.pathMap = new PathMap()
	/* The parent element**/
	this.parent = parent

	/* The diagram element indexed by their bpmn element id **/
	this.svgElements = {}

	/* Keep the reference to the bpmn element to... **/
	this.bpmnElements = {}

	/* The instances list view **/
	this.instanceListView = null

	/* Contain the data **/
	this.bpmnDiagram = bpmnDiagram

	/* Now I will create the svg canvas. **/
	this.canvas = new SVG_Canvas(this.parent, bpmnDiagram.M_id, "bpmndi_canvas", parent.element.clientWidth, parent.element.clientHeight);

	/* Now the svg element itself **/
	this.plane = new SVG_Group(this.canvas, bpmnDiagram.M_BPMNPlane.M_id, "bpmndi_plane", 0, 0, 0, 0)

	// Here I will connect the resize event of the parent with the parent.
	parent.element.onresize = function (parent, svg) {
		return function () {
			svg.setSvgAttribute("width", parent.element.clientWidth)
			svg.setSvgAttribute("height", parent.element.clientHeight)
		}
	} (this.parent, this.canvas)

	// The scale factor of the diagram...
	this.canvas.scaleFactor = 1

	// That function is use to refresh the size of the svg canva.
	this.canvas.resize = function (canvas) {
		return function (widht, height) {
			canvas.setSvgAttribute("width", widht)
			canvas.setSvgAttribute("height", height)

			var viewBox = canvas.element.viewBox
			// TODO the pan here

			// Set the new width and height with respect to the scaleFactor...
			viewBox.baseVal.width = widht * (1 / canvas.scaleFactor)
			viewBox.baseVal.height = height * (1 / canvas.scaleFactor)
		}
	} (this.canvas)

	// Now the zoom in and out ...
	this.canvas.element.addEventListener("wheel",
		function (canvas) {
			return function (evt) {
				if (evt.deltaY > 0) {
					canvas.scaleFactor -= .1
				} else {
					canvas.scaleFactor += .1
				}
				canvas.initWorkspace()
			}
		} (this.canvas));

	return this
}

/*
 * Recursively draw diagram element.
 */
SvgDiagram.prototype.init = function () {
	var diagramElements = this.bpmnDiagram.M_BPMNPlane.M_DiagramElement
	for (var index = 0; index < diagramElements.length; index++) {
		// Here I will call the same function on the next element.
		var diagramElement = diagramElements[index]
		// Here I will create a function to get the diagram element...
		var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
		bpmnElement.getDiagramElement = function (diagramElement) {
			return function () {
				return diagramElement
			}
		} (diagramElement)
	}
}

/*
 * Recursively draw diagram element.
 */
SvgDiagram.prototype.drawDiagramElements = function () {
	var diagramElements = this.bpmnDiagram.M_BPMNPlane.M_DiagramElement
	for (var index = 0; index < diagramElements.length; index++) {
		// Here I will call the same function on the next element.
		var diagramElement = diagramElements[index]
		// Here I will create a function to get the diagram element...
		this.drawDiagramElement(diagramElement)
	}
}


SvgDiagram.prototype.drawDiagramElement = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var typeName = bpmnElement.TYPENAME
	var svgElement = null

	if (typeName == "BPMN20.ScriptTask") {
		svgElement = this.drawScriptTask(diagramElement)
	} else if (typeName == "BPMN20.Task") {
		svgElement = this.drawTask(diagramElement, " bpmndi_task")
	} else if (typeName == "BPMN20.UserTask") {
		svgElement = this.drawUserTask(diagramElement)
	} else if (typeName == "BPMN20.ReceiveTask") {
		svgElement = this.drawReceiveTask(diagramElement)
	} else if (typeName == "BPMN20.SendTask") {
		svgElement = this.drawSendTask(diagramElement)
	} else if (typeName == "BPMN20.ServiceTask") {
		svgElement = this.drawServiceTask(diagramElement)
	} else if (typeName == "BPMN20.StartEvent") {
		svgElement = this.drawStartEvent(diagramElement)
	} else if (typeName == "BPMN20.EndEvent") {
		svgElement = this.drawEndEvent(diagramElement)
	} else if (typeName == "BPMN20.IntermediateCatchEvent" || typeName == "BPMN20.IntermediateThrowEvent") {
		svgElement = this.drawIntermediateEvent(diagramElement)
	} else if (typeName == "BPMN20.EventBasedGateway") {
		svgElement = this.drawEventBasedGateway(diagramElement)
	} else if (typeName == "BPMN20.InclusiveGateway") {
		svgElement = this.drawInclusiveGateway(diagramElement)
	} else if (typeName == "BPMN20.ExclusiveGateway") {
		svgElement = this.drawExclusiveGateway(diagramElement)
	} else if (typeName == "BPMN20.ParallelGateway") {
		svgElement = this.drawParallelGateway(diagramElement)
	} else if (typeName == "BPMN20.ComplexGateway") {
		svgElement = this.drawComplexGateway(diagramElement)
	} else if (typeName == "BPMN20.DataObjectReference") {
		svgElement = this.drawDataObjectReference(diagramElement)
	} else if (typeName == "BPMN20.DataStoreReference") {
		svgElement = this.drawDataStoreReference(diagramElement)
	} else if (typeName == "BPMN20.SequenceFlow") {
		svgElement = this.drawSequenceFlow(diagramElement)
	} else if (typeName == "BPMN20.DataInputAssociation") {
		svgElement = this.drawDataInputAssociation(diagramElement)
	} else if (typeName == "BPMN20.DataInput") {
		svgElement = this.drawDataInput(diagramElement)
	} else if (typeName == "BPMN20.DataOutputAssociation") {
		svgElement = this.drawDataOutputAssociation(diagramElement)
	} else if (typeName == "BPMN20.DataOutput") {
		svgElement = this.drawDataOutput(diagramElement)
	} else if (typeName == "BPMN20.TextAnnotation") {
		svgElement = this.drawTextAnnotation(diagramElement)
	} else if (typeName == "BPMN20.Participant" || typeName == "BPMN20.Lane") {
		svgElement = this.drawParticipantLane(diagramElement)
	} else if (typeName == "BPMN20.MessageFlow") {
		svgElement = this.drawMessageFlow(diagramElement)
	} else if (typeName == "BPMN20.Association") {
		svgElement = this.drawAssociation(diagramElement)
	} else if (typeName == "BPMN20.BoundaryEvent") {
		svgElement = this.drawBoundaryEvent(diagramElement)
	} else if (typeName == "BPMN20.CallActivity" || typeName == "BPMN20.SubProcess" || typeName == "BPMN20.AdHocSubProcess") {
		svgElement = this.drawCallableActivity(diagramElement)
	} else if (typeName == "BPMN20.Group") {
		svgElement = this.drawCallableActivity(diagramElement)
	} else if (typeName == "BPMN20.Process") {
		svgElement = this.plane
	} else {
		console.log(diagramElement + " has undefined typeName")
	}

	// Keep the reference of the svg element...
	if (svgElement != null) {
		diagramElement.getSvgElement = function (svgElement) {
			return function () {
				return svgElement
			}
		} (svgElement)
	}

	// Now if the element contain a label...
	// Pool text are svg text element, not html tag...
	if (diagramElement.M_BPMNLabel != undefined && typeName != "BPMN20.Participant" && typeName != "BPMN20.Lane") {
		if (diagramElement.M_BPMNLabel.M_Bounds != null && diagramElement.M_BPMNLabel.M_Bounds != "") {
			if (bpmnElement.M_dataObjectRef != undefined) {
				bpmnElement["set_M_dataObjectRef_" + bpmnElement.M_dataObjectRef + "_ref"](function (diagram, diagramElement) {
					return function (ref) {
						var label = diagram.drawText(ref.M_name, diagramElement, diagramElement.M_BPMNLabel.M_Bounds)
						diagramElement.getLabel = function (label) {
							return function () {
								return label
							}
						} (label)
					}
				} (this, diagramElement))
			} else if (bpmnElement.M_dataStoreRef != undefined) {
				bpmnElement["set_M_dataStoreRef_" + bpmnElement.M_dataStoreRef + "_ref"](function (diagram, diagramElement) {
					return function (ref) {
						var label = diagram.drawText(ref.M_name, diagramElement, diagramElement.M_BPMNLabel.M_Bounds)
						diagramElement.getLabel = function (label) {
							return function () {
								return label
							}
						} (label)
					}
				} (this, diagramElement))
			} else {
				var label = this.drawText(bpmnElement.M_name, diagramElement, diagramElement.M_BPMNLabel.M_Bounds)
				diagramElement.getLabel = function (label) {
					return function () {
						return label
					}
				} (label)
			}

		}
	}

	// Set the expended property
	if (typeName == "BPMN20.CallActivity" || typeName == "BPMN20.SubProcess" || typeName == "BPMN20.AdHocSubProcess") {
		if (diagramElement.M_isExpanded) {
			var label = diagramElement.getLabel()
			label.element.style.display = "none"
			label.element.style.zIndex = 0
		} else {
			svgElement.element.firstChild.style.fill = "white"
		}
	}

}

/*
 * So here i will set the the collapsed element...
 * must be call after init and drawElements...
 */
SvgDiagram.prototype.setCollapsedElement = function () {

	var diagramElements = this.bpmnDiagram.M_BPMNPlane.M_DiagramElement
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	for (var index = 0; index < diagramElements.length; index++) {
		var diagramElement = diagramElements[index]
		if (bpmnElement.TYPENAME == "BPMN20.SubProcess" || bpmnElement.TYPENAME == "BPMN20.CallActivity") {

			// The expension button...
			var expandButton = diagramElement.getSvgElement().getChildById("bpmndi_expand_marker")
			var expandMarker = diagramElement.getSvgElement().getChildById("marker_path")

			expandMarker.element.onclick = expandButton.element.onclick = function (diagramElement) {
				return function () {
					var diagramElements = this.bpmnDiagram.M_BPMNPlane.M_DiagramElement
					var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

					if (bpmnElement.TYPENAME == "BPMN20.SubProcess") {
						if (diagramElement.M_isExpanded == true) {
							// The first child of the group must be a rect element.
							diagramElement.getSvgElement().element.firstChild.style.fill = "white"
							if (bpmnElement.M_flowElement != undefined) {
								for (var i = 0; i < bpmnElement.M_flowElement.length; i++) {
									var flowDiagramElement = bpmnElement.M_flowElement[i].getDiagramElement()
									var svgElement = flowDiagramElement.getSvgElement()
									if (svgElement != undefined) {
										svgElement.element.style.display = "none"
									}
									if (flowDiagramElement.getLabel != undefined) {
										flowDiagramElement.getLabel().element.style.display = "none"
									}
								}
							}
							diagramElement.M_isExpanded = false
							diagramElement.getLabel().element.style.display = ""
						} else {

							diagramElement.getSvgElement().element.firstChild.style.fill = "none"
							if (bpmnElement.M_flowElement != undefined) {
								for (var i = 0; i < bpmnElement.M_flowElement.length; i++) {
									var flowDiagramElement = bpmnElement.M_flowElement[i].getDiagramElement()
									var svgElement = flowDiagramElement.getSvgElement()
									if (flowDiagramElement.getLabel != undefined) {
										flowDiagramElement.getLabel().element.style.display = ""
									}
									if (svgElement != undefined) {
										svgElement.element.style.display = ""
									}
								}
							}
							diagramElement.getLabel().element.style.display = "none"
							diagramElement.M_isExpanded = true

						}
					} else if (bpmnElement.TYPENAME == "BPMN20.CallActivity") {
						if (diagramElement.M_isExpanded == true) {
							// The first child of the group must be a rect element.
							diagramElement.getSvgElement().element.firstChild.style.fill = "white"
							var calledElement = bpmnElement.M_calledElement
							if (calledElement != null) {
								if (calledElement.M_flowElement != undefined) {
									if (calledElement.M_flowElement != undefined) {
										for (var i = 0; i < calledElement.M_flowElement.length; i++) {
											var flowDiagramElement = calledElement.M_flowElement[i].getDiagramElement()
											//flowDiagramElement.element.style.display = "none"
											var svgElement = flowDiagramElement.getSvgElement()
											if (flowDiagramElement.getLabel != undefined) {
												flowDiagramElement.getLabel().element.style.display = "none"
											}
											if (svgElement != undefined) {
												svgElement.element.style.display = "none"
											}
										}
									}
								}
								diagramElement.getLabel().element.style.display = ""
								diagramElement.M_isExpanded = false
							}
						} else {
							diagramElement.getSvgElement().element.firstChild.style.fill = "none"
							var calledElement = bpmnElement.M_calledElement
							if (calledElement == null) {
								calledElement = diagramElement
							}
							if (calledElement != null) {
								if (calledElement.M_flowElement != undefined) {
									if (calledElement.M_flowElement != undefined) {
										for (var i = 0; i < calledElement.M_flowElement.length; i++) {
											var flowDiagramElement = calledElement.M_flowElement[i].getDiagramElement()
											if (flowDiagramElement.getLabel != undefined) {
												flowDiagramElement.getLabel().element.style.display = ""
											}
											var svgElement = flowDiagramElement.getSvgElement()
											if (svgElement != undefined) {
												svgElement.element.style.display = ""
											}
										}
									}
								}
								diagramElement.getLabel().element.style.display = "none"
								diagramElement.M_isExpanded = true
							}
						}
					}
				}
			} (diagramElement)

			expandButton.element.onmouseout = function () {
				this.style.cursor = "default"
			}


			expandMarker.element.onmouseover = expandButton.element.onmouseover = function () {
				this.style.cursor = "pointer"
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Task elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawTask = function (diagramElement, className) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)

	var shape = new SVG_Rectangle(group, diagramElement.M_id, "bpmndi_shape " + className,
		0, 0,
		diagramElement.M_Bounds.M_width,
		diagramElement.M_Bounds.M_height,
		5, 5)

	// Here I will set the name with the id of the bpmn element...
	shape.setSvgAttribute("name", bpmnElement.M_id)

	this.svgElements[diagramElement.M_bpmnElement.UUID] = shape

	if (diagramElement.M_BPMNLabel == undefined) {
		var label = this.drawText(bpmnElement.M_name, diagramElement, diagramElement.M_Bounds)
		label.element.style.minHeight = "70px"
		diagramElement.getLabel = function (label) {
			return function () {
				return label
			}
		} (label)
	}

	var position;
	if (bpmnElement.TYPENAME == "BPMN20.SubProcess") {
		position = {
			seq: -21,
			parallel: -22,
			compensation: -42,
			loop: -18,
			adhoc: 10
		};
	} else {
		position = {
			seq: -3,
			parallel: -6,
			compensation: -27,
			loop: 0,
			adhoc: 10
		};
	}

	// Now the symbol...
	// Is for compensation...
	if (bpmnElement.M_isForCompensation) {
		var markerMath = this.pathMap.getScaledPath('MARKER_COMPENSATION', {
			xScaleFactor: 1,
			yScaleFactor: 1,
			containerWidth: bpmnElement.M_width,
			containerHeight: bpmnElement.M_height,
			position: {
				mx: ((bpmnElement.M_width / 2 + position.compensation) / bpmnElement.M_width),
				my: (bpmnElement.M_height - 13) / bpmnElement.M_height
			}
		});
		new SVG_Path(group, diagramElement.M_id + "_loop_characteristic", "marker", [markerPath])
	}

	// It's ad hoc subprocess...
	if (bpmnElement.TYPENAME == "BPMN20.AdHocSubProcess") {
		var markerPath = this.pathMap.getScaledPath('MARKER_ADHOC', {
			xScaleFactor: 1,
			yScaleFactor: 1,
			containerWidth: diagramElement.M_Bounds.M_width,
			containerHeight: diagramElement.M_Bounds.M_height,
			position: {
				mx: ((diagramElement.M_Bounds.M_width / 2 + position.adhoc) / diagramElement.M_Bounds.M_width),
				my: (diagramElement.M_Bounds.M_height - 15) / diagramElement.M_Bounds.M_height
			}
		});
		new SVG_Path(group, diagramElement.M_id + "_loop_characteristic", "marker", [markerPath])
	}

	// The loop characteristics.
	if (bpmnElement.M_loopCharacteristics != undefined) {
		var loopCharacteristics = bpmnElement.M_loopCharacteristics
		if (loopCharacteristics.M_isSequential == undefined) {
			var markerPath = this.pathMap.getScaledPath('MARKER_LOOP', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: ((diagramElement.M_Bounds.M_width / 2 + position.loop) / diagramElement.M_Bounds.M_width),
					my: (diagramElement.M_Bounds.M_height - 7) / diagramElement.M_Bounds.M_height
				}
			});

			new SVG_Path(group, diagramElement.M_id + "_loop_characteristic", "marker", [markerPath])

		}

		if (loopCharacteristics.M_isSequential == true) {
			var markerPath = this.pathMap.getScaledPath('MARKER_PARALLEL', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: ((diagramElement.M_Bounds.M_width / 2 + position.parallel) / diagramElement.M_Bounds.M_width),
					my: (diagramElement.M_Bounds.M_height - 20) / diagramElement.M_Bounds.M_height
				}
			});
			new SVG_Path(group, diagramElement.M_id + "_loop_characteristic", "marker", [markerPath])
		}

		if (loopCharacteristics.M_isSequential == false) {
			var markerPath = this.pathMap.getScaledPath('MARKER_SEQUENTIAL', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: ((diagramElement.M_Bounds.M_width / 2 + position.seq) / diagramElement.M_Bounds.M_width),
					my: (diagramElement.M_Bounds.M_height - 19) / diagramElement.M_Bounds.M_height
				}
			});
			new SVG_Path(group, diagramElement.M_id + "_loop_characteristic", "marker", [markerPath])
		}
	}

	// Set the position of the task...
	group.setSvgAttribute("transform", "translate(" + diagramElement.M_Bounds.M_x + "," + diagramElement.M_Bounds.M_y + ")")

	return group
}

SvgDiagram.prototype.drawUserTask = function (diagramElement) {
	// Here this is a rectange...
	var group = this.drawTask(diagramElement, "bpmndi_task bpmndi_user_task")
	var userIco = new SVG_Image(group, "user_ico", "", "/Cargo/svg/user-ico.svg", 3, 3, 20, 20)
	return group
}

SvgDiagram.prototype.drawScriptTask = function (diagramElement) {
	// Here this is a rectange...
	var group = this.drawTask(diagramElement, "bpmndi_task bpmndi_script_task")
	var scriptIco = new SVG_Image(group, "script_ico", "", "/Cargo/svg/script-ico.svg", 3, 3, 20, 20)

	return group
}

SvgDiagram.prototype.drawReceiveTask = function (diagramElement) {
	var group = this.drawTask(diagramElement, "bpmndi_task bpmndi_receive_task")
    pathStr = this.pathMap.getScaledPath('TASK_TYPE_SEND', {
		xScaleFactor: 0.9,
		yScaleFactor: 0.9,
		containerWidth: 21,
		containerHeight: 14,
		position: {
            mx: 0.3,
            my: 0.4
		}
	});

	var enveloppePath = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_receive_enveloppe", [pathStr])

	enveloppePath.setSvgAttribute("fill", 'black')
	enveloppePath.setSvgAttribute("stroke", 'white')

	return group
}

SvgDiagram.prototype.drawSendTask = function (diagramElement) {
	var group = this.drawTask(diagramElement, "bpmndi_task bpmndi_send_task")
	var pathStr = this.pathMap.getScaledPath('TASK_TYPE_SEND', {
        xScaleFactor: 1,
        yScaleFactor: 1,
        containerWidth: 21,
        containerHeight: 14,
        position: {
			mx: 0.285,
			my: 0.357
        }
	});

	var enveloppePath = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_send_enveloppe", [pathStr])
	return group
}

SvgDiagram.prototype.drawServiceTask = function (diagramElement) {
	var group = this.drawTask(diagramElement, "bpmndi_task bpmndi_service_task")
	var serviceIco = new SVG_Image(group, "service_ico", "", "/Cargo/svg/service-ico.svg", 3, 3, 32, 32)
	return group
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Lane elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawParticipantLane = function (diagramElement) {
	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	// The externat shape...
	var shape = new SVG_Rectangle(group, diagramElement.M_id, "bpmndi_participant",
		0, 0,
		diagramElement.M_Bounds.M_width,
		diagramElement.M_Bounds.M_height,
		0, 0)

	// Here I will set the name with the id of the bpmn element...
	shape.setSvgAttribute("name", bpmnElement.M_id)

	// Only participant has header box...
	var headerHeight = 30
	var header = null
	if (bpmnElement.TYPENAME == "BPMN20.Participant") {
		if (diagramElement.M_isHorizontal == true) {
			header = new SVG_Rectangle(group, diagramElement.M_id, "bpmndi_participant_header",
				0, 0,
				headerHeight,
				diagramElement.M_Bounds.M_height,
				0, 0)
		} else {
			header = new SVG_Rectangle(group, diagramElement.M_id, "bpmndi_participant_header",
				0, 0,
				diagramElement.M_Bounds.M_width,
				headerHeight,
				0, 0)
		}

		// Now the participant multiplicity...
		if (bpmnElement.M_participantMultiplicity != null) {
			var markerPath = this.pathMap.getScaledPath('MARKER_PARALLEL', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: ((/*element.width*/ diagramElement.M_Bounds.M_width / 2) / diagramElement.M_Bounds.M_width),
					my: (/*element.height*/diagramElement.M_Bounds.M_height - 15) / diagramElement.M_Bounds.M_height
				}
			});
			new SVG_Path(group, "participant_multiplicity_symbol", "participant_multiplicity_symbol", [markerPath])
		}
	}

	// The text header...
	var title = new SVG_Text(group, diagramElement.M_id + "_title", "bpmndi_participant_header_title", bpmnElement.M_name, 100, 100);
	if (diagramElement.M_isHorizontal == true) {
		title.setSvgAttribute("transform", "rotate(-90, 0,0)")
		title.setSvgAttribute("x", - 1 * diagramElement.M_Bounds.M_height / 2)
	} else {
		title.setSvgAttribute("x", diagramElement.M_Bounds.M_width / 2)
	}

	// y depand of the header heigth only...
	title.setSvgAttribute("y", headerHeight / 2)
	this.svgElements[bpmnElement.UUID] = shape
	group.setSvgAttribute("transform", "translate(" + diagramElement.M_Bounds.M_x + "," + diagramElement.M_Bounds.M_y + ")")


	return group
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Text elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawTextAnnotation = function (diagramElement) {
	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	var pathStr = "M 0 0 L 10 0 M 0 " + diagramElement.M_Bounds.M_height + " L 0 0 M 10 " + diagramElement.M_Bounds.M_height + " L 0 " + diagramElement.M_Bounds.M_height
	var dtextAnnotationPath = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_text_annotation", [pathStr])
	dtextAnnotationPath.setSvgAttribute("name", bpmnElement.M_id)

	group.setSvgAttribute("transform", "translate(" + diagramElement.M_Bounds.M_x + "," + diagramElement.M_Bounds.M_y + ")")
	this.svgElements[bpmnElement.UUID] = group

	textElement = this.drawText(bpmnElement.M_text.M_text, diagramElement, diagramElement.M_Bounds)
	textElement.element.className += " bpmndi_text_annotation"
	return group
}

SvgDiagram.prototype.drawText = function (text, diagramElement, bounds) {
	var htmlDiv = null

	htmlDiv = new SVG_ForeignObject(this.plane, diagramElement.M_id, "",
		bounds.M_x,
		bounds.M_y,
		bounds.M_width,
		bounds.M_height)

	// Here I'm in plain html...
	var textElement = new Element(null, { "tag": "div", "id": "task_txt", "class": "bpmndi_text_box" })
	textElement.element.style.width = diagramElement.M_width * 1.1 + "px"
	textElement.element.style.height = diagramElement.M_height + "px"
	textElement.appendElement({ "tag": "span", "innerHtml": text })

	// Here I will append the element...
	htmlDiv.appendElement(textElement)

	return textElement
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Definition elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawEventDefintion = function (diagramElement, eventGroup) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var eventDefinitionElements = bpmnElement.M_eventDefinition

	var group = new SVG_Group(eventGroup, diagramElement.M_id, "", 100, 0)

	for (var i = 0; i < eventDefinitionElements.length; i++) {
		var def = eventDefinitionElements[0]
		var pathStr = ""
		var evtDefPath = null
		var isThrowing = def.M_throwEventPtr != undefined

		if (def.TYPENAME == "BPMN20.TimerEventDefinition") {
			var r = diagramElement.M_Bounds.M_width / 2
			var shape = new SVG_Circle(group, diagramElement.M_id, "", r, r, r * .70)
			shape.setSvgAttribute("stroke", "#808080")
			shape.setSvgAttribute("fill", "none")
			shape.setSvgAttribute("stroke-width", "1.5px")
			pathStr = this.pathMap.getScaledPath('EVENT_TIMER_WH', {
				xScaleFactor: 0.75,
				yScaleFactor: 0.75,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.5,
					my: 0.5
				}
			})
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_timer_event_definition", [pathStr])

			for (var i = 0; i < 12; i++) {

				var linePathData = this.pathMap.getScaledPath('EVENT_TIMER_LINE', {
					xScaleFactor: 0.75,
					yScaleFactor: 0.75,
					containerWidth: diagramElement.M_Bounds.M_width,
					containerHeight: diagramElement.M_Bounds.M_height,
					position: {
						mx: 0.5,
						my: 0.5
					}
				});

				var width = diagramElement.M_Bounds.M_width / 2;
				var height = diagramElement.M_Bounds.M_height / 2;

				var clockMarker = new SVG_Path(group, def.M_id + "_path", "bpmndi_timer_event_clock_hour_marker", [linePathData])
				clockMarker.setSvgAttribute("transform", 'rotate(' + (i * 30) + ',' + height + ',' + width + ')')
			}

		} else if (def.TYPENAME == "BPMN20.MessageEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_MESSAGE', {
				xScaleFactor: 0.9,
				yScaleFactor: 0.9,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.235,
					my: 0.315
				}
			})

			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_message_event_definition", [pathStr])

		} else if (def.TYPENAME == "BPMN20.ConditionalEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_CONDITIONAL', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.5,
					my: 0.222
				}
			});
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_conditional_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.LinkEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_LINK', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.57,
					my: 0.263
				}
			});
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_link_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.SignalEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_SIGNAL', {
				xScaleFactor: 0.9,
				yScaleFactor: 0.9,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.5,
					my: 0.2
				}
			});
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_signal_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.ErrorEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_ERROR', {
				xScaleFactor: 1.1,
				yScaleFactor: 1.1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.2,
					my: 0.722
				}
			});
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_error_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.EscalationEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_ESCALATION', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.5,
					my: 0.555
				}
			})
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_escalation_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.TerminateEventDefinition") {
			var r = diagramElement.M_Bounds.M_width / 2
			var shape = new SVG_Circle(group, diagramElement.M_id, "", r, r, r * .75)
			shape.setSvgAttribute("stroke-width", "4px")
			shape.setSvgAttribute("fill", "black")
		} else if (def.TYPENAME == "BPMN20.CompensationEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_COMPENSATION', {
				xScaleFactor: 1,
				yScaleFactor: 1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.22,
					my: 0.5
				}
			});
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_compensation_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.CancelEventDefinition") {
			var pathStr = this.pathMap.getScaledPath('EVENT_CANCEL_45', {
				xScaleFactor: 1.0,
				yScaleFactor: 1.0,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.638,
					my: -0.055
				}
			});
			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_cancel_event_definition", [pathStr])
		} else if (def.TYPENAME == "BPMN20.ComplexBehaviorDefinition") {
			var pathStr = pathMap.getScaledPath('EVENT_MULTIPLE', {
				xScaleFactor: 1.1,
				yScaleFactor: 1.1,
				containerWidth: diagramElement.M_Bounds.M_width,
				containerHeight: diagramElement.M_Bounds.M_height,
				position: {
					mx: 0.222,
					my: 0.36
				}
			});

			evtDefPath = new SVG_Path(group, def.M_id + "_path", "bpmndi_event_definition bpmndi_complex_behavior_event_definition", [pathStr])

		}

		if (evtDefPath != null) {

			evtDefPath.setSvgAttribute("name", def.M_id)

			if (isThrowing) {
				evtDefPath.setSvgAttribute("fill", "black")
				evtDefPath.setSvgAttribute("stroke", "white")
			} else {
				evtDefPath.setSvgAttribute("fill", "white")
				evtDefPath.setSvgAttribute("stroke", "black")
			}
			evtDefPath.setSvgAttribute("x", diagramElement.M_Bounds.M_x)
			evtDefPath.setSvgAttribute("y", diagramElement.M_Bounds.M_y)
		}

	}

	return group
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Event elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawEvent = function (diagramElement) {
	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	// Here I will set the name with the id of the bpmn element...
	group.setSvgAttribute("transform", "translate(" + diagramElement.M_Bounds.M_x + "," + diagramElement.M_Bounds.M_y + ")")

	var r = diagramElement.M_Bounds.M_width / 2
	var className = "bpmndi_shape bpmndi_event"
	if (bpmnElement.TYPENAME == "BPMN20.BoundaryEvent") {
		className = "bpmndi_boundary_event"
	}
	var shape = new SVG_Circle(group, diagramElement.M_id, className, r, r, r)
	shape.setSvgAttribute("name", bpmnElement.M_id)

	this.svgElements[bpmnElement.UUID] = group

	return group
}

SvgDiagram.prototype.drawBoundaryEvent = function (diagramElement) {
	var group = this.drawEvent(diagramElement)
	var outerCircle = group.lastChild
	outerCircle.setSvgAttribute("stroke", "black")

	// Now the boundary event can be interupting or non interupting...
	var r = diagramElement.M_Bounds.M_width / 2
	var innerCircle = new SVG_Circle(group, diagramElement.M_id, "", r, r, r * .85)
	innerCircle.setSvgAttribute("stroke", "black")
	innerCircle.setSvgAttribute("fill", "none")

	this.drawEventDefintion(diagramElement, group)

	if (diagramElement.M_bpmnElement.M_cancelActivity == false) {
		innerCircle.setSvgAttribute("stroke-dasharray", "4 2")
		outerCircle.setSvgAttribute("stroke-dasharray", "4 2")
	}

	return group
}

SvgDiagram.prototype.drawStartEvent = function (diagramElement) {
	var group = this.drawEvent(diagramElement)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	if (bpmnElement.M_eventDefinition != undefined) {
		// Here I will draw the definition element...
		this.drawEventDefintion(diagramElement, group)
	} else {
		group.element.firstElementChild.className.baseVal += " bpmndi_startEvent"
	}

	// Here i will create the event...
	group.element.firstElementChild.onclick = function (parent, diagramElement) {
		return function (event) {

			// Get the menu position
			var x = event.offsetX - 5;     // Get the horizontal coordinate
			var y = event.offsetY - 5;     // Get the vertical coordinate

			var popup = parent.appendElement({ "tag": "div", "class": "popupMenu", "style": "top:" + y + "px; left:" + x + "px;" }).down()

			popup.element.onmouseout = function () {
				if (this.parentNode != null) {
					this.parentNode.removeChild(this)
				}
			}

			// Now I will append the menu element...
			startMenu = popup.appendElement({ "tag": "div", "class": "popupMenuItem", "innerHtml": "Start" }).down()
			startMenu.element.onclick = function (popup, parent, diagramElement) {
				return function () {
					var startEvent = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
					// Remove the menu.
					popup.element.parentNode.removeChild(popup.element)

					// Create a new process instance wizard...
					new ProcessWizard(parent, startEvent)
				}
			} (popup, parent, diagramElement)

		}
	} (this.parent, diagramElement)

	return group
}

SvgDiagram.prototype.drawEndEvent = function (diagramElement) {
	var group = this.drawEvent(diagramElement)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	group.element.childNodes[0].className.baseVal += " bpmndi_endEvent"
	group.element.childNodes[0].className.animVal += " bpmndi_endEvent"

	if (bpmnElement.M_eventDefinition != undefined) {
		// Here I will draw the definition element...
		this.drawEventDefintion(diagramElement, group)
	}


	return group
}

SvgDiagram.prototype.drawIntermediateEvent = function (diagramElement) {
	var group = this.drawEvent(diagramElement)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var r = diagramElement.M_Bounds.M_width / 2
	var shape = new SVG_Circle(group, diagramElement.M_id, "bpmndi_inner_circle", r, r, r * .85)

	group.element.firstElementChild.className.baseVal += " bpmndi_intermediate_event"
	if (bpmnElement.M_eventDefinition != undefined) {
		// Here I will draw the definition element...
		this.drawEventDefintion(diagramElement, group)
	}

	return group
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Connector elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawConnector = function (id, className, waypoints, attrs) {
	var group = new SVG_Group(this.plane, id, "", 100, 0)

	// Here I will set the name with the id of the bpmn element...
   	var pathStr = "M"
   	var lastPoint = null
   	var startX = 0
   	var startY = 0

   	for (var i = 0; i < waypoints.length; i++) {
		var wayPoint = waypoints[i]
		if (i == 0) {
			startX = wayPoint.M_x
			startY = wayPoint.M_y
		}

		if (i > 0) {
			pathStr += " L "
		}

		pathStr += " " + (wayPoint.M_x - startX)
		pathStr += " " + (wayPoint.M_y - startY)
		lastPoint = wayPoint
   	}

   	var sequenceFlow = new SVG_Path(group, id, className, [pathStr])
   	sequenceFlow.setSvgAttribute("name", id)

   	group.setSvgAttribute("transform", "translate(" + startX + "," + startY + ")")
   	return group
}

SvgDiagram.prototype.drawAssociation = function (diagramElement, isInput) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
   	connector = this.drawConnector(diagramElement.M_id, "bpmndi_association", diagramElement.M_waypoint)
   	connector.setSvgAttribute("name", bpmnElement.M_id)

   	if (bpmnElement.M_associationDirection == 0) {
		/* One **/
		connector.element.firstElementChild.className.baseVal += " bpmndi_data_association_one"
   	} else if (bpmnElement.M_associationDirection == 1) {
		/* Both **/
		connector.element.firstElementChild.className.baseVal += " bpmndi_data_association_both"
   	}

   	return connector
}

SvgDiagram.prototype.drawMessageFlow = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var connector = this.drawConnector(diagramElement.M_id, "bpmndi_association bpmndi_message_flow", diagramElement.M_waypoint)
	connector.setSvgAttribute("name", bpmnElement.M_id)
	return connector
}

SvgDiagram.prototype.drawSequenceFlow = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
   	connector = this.drawConnector(diagramElement.M_id, "bpmndi_association bpmndi_sequence_flow", diagramElement.M_waypoint)
   	connector.setSvgAttribute("name", bpmnElement.M_id)

   	// Now I will draw the marker...
   	var source = bpmnElement.M_sourceRef
   	var isActivityInstance = source.TYPENAME == "BPMN20.AdHocSubProcess" ||
		source.TYPENAME == "BPMN20.BusinessRuleTask" ||
		source.TYPENAME == "BPMN20.CallActivity" ||
		source.TYPENAME == "BPMN20.ManualTask" ||
		source.TYPENAME == "BPMN20.ReceiveTask" ||
		source.TYPENAME == "BPMN20.ScriptTask" ||
		source.TYPENAME == "BPMN20.SendTask" ||
		source.TYPENAME == "BPMN20.ServiceTask" ||
		source.TYPENAME == "BPMN20.Task" ||
		source.TYPENAME == "BPMN20.UserTask" ||
		source.TYPENAME == "BPMN20.SubProcess"

   	if (bpmnElement.M_conditionExpression != undefined && isActivityInstance) {
		connector.element.firstElementChild.className.baseVal += " bpmndi_sequence_flow_condition"
   	}

   	var isGatewayInstance = source.TYPENAME == "BPMN20.ComplexGateway" ||
		source.TYPENAME == "BPMN20.EventBasedGateway" ||
		source.TYPENAME == "BPMN20.ExclusiveGateway" ||
		source.TYPENAME == "BPMN20.InclusiveGateway" ||
		source.TYPENAME == "BPMN20.ParallelGateway"

   	// Now the default marker...
   	if (isObject(source.M_default) == true && (isGatewayInstance || isActivityInstance)) {
		if (source.M_default.M_id == bpmnElement.M_id) {
			connector.element.firstElementChild.className.baseVal += " bpmndi_sequence_flow_default"
		}
   	}

   	return connector
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Data elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawDataStoreReference = function (diagramElement) {
	// Data store reference...
	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	// Here I will set the name with the id of the bpmn element...
	group.setSvgAttribute("name", bpmnElement.M_id)
	group.setSvgAttribute("transform", "translate(" + diagramElement.M_Bounds.M_x + "," + diagramElement.M_Bounds.M_y + ")")

	var pathStr = this.pathMap.getScaledPath('DATA_STORE', {
        xScaleFactor: 1,
        yScaleFactor: 1,
        containerWidth: diagramElement.M_Bounds.M_width,
        containerHeight: diagramElement.M_Bounds.M_height,
        position: {
			mx: 0,
			my: 0.133
        }
	});

	var dataStoreReferencePath = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_shape bpmndi_data_store_reference", [pathStr])
	this.svgElements[bpmnElement.UUID] = group

	return group
}

SvgDiagram.prototype.drawDataInput = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var group = new SVG_Group(this.drawDataObjectReference(diagramElement), diagramElement.M_id, "", 100, 0) 

	var pathStr = this.pathMap.getScaledPath('DATA_ARROW', {
		xScaleFactor: 1,
		yScaleFactor: 1,
		containerWidth: diagramElement.M_Bounds.M_width,
		containerHeight: diagramElement.M_Bounds.M_height,
		position: {
            mx: 0,
            my: 0
		}
	});

	evtDefPath = new SVG_Path(group, bpmnElement.M_id + "_path", "bpmndi_data_input", [pathStr])
	group.setSvgAttribute("name", bpmnElement.M_id)
	group.setSvgAttribute("transform", "translate(" + (diagramElement.M_Bounds.M_x - 2) + "," + diagramElement.M_Bounds.M_y + ")")

   	return group
}

SvgDiagram.prototype.drawDataInputAssociation = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
   	connector = this.drawConnector(diagramElement.M_id, "bpmndi_association bpmndi_data_association bpmndi_data_association_one", diagramElement.M_waypoint)
   	connector.setSvgAttribute("name", bpmnElement.M_id)
   	return connector
}

SvgDiagram.prototype.drawDataOutput = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var group = new SVG_Group(this.drawDataObjectReference(diagramElement), diagramElement.M_id, "", 100, 0) 

	var pathStr = this.pathMap.getScaledPath('DATA_ARROW', {
		xScaleFactor: 1,
		yScaleFactor: 1,
		containerWidth: diagramElement.M_Bounds.M_width,
		containerHeight: diagramElement.M_Bounds.M_height,
		position: {
            mx: 0,
            my: 0
		}
	});

	evtDefPath = new SVG_Path(group, bpmnElement.M_id + "_path", "", [pathStr])
	group.setSvgAttribute("name", bpmnElement.M_id)
	group.setSvgAttribute("transform", "translate(" + (diagramElement.M_Bounds.M_x - 2) + "," + diagramElement.M_Bounds.M_y + ")")

   	return group
}

SvgDiagram.prototype.drawDataOutputAssociation = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
   	connector = this.drawConnector(diagramElement.M_id, "bpmndi_association bmndi_data_association bpmndi_data_association_one", diagramElement.M_waypoint)
   	connector.setSvgAttribute("name", bpmnElement.M_id)
   	return connector
}

SvgDiagram.prototype.drawDataObjectReference = function (diagramElement) {
	// So here I will create a new group...
	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	var pathStr = "M " + diagramElement.M_Bounds.M_x + " " + diagramElement.M_Bounds.M_y
	// The first line...
	var pt1_x = diagramElement.M_Bounds.M_x
	var pt1_y = diagramElement.M_Bounds.M_y
	var corner = diagramElement.M_Bounds.M_width * .35
	var pt2_x = pt1_x + diagramElement.M_Bounds.M_width - corner
	var pt2_y = pt1_y
	var pt3_x = pt2_x + corner
	var pt3_y = pt2_y + corner
	var pt4_x = pt2_x + corner
	var pt4_y = pt1_y + diagramElement.M_Bounds.M_height
	var pt5_x = pt1_x
	var pt5_y = pt1_y + diagramElement.M_Bounds.M_height
	// Now the little inside corner...
	var pt6_x = pt2_x
	var pt6_y = pt1_y + corner

	pathStr += " M" + pt1_x + " " + pt1_y
	pathStr += " L" + pt2_x + " " + pt2_y
	pathStr += " L" + pt3_x + " " + pt3_y
	pathStr += " L" + pt4_x + " " + pt4_y
	pathStr += " L" + pt5_x + " " + pt5_y
	pathStr += " L" + pt1_x + " " + pt1_y

	//Now the little corner...
	pathStr += " M" + pt2_x + " " + pt2_y
	pathStr += " L" + pt6_x + " " + pt6_y
	pathStr += " M" + pt6_x + " " + pt6_y
	pathStr += " L" + pt3_x + " " + pt3_y

	var dataObjectRefPath = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_shape  bpmndi_data_object_reference", [pathStr])
	dataObjectRefPath.setSvgAttribute("name", bpmnElement.M_id)

	this.svgElements[bpmnElement.UUID] = group

	if (diagramElement.M_BPMNLabel == undefined) {
		var label = this.drawText(bpmnElement.M_name, diagramElement, diagramElement.M_Bounds)
		label.element.style.minHeight = "70px"
		diagramElement.getLabel = function (label) {
			return function () {
				return label
			}
		} (label)
	}
	return group
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Gateway elements.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
SvgDiagram.prototype.drawGateway = function (diagramElement) {

	var group = new SVG_Group(this.plane, diagramElement.M_id, "", 100, 0)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	group.setSvgAttribute("name", bpmnElement.M_id)

	var x = 0
	var y = 0
	var w = diagramElement.M_Bounds.M_width
	var h = diagramElement.M_Bounds.M_height

	pt0_x = x
	pt0_y = y + (h / 2)
	pt1_x = x + (w / 2)
	pt1_y = y + h
	pt2_x = x + w
	pt2_y = y + (h / 2)
	pt3_x = x + (w / 2)
	pt3_y = y

   	var polygon = new SVG_Polygon(group, diagramElement.M_id, "bpmndi_shape bpmndi_gateway", [pt0_x, pt0_y, pt1_x, pt1_y, pt2_x, pt2_y, pt3_x, pt3_y])
   	this.svgElements[bpmnElement.UUID] = polygon
   	group.setSvgAttribute("transform", "translate(" + diagramElement.M_Bounds.M_x + "," + diagramElement.M_Bounds.M_y + ")")

   	return group
}

SvgDiagram.prototype.drawEventBasedGateway = function (diagramElement) {
	var group = this.drawGateway(diagramElement)
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]

	var r = diagramElement.M_Bounds.M_width / 2

	// The tow circle...
	var shape = new SVG_Circle(group, diagramElement.M_id, "bpmndi_inner_circle", r, r, r * .55)
	var shape_ = new SVG_Circle(group, diagramElement.M_id, "bpmndi_inner_circle", r, r, r * .45)

	// Now the pentagone...
	var points = []

	// Here the polygon are calculated.
	var Base = { x: r, y: r - (r * .35) };

	var radius = r * .35;
	for (var i = 1; i <= 5; ++i) {
		var th = i * 2 * Math.PI / 5;
		var x = Base.x + radius * Math.sin(th);
		points.push(x)
		var y = Base.y + radius - radius * Math.cos(th);
		points.push(y)
	}

	var polygon = new SVG_Polygon(group, diagramElement.M_id, "bpmndi_event_base_gateway_polygone", points)
   	this.svgElements[bpmnElement.UUID] = polygon

   	return group
}

SvgDiagram.prototype.drawExclusiveGateway = function (diagramElement) {
	var group = this.drawGateway(diagramElement)

	var pathData = this.pathMap.getScaledPath('GATEWAY_EXCLUSIVE', {
        xScaleFactor: 0.4,
        yScaleFactor: 0.4,
        containerWidth: diagramElement.M_Bounds.M_width,
        containerHeight: diagramElement.M_Bounds.M_height,
        position: {
			mx: 0.32,
			my: 0.3
        }
	});

	var path = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_exclusive_gateway", [pathData])
	path.setSvgAttribute("x", diagramElement.M_Bounds.M_x)
	path.setSvgAttribute("y", diagramElement.M_Bounds.M_y)

	return group
}

SvgDiagram.prototype.drawInclusiveGateway = function (diagramElement) {
	var group = this.drawGateway(diagramElement)
	var r = diagramElement.M_Bounds.M_width / 2
	var shape = new SVG_Circle(group, diagramElement.M_id, "bpmndi_inclusive_gateway", r, r, r * .55)

	return group
}

SvgDiagram.prototype.drawParallelGateway = function (diagramElement) {
	var group = this.drawGateway(diagramElement)
	var pathStr = this.pathMap.getScaledPath('GATEWAY_PARALLEL', {
        xScaleFactor: 0.6,
        yScaleFactor: 0.6,
        containerWidth: diagramElement.M_Bounds.M_width,
        containerHeight: diagramElement.M_Bounds.M_height,
        position: {
			mx: 0.46,
			my: 0.2
        }
	});
	var path = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_parallel_gateway", [pathStr])
	path.setSvgAttribute("x", diagramElement.M_Bounds.M_x)
	path.setSvgAttribute("y", diagramElement.M_Bounds.M_y)
	return group
}

SvgDiagram.prototype.drawComplexGateway = function (diagramElement) {
	var group = this.drawGateway(diagramElement)
	var pathStr = this.pathMap.getScaledPath('GATEWAY_COMPLEX', {
        xScaleFactor: 0.5,
        yScaleFactor: 0.5,
        containerWidth: diagramElement.M_Bounds.M_width,
        containerHeight: diagramElement.M_Bounds.M_height,
        position: {
			mx: 0.46,
			my: 0.26
        }
	});
	var path = new SVG_Path(group, diagramElement.M_id + "_path", "bpmndi_complex_gateway", [pathStr])
	path.setSvgAttribute("x", diagramElement.M_Bounds.M_x)
	path.setSvgAttribute("y", diagramElement.M_Bounds.M_y)
	return group
}


/*
 * The process or call activity...
 */
SvgDiagram.prototype.drawCallableActivity = function (diagramElement) {
	var bpmnElement = server.workflowManager.bpmnElements[diagramElement.M_bpmnElement]
	var className = ""
	if (bpmnElement.TYPENAME == "BPMN20.SubProcess") {
		className += " bpmndi_sub_process"
	} else if (bpmnElement.TYPENAME == "BPMN20.Group") {
		className += " bpmndi_group"
	} else if (bpmnElement.TYPENAME == "BPMN20.CallActivity") {
		className += " bpmndi_call_activity"
	}

	var group = this.drawTask(diagramElement, className)
	if (bpmnElement.TYPENAME != "BPMN20.Group") {
		// Here I'm in plain html...
		var markerRect = new SVG_Rectangle(group, "bpmndi_expand_marker", "bpmndi_expand_marker",
			0, 0,
			14, 14,
			0, 0)
		markerRect.element.style.zIndex = 10

		// Process marker is placed in the middle of the box
		// therefore fixed values can be used here
		markerRect.setSvgAttribute("transform", 'translate(' + (diagramElement.M_Bounds.M_width / 2 - 7.5) + ',' + (diagramElement.M_Bounds.M_height - 20) + ')')
		var markerPath = this.pathMap.getScaledPath('MARKER_SUB_PROCESS', {
			xScaleFactor: 1.5,
			yScaleFactor: 1.5,
			containerWidth: diagramElement.M_Bounds.M_width,
			containerHeight: diagramElement.M_Bounds.M_height,
			position: {
				mx: (diagramElement.M_Bounds.M_width / 2 - 7.5) / diagramElement.M_Bounds.M_width,
				my: (diagramElement.M_Bounds.M_height - 20) / diagramElement.M_Bounds.M_height
			}
		});

		var path = new SVG_Path(group, "marker_path", "marker_path", [markerPath])
	}

	return group
}
