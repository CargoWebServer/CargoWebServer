
var EntityPanel = function (parent, typeName, callback) {
	this.typeName = typeName

	// The main entity panel.
	this.panel = new Element(parent, { "tag": "div", "class": "entity entity_panel" })

	// fields panels.
	this.fields = {}

	// Keep track of current display entity.
	this.entityUuid = ""

	// The panel header.
	this.header = null

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
		new FieldPanel(entityPanel, index,
			function (entityPanel, prototype, index, callback) {
				return function (fieldPanel) {
					entityPanel.fields[prototype.Fields[index]] = fieldPanel
					index += 1
					if (index < prototype.Fields.length) {
						if (prototype.FieldsVisibility[index] == false) {
							index += 1
						}
						if (index < prototype.Fields.length) {
							initField(entityPanel, prototype, index, callback)
						} else {
							entityPanel.header = new EntityPanelHeader(entityPanel)
							callback(entityPanel)
						}
					} else {
						entityPanel.header = new EntityPanelHeader(entityPanel)
						callback(entityPanel)
					}
				}
			}(entityPanel, prototype, index, callback)
		)
	}
	// The first tree field are not display.
	if (prototype != undefined) {
		var index = 3
		if (prototype.Fields.length > 3) {
			initField(this, prototype, index, callback)
		}
	}
}

/**
 * Display an entity in the panel.
 * @param {*} entity 
 */
EntityPanel.prototype.setEntity = function (entity) {

	// Here I will associate the panel and the entity.
	if (this.getEntity != null) {
		server.entityManager.detach(this, NewEntityEvent)
		server.entityManager.detach(this, UpdateEntityEvent)
		server.entityManager.detach(this, DeleteEntityEvent)

		this.getEntity().getPanel = null
		this.clear()
	}

	// In that case I will set the value of the field renderer.
	var prototype = getEntityPrototype(entity.TYPENAME)
	this.entityUuid = entity.UUID
	this.entity = entity // use it as read only...

	this.getEntity = function (entity) {
		return function () {
			return entity
		}
	}(entity)

	if (entity.getTitles != undefined) {
		if (entity.getTitles().length > 0) {
			this.header.setTitle(entity.getTitles())
		} else {
			this.header.setTitle([entity.TYPENAME])
		}
	}

	for (var i = 3; i < prototype.Fields.length; i++) {
		if (this.fields[prototype.Fields[i]] != undefined && entity[prototype.Fields[i]] != undefined) {
			var value = entity[prototype.Fields[i]]
			if (value == "false") {
				value = false
			} else if (value == "true") {
				value = true
			}

			this.fields[prototype.Fields[i]].setValue(value)
		}

		// If the field is a table I will set it parent uuid.
		if (this.fields[prototype.Fields[i]] != undefined) {
			if (this.fields[prototype.Fields[i]].renderer != undefined) {
				if (this.fields[prototype.Fields[i]].renderer.renderer != undefined) {
					if (this.fields[prototype.Fields[i]].renderer.renderer.getModel != undefined) {
						this.fields[prototype.Fields[i]].renderer.renderer.getModel().ParentUuid = entity.UUID
						this.fields[prototype.Fields[i]].renderer.renderer.getModel().ParentLnk = this.fields[prototype.Fields[i]].fieldName
					}
				}
			}
		}
	}

	entity.getPanel = function (entityPanel) {
		return function () {
			return entityPanel
		}
	}(this)

	/////////////////////////////////////////////////////////////////////////////////
	// Now the event listener...
	/////////////////////////////////////////////////////////////////////////////////

	// The delete entity event.
	server.entityManager.attach(this, DeleteEntityEvent, function (evt, entityPanel) {
		if (evt.dataMap["entity"].TYPENAME == entityPanel.typeName) {
			if (evt.dataMap["entity"] != undefined && entityPanel.getEntity() != null) {
				if (evt.dataMap["entity"].UUID == entityPanel.getEntity().UUID) {
					// so here i will remove the panel from it parent.
					if (entityPanel.removeOnDelete) {
						try {
							entityPanel.panel.element.parentNode.removeChild(entityPanel.panel.element)
						} catch (err) {
							// Nothing to do here.
						}
					} else {
						// Clear the entity value
						entityPanel.clear()
					}
				}
			}
		}
	})

	// The new entity event.
	server.entityManager.attach(this, NewEntityEvent, function (evt, entityPanel) {
		// I will reinit the panel here...
		if (evt.dataMap["entity"].TYPENAME == entityPanel.typeName) {
			if (evt.dataMap["entity"] && entityPanel.getEntity() != null) {
				if (entityPanel.getEntity().UUID == evt.dataMap["entity"].UUID) {
					entityPanel.init(entityPanel.proto)
					entityPanel.setEntity(evt.dataMap["entity"])
				}
			}
		}
	})

	// The update entity event.
	server.entityManager.attach(this, UpdateEntityEvent, function (evt, entityPanel) {
		if (evt.dataMap["entity"].TYPENAME == entityPanel.typeName) {
			if (evt.dataMap["entity"] && entityPanel.getEntity() != null) {
				// I will reinit the panel here...
				if (entityPanel.getEntity().UUID == evt.dataMap["entity"].UUID) {
					entityPanel.setEntity(evt.dataMap["entity"])
				}
			}
		}
	})
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

	if (this.getEntity() == null) {
		return
	}

	this.entity = null

	// Remove actual panel element befor reinit it.
	this.panel.removeAllChilds()

	this.init(getEntityPrototype(this.typeName),
		// init callback.
		function () {
		})
	if (this.initCallback != undefined) {
		this.initCallback(this)
	}
}

var EntityPanelHeader = function (parent) {
	this.panel = parent.panel.prependElement({ "tag": "div", "class": "entity_panel_header", "style": "display: none;" }).down()
	this.expandBtn = this.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button", "style": "display: none;" }).down()
	this.expandBtn.appendElement({ "tag": "i", "class": "fa fa-caret-right" }).down()
	this.shrinkBtn = this.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
	this.shrinkBtn.appendElement({ "tag": "i", "class": "fa fa-caret-down" }).down()
	this.title = this.panel.appendElement({ "tag": "div", "class": "entity_panel_header_title" }).down()

	// Now the event...
	this.expandBtn.element.onclick = function (header, entityPanel) {
		return function (evt) {
			evt.stopPropagation(true)
			header.expandBtn.element.style.display = "none"
			header.shrinkBtn.element.style.display = ""
			for (var field in entityPanel.fields) {
				entityPanel.fields[field].panel.element.style.display = ""
			}
		}
	}(this, parent)

	this.shrinkBtn.element.onclick = function (header, entityPanel) {
		return function (evt) {
			evt.stopPropagation(true)
			header.expandBtn.element.style.display = ""
			header.shrinkBtn.element.style.display = "none"
			for (var field in entityPanel.fields) {
				entityPanel.fields[field].panel.element.style.display = "none"
			}
		}
	}(this, parent)

	return this;
}

EntityPanelHeader.prototype.display = function () {
	this.panel.element.style.display = ""
	this.shrinkBtn.element.click()
}

EntityPanelHeader.prototype.setTitle = function (titles) {
	var title = ""
	for (var i = 0; i < titles.length; i++) {
		title += titles[i]
		if (i < titles.length - 1) {
			title += " "
		}
	}
	this.title.element.innerHTML = title
}

var FieldPanel = function (entityPanel, index, callback) {
	this.panel = entityPanel.panel.appendElement({ "tag": "div", "class": "field_panel", "style": "" }).down()
	this.parent = entityPanel;
	this.storeId = getEntityPrototype(entityPanel.typeName).PackageName
	this.fieldName = getEntityPrototype(entityPanel.typeName).Fields[index]
	this.fieldType = getEntityPrototype(entityPanel.typeName).FieldsType[index]

	var title = this.fieldName.replace("M_", "").replaceAll("_", " ")

	// Display label if is not valueOf or listOf...
	if (this.fieldName != "M_valueOf" && this.fieldName != "M_listOf") {
		this.label = this.panel.appendElement({ "tag": "div", "innerHtml": title, "style": "min-width: 100px;" }).down();
	}
	this.value = this.panel.appendElement({ "tag": "div", "class": "field_panel_value" }).down();

	// Here I will create the field renderer.
	this.renderer = null;

	// init the renderer
	this.init(callback)

	this.editor = null

	this.value.element.onclick = function (fieldPanel) {
		return function (evt) {
			evt.stopPropagation(true)
			if (fieldPanel.editor != null || fieldPanel.fieldType.startsWith("[]")) {
				return // editor already exist
			}

			new FieldEditor(fieldPanel, function (value) {
				return function (fieldPanel, editor) {
					if (editor.editor != null) {
						// clear the content.
						fieldPanel.value.element.style.display = "none"
						fieldPanel.editor = editor
						fieldPanel.editor.setValue(value)
					}
				}
			}(this.innerText))
		}
	}(this)

	return this
}

FieldPanel.prototype.init = function (callback) {
	new FieldRenderer(this, function (callback, fieldPanel) {
		return function (fieldRenderer) {
			fieldPanel.renderer = fieldRenderer;
			callback(fieldPanel)
		}
	}(callback, this))
}

FieldPanel.prototype.setValue = function (value) {
	// If the value is an entity reference... or an array of entity
	// and the type is not a reference I will try to get it value 
	// from the entitie map.
	if (!this.fieldType.endsWith(":Ref")) {
		if (this.fieldType.startsWith("[]")) {
			var values = []
			for (var i = 0; i < value.length; i++) {
				if (isObjectReference(value[i])) {
					values.push(entities[value[i]])
				} else {
					values.push(value[i])
				}
			}
			value = values
		} else {
			if (isObjectReference(value)) {
				value = entities[value]
			}
		}
	}

	this.value.element.style.display = ""
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
	this.renderer = null

	// Init the renderer.
	this.init(callback)

	return this;
}

FieldRenderer.prototype.init = function (callback) {
	var typeName = this.parent.fieldType.replace("[]", "").replace(":Ref", "")
	if (typeName.startsWith("enum:")) {
		this.renderer = this.parent.value
		callback(this)
	} else {
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
}

/**
 * Create html element to render the entity value.
 * @param {*} prototype 
 */
FieldRenderer.prototype.render = function (prototype, callback) {
	if (this.isArray) {
		// Array value use a table to display the entity.
		var div = this.parent.value.appendElement({ "tag": "div", "style": "display: flex; justify-content: left; align-items: flex-start;" }).down()
		this.renderer = new Table(randomUUID(), div)

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
			callback(this)
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
	if (this.renderer != null) {
		if (this.isArray) {
			if (this.parent.fieldName != "M_listOf" && !this.parent.fieldType.startsWith("[]xs.")) {
				// Here we got an array of entities
				function setValue(value, fieldRenderer) {
					var row = fieldRenderer.renderer.appendRow(value, value.UUID)
					if (row == undefined) {
						var data = fieldRenderer.renderer.getModel().appendRow(value)
						// The row is not append in the table rows collection, but display.
						var row = new TableRow(fieldRenderer.renderer, fieldRenderer.renderer.rows.length, data, undefined)
						var lastRowIndex = fieldRenderer.renderer.rows.length
						simulate(row.cells[lastRowIndex, 0].div.element, "dblclick");
					}

					if (fieldRenderer.parent.parent.entityUuid.length > 0) {
						value.ParentUuid = fieldRenderer.parent.parent.entityUuid
						value.ParentLnk = fieldRenderer.parent.fieldName
					}

					fieldRenderer.renderer.header.maximizeBtn.element.click();
				}
				if (isArray(value)) {
					for (var i = 0; i < value.length; i++) {
						setValue(value[i], this)
					}
				} else {
					setValue(value, this)
				}
			} else {
				function setValue(value, index, fieldRenderer) {
					// simply append the values with there index in that case.
					var v = formatValue(value, fieldRenderer.parent.fieldType.replace("[]", ""))
					var row = fieldRenderer.renderer.appendRow([index + 1, v], index)
					var model = fieldRenderer.renderer.getModel()

					// set the save value function for the row.
					model.saveValue = function (uuid, field) {
						return function (row) {
							// Here I will simply remove the element 
							// The entity must contain a list of field...
							if (entities[uuid] != undefined) {
								entity = entities[uuid]
							}

							if (entity[field] != undefined) {
								var values = row.table.getModel().values
								entity[field] = []
								for (var i = 0; i < values.length; i++) {
									entity[field][i] = row.table.getModel().getValueAt(i, 1)
								}

								if (entity.UUID != "") {
									server.entityManager.saveEntity(entity,
										function (result, caller) {
										},
										function () {

										}, this)
								} else {
									// Here the entity dosent exist...
									server.entityManager.createEntity(entity.ParentUuid, entity.ParentLnk, entity,
										function (result, caller) {
										},
										function () {

										}, this)
								}
							}
						}
					}(fieldRenderer.parent.parent.entityUuid, fieldRenderer.parent.fieldName)
				}
				// Here we got an array of basic types.
				if (isArray(value)) {
					for (var i = 0; i < value.length; i++) {
						setValue(value[i], i, this)
					}
				} else {
					setValue(value.value, value.index, this)
				}
			}
		} else {
			// not array...
			var fieldType = this.parent.fieldType
			if (fieldType.startsWith("xs.")) {
				if (isXsId(fieldType) || isXsString(fieldType || isXsRef(fieldType))) {
					this.parent.value.element.innerText = value
				} else if (isXsNumeric(fieldType)) {
					this.parent.value.element.innerText = parseFloat(value)
				} else if (isXsInt(fieldType)) {
					this.parent.value.element.innerText = parseInt(value)
				} else if (isXsTime(fieldType)) {
					var value = moment(value).unix()
				} else if (isXsBoolean(fieldType)) {
					this.parent.value.element.innerText = value
				}
			} else if (value.TYPENAME != undefined) {
				// In that case I got a subpanel....
				this.renderer.setEntity(value)
			} else if (fieldType.startsWith("enum:")) {
				var values = fieldType.replace("enum:", "").split(":")
				if (value - 1 >= 0) {
					// keep the value in the element.
					this.parent.value.element.value = value
					this.parent.value.element.innerText = values[value - 1].substring(values[value - 1].indexOf("_") + 1).replace("_", " ")
				}
			} else {
				this.parent.value.element.innerText = value
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
var FieldEditor = function (fieldPanel, callback) {
	var fieldType = fieldPanel.fieldType
	this.editor = null
	this.parent = fieldPanel
	var id = this.parent.parent.entityUuid + "_" + this.parent.fieldName

	if (fieldType.startsWith("xs.")) {
		if (isXsId(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "id": id }).down()
		} else if (isXsRef(fieldType)) {
			// Reference here... autocomplete...
			this.editor = this.parent.panel.appendElement({ "tag": "input", "id": id, "type": "url" }).down()
		} else if (isXsInt(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "1", "id": id }).down()
		} else if (isXsDate(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "date", "id": id }).down()
		} else if (isXsTime(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "datetime-local", "id": id }).down()
		} else if (isXsString(fieldType)) {
			if (this.parent.fieldName == "M_description") {
				this.parent.panel.element.style.verticalAlign = "top"
				this.editor = this.parent.panel.appendElement({ "tag": "textarea", "id": id }).down()
			} else {
				this.editor = this.parent.panel.appendElement({ "tag": "input", "id": id }).down()
				if (this.parent.fieldName.indexOf("pwd") > -1) {
					this.editor.element.type = "password"
				}
			}
		} else if (isXsBinary(fieldType)) {

		} else if (isXsBoolean(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "checkbox", "id": id }).down()
		} else if (isXsNumeric(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "0.01", "id": id }).down()
		} else if (isXsMoney(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "0.01", "id": id }).down()
		}

		// Cancel action on esc, set value on enter.
		this.editor.element.onkeyup = function (fieldPanel) {
			return function (e) {
				e.stopPropagation(true)
				if (e.keyCode === 27) {
					fieldPanel.value.element.style.display = ""
					this.onblur()
				} else if (e.keyCode === 13) {
					fieldPanel.value.element.style.display = ""
					this.onblur()
				}
			}
		}(fieldPanel)
	} else if (fieldType.startsWith("enum:")) {
		this.editor = this.parent.panel.appendElement({ "tag": "select", "id": id }).down()
		var values = fieldType.replace("enum:", "").split(":")

		// keep the value in the element.
		for (var i = 0; i < values.length; i++) {
			var text = values[i].substring(values[i].indexOf("_") + 1).replace("_", " ")
			this.editor.appendElement({ "tag": "option", "innerHtml": text, "value": text })
		}
	}

	if (this.editor != null) {
		// Now the event...
		this.editor.element.onblur = function (fieldPanel) {
			return function (evt) {
				var value
				fieldPanel.value.element.innerText = this.value

				if (this.type == "checkbox") {
					value = this.checked
				} else if (this.tagName == "SELECT") {

					value = this.selectedIndex + 1
				} else {
					value = this.value
				}

				// Now I will modify the value in the local entity.
				var entity = entities[fieldPanel.parent.entityUuid]
				if (entity == undefined) {
					// newly created entity not alread
					entity = fieldPanel.parent.entity
					entity[fieldPanel.fieldName] = value
					if (entity.getParent != undefined) {
						entity = entity.getParent()
					}
				} else {
					entity[fieldPanel.fieldName] = value

				}
				// Remove the editor
				try {
					this.parentNode.removeChild(this)
				} catch (error) {

				}
				fieldPanel.editor = null // remove the edior.
				fieldPanel.value.element.style.display = ""

				// In that case I will save the entity directly.
				/*server.entityManager.saveEntity(entity,
					function (result, caller) {
					},
					// error callback
					function () {

					}, {})*/


			}
		}(fieldPanel)

		this.editor.element.focus()
	}

	callback(fieldPanel, this)
	return this;
}

FieldEditor.prototype.setValue = function (value) {
	if (this.editor.element.tagName == "INPUT") {
		if (this.editor.element.type == "checkbox") {
			if (value == "true") {
				this.editor.element.checked = true
			} else {
				this.editor.element.checked = false
			}
		} else {
			this.editor.element.value = value
		}
	} else if (this.editor.element.tagName == "SELECT") {
		this.editor.element.value = value
	}
}

FieldEditor.prototype.clear = function () {

}