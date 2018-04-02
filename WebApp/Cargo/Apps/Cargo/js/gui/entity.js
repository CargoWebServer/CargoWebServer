

var EntityPanel = function (parent, typeName, callback) {
	this.typeName = typeName

	// The main entity panel.
	this.panel = new Element(parent, { "tag": "div", "class": "entity_panel" })

	// fields panels.
	this.fields = {}

	// Keep track of current display entity.
	this.entityUuid = ""

	// The panel header.
	this.header = this.panel.appendElement({ "tag": "div", "class": "entity_panel_header" }).down()

	// Get the prototype from the server.
	this.getEntityPrototype(callback)

	return this;
}

EntityPanel.prototype.getEntityPrototype = function (callback) {
	server.entityManager.getEntityPrototype(this.typeName, this.typeName.split(".")[0],
		// success callback
		function (prototype, caller) {
			// Initialyse the entity panel from it content.
			caller.entityPanel.init(prototype, caller.callback)
		},
		// error callbacak
		function () {

		}, { "entityPanel": this, "callback": callback }
	)
}

EntityPanel.prototype.init = function (prototype, callback) {
	// Now I will set the field panel.
	var initField = function (entityPanel, prototype, index, callback) {
		if (prototype.FieldsVisibility[index] == true) {
			new FieldPanel(entityPanel, index,
				function (entityPanel, prototype, index, callback) {
					return function (fieldPanel) {
						entityPanel.fields[prototype.Fields[index]] = fieldPanel
						index += 1
						if (index < prototype.Fields.length) {
							initField(entityPanel, prototype, index, callback)
						} else {
							callback(entityPanel)
						}
					}
				}(entityPanel, prototype, index, callback)
			)
		}
	}

	// The first tree field are not display.
	var index = 3
	if (prototype.Fields.length > 3) {
		initField(this, prototype, index, callback)
	}
}

/**
 * Display an entity in the panel.
 * @param {*} entity 
 */
EntityPanel.prototype.setEntity = function (entity) {
	// In that case I will set the value of the field renderer.
	var prototype = getEntityPrototype(entity.TYPENAME)
	this.entityUuid = entity.UUID

	for (var i = 3; i < prototype.Fields.length; i++) {
		if (this.fields[prototype.Fields[i]] != undefined && entity[prototype.Fields[i]] != undefined) {
			this.fields[prototype.Fields[i]].setValue(entity[prototype.Fields[i]])
		}
	}
}

/**
 * Remove the entity entity from the panel.
 * @param {*} entity 
 */
EntityPanel.prototype.clear = function () {
	this.entityUuid = ""
	for (var id in this.fields) {
		this.fields[id].clear()
	}
}

var FieldPanel = function (entityPanel, index, callback) {
	this.panel = entityPanel.panel.appendElement({ "tag": "div", "class": "field_panel" }).down()

	this.parent = entityPanel;

	this.fieldName = getEntityPrototype(entityPanel.typeName).Fields[index]
	this.fieldType = getEntityPrototype(entityPanel.typeName).FieldsType[index]

	var title = this.fieldName.replace("M_", "").replaceAll("_", " ")
	
	// Display label if is not valueOf or listOf...
	if(this.fieldName != "M_valueOf" && this.fieldName != "M_listOf"){
		this.label = this.panel.appendElement({ "tag": "div", "innerHtml": title }).down();
	}
	this.value = this.panel.appendElement({ "tag": "div" }).down();

	// Here I will create the field renderer.
	this.renderer = null;
	new FieldRenderer(this, function (callback, fieldPanel) {
		return function (fieldRenderer) {
			fieldPanel.renderer = fieldRenderer;
			callback(fieldPanel)
		}
	}(callback, this))

	this.editor = null

	return this
}

FieldPanel.prototype.setValue = function (value) {
	this.renderer.setValue(value)
	// clear the editor
	if (this.editor != null) {
		this.editor.setValue(value)
	}
}

FieldPanel.prototype.clear = function () {
	// clear the editor
	if (this.editor != null) {
		this.editor.clear()
	}

	// clear the renderer
	this.renderer.clear()
}

// The field renderer.
var FieldRenderer = function (fieldPanel, callback) {
	this.parent = fieldPanel;
	this.isArray = this.parent.fieldType.startsWith("[]")
	this.isRef = this.parent.fieldType.endsWith(":Ref")
	// Init the renderer.
	this.init(callback)

	this.renderer = null

	return this;
}

FieldRenderer.prototype.init = function (callback) {
	var typeName = this.parent.fieldType.replace("[]", "").replace(":Ref", "")
	server.entityManager.getEntityPrototype(typeName, typeName.split(".")[0],
		function (prototype, caller) {
			// Here I will render the panel, create sub-panel render etc...
			caller.fieldRenderer.render(prototype, caller.callback)
		},
		function () {

		},
		{ "fieldRenderer": this, "callback": callback }
	)
}

/**
 * Create html element to render the entity value.
 * @param {*} prototype 
 */
FieldRenderer.prototype.render = function (prototype, callback) {
	if (this.isArray) {
		// Array value use a table to display the entity.
		this.renderer = new Table(randomUUID(), this.parent.value)
		var model = undefined
		if (this.parent.fieldName != "M_listOf" && !this.parent.fieldType.startsWith("[]xs.")) {
			model = new EntityTableModel(prototype)
		} else {
			model = new TableModel(["index", "values"])
			model.fields = ["xs.int", this.parent.fieldType.replace("[]", "")]
		}
		this.renderer.setModel(model,
			function (table, callback, fieldRenderer) {
				return function () {
					table.init()
					table.refresh()
					callback(fieldRenderer)
				}
			}(this.renderer, callback, this))
	} else {
		if (this.isRef) {
			// Render a reference....
			this.renderer = this.parent.value
		} else {
			// simply set the value of parent field panel.
			if (this.parent.fieldType.startsWith("xs.")) {
				this.renderer = this.parent.value
				callback(this)
			} else {
				new EntityPanel(this.parent.value, prototype.TypeName,
					function (callback, fieldRenderer) {
						return function (entityPanel) {
							fieldRenderer.renderer = entityPanel
							callback(fieldRenderer)
						}
					}(callback, this))
			}
		}
	}
}

FieldRenderer.prototype.setValue = function (value) {
	//console.log("---> 209", value)
	if (this.renderer != null) {
		if (this.isArray) {
			if (this.parent.fieldName != "M_listOf" && !this.parent.fieldType.startsWith("[]xs.")) {
				// Here we got an array of entities
				for (var i = 0; i < value.length; i++) {
					var data = this.renderer.getModel().appendRow(value[i])
					var row = new TableRow(this.renderer, this.renderer.rows.length, data, undefined)
					row.saveBtn.element.style.visibility = "visible";
					this.renderer.header.maximizeBtn.element.click();
				}
			} else {
				// Here we got an array of basic types.
				for (var i = 0; i < value.length; i++) {
					// simply append the values with there index in that case.
					var row = this.renderer.appendRow([i + 1, value[i]], i)

					// The delete row action...
					row.deleteBtn.element.onclick = function (uuid, field, row) {
						return function () {
							// Here I will simply remove the element 
							// The entity must contain a list of field...
							if (entities[uuid] != undefined) {
								entity = entities[uuid]
							}

							if (entity[field] != undefined) {
								entity[field].splice(row.index, 1)
								entity.NeedSave = true
								server.entityManager.saveEntity(entity)
							}
						}
					}(this.parent.parent.entityUuid, this.parent.fieldName, row)

					// The save row action
					row.saveBtn.element.onclick = function (uuid, field, row) {
						return function () {
							// Here I will simply remove the element 
							// The entity must contain a list of field...
							if (entities[uuid] != undefined) {
								entity = entities[uuid]
							}

							if (entity[field] != undefined) {
								entity[field][row.index] = row.table.getModel().getValueAt(row.index, 1)
								entity.NeedSave = true
								if (entity.UUID != "") {
									server.entityManager.saveEntity(entity,
										function (result, caller) {
											caller.style.visibility = "hidden"
										},
										function () {

										}, this)
								} else {
									// Here the entity dosent exist...
									server.entityManager.createEntity(entity.ParentUuid, entity.parentLnk, entity.TYPENAME, "", entity,
										function (result, caller) {
											caller.style.visibility = "hidden"
										},
										function () {

										}, this)
								}
							}
						}
					}(this.parent.parent.entityUuid, this.parent.fieldName, row)
				}
			}
		} else {
			// not array...
			var fieldType = this.parent.fieldType
			if (fieldType.startsWith("xs.")) {
				if (isXsId(fieldType) || isXsString(fieldType || isXsRef(fieldType))) {
					this.parent.value.element.innerHTML = value
				} else if (isXsNumeric(fieldType)) {
					this.parent.value.element.innerHTML = parseFloat(value)
				} else if (isXsInt(fieldType)) {
					this.parent.value.element.innerHTML = parseInt(value)
				} else if (isXsTime(fieldType)) {
					var value = moment(value).unix()
				} else if (isXsBoolean(fieldType)) {
					this.parent.value.element.innerHTML = value
				} else if (fieldType.startsWith("enum:")) { // Cargo enum not xsd extention.

				}
			} else if (value.TYPENAME != undefined) {
				// In that case I got a subpanel....
				this.renderer.setEntity(value)
			} else {
				this.parent.value.element.innerHTML = value
			}
		}
	} else {
		// In that case the renderer was not completely initialysed so I will intialyse it and set it value 
		// after.
		this.renderer = this.render(getEntityPrototype(this.parent.fieldType.replace("[]", "").replace(":Ref", "")),
			function (value, fieldRenderer) {
				return function () {
					fieldRenderer.setValue(value)
				}
			}(value, this))
	}
}

FieldRenderer.prototype.clear = function () {

}

// The field editor.
var FieldEditor = function (fieldPanel, field, callback) {

	return this;
}

FieldEditor.prototype.setValue = function (value) {

}

FieldEditor.prototype.clear = function () {

}