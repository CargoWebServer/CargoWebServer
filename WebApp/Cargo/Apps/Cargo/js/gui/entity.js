
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
							if (entityPanel.header == null) {
								entityPanel.header = new EntityPanelHeader(entityPanel)
							}
							callback(entityPanel)
						}
					} else {
						if (entityPanel.header == null) {
							entityPanel.header = new EntityPanelHeader(entityPanel)
						}
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
	var displayHeader = this.header.panel.element.style.display == ""
	if (this.getEntity != null) {
		server.entityManager.detach(this, NewEntityEvent)
		server.entityManager.detach(this, UpdateEntityEvent)
		server.entityManager.detach(this, DeleteEntityEvent)
		this.getEntity().getPanel = null
		this.clear()

		// If the entity is a new one I will remove all it nodes except the header.
		if (this.getEntity().UUID == undefined) {
			for (var childId in this.panel.childs) {
				var child = this.panel.childs[childId]
				if (child.element.className != "entity_panel_header") {
					this.panel.removeElement(child)
				}
			}
		}
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

	if (displayHeader) {
		this.header.display()
		this.header.expandBtn.element.click()
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

	// remove the function.
	this.entity.getPanel = undefined

	// set the entity to nill
	this.entity = null

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
	this.deleteBtn = this.panel.appendElement({ "tag": "div", "class": "entity_panel_header_button enabled"/*, "style": "display: none;"*/ }).down()
	this.deleteBtn.appendElement({ "tag": "i", "class": "fa fa-trash" })

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

	this.deleteBtn.element.onclick = function (entityPanel) {
		return function () {
			// Here I will save the entity...
			if (entityPanel.getEntity() != null) {
				// Here I will ask the user if here realy want to remove the entity...
				var confirmDialog = new Dialog(randomUUID(), undefined, true)
				confirmDialog.div.element.style.maxWidth = "450px"
				confirmDialog.setCentered()
				server.languageManager.setElementText(confirmDialog.title, "delete_dialog_entity_title")
				confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete entity " + entityPanel.getEntity().getTitles() + "?" })

				confirmDialog.ok.element.onclick = function (dialog, entityPanel) {
					return function () {
						// I will call delete file
						server.entityManager.removeEntity(entityPanel.getEntity().UUID,
							// Success callback 
							function (result, caller) {
								/** The action will be done in the event listener */
								if (caller.deleteCallback != undefined) {
									caller.deleteCallback(caller.getEntity())
								}
							},
							// Error callback
							function (errMsg, caller) {

							}, entityPanel)
						dialog.close()
					}
				}(confirmDialog, entityPanel)
			}
		}
	}(parent)

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
			var entity = fieldPanel.parent.getEntity()
			new FieldEditor(fieldPanel, function (value) {
				return function (fieldPanel, editor) {
					if (editor.editor != null) {
						// clear the content.
						fieldPanel.value.element.style.display = "none"
						fieldPanel.editor = editor
						fieldPanel.editor.setValue(value)
					}
				}
			}(entity[fieldPanel.fieldName]))
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

	// Remove the panel.
	this.value.removeAllChilds()
}

// The field renderer.
var FieldRenderer = function (fieldPanel, callback) {
	this.parent = fieldPanel;
	this.isArray = this.parent.fieldType.startsWith("[]")
	this.isRef = this.parent.fieldType.endsWith(":Ref")
	this.renderer = null

	// Init the renderer.
	if (this.renderer == null) {
		this.init(callback)
	} else {
		callback(this)
	}
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
			model = new TableModel(["values"])
			model.fields = [this.parent.fieldType.replace("[]", "")]
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
					var row = fieldRenderer.renderer.appendRow([v], index)
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
			var enumarations = []
			var entity = this.parent.parent.getEntity()
			if (entity != undefined) {
				prototype = getEntityPrototype(entity.TYPENAME)
				if (prototype.Restrictions != undefined) {
					for (var i = 0; i < prototype.Restrictions.length; i++) {
						if (prototype.Restrictions[i].Type == 1) {
							enumarations.push(prototype.Restrictions[i].Value)
						}
					}
				}
			}

			if (fieldType.startsWith("enum:") || enumarations.length > 0) {
				if (enumarations.length > 0) {
					this.parent.value.element.value = value
					this.parent.value.element.innerText = enumarations[value - 1]
				} else {
					var values = fieldType.replace("enum:", "").split(":")
					if (value - 1 >= 0) {
						// keep the value in the element.
						this.parent.value.element.value = value
						this.parent.value.element.innerText = values[value - 1].substring(values[value - 1].indexOf("_") + 1).replace("_", " ")
					}
				}
			} else if (fieldType.startsWith("xs.")) {
				if (isXsId(fieldType) || isXsString(fieldType || isXsRef(fieldType))) {
					this.parent.value.element.innerText = value
				} else if (isXsNumeric(fieldType)) {
					this.parent.value.element.innerText = parseFloat(value)
				} else if (isXsInt(fieldType)) {
					this.parent.value.element.innerText = parseInt(value)
				} else if (isXsTime(fieldType)) {
					var date
					if (isString(value)) {
						date = new Date(value)
					} else {
						date = new Date(value * 1000)
					}
					this.parent.value.element.innerText = date.toLocaleDateString() + " " + date.toLocaleTimeString()
				} else if (isXsBoolean(fieldType)) {
					this.parent.value.element.innerText = value
				}
			} else if (value.TYPENAME != undefined) {
				// In that case I got a subpanel....
				this.renderer.setEntity(value)
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
	if (this.renderer != null) {
		if (this.renderer.removeAllChilds != undefined) {
			this.renderer.removeAllChilds() // clear it content.
		} else if (this.renderer.clear != undefined) {
			this.renderer.clear() // table
		}
	}
}

// The field editor.
var FieldEditor = function (fieldPanel, callback) {
	var fieldType = fieldPanel.fieldType
	this.editor = null
	this.parent = fieldPanel
	var id = this.parent.parent.entityUuid + "_" + this.parent.fieldName
	var entity = this.parent.parent.getEntity()

	// Now the restrictions...
	var enumarations = []

	if (entity != undefined) {
		prototype = getEntityPrototype(entity.TYPENAME)
		if (prototype.Restrictions != undefined) {
			for (var i = 0; i < prototype.Restrictions.length; i++) {
				if (prototype.Restrictions[i].Type == 1) {
					enumarations.push(prototype.Restrictions[i].Value)
				}
			}
		}
	}

	if (fieldType.startsWith("enum:") || enumarations.length > 0) {
		this.editor = this.parent.panel.appendElement({ "tag": "select", "id": id }).down()

		if (enumarations.length > 0) {
			for (var i = 0; i < enumarations.length; i++) {
				this.editor.appendElement({ "tag": "option", "innerHtml": enumarations[i], "value": enumarations[i] })
			}
		} else {
			var values = fieldType.replace("enum:", "").split(":")
			// keep the value in the element.
			for (var i = 0; i < values.length; i++) {
				var text = values[i].substring(values[i].indexOf("_") + 1).replace("_", " ")
				this.editor.appendElement({ "tag": "option", "innerHtml": text, "value": text })
			}
		}
	} else if (fieldType.startsWith("xs.")) {
		if (isXsId(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "id": id }).down()
		} else if (isXsRef(fieldType)) {
			// Reference here... autocomplete...
			this.editor = this.parent.panel.appendElement({ "tag": "input", "id": id, "type": "url" }).down()
		} else if (isXsInt(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "1", "id": id }).down()
		} else if (isXsTime(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "datetime-local", "id": id }).down()
		} else if (isXsDate(fieldType)) {
			this.editor = this.parent.panel.appendElement({ "tag": "input", "type": "date", "id": id }).down()
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

				// In case of date i need to transform the value into a unix time.
				if (isXsTime(fieldPanel.fieldType) || isXsDate(fieldPanel.fieldType)) {
					value = moment(value).unix()
				}

				// Now I will modify the value in the local entity.
				var entity = entities[fieldPanel.parent.entityUuid]
				if (entity == undefined) {
					// The entity not already exist on the sever side.
					entity = fieldPanel.parent.entity
					entity[fieldPanel.fieldName] = value
					// if it has a parent, the parent will be save in that case.
					if (entity.getParent != undefined) {
						entity = entity.getParent()
					}

					if (entity.getPanel != undefined) {
						var panel = entity.getPanel()
						// If the panel is a table row...
						if (panel.table != undefined) {
							panel.table.getModel().saveValue(panel.index)
						}
					} else {
						// In that case the entity does not exist and are not in a table.
						server.entityManager.createEntity(fieldPanel.parent.entity.ParentUuid, fieldPanel.parent.entity.ParentLnk, fieldPanel.parent.entity,
							function (result, caller) {
							},
							// error callback
							function () {

							}, {})

					}
				} else {
					// Here the entity exist and the entity panel is use to edit it values.
					entity[fieldPanel.fieldName] = value
					// Remove the editor
					try {
						this.parentNode.removeChild(this)
					} catch (error) {

					}
					fieldPanel.editor = null // remove the edior.
					fieldPanel.value.element.style.display = ""

					// In that case I will save the entity directly.
					server.entityManager.saveEntity(entity,
						function (result, caller) {
						},
						// error callback
						function () {

						}, {})
				}


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
		} else if (this.editor.element.type == "date" || this.editor.element.type == "datetime-local") {
			if (isXsTime(this.parent.fieldType)) {
				if (isString(value)) {
					value = moment(value).unix()
				}
				this.editor.element.value = moment.unix(value).format("YYYY-MM-DDTHH:mm:ss")

			} else if (isXsDate(this.parent.fieldType)) {
				if (isString(value)) {
					value = moment(value).unix()
				}
				this.editor.element.value = moment(value).format('YYYY-MM-DD');
			}
		} else {
			this.editor.element.value = value
		}
	} else if (this.editor.element.tagName == "SELECT") {
		this.editor.element.selectedIndex = value - 1
	}
}

FieldEditor.prototype.clear = function () {

}