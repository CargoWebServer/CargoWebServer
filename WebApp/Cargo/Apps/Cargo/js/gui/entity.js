/**
 * Code that display an entity.
 */

/**
 * That map contain word use by the propertie panel
 * to display the name of a propertie in more readeable
 * fashion. If the name is not found in that list, the original
 * name is use instead... 
 */
var languageInfo = {
	"en": {
		"M_firstName": "First Name",
		"M_lastName": "Last name",
		"M_id": "Id",
		"M_name": "Name",
		"delete_dialog_entity_title": "Delete entity",
	},
	"fr": {
		"M_firstName": "Prénom",
		"M_lastName": "Nom",
		"M_id": "Id",
		"M_name": "Nom",
		"delete_dialog_entity_title": "Supprimer",
	}
}

// Set the language informtion here...
server.languageManager.appendLanguageInfo(languageInfo)

/**
 * That panel i use to navigate in information.
 * @param parent The parent element where the panel will be created.
 * @param typeName The type name of the entity to display
 * @param initCallback The callback function to be call when the entity is created.
 * @param parentEntityPanel If the entity is a sub-entity there is the parent entity panel.
 * @param removeOnDelete If is set to true the panel will be remove on delete event, otherwsie it will be clear.
 * @param parentEntity this is the parent entity of the entity display by the panel.
 * @param parentLnk The name of the reference in the parent entity to the entity of the panel.
 */
var EntityPanel = function (parent, typeName, initCallback, parentEntityPanel, removeOnDelete, parentEntity, parentLnk) {
	// Only know enity can have an panel.
	if (entityPrototypes[typeName] == undefined) {
		return
	}

	// if is set to true that will remove the panel when deleteted...
	this.removeOnDelete = removeOnDelete

	// The default value will be true.
	if (this.removeOnDelete == undefined) {
		this.removeOnDelete = true
	}

	// I will keep the type name
	this.typeName = typeName

	// the parent entity panel, can be undefined...
	this.parentEntityPanel = parentEntityPanel

	// the parent entity if there one.
	this.parentEntity = parentEntity

	// the parent link if there one.
	this.parentLnk = parentLnk

	// unique id
	this.id = randomUUID()

	// The panel where the entity is display.
	this.parent = parent

	// This contain a list of entity panel...
	this.subEntityPanel = null

	// The panel.
	this.panel = parent.appendElement({ "tag": "div", "class": "panel entities_panel" }).down()

	// the save button
	this.saveBtn = null

	// the save callback will be call after success
	this.saveCallback = null

	// The header serction.
	this.header = null

	// The entity display by the panel.
	this.entity = null

	// The prototype.
	this.proto = null

	// Reference to control for a given field
	this.controls = {}

	// The div here controls are display
	this.entitiesDiv = null

	// Now the navigation buttons...
	this.spacer = null

	// The navigations button (move up, move down)
	this.moveUp = null
	this.moveDown = null

	// save button
	this.saveBtn = null

	// the delete button
	this.deleteBtn = null

	// the delete callback will be call after success
	this.deleteCallback = null

	// Keep track of the init callback.
	this.initCallback = initCallback

	// Finish the intialysation...
	this.init(entityPrototypes[typeName], initCallback)

	// retur the pointer to the entity panel.
	return this
}

/**
 * Initialyse the entity for a given prototype.
 * @param {EntityPrototype} proto The entity prototype.
 * @param {function} initCallback The function to call after the initialysation is done.
 */
EntityPanel.prototype.init = function (proto, initCallback) {
	this.proto = proto
	this.controls = {}
	if (this.proto == null) {
		return
	}

	this.typeName = proto.TypeName

	// First I will clear the current element.
	//this.panel.removeAllChilds()
	this.panel.element.innerHTML = ""
	this.panel.childs = {}
	// Also remove sub entity panel here.
	if (this.subEntityPanel != null) {
		this.subEntityPanel.panel.element.innerHTML = ""
		this.subEntityPanel.panel.childs = {}
	}

	// Init the header.
	this.initHeader()

	for (var i = 0; i < proto.FieldsOrder.length; i++) {
		var index = proto.FieldsOrder[i]
		var field = proto.Fields[index]
		var fieldType = proto.FieldsType[index]
		var fieldVisibility = proto.FieldsVisibility[index]
		if (fieldVisibility == true) {
			this.initField(this.entitiesDiv, field, fieldType, proto.Restrictions)
		}
	}

	// Call after the initialisation....
	if (initCallback != undefined) {

		initCallback(this)
	}
}

/**
 * Hide the panel
 */
EntityPanel.prototype.hide = function () {
	this.panel.element.style.display = "none"
	if (this.subEntityPanel != null) {
		this.subEntityPanel.hide()
	}
}

/**
 * Show the panel
 */
EntityPanel.prototype.show = function () {
	if (this.subEntityPanel != null) {
		this.subEntityPanel.show()
	} else {
		this.panel.element.style.display = ""
	}
}

/**
 * Set the entity to display in the panel.
 * @param {Entity} entity The entity to display.
 */
EntityPanel.prototype.setEntity = function (entity) {
	if (entity == undefined) {
		return
	}

	if (this.typeName != entity.TYPENAME) {
		server.entityManager.getEntityPrototype(entity.TYPENAME, entity.TYPENAME.split(".")[0],
			// success callback
			function (proto, caller) {
				caller.entityPanel.init(proto, function (entity) {
					return function (entityPanel) {
						entityPanel.setEntity(entity)
						entityPanel.setTitle(entity.TYPENAME)
					}
				}(caller.entity))
			},
			// error callback
			function () {

			},
			{ "entityPanel": this, "entity": entity })
		return
	}

	// Set the panel id with the entity id.
	this.panel.element.id = entity.UUID

	// Here I will associate the panel and the entity.
	if (this.entity != null) {

		server.entityManager.detach(this, NewEntityEvent)
		server.entityManager.detach(this, UpdateEntityEvent)
		server.entityManager.detach(this, DeleteEntityEvent)

		this.entity.panel = null
		this.clear()
	}

	// Set the reference to the panel inside the entity.
	this.entity = entity
	this.entity.panel = this

	if (this.substitutionGroupSelect != undefined) {
		this.substitutionGroupSelect.element.style.display = "none"
	}

	/////////////////////////////////////////////////////////////////////////////////
	// Now the event listener...
	/////////////////////////////////////////////////////////////////////////////////

	// The delete entity event.
	server.entityManager.attach(this, DeleteEntityEvent, function (evt, entityPanel) {
		if (evt.dataMap["entity"].TYPENAME == entityPanel.typeName) {
			if (evt.dataMap["entity"] != undefined && entityPanel.entity != null) {
				if (evt.dataMap["entity"].UUID == entityPanel.entity.UUID) {
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
			if (evt.dataMap["entity"] && entityPanel.entity != null) {
				if (entityPanel.entity.UUID == evt.dataMap["entity"].UUID) {
					entityPanel.init(entityPanel.proto)
					entityPanel.setEntity(entities[evt.dataMap["entity"].UUID])
				}
			}
		}
	})

	// The update entity event.
	server.entityManager.attach(this, UpdateEntityEvent, function (evt, entityPanel) {
		if (evt.dataMap["entity"].TYPENAME == entityPanel.typeName) {
			if (evt.dataMap["entity"] && entityPanel.entity != null) {
				// I will reinit the panel here...
				if (entityPanel.entity.UUID == evt.dataMap["entity"].UUID) {
					entityPanel.setEntity(entities[evt.dataMap["entity"].UUID])
				}
			}
		}
	})


	// So here I will set the propertie of the object.
	for (var i = 0; i < this.proto.FieldsOrder.length; i++) {
		var index = this.proto.FieldsOrder[i]
		var field = this.proto.Fields[index]
		var fieldType = this.proto.FieldsType[index]
		var fieldVisibility = this.proto.FieldsVisibility[index]

		if (fieldVisibility == true) {
			var control = this.controls[this.proto.TypeName + "_" + field]
			var value = this.entity[field]
			if (control != null) {
				if (control.constructor.name == "EntityPanel") {
					control.parentEntity = entity
				}
			}

			if (value != null && control != null && value != "") {
				if (control.setFieldValue == undefined) {
					if (fieldType == "xs.base64Binary") {
						this.setGenericFieldValue(control, field, value, entity.UUID)
					} else {
						this.setFieldValue(control, field, fieldType, value, entity.UUID)
					}
				} else {
					// Here the control is a entity panel so i will redirect 
					// the entity to the control.
					if (value.TYPENAME != undefined) {
						control.setEntity(value)
					}
				}
			} else if (control == null) {
				console.log("No control found for display value " + value + " with type name " + fieldType)
			} else if (value == null || value == "") {
				console.log("The value is null or empty.")
			}
		}

		// Display the append ref button.
		var plusBtn = this.panel.getChildById(this.proto.TypeName + "_" + fieldType.replace(":Ref", "").replace("[]", "") + "_" + field + "_plus_btn")
		if (plusBtn != null) {
			plusBtn.element.style.display = "table-cell"
		}

		var newInput = this.panel.getChildById(this.proto.TypeName + "_" + fieldType.replace(":Ref", "").replace("[]", "") + "_" + field + "_new")
		if (newInput != null) {
			newInput.element.style.display = "none"
		}
	}
	this.saveBtn.element.id = entity.UUID + "_save_btn"
	this.saveBtn.element.style.display = ""
	this.deleteBtn.element.style.display = "table-cell"
}

/**
 * Reset the content of the panel.
 */
EntityPanel.prototype.clear = function () {
	if(this.entity == null){
		return
	}
	
	this.entity = null
	this.init(this.proto)
	if (this.initCallback != undefined) {
		this.initCallback(this)
	}

	this.maximizeBtn.element.click()
	this.saveBtn.element.style.display = ""
	this.deleteBtn.element.style.display = "none"

	if (this.substitutionGroupSelect != undefined) {
		this.substitutionGroupSelect.element.style.display = "table-cell"
	}

}

/**
 * return true if the panel not display an entity.
 */
EntityPanel.prototype.isEmpty = function () {
	return this.entity == null
}

/**
 * This will display the entity type...
 */
EntityPanel.prototype.initHeader = function () {

	// Set the header section.
	this.header = this.panel.appendElement({ "tag": "div", "class": "entities_panel_header" }).down()
		.appendElement({ "tag": "div" }).down()

	this.entitiesDiv = this.panel.appendElement({ "tag": "div", "class": "entity_panel" }).down()
		.appendElement({ "tag": "div" }).down()

	this.maximizeBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled", "style": "display: none;" }).down()
	this.maximizeBtn.appendElement({ "tag": "i", "class": "fa fa-plus-square-o" })

	this.minimizeBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled", "style": "display: table-cell;" }).down()
	this.minimizeBtn.appendElement({ "tag": "i", "class": "fa fa-minus-square-o" })

	this.maximizeBtn.element.onclick = function (entitiesDiv, minimizeBtn) {
		return function () {
			this.style.display = "none"
			entitiesDiv.element.style.display = "table-row"
			minimizeBtn.element.style.display = "table-cell"
			fireResize()
		}
	}(this.entitiesDiv, this.minimizeBtn)

	this.minimizeBtn.element.onclick = function (entitiesDiv, maximizeBtn) {
		return function () {
			this.style.display = "none"
			entitiesDiv.element.style.display = "none"
			maximizeBtn.element.style.display = "table-cell"
			fireResize()
		}
	}(this.entitiesDiv, this.maximizeBtn)
	this.minimizeBtn.element.click()

	// The save button.
	this.saveBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled" }).down()
	this.saveBtn.appendElement({ "tag": "i", "class": "fa fa-save" })

	this.saveBtn.element.addEventListener("click", function (entityPanel) {
		return function () {
			this.style.display = "none"
			// Here I will save the entity...
			entityPanel.entity.NeedSave = true
			if (entityPanel.entity != null) {
				entityPanel.saveBtn.element.id = entityPanel.entity.UUID + "_save_btn"
				if (entityPanel.entity.exist == false && entityPanel.parentEntity != null) {
					// The parent entity is know and the entity does not exist.
					server.entityManager.createEntity(entityPanel.parentEntity.UUID, entityPanel.parentLnk, entityPanel.entity.TYPENAME, entityPanel.entity.UUID, entityPanel.entity,
						// Success callback
						function (entity, entityPanel) {
							entityPanel.setEntity(entity)
							if (entityPanel.saveCallback != undefined) {
								entityPanel.saveCallback(entity)
								if (entityPanel.parentEntity != null) {
									if (entityPanel.parentEntity.panel != null) {
										if (entityPanel.parentEntity.panel.saveCallback != undefined) {
											entityPanel.parentEntity.panel.saveCallback(entityPanel.parentEntity)
										}
									}
								}
							}
						},
						// Error callback.
						function (result, caller) {

						}, entityPanel)
				} else {
					// Here the entity will be created of save...
					server.entityManager.saveEntity(entityPanel.entity,
						// Success callback
						function (entity, entityPanel) {
							if (entityPanel.saveCallback != undefined) {
								entityPanel.saveCallback(entity)
								if (entityPanel.parentEntity != null) {
									if (entityPanel.parentEntity.panel != null) {
										if (entityPanel.parentEntity.panel.saveCallback != undefined) {
											entityPanel.parentEntity.panel.saveCallback(entityPanel.parentEntity)
										}
									}
								}
							}
						},
						// Error callback.
						function () {

						}, entityPanel)
				}

			}
		}
	}(this))

	// The remove button.
	this.deleteBtn = this.header.appendElement({ "tag": "div", "class": "entities_header_btn enabled", "style": "display: none;" }).down()
	this.deleteBtn.appendElement({ "tag": "i", "class": "fa fa-trash" })

	this.deleteBtn.element.onclick = function (entityPanel) {
		return function () {
			// Here I will save the entity...
			if (entityPanel.entity != null) {
				// Here I will ask the user if here realy want to remove the entity...
				var confirmDialog = new Dialog(randomUUID(), undefined, true)
				confirmDialog.div.element.style.maxWidth = "450px"
				confirmDialog.setCentered()
				server.languageManager.setElementText(confirmDialog.title, "delete_dialog_entity_title")
				confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete entity " + entityPanel.entity.UUID + "?" })

				confirmDialog.ok.element.onclick = function (dialog, entityPanel) {
					return function () {
						// I will call delete file
						server.entityManager.removeEntity(entityPanel.entity.UUID,
							// Success callback 
							function (result, caller) {
								/** The action will be done in the event listener */
								if (caller.deleteCallback != undefined) {
									caller.deleteCallback(caller.entity)
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
	}(this)

	// Set the title div, the type is the default title.
	var typeDiv = this.header.appendElement({ "tag": "div", "class": "entity_type" }).down()
	var titleDiv = typeDiv.appendElement({ "tag": "div", "style": "display: table-row;" }).down()
	var titleSpan = titleDiv.appendElement({ "tag": "span", "style": "display: table-cell;", "id": this.typeName }).down()

	// Now I will set the list of substitution group.
	if (this.proto.SubstitutionGroup != undefined) {
		this.substitutionGroupSelect = titleDiv.appendElement({ "tag": "select", "style": "display: none;" }).down()

		if (this.proto.SubstitutionGroup.length > 0) {

			var substitutionGroup = this.proto.SubstitutionGroup.sort()
			if (substitutionGroup.indexOf("") == -1) {
				substitutionGroup.unshift("")
			}

			// Here I will append the list of substitution group in the result...
			for (var i = 0; i < substitutionGroup.length; i++) {
				this.substitutionGroupSelect.appendElement({ "tag": "option", "value": substitutionGroup[i], "innerHtml": substitutionGroup[i] })
			}

			this.substitutionGroupSelect.element.style.display = "table-cell"
			this.substitutionGroupSelect.element.onchange = function (entityPanel) {
				return function () {
					// In that case I will set the entity content with the given type.
					server.entityManager.getEntityPrototype(this.value, this.value.split(".")[0],
						function (proto, entityPanel) {
							entityPanel.init(proto, function (entityPanel) {
								entityPanel.maximizeBtn.element.click()
								entityPanel.setTitle(proto.TypeName)
							})
						},
						function () {

						}, entityPanel)
				}
			}(this)
		}
	}

	// Back to the sypertype 
	if (this.proto.SuperTypeNames != undefined) {
		if (this.proto.SuperTypeNames.length > 0) {
			var backButon = titleDiv.prependElement({ "tag": "div", "class": "entities_btn" }).down()
				.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-left entities_header_btn", "style": "padding-bottom: 3px;" }).down()

			backButon.element.onclick = function (superTypeName, entityPanel) {
				return function () {
					server.entityManager.getEntityPrototype(superTypeName, superTypeName.split(".")[0],
					function (proto, entityPanel) {
						entityPanel.init(proto, function (entityPanel) {
							entityPanel.maximizeBtn.element.click()
							entityPanel.setTitle(proto.TypeName)
						})
					},
					function () {

					}, entityPanel)
				}
			}(this.proto.SuperTypeNames[this.proto.SuperTypeNames.length - 1], this)
		}
	}

	this.spacer = this.header.appendElement({ "tag": "div", "style": "display: table-cell; width: 100%;" }).down()
	this.moveUp = this.header.appendElement({ "tag": "div", "class": "entities_header_btn", "style": "display: table-cell;" }).down()
	this.moveUp.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-left entities_header_btn" })

	this.moveDown = this.header.appendElement({ "tag": "div", "class": "entities_header_btn", "style": "display: table-cell;" }).down()
	this.moveDown.appendElement({ "tag": "i", "class": "fa fa-caret-square-o-right entities_header_btn" })

	if (this.parentEntityPanel != null) {
		this.moveUp.element.className += " enabled"
		this.moveUp.element.firstChild.style.color = "white"
		// Now the action...
		this.moveUp.element.onclick = function (entityPanel) {
			return function () {
				entityPanel.panel.element.style.display = "none"
				entityPanel.parentEntityPanel.panel.element.style.display = ""
			}
		}(this)
	}
}

/**
 * Hide the navigation button...
 */
EntityPanel.prototype.hideNavigationButtons = function () {
	this.spacer.element.style.display = "none"
	this.moveUp.element.style.display = "none"
	this.moveDown.element.style.display = "none"
}

/**
 * Show the navigation button.
 */
EntityPanel.prototype.showNavigationButtons = function () {
	this.spacer.element.style.display = ""
	this.moveUp.element.style.display = ""
	this.moveDown.element.style.display = ""
}

/**
 * Return the label for a given field name.
 */
EntityPanel.prototype.getFieldLabel = function (fieldName) {
	var label = this.panel.getChildById(this.proto.TypeName + "_" + fieldName + "_lbl")
	return label
}

/**
 * Retreive a field control with a given id.
 */
EntityPanel.prototype.getFieldControl = function (fieldName) {
	var control = this.panel.getChildById(this.proto.TypeName + "_" + fieldName)
	return control
}

/**
 * Set the title.
 */
EntityPanel.prototype.setTitle = function (title) {
	var titleDiv = this.panel.getChildById(this.typeName)
	titleDiv.element.innerHTML = title
}

/**
 * Create a control for a XSD type.
 */
EntityPanel.prototype.createXsControl = function (id, valueDiv, field, fieldType, restrictions) {
	// Create a select box from a xs restriction.
	function appendSelect(restrictions, id, valueDiv) {
		for (var i = 0; i < restrictions.length; i++) {
			var restriction = restrictions[i]
			if (restriction.Type == 1) {
				if (control == undefined) {
					control = valueDiv.appendElement({ "tag": "select", "id": id }).down()
				}
				control.appendElement({ "tag": "option", "value": restriction.Value, "innerHtml": restriction.Value })
			}
		}
		return control
	}

	var control = null
	var isEnum = false
	if (restrictions.length > 0) {
		isEnum = restrictions[0].Type == 1
	}

	if (isEnum) { // The restriction represent a list of values here.
		if (restrictions[0].Type == 1) {
			control = appendSelect(restrictions, id, valueDiv)
		}
	} else {
		if (isXsId(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "id": id }).down()
		} else if (isXsRef(fieldType)) {
			// Reference here... autocomplete...
			control = valueDiv.appendElement({ "tag": "input", "id": id }).down()
		} else if (isXsInt(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "1", "id": id }).down()
		} else if (isXsDate(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "type": "date", "id": id }).down()
		} else if (isXsTime(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "type": "datetime-local", "id": id }).down()
		} else if (isXsString(fieldType)) {
			if (field == "M_description") {
				valueDiv.element.style.verticalAlign = "top"
				control = valueDiv.appendElement({ "tag": "textarea", "id": id }).down()
			} else {
				control = valueDiv.appendElement({ "tag": "input", "id": id }).down()
				if (field.indexOf("pwd") > -1) {
					control.element.type = "password"
				}
			}
		} else if (isXsBinary(fieldType)) {

		} else if (isXsBoolean(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "type": "checkbox", "id": id }).down()
		} else if (isXsNumeric(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "0.01", "id": id }).down()
		} else if (isXsMoney(fieldType)) {
			control = valueDiv.appendElement({ "tag": "input", "type": "number", "min": "0", "step": "0.01", "id": id }).down()
		} else {

		}
	}
	return control
}

/**
 * Initialyse the entity panel content.
 *  @param {string} parent The parent element where the panel belong to.
 *  @param {string} field The name of property 
 *  @param {string} fieldType The type name of the property
 */
EntityPanel.prototype.initField = function (parent, field, fieldType, restrictions) {

	// The entity div.
	var entityDiv = parent.appendElement({ "tag": "div", "class": "entity" }).down()
	var id = this.proto.TypeName + "_" + field


	// In case of a generic value....
	if (fieldType == "xs.base64Binary") {
		this.controls[id] = parent
		return
	}

	// The entity label here...
	var label = entityDiv.appendElement({ "tag": "div", id: this.proto.TypeName + "_" + field + "_lbl" }).down()
	label.appendElement({ "tag": "span", "innerHtml": field.replace("M_", ""), "id": field.replace("M_", "") }).down()

	server.languageManager.setElementText(label, field)

	// Now the entity value...
	var valueDiv = entityDiv.appendElement({ "tag": "div" }).down()
	var control = null
	var isArray = fieldType.startsWith("[]")
	var isRef = fieldType.endsWith(":Ref")

	fieldType = fieldType.replace("[]", "").replace(":Ref", "")
	var prototype = entityPrototypes[fieldType]

	// if there is no restriction
	if (restrictions == undefined) {
		restrictions = []
	}

	var baseType = getBaseTypeExtension(fieldType)
	// If the value is simple string...
	if (!isArray && (isXsBaseType(baseType) || fieldType.startsWith("sqltypes.") || fieldType.startsWith("xs.") || fieldType.startsWith("enum:"))) {
		// in case of an enum
		if (fieldType.startsWith("enum:")) {
			var values = fieldType.replace("enum:", "").split(":")
			for (var i = 0; i < values.length; i++) {
				restrictions.push({ "Type": 1, "Value": values[i] })
			}
		} else {
			var fieldPrototype = entityPrototypes[fieldType]
			if (fieldPrototype != undefined) {
				if (fieldPrototype.Restrictions != undefined) {
					// Set the restriction here.
					restrictions = fieldPrototype.Restrictions
				}
			}
		}
		control = this.createXsControl(id, valueDiv, field, fieldType, restrictions)
	} else if (!isArray && (fieldType == "interface{}")) {
		// Here it can be plain text or xml text... 
		control = valueDiv.appendElement({ "tag": "textarea", "id": id }).down()
	} else {
		// Here it must be a reference to other entity...
		var newLnkButton = null
		if (isArray) {
			// Here I will create the add button...
			newLnkButton = label.appendElement({ "tag": "div", "class": "entities_btn", "style": "display: none;", "id": this.proto.TypeName + "_" + fieldType + "_" + field + "_plus_btn" }).down()
			newLnkButton.appendElement({ "tag": "i", "class": "fa fa-plus" }).down()
			newLnkButton.element.style.verticalAlign = "middle"
			if (!isRef) {
				var itemTable = undefined
				if (field != "M_listOf" && !fieldType.startsWith("xs.")) {
					// I will create the entity table.
					var itemPrototype = entityPrototypes[fieldType.replace("[]", "")]
					var itemsTableModel = new EntityTableModel(itemPrototype)
					var itemTable = new Table(randomUUID(), valueDiv)

					itemTable.setModel(itemsTableModel, function (table) {
						return function () {
							table.init()
							table.refresh()
						}
					}(itemTable))

				} else {
					var tableModel = new TableModel(["index", "values"])
					tableModel.editable[1] = true

					var baseType = fieldType
					if (!baseType.startsWith('xs.') && !baseType.startsWith('sqltypes.')) {
						baseType = getBaseTypeExtension(fieldType)
					}

					tableModel.fields = ["xs.int", baseType.replace("[]", "")]
					itemTable = new Table(randomUUID(), valueDiv)
					itemTable.setModel(tableModel, function () { })
				}
				control = itemTable
			} else {
				// Simply create a div where the list of hyper link will be display.
				control = valueDiv.appendElement({ "tag": "div", "class": "scrolltable", "id": id }).down()
			}
		} else {
			if (isRef) {
				// Here i will create the edit button.
				newLnkButton = label.appendElement({ "tag": "div", "class": "entities_btn", "style": "display: none;", "id": this.proto.TypeName + "_" + fieldType + "_" + field + "_plus_btn" }).down()
				newLnkButton.appendElement({ "tag": "i", "class": "fa fa-pencil-square-o" }).down()
				control = valueDiv.appendElement({ "tag": "div", "id": id }).down()
			} else {
				// Here I have a item inside another item...
				var subentityPanel = new EntityPanel(valueDiv, fieldType, function () { }, undefined, true, this.entity, field)
				// The control will be the sub-entity panel.
				this.controls[id] = subentityPanel
			}
		}

		if (newLnkButton != null) {
			newLnkButton.element.onclick = function (entityPanel, field, fieldType, valueDiv) {
				return function () {
					if (isRef) {
						// So here I will create an auto completion list panel...
						if (entityPanel.controls[id + "_new"] != undefined) {
							// remove from the layout
							entityPanel.controls[id + "_new"].element.parentNode.removeChild(entityPanel.controls[id + "_new"].element)
						}

						if (isArray) {
							entityPanel.controls[id + "_new"] = valueDiv.appendElement({ "tag": "input", "id": id + "_new" }).down()
						} else {
							var parentElement
							if (valueDiv.element.firstChild.firstChild == undefined) {
								parentElement = valueDiv.element
							} else {
								parentElement = valueDiv.element.firstChild.firstChild.firstChild
							}

							entityPanel.controls[id + "_new"] = new Element(parentElement, { "tag": "input", "id": id + "_new" })
							if (parentElement.firstChild != null) {
								parentElement.firstChild.style.display = "none"
								entityPanel.controls[id + "_new"].element.value = parentElement.firstChild.innerHTML
							}
						}

						// Display it
						entityPanel.controls[id + "_new"].element.style.display = ""

						// Get the pacakge prototypes...
						var typeName = fieldType.replace("[]", "").replace(":Ref", "")
						server.entityManager.getEntityPrototypes(typeName.split(".")[0],
							function (result, caller) {

								// Set variables...
								var entityPanel = caller.entityPanel
								var id = caller.id
								var field = caller.field
								var prototype = entityPrototypes[caller.typeName]

								// Now i will set it autocompletion list...
								attachAutoCompleteInput(entityPanel.controls[id + "_new"], fieldType, field, entityPanel, prototype.getTitles(),
									function (entityPanel, field, fieldsType) {
										return function (value) {
											// Here I will set the field of the entity...
											if (entityPanel.entity != undefined) {
												// Set the new object value.
												appendObjectValue(entityPanel.entity, field, value)
												if (entityPanel.entity.UUID != "") {
													// Automatically saved...
													server.entityManager.saveEntity(entityPanel.entity)
												} else {
													// Here the entity dosent exist...
													server.entityManager.createEntity(entityPanel.entity.ParentUuid, entityPanel.entity.parentLnk, entityPanel.entity.TYPENAME, "", entityPanel.entity,
														function (result, caller) {
															// Set the result inside the panel.
															entityPanel.setEntity(result)
														},
														function () {

														}, entityPanel)
												}
											}
										}
									}(entityPanel, field))

								entityPanel.controls[id + "_new"].element.focus()
								entityPanel.controls[id + "_new"].element.select();
							},
							// The error callback.
							function () {
							}, { "entityPanel": entityPanel, "field": field, "id": id, "typeName": typeName })

					} else {
						// In that case I will create a new entity
						if (itemPrototype != undefined) {
							var item = eval("new " + itemPrototype.TypeName + "()")
							item.TYPENAME = itemPrototype.TypeName

							// Set the parent uuid.
							item.ParentUuid = entityPanel.entity.UUID
							item.parentLnk = field

							if (isArray) {
								var itemTable = entityPanel.controls[id]
								var row = itemTable.appendRow(item, item.UUID)
								row.saveBtn.element.style.visibility = "visible"
							} else {

							}
						} else {
							// Here I will try to append a new value inside the table...
							if (isArray) {
								var itemTable = entityPanel.controls[id]
								var newRow = itemTable.appendRow([entityPanel.entity[field].length + 1, "0"], entityPanel.entity[field].length)
								newRow.saveBtn.element.style.visibility = "visible"

								simulate(newRow.cells[entityPanel.entity[field].length, 1].div.element, "dblclick");

								newRow.deleteBtn.element.onclick = function (entity, field, row) {
									return function () {
										// Here I will simply remove the element 
										// The entity must contain a list of field...
										if (entity[field] != undefined) {
											entity[field].splice(row.index, 1)
											entity.NeedSave = true
											server.entityManager.saveEntity(entity)
										}
									}
								}(entityPanel.entity, field, newRow)

								// The save row action
								newRow.saveBtn.element.onclick = function (entity, field, row) {
									return function () {
										// Here I will simply remove the element 
										// The entity must contain a list of field...
										if (entity[field] != undefined) {
											entity[field][row.index] = row.table.model.getValueAt(row.index, 1)
											entity.NeedSave = true
											server.entityManager.saveEntity(entity,
												function (result, caller) {
													caller.style.visibility = "hidden"
												},
												function () {

												}, this)
										}
									}
								}(entityPanel.entity, field, newRow)

								//itemTable.refresh()
							}
						}

					}
				}
			}(this, field, fieldType, valueDiv)
		}
	}

	if (control != null) {
		this.controls[id] = control
		if (control.element != null) {
			// Now the change event...
			control.element.addEventListener("change", function (entityPanel, attribute, fieldType) {
				return function () {
					var entity = null
					if (entityPanel.entity == null) {
						// Here I will create a new entity.
						var entity = eval("new " + entityPanel.typeName)

						// set basic values.
						if (entityPanel.parentEntity != null) {
							entity.ParentUuid = entityPanel.parentEntity.UUID
							entity.parentLnk = entityPanel.parentLnk
						}

					} else {
						// Get the existing entity.
						entity = entityPanel.entity
					}

					/** Here it's a string **/
					var baseType = getBaseTypeExtension(fieldType)
					if (isXsString(baseType) || isXsId(fieldType) || isXsString(fieldType || isXsRef(fieldType))) {
						if (entity[attribute] != this.value) {
							entity[attribute] = this.value
							entity.NeedSave = true
						}
					} else if (isXsNumeric(fieldType)) {
						if (entity[attribute] != parseFloat(this.value)) {
							entity[attribute] = parseFloat(this.value)
							entity.NeedSave = true
						}
					} else if (isXsInt(fieldType)) {
						if (entity[attribute] != parseInt(this.value)) {
							entity[attribute] = parseInt(this.value)
							entity.NeedSave = true
						}
					} else if (isXsTime(fieldType)) {
						var value = moment(this.value).unix()
						if (entity[attribute] != value) {
							entity[attribute] = value
							entity.NeedSave = true
						}
					} else if (isXsBoolean(fieldType)) {
						if (entity[attribute] != this.checked) {
							entity[attribute] = this.checked
							entity.NeedSave = true
						}
					} else if (fieldType.startsWith("enum:")) {
						if (entity[attribute] != this.selectedIndex + 1) {
							entity[attribute] = this.selectedIndex + 1
							entity.NeedSave = true
						}
					} else {
						// The field is a reference...

					}

					// Display the save button if the entity has changed.
					if (entity.NeedSave) {
						entityPanel.saveBtn.element.style.display = "table-cell"
					}

					// Set the entity in that case.
					if (entityPanel.entity == null) {
						entityPanel.setEntity(entity)
						// Keep on the entity manager.
						server.entityManager.setEntity(entity)
					}
				}
			}(this, field, fieldType))

			// Append the listener to display the save button.
			control.element.addEventListener("keyup", function (entityPanel, field) {
				return function () {
					if (entityPanel.entity != null) {
						if (this.value.length > 0) {
							if (entityPanel.entity[field] != this.value) {
								entityPanel.saveBtn.element.style.display = "table-cell"
							}
						} else {
							// reset the panel value.
							entityPanel.clear()

							// Set back to the most generic type.
							if (entityPanel.proto.SuperTypeNames != undefined) {
								if (entityPanel.proto.SuperTypeNames.length > 0) {
									var typeName = entityPanel.proto.SuperTypeNames[0]
									server.entityManager.getEntityPrototype(typeName, typeName.split(".")[0],
										// success callback
										function (proto, caller) {
											caller.entityPanel.init(proto, function (proto) {
												return function (entityPanel) {
													entityPanel.setTitle(proto.TypeName)
												}
											}(proto))
										},
										function () {

										},
										{ "entityPanel": entityPanel })
								}
							}

						}
					}
				}
			}(this, field))
		}

		// if the field is an index key i will set the auto complete on it...
		if ((contains(this.proto.Indexs, field) || contains(this.proto.Ids, field)) && !isArray) {
			attachAutoCompleteInput(control, this.typeName, field, this, [],
				function (entityPanel) {
					return function (value) {
						if (entityPanel.entity != undefined) {
							if (entityPanel.entity.UUID != value.UUID) {
								entityPanel.setEntity(value)
							}
						} else {
							entityPanel.setEntity(value)
						}
					}
				}(this))
		}
	}
}

/**
 * That function is use to display a field when there is no hint about the way to display it,
 * so introspection will be use instead of entity prototype.
 */
EntityPanel.prototype.setGenericFieldValue = function (control, field, value, parentUuid) {

	// First I will get the parent entity.
	var parentEntity = entities[parentUuid]
	var id = parentEntity.TYPENAME + "_" + field
	var fieldType

	if (isString(value)) {
		// Here I will create a xs control...
		fieldType = "xs.string"
	} else {
		// Here I will create a xs control to display an object reference.
		if (!isArray(value)) {
			fieldType = value.TYPENAME + ":Ref"
			setRef(parentEntity, field, value, false)
		} else {
			for (var i = 0; i < value.length; i++) {
				// In the case of an array...
				if (isObject(value[i])) {
					fieldType = "[]" + value[0].TYPENAME + ":Ref"
					setRef(parentEntity, field, value, true)
				} else {
					if (isString(value[i])) {
						// an array of string...
						fieldType = "[]xs.string"
					}
					// TODO implement other field type here.
				}
			}
		}

	}

	this.initField(control, field, fieldType, [])
	this.setFieldValue(this.controls[id], field, fieldType, value, parentUuid)
}

/**
 * (Multiple value) If the object is a sub-object I will display a table with the value of the sub-object.
 * Composition
 */
EntityPanel.prototype.appendObjects = function (itemsTable, values, field, fieldType, parentUuid) {
	if (itemsTable != undefined) {
		itemsTable.clear()
		for (var i = 0; i < values.length; i++) {
			// keep information about the parent entity...
			values[i].ParentUuid = this.entity.UUID
			values[i].parentLnk = field
			if (values[i].UUID != undefined) {
				var row = itemsTable.appendRow(values[i], values[i].UUID)
			} else {
				// Append a row with a single value inside it...
				var row = itemsTable.appendRow([i + 1, values[i]], i)

				// The delete row action...
				row.deleteBtn.element.onclick = function (entity, field, row) {
					return function () {
						// Here I will simply remove the element 
						// The entity must contain a list of field...
						if (entity[field] != undefined) {
							entity[field].splice(row.index, 1)
							entity.NeedSave = true
							server.entityManager.saveEntity(entity)
						}
					}
				}(entities[parentUuid], field, row)

				// The save row action
				row.saveBtn.element.onclick = function (entity, field, row) {
					return function () {
						// Here I will simply remove the element 
						// The entity must contain a list of field...
						if (entity[field] != undefined) {
							entity[field][row.index] = row.table.model.getValueAt(row.index, 1)
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
				}(entities[parentUuid], field, row)
			}
		}
	}
}

/**
 * Same as appendObject for a single value.
 */
EntityPanel.prototype.appendObject = function (object, valueDiv, field, fieldType) {
	var subEntityPanel = new EntityPanel(valueDiv, fieldType,
		// The init callback. 
		function (value) {
			return function (panel) {
				panel.setEntity(value)
			}
		}(object),
		undefined, true, object, field)
}

/**
 * If the the object is a reference I will display a link to the other oject entity panel.
 * Aggregation.
 */
EntityPanel.prototype.appendObjectRef = function (object, valueDiv, field, fieldType) {

	var prototype = entityPrototypes[object.TYPENAME]
	var titles = object.getTitles()
	var refName = ""
	for (var j = 0; j < titles.length; j++) {

		refName += titles[j]
		if (j < titles.length - 1) {
			refName += " "
		}
	}

	// Here the object must be init...
	if (refName != undefined && refName.length > 0) {
		valueDiv.element.style.width = "auto"
		var ln = valueDiv.appendElement({ "tag": "div", "class": "entities_btn_container" }).down()
		var ref = ln.appendElement({ "tag": "div" }).down().appendElement({ "tag": "a", "href": "#", "title": object.TYPENAME, "innerHtml": refName }).down()
		ref.element.id = object.UUID
		var deleteLnkButton = ln.appendElement({ "tag": "div", "class": "entities_btn" }).down().appendElement({ "tag": "i", "class": "fa fa-trash" }).down()

		// Now the action...
		ref.element.onclick = function (object, propertiePanel) {
			return function () {
				// alert(object.UUID)
				// so here I will get the entity value and display it on the server...
				propertiePanel.panel.element.style.display = "none"

				// In case of cyclic reference...
				var parentEntityPanel = propertiePanel.parentEntityPanel
				while (parentEntityPanel != undefined) {
					if (parentEntityPanel.typeName == object.TYPENAME) {
						break
					} else {
						if (parentEntityPanel != propertiePanel.parentEntityPanel) {
							parentEntityPanel = propertiePanel.parentEntityPanel
						} else {
							parentEntityPanel = undefined
						}
					}
				}

				var subentityPanel
				if (parentEntityPanel != undefined) {
					subPropertiePanel = parentEntityPanel
					parentEntityPanel.panel.element.style.display = ""
					var entity = entities[object.UUID]
					if (entity == undefined) {
						entity = object
					}
					parentEntityPanel.setEntity(entity)
				} else {
					subPropertiePanel = new EntityPanel(propertiePanel.parent, object.TYPENAME, function (object) {
						return function (panel) {
							// Set the object.
							var entity = entities[object.UUID]
							if (entity == undefined) {
								entity = object
							}
							panel.setEntity(entity)
							panel.setTitle(object.TYPENAME)

						}
					}(object), propertiePanel)
				}

				if (propertiePanel.subEntityPanel != null) {
					// Remove the existing panel here.
					propertiePanel.subEntityPanel.parent.removeElement(propertiePanel.subEntityPanel.panel)
				}

				propertiePanel.subEntityPanel = subPropertiePanel

				propertiePanel.moveDown.element.className += " enabled"
				propertiePanel.moveDown.element.firstChild.style.color = "white"

				// Now the action...
				propertiePanel.moveDown.element.onclick = function (entityPanel) {
					return function () {
						entityPanel.panel.element.style.display = "none"
						entityPanel.subEntityPanel.panel.element.style.display = ""
					}
				}(propertiePanel)

			}
		}(object, this)

		ref.element.onmouseover = function (object) {
			return function () {
				var toActivated = []
				if (document.getElementById(object.UUID) != undefined) {
					toActivated.push(document.getElementById(object.UUID))
				}

				var toActivated_ = document.getElementsByName(object.UUID)
				for (var i = 0; i < toActivated_.length; i++) {
					toActivated.push(toActivated_[i])
				}

				for (var i = 0; i < toActivated.length; i++) {
					if (toActivated[i].className.baseVal != undefined) {
						toActivated[i].className.baseVal += " active"
					}
				}
			}
		}(object)

		ref.element.onmouseout = function (object) {
			return function () {

				var toDectivated = []
				if (document.getElementById(object.UUID) != undefined) {
					toDectivated.push(document.getElementById(object.UUID))
				}

				var toDectivated_ = document.getElementsByName(object.UUID)
				for (var i = 0; i < toDectivated_.length; i++) {
					toDectivated.push(toDectivated_[i])
				}

				for (var i = 0; i < toDectivated.length; i++) {
					if (toDectivated[i].className.baseVal != undefined) {
						toDectivated[i].className.baseVal = toDectivated[i].className.baseVal.replace(" active", "")
					}
				}
			}
		}(object)

		deleteLnkButton.element.onclick = function (entityUUID, object, field) {
			return function () {
				var entity = entities[entityUUID]
				removeObjectValue(entity, field, object)
				server.entityManager.saveEntity(entity,
					function (result, caller) {
						// nothing to do here.
					},
					function (result, caller) {
						// nothing to here.
					},
					undefined)
				// in case of local object.
				if (entity.onChange != undefined) {
					entity.onChange(entity)
				}
			}
		}(this.entity.UUID, object, field)
	}

}

/**
 * Exit edit mode for a given field.
 */
EntityPanel.prototype.resetFieldValue = function (field, input) {
	if (!isArray(this.entity[field])) {
		input.element.parentNode.firstChild.style.display = ""
	}
	input.element.parentNode.removeChild(input.element)
	input.autocompleteDiv.element.parentNode.removeChild(input.autocompleteDiv.element)
	delete this.controls[input.element.id]
}

EntityPanel.prototype.setFieldValue = function (control, field, fieldType, value, parentUuid) {
	if (control == undefined) {
		return
	}

	// In case of a reference string...
	if (fieldType == "xs.string") {
		if (isObjectReference(value)) {
			// In that case I will create the link object to reach the 
			// reference...
			fieldType = value.split("%")[0] + ":Ref"

			// Here I will remove the default control and replace it by a div where the reference will be placed.
			control.element.parentNode.removeChild(control.element)
			control = control.parentElement.appendElement({ "tag": "div" }).down()
		}
	}

	this.maximizeBtn.element.click()
	// Here the entity is a reference to a complex type...
	var isRef = fieldType.endsWith(":Ref")

	// Here I will see if the type is derived basetype...
	if (!fieldType.startsWith("[]") && !isRef) {
		var baseType = getBaseTypeExtension(fieldType)
		if (fieldType.startsWith("enum:")) {
			// Here the value is an enumeration...
			control.element.selectedIndex = parseInt(value) - 1
		} else if (isXsString(baseType) || isXsString(fieldType) || fieldType == "interface{}") {
			control.element.value = value
		} else if (isXsNumeric(fieldType)) {
			if (value != "") {
				control.element.value = parseFloat(value)
			} else {
				control.element.value = ""
			}
		} else if (isXsInt(fieldType)) {
			if (value != "") {
				control.element.value = parseInt(value)
			} else {
				control.element.value = ""
			}
		} else if (isXsBoolean(fieldType)) {
			control.element.checked = value
		} else if (isXsDate(fieldType)) {
			if (value != "") {
				control.element.value = moment(value).format('YYYY-MM-DD');
			} else {
				control.element.value = ""
			}
		} else if (isXsTime(fieldType)) {
			if (value != "") {
				control.element.value = moment.unix(value).format("YYYY-MM-DDThh:mm:ss")
			} else {
				control.element.value = ""
			}
		} else if (isXsId(fieldType)) {
			control.element.value = value
		} else if (isXsRef(fieldType)) {
			control.element.innerHTML = value.replace("#", "")
			control.element.href = value
		}
	} else {

		if (fieldType.startsWith("[]")) {
			// The propertie is an array...
			if (value != null && value != "") {
				// An array of reference.
				if (isRef) {
					// append a new container for link's...
					for (var i = 0; i < value.length; i++) {
						// append link's
						// Here I will call the set entity reference function.
						var uuid
						if (isObject(value[i])) {
							uuid = value[i].UUID
						} else {
							uuid = value[i]
						}
						if (uuid.length > 0 && isObjectReference(uuid)) {
							if (this.entity["set_" + field + "_" + uuid + "_ref"] == undefined) {
								setRef(this.entity, field, uuid, true)
							}
							this.entity["set_" + field + "_" + uuid + "_ref"](
								function (panel, control, field, fieldType) {
									return function (ref) {
										panel.appendObjectRef(ref, control, field, fieldType)
									}
								}(this, control, field, fieldType)
							)
						}
					}
				} else {
					// I will display a table.
					this.appendObjects(control, value, field, fieldType, parentUuid)
				}
			} else {
				// Empty the table.
				if (isRef) {
					control.removeAllChilds() // Empty all current element...
				} else {
					control.clear()
				}
			}
		} else {
			if (value != null) {
				// a single reference...
				if (isRef) {
					control.removeAllChilds()
					var uuid
					if (isObject(value)) {
						uuid = value.UUID
					} else {
						uuid = value
					}
					if (uuid.length > 0 && isObjectReference(uuid)) {
						if (this.entity["set_" + field + "_" + uuid + "_ref"] == undefined) {
							setRef(this.entity, field, uuid, false)
						}
						this.entity["set_" + field + "_" + uuid + "_ref"](
							function (panel, control, field, fieldType) {
								return function (ref) {
									panel.appendObjectRef(ref, control, field, fieldType)
								}
							}(this, control, field, fieldType)
						)
					}
				} else {
					// I will display an entity panel inside the existing one.
					if (control != undefined) {
						if (control.removeAllChilds != undefined) {
							control.removeAllChilds()
							this.appendObject(value, control, field, fieldType)
						}
					}

				}
			}
		}
	}
}

/**
 * Create an autocompletion list to an input box.
 * @param input The input to attach
 * @param typeName Element in the list are object of a given type.
 * @param entityPanel the parent panel of the entity.
 * @param ids a list of ids to remove from the list of choice.
 * @param onSelect the function to call when the value change in the list.
 */
function attachAutoCompleteInput(input, typeName, field, entityPanel, ids, onSelect) {
	var objMap = {}
	if (ids == undefined) {
		// must be an array...
		ids = []
	}

	input.element.readOnly = true
	input.element.style.cursor = "progress"
	input.element.style.width = "auto"

	// TODO use query instead of download all elements.

	server.entityManager.getEntities(typeName, typeName.substring(0, typeName.indexOf(".")), "", 0, -1, [], true,
		// Progress...
		function () {

		},
		// Sucess...
		function (results, caller) {
			// Set local variables.
			var input = caller.input
			var values = caller.values
			var objMap = caller.objMap
			var panel = caller.panel
			var onSelect = caller.onSelect
			var entityPanel = caller.entityPanel
			var ids = caller.ids
			var field = caller.field
			var lst = []

			input.element.readOnly = false
			input.element.style.cursor = "default"

			if (results.length > 0) {
				var prototype = entityPrototypes[results[0].TYPENAME]
				// get title display a readable name for the end user
				// or the first entity id.
				for (var i = 0; i < results.length; i++) {
					var result = results[i]
					if (result.getTitles != undefined) {
						var titles = result.getTitles()
						for (var j = 0; j < titles.length; j++) {
							if (ids.indexOf(titles[j]) == -1) {
								objMap[titles[j]] = result
								if (lst.indexOf(titles[j]) == -1) {
									lst.push(titles[j])
								}
							}
						}
					}
				}
			}

			attachAutoComplete(input, lst, false)

			input.element.addEventListener("keyup", function (entityPanel, field, input, objMap) {
				return function (e) {
					// If the key is escape...
					if (e.keyCode === 27) {
						entityPanel.resetFieldValue(field, input)
					}
					// Only index selection will erase the panel if the input is empty.
					if (entityPanel.proto != undefined) {
						if (entityPanel.proto.Indexs.indexOf(field) > -1 || entityPanel.proto.Ids.indexOf(field) > -1) {
							if (this.value.length == 0) {
								entityPanel.clear()
								this.focus()
								this.select()
							}
						}
					}
				}
			}(entityPanel, field, input, objMap))

			input.element.onblur = input.element.onchange = function (objMap, values, entityPanel, onSelect) {
				return function (evt) {
					var value = objMap[this.value]
					if (value != undefined) {
						var entity = entities[value.UUID]
						onSelect(entity)
					}
				}
			}(objMap, values, entityPanel, onSelect)

		},
		function (errMsg, caller) {
			// here there's no indexation... so what!
			var input = caller.input
			input.element.readOnly = false
			input.element.style.cursor = "default"
		},
		{ "input": input, "field": field, "objMap": objMap, "values": this.values, "entityPanel": entityPanel, "onSelect": onSelect, "ids": ids })

	// If a new entity is created related to the autocomplete type I need to append it into the list of choice.
	server.entityManager.attach({ "input": input, "field": field, "objMap": objMap, "values": this.values, "entityPanel": entityPanel, "onSelect": onSelect, "ids": ids }, NewEntityEvent, function (evt, caller) {
		// I will reinit the panel here...
		if (evt.dataMap["entity"].TYPENAME == caller.entityPanel.typeName) {
			var entity = entities[evt.dataMap["entity"].UUID]
			var entityPanel = caller.entityPanel
			var titles = entity.getTitles()
			var ids = caller.ids
			var lst = []
			var objMap = caller.objMap
			var input = caller.input
			var field = caller.field
			var onselect = caller.onSelect
			var values = caller.values

			for (var j = 0; j < titles.length; j++) {
				if (ids.indexOf(titles[j]) == -1) {
					objMap[titles[j]] = entity
					if (lst.indexOf(titles[j]) == -1) {
						lst.push(titles[j])
					}
				}
			}

			input.element.readOnly = false
			input.element.style.cursor = "default"

			attachAutoComplete(input, lst, false)

			input.element.addEventListener("keyup", function (entityPanel, field, input, objMap) {
				return function (e) {
					// If the key is escape...
					if (e.keyCode === 27) {
						entityPanel.resetFieldValue(field, input)
					}
					// Only index selection will erase the panel if the input is empty.
					if (entityPanel.proto != undefined) {
						if (entityPanel.proto.Indexs.indexOf(field) > -1 || entityPanel.proto.Ids.indexOf(field) > -1) {
							if (this.value.length == 0) {
								entityPanel.clear()
								this.focus()
								this.select()
							}
						}
					}
				}
			}(entityPanel, field, input, objMap))

			input.element.onblur = input.element.onchange = function (objMap, values, entityPanel, onSelect) {
				return function (evt) {
					var value = objMap[this.value]
					if (value != undefined) {
						var entity = entities[value.UUID]
						onSelect(entity)
					}
				}
			}(objMap, values, entityPanel, onSelect)
		}
	})

}