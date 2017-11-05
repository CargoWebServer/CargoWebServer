
/*
 * (C) Copyright 2016 Mycelius SA (http://mycelius.com/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * @fileOverview Table model functionality.
 * @author Dave Courtois
 * @version 1.0
 */


var languageInfo = {
	"en": {
		"month_1": "january",
		"month_2": "febuary",
		"month_3": "March",
		"month_4": "april",
		"month_5": "may",
		"month_6": "june",
		"month_7": "july",
		"month_8": "august",
		"month_9": "september",
		"month_10": "october",
		"month_11": "november",
		"month_12": "december"
	},
	"fr": {
		"month_1": "janvier",
		"month_2": "février",
		"month_3": "mars",
		"month_4": "avril",
		"month_5": "mai",
		"month_6": "juin",
		"month_7": "juillet",
		"month_8": "août",
		"month_9": "september",
		"month_10": "octobre",
		"month_11": "novembre",
		"month_12": "décembre"
	}
}

// Set the text...
server.languageManager.appendLanguageInfo(languageInfo)

/**
 * A table to display tabular data, extend the html table functinality.
 * 
 * @constructor
 * @param titles The list of table headers
 */
var Table = function (id, parent) {

	this.id = id
	if (this.id.length == 0) {
		this.id = randomUUID()
	}

	// The element where the table is contained
	this.parent = parent

	// The div...
	this.div = parent.appendElement({ "tag": "div", "class": "scrolltable", id: id }).down()

	// The header...
	this.header = null

	// The model contain the data to be display.
	this.model = null

	// The cells editor, date, int, string... 
	this.cellEditors = []

	// The row...
	this.rowGroup = null

	// The array of row...
	this.rows = []
	this.rowsId = {}

	/* the order of row **/
	this.orderedRows = []

	// The column formater formater...
	this.columnFormater = new ColumnFormater()

	// Format
	this.filterFormat = "YYYY-MM-DD"

	// The column sorters...
	this.sorters = []

	// The column filters
	this.filters = []

	// The delete row callback...
	this.deleteRowCallback = null

	// Attach the window resize event...
	window.addEventListener('resize',
		function (table) {
			return function (event) {
				table.refresh()
			}
		}(this)
	);

	return this
}


/**
 * Set the model of the table, different model are availables. (Sql, Key value etc...)
 * @param model
 * @param {function} initCallback The function to call when the intialisation is completed.
 */
Table.prototype.setModel = function (model, initCallback) {

	this.model = model

	// Initialyse the table model.
	this.model.init(
		// Success callback...
		function (results, caller) {
			//caller.caller.init()
			var table = caller.caller

			if (caller.initCallback != undefined) {
				caller.initCallback()
				caller.initCallback = undefined
				// Because html table suck with the header position vs scroll 
				// I do some little hack to fix it without weird stuff like 
				// copy another header etc...
				var widths = []
				// Now I will wrote the code for the layout...
				if(table.header != null){
					table.header.maximizeBtn.element.click()
				}
				
				if (table.parent.element.style.position == "relative") {
					var w = table.header.buttonDiv.element.offsetWidth
					table.header.buttonDiv.element.style.width = w + "px"
					table.header.buttonDiv.element.style.minWidth = w + "px"

					for (var i = 0; i < table.header.cells.length; i++) {
						var w = table.header.cells[i].element.offsetWidth
						table.header.cells[i].element.style.width = w + "px"
						table.header.cells[i].element.style.minWidth = w + "px"
						widths.push(w)
					}

					// Now the table body...
					for (var i = 0; i < table.rows.length; i++) {
						for (var j = 0; j < table.rows[i].cells.length; j++) {
							var cell = table.rows[i].cells[j]
							cell.div.element.style.width = widths[j] + "px"
							cell.div.element.style.minWidth = widths[j] + "px"
						}
					}

					// The table header.
					table.header.div.element.style.position = "absolute"
					table.header.div.element.style.left = "2px"

					// Now the table body
					table.rowGroup.element.style.position = "absolute"
					table.rowGroup.element.style.overflowX = "hidden"
					table.rowGroup.element.style.overflowY = "auto"
					table.rowGroup.element.style.left = "10px"
					table.rowGroup.element.style.top = table.header.div.element.offsetHeight + 2 + "px"

					// Now the height of the panel...
					table.rowGroup.element.style.height = (table.parent.element.offsetHeight - table.header.div.element.offsetHeight) - 15 + "px"

					// Here the scrolling event.
					table.rowGroup.element.onscroll = function (header) {
						return function () {
							var position = this.scrollTop;
							if (this.scrollTop > 0) {
								header.className = "table_header scrolling"
							} else {
								header.className = "table_header"
							}
						}
					}(table.header.div.element)

					// Now the resize event.
					window.addEventListener('resize',
						function (table) {
							return function () {

								table.rowGroup.element.style.height = (table.parent.element.offsetHeight - table.header.div.element.offsetHeight) - 15 + "px"
							}
						}(table), true);
				}

			}
		},
		// The progress callback...
		function (index, total, caller) {

		},
		// Error callback
		function (errMsg, caller) {
			if (caller.initCallback != undefined) {
				caller.initCallback()
				caller.initCallback = undefined
			}
		},
		{ "caller": this, "initCallback": initCallback })

}

/**
 * Initialyse the table.
 */
Table.prototype.init = function () {
	this.clear()

	if (this.rowGroup == null) {
		this.rowGroup = this.div.appendElement({ "tag": "div", "class": "table_body" }).down()
	}

	for (var i = 0; i < this.model.getRowCount(); i++) {
		var data = []
		for (var j = 0; j < this.model.getColumnCount(); j++) {
			if (this.header == null) {
				this.setHeader()
			}
			var value = this.model.getValueAt(i, j)
			data.push(value)
		}

		// If the model contain entities I will set the row id and set map entry.
		if (this.model.entities != undefined) {
			this.rows[i] = new TableRow(this, i, data, this.model.entities[i].UUID)
			this.rowsId[this.model.entities[i].UUID] = this.rows[i]
		} else {
			this.rows[i] = new TableRow(this, i, data)
			this.rowsId[data[0]] = this.rows[i]
		}
	}
	// In the case of sql table I will connect the listener here...
	// New connection of the event listeners...
	if (this.model.constructor === EntityTableModel) {

		/* The delete entity event **/
		server.entityManager.attach(this, DeleteEntityEvent, function (evt, table) {
			// So here I will remove the line from the table...
			var entity = evt.dataMap["entity"]
			if (entities[entity.UUID] != undefined) {
				entity = entities[entity.UUID]
			}

			if (entity.TYPENAME == table.model.proto.TypeName || table.model.proto.SubstitutionGroup.indexOf(entity.TYPENAME) != -1) {
				// So here I will remove the line from the model...
				var orderedRows = []
				for (var i = 0; i < table.orderedRows.length; i++) {
					var row = table.orderedRows[i]
					if (row.id != entity.UUID) {
						orderedRows.push(row)
					} else {
						// remove from the display...
						row.div.element.parentNode.removeChild(row.div.element)
					}
				}

				// Now set the model values and entities.
				var values = []
				var entities_ = []
				for (var i = 0; i < table.model.entities.length; i++) {
					if (table.model.entities[i].UUID != entity.UUID) {
						var entity = table.model.entities[i]
						if (entities[entity.UUID] != undefined) {
							entity = table.model.entities[i] = entities[entity.UUID]
						}
						entities_.push(entity)
						values.push(table.model.values[i])
					}
				}
				// Set the values...
				table.model.values = values
				table.model.entities = entities_

				table.orderedRows = orderedRows

				var rows = []
				var rowsId = {}

				for (var i = 0; i < table.rows.length; i++) {
					if (table.header == null) {
						table.setHeader()
					}
					var row = table.rows[i]
					if (row.id != entity.UUID) {
						row.index = rows.length
						rows.push(row)
						rowsId[row.id] = row
					} else {
						// remove from the display...
						if (row.div.element.parentNode != null) {
							row.div.element.parentNode.removeChild(row.div.element)
						}
					}
				}

				table.rowsId = rowsId
				table.rows = rows
				table.refresh()
			}
		})

		// The new entity event...
		server.entityManager.attach(this, NewEntityEvent, function (evt, table) {
			if (evt.dataMap["entity"] != undefined) {
				var entity = entities[evt.dataMap["entity"].UUID]
				if (entity != undefined) {
					if (entity.TYPENAME == table.model.proto.TypeName || table.model.proto.SubstitutionGroup.indexOf(entity.TYPENAME) != -1) {
						if (entity.ParentUuid != undefined && table.model.getParentUuid() != undefined) {
							if (table.model.getParentUuid() == entity.ParentUuid) {
								var row = table.appendRow(entity, entity.UUID)
								row.table.model.entities[row.index] = entity
								row.saveBtn.element.style.visibility = "hidden"
							}
						}
					}
				}
			}
		})

		// The update entity event.
		server.entityManager.attach(this, UpdateEntityEvent, function (evt, table) {
			if (evt.dataMap["entity"] != undefined) {
				var entity = entities[evt.dataMap["entity"].UUID]
				if (entity != undefined) {
					if (entity.TYPENAME == table.model.proto.TypeName || table.model.proto.SubstitutionGroup.indexOf(entity.TYPENAME) != -1) {
						if (entity.ParentUuid != undefined && table.model.getParentUuid() != undefined) {
							if (table.model.getParentUuid() == entity.ParentUuid) {
								var row = table.appendRow(entity, entity.UUID)
								row.table.model.entities[row.index] = entity
								row.saveBtn.element.style.visibility = "hidden"
							}
						}
					}
				}
			}
		})

	} else if (this.model.constructor === SqlTableModel) {

		/* The delete row event **/
		server.dataManager.attach(this, DeleteRowEvent, function (evt, table) {
			// So here I will remove the line from the table...
			if (evt.dataMap.tableName == table.id) {
				// So here I will remove the line from the model...F
				table.model.removeRow(evt.dataMap.id_0)

				var orderedRows = []
				for (var rowIndex in table.orderedRows) {
					var row = table.orderedRows[rowIndex]
					if (row.id != evt.dataMap.id_0) {
						orderedRows.push(row)
					} else {
						// remove from the display...
						row.div.element.parentNode.removeChild(row.div.element)
					}
				}

				table.orderedRows = orderedRows

				var rows = []
				for (var rowIndex in table.rows) {
					var row = table.rows[rowIndex]
					if (row.id != evt.dataMap.id_0) {
						row.index = rows.length
						rows.push(row)
					} else {
						// remove from the display...
						if (row.div.element.parentNode != null) {
							row.div.element.parentNode.removeChild(row.div.element)
						}
					}
				}
				table.rows = rows
				table.refresh()
			}
		})
	}

	// Refresh the parent table.
	this.refresh()
}

/**
 * Remove all items from table.
 */
Table.prototype.clear = function () {

	if (this.rows.length > 0) {
		this.model.removeAllValues()
		this.rows = []
		this.rowsId = {}
		this.rowGroup.element.innerHTML = ""
		this.div.element.style.display = "none"
		this.header.minimizeBtn.element.click()

		// Detach the listener.
		server.entityManager.detach(this, UpdateEntityEvent)
		server.entityManager.detach(this, NewEntityEvent)
		server.entityManager.detach(this, DeleteEntityEvent)
	}

	if (this.header != null) {
		this.header.div.element.parentNode.removeChild(this.header.div.element)
		this.header = null
	}

}

/**
 * Sort the table in respect of sorter's state.
 */
Table.prototype.sort = function () {
	// Here I will sort the table row...
	// I will create an ordered array of active sorters...
	var sorters = new Array()
	for (var s in this.sorters) {
		var sorter = this.sorters[s]
		sorter.childSorter = null
		if (sorter.state != undefined) {
			if (sorter.state != 0) {
				sorters[sorter.order - 1] = sorter
			}
		}
	}

	// I will copy values of rows to keep the original order...
	var values = new Array()
	for (var i = 0; i < this.rows.length; i++) {
		values[i] = this.rows[i]
	}

	// reset to default...
	if (sorters.length > 0) {
		// Link the sorter with each other...
		for (var i = 0; i < sorters.length - 1; i++) {
			sorters[i].childSorter = sorters[i + 1]
		}
		// Now I will call sort on the first sorter...
		sorters[0].sortValues(values)
	}

	for (var i = 0; i < values.length; i++) {
		this.orderedRows[i] = values[i]
	}

	this.refresh()
}

/**
 * Filter the table in respect with the filter's value.
 */
Table.prototype.filterValues = function () {
	// First I will display the rows...
	for (var i = 0; i < this.rows.length; i++) {
		this.rows[i].div.element.style.display = ""
	}

	// Now I will apply the filters...
	for (var i = 0; i < this.filters.length; i++) {
		this.filters[i].filterValues()
	}
}

/**
 * Refresh the table content.
 */
Table.prototype.refresh = function () {
	// Add items to HTML table element
	for (var rowIndex in this.orderedRows) {
		if (this.header != null) {
			this.header.div.element.style.display = ""
		}
		var item = this.orderedRows[rowIndex].div
		if (item != undefined) {
			this.rowGroup.element.appendChild(item.element)
		}
	}
}

/**
 * Return a row with a given id.
 */
Table.prototype.getRow = function (id) {
	if (this.rowsId[id] == undefined) {
		// Here it can be a new row...
		for (i = 0; i < this.rows.length; i++) {
			if (this.rows[i].id === "") {
				this.rows[i].id = id
				this.rowsId[id] = this.rows[i]
				break
			}
		}
	}
	return this.rowsId[id]
}

/**
 * Append a new row inside the table. If the row already exist it update it value.
 */
Table.prototype.appendRow = function (values, id) {
	if(id == undefined){
		id = 0
	}
	
	if (this.rowGroup == null) {
		//this.init()
		this.rowGroup = this.div.appendElement({ "tag": "div", "class": "table_body" }).down()
	}

	this.div.element.style.display = ""
	var row = this.getRow(id)
	if (row == undefined) {
		// append the value in the model and update the table.
		var data = this.model.appendRow(values)
		row = new TableRow(this, this.rows.length, data, id)
		this.rows[this.rows.length] = row
		this.rowsId[id] = row

		// Set the update listener for each row entity...
		if (id.length > 0) {
			if (this.model.constructor === EntityTableModel) {
				server.entityManager.attach(this, UpdateEntityEvent, function (evt, table) {
					if (evt.dataMap["entity"] != undefined) {
						var entity = entities[evt.dataMap["entity"].UUID]
						if (entity != undefined) {
							for (var i = 0; i < table.model.entities.length; i++) {
								if (table.model.entities[i] != undefined) {
									if (table.model.entities[i].UUID == entity.UUID) {
										row = row.table.appendRow(entity, entity.UUID)
										entity.ParentLnk = row.table.model.entities[row.index].ParentLnk
										row.table.model.entities[row.index] = entity
										row.saveBtn.element.style.visibility = "hidden"
									}
								}
							}
						}
					}
				})
			}
		}
	} else {
		// Here i will update the values...
		for (var i = 0; i < row.cells.length; i++) {
			var cell = row.cells[i]
			this.model.setValueAt(values, this.rows.indexOf(row), i)
			cell.setValue(this.model.values[this.rows.indexOf(row)][i])
		}
	}

	// Display the table.
	if (this.header == null) {
		this.setHeader()
	}
	this.header.maximizeBtn.element.click()

	return row
}

/**
 * Create a new header, this is call by set model, where the header data belong...
 */
Table.prototype.setHeader = function () {
	if (this.header == null) {
		this.header = new TableHeader(this)
	}
}

/**
 * return true if the panel is empty.
 */
Table.prototype.isEmpty = function () {
	return this.rows.length == 0
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  									Header/Row/Cell
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * The header contain the name of columns
 * @param table the parent table.
 * @constructor
 */
var TableHeader = function (table) {
	this.table = table
	this.cells = []
	this.div = table.div
		.prependElement({ "tag": "div", "class": "table_header", "style": "" }).down()

	// The first cell will be there to match the save button...
	this.buttonDiv = this.div.appendElement({ "tag": "div", "class": "table_header_btn_div" }).down()

	this.exportBtn = this.buttonDiv.appendElement({ "tag": "div", "class": "table_header_size_btn", "style": "width: 100%;" }).down()
	this.exportBtn.appendElement({ "tag": "i", "class": "	fa fa-download", "title": "download table data file." })

	this.maximizeBtn = this.buttonDiv.appendElement({ "tag": "div", "class": "table_header_size_btn", "style": "display: none;" }).down()
	this.maximizeBtn.appendElement({ "tag": "i", "class": "fa fa-plus-square-o" })

	this.minimizeBtn = this.buttonDiv.appendElement({ "tag": "div", "class": "table_header_size_btn" }).down()
	this.minimizeBtn.appendElement({ "tag": "i", "class": "fa fa-minus-square-o" })

	this.numberOfRowLabel = this.buttonDiv.appendElement({ "tag": "div", "class": "number_of_row_label", "style": "display: none;" }).down()

	this.maximizeBtn.element.onclick = function (rowGroup, minimizeBtn, numberOfRowLabel) {
		return function () {
			this.style.display = "none"
			rowGroup.element.style.display = "table-row-group"
			minimizeBtn.element.style.display = "table-cell"
			numberOfRowLabel.element.style.display = "none"
		}
	}(this.table.rowGroup, this.minimizeBtn, this.numberOfRowLabel)

	this.minimizeBtn.element.onclick = function (rowGroup, maximizeBtn, numberOfRowLabel, table) {
		return function () {
			this.style.display = "none"
			if (rowGroup != null) {
				rowGroup.element.style.display = "none"
			}
			maximizeBtn.element.style.display = "table-cell"
			if (table.rows.length > 0) {
				numberOfRowLabel.element.style.display = "table-cell"
				numberOfRowLabel.element.innerHTML = table.rows.length
			} else {
				numberOfRowLabel.element.style.display = "none"
			}
		}
	}(this.table.rowGroup, this.maximizeBtn, this.numberOfRowLabel, this.table)

	this.minimizeBtn.element.click()

	// I will create the header cell...
	for (var i = 0; i < table.model.getColumnCount(); i++) {
		var title = table.model.getColumnName(i)
		var cell = this.div.appendElement({ "tag": "div", "class": "header_cell" }).down()

		var cellContent = cell.appendElement({ "tag": "div", "class": "cell_content" }).down()
		// The column sorter...
		var sorter = new ColumnSorter(i, table)
		cellContent.appendElement(sorter.div)

		// Set the title
		cellContent.appendElement({ "tag": "div", "class": "cell_content", "innerHtml": title })

		// The colum filters...
		var filter = new ColumnFilter(i, table)
		cellContent.appendElement(filter.div)

		this.table.filters.push(filter)
		this.table.sorters.push(sorter)
		this.cells.push(cell)
	}

	// Expost callback will be call with data so it's possible to 
	// do what we want at export time...
	this.table.exportCallback = null

	// The download csv action.
	this.exportBtn.element.onclick = function (table) {
		return function () {
			// I will create csv file from the model values.
			var rows = table.model.values
			if (table.exportCallback != null) {
				rows.unshift(table.model.titles)
				table.exportCallback(rows)
			}
		}
	}(this.table)

	return this
}

/**
 * Return the column width
 * @param index the index of the column that we want the width.
 */
TableHeader.prototype.getColumnWidth = function (index) {
	return this.cells[index].element.firstChild.offsetWidth
}

/**
 * Set the width of a column.
 * @param width the new width.
 */
TableHeader.prototype.setColumnWidth = function (width) {
	this.cells[index].element.style.width = width
}

/** 
 * The table row.
 * @param table The parent table.
 * @param index The row index
 * @param data The value to display in the row.
 * @param id The id of the row must be unique in the table.
 * @constructor
 */
var TableRow = function (table, index, data, id) {
	if (table.rowGroup == null) {
		return
	}

	this.index = index
	this.table = table
	this.cells = []
	this.id = id
	this.saveBtn = null
	this.deleteBtn = null

	if (this.id == undefined) {
		this.id = data[0] // The first data must be the id...
	}

	this.div = table.rowGroup.appendElement({ "tag": "div", "class": "table_row", "id": this.id }).down()

	this.saveBtn = this.div.appendElement({ "tag": "div", "class": "row_button", "style": "visibility: hidden;", "id": id + "_delete_btn" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-floppy-o" }).down()

	this.saveBtn.element.onclick = function (row) {
		return function () {
			this.style.visibility = "hidden"
			row.table.model.saveValue(row)
		}
	}(this)

	// I will create the header cell...
	for (var i = 0; i < data.length; i++) {
		var cell = new TableCell(this, i, data[i])
		this.cells.push(cell)
	}

	// The delete button.
	this.deleteBtn = this.div.appendElement({ "tag": "div", "class": "row_button", "id": id + "_delete_btn" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-trash-o" }).down()

	// Now the action...
	this.deleteBtn.element.onclick = function (self, id, deleteCallback) {
		return function () {
			// here I will delete the row...
			self.table.model.removeRow(self.index)

			// Call the delete callback function...
			if (deleteCallback != null) {
				deleteCallback(id)
			}
		}
	}(this, data[0], this.table.deleteRowCallback)

	return this
}

/**
 * The table cell.
 * @param {Row} row The parent Row
 * @param {int} index The cell index.
 * @param {*} value The value of the cell. 
 * @constructor
 */
var TableCell = function (row, index, value) {
	this.index = index
	this.row = row
	this.div = row.div.appendElement({ "tag": "div", "class": "body_cell" }).down()
	this.valueDiv = this.div.appendElement({ "tag": "div", "class": "cell_value" }).down()

	if (value != null) {
		var formatedValue
		if (isObject(value)) {
			if (value.M_valueOf != undefined) {
				formatedValue = this.formatValue(value.M_valueOf)
			} else {
				formatedValue = this.formatValue(value)
			}
		} else {
			formatedValue = this.formatValue(value)
		}

		if (formatedValue == undefined) {
			return
		}

		if (formatedValue.constructor === Table) {
			this.valueDiv.appendElement(formatedValue.div)
		} else if (formatedValue.constructor === EntityPanel) {
			this.valueDiv.appendElement(formatedValue.panel)
		} else if (formatedValue.element != null) {
			this.valueDiv.appendElement(formatedValue)
		} else {
			this.valueDiv.element.innerHTML = formatedValue
		}
	}

	// Now the click event...
	this.div.element.ondblclick = function (cell) {
		return function (e) {
			e.stopPropagation()
			cell.appendCellEditor(this.offsetWidth, this.offsetHeight)
		}
	}(this)
}

// create an item link...
function createItemLnk(entity, value, field, valueDiv) {
	// Be sure we point to the map entity
	if (entities[entity.UUID] != undefined) {
		entity = entities[entity.UUID]
	}

	// Remove the content if any...
	valueDiv.element.innerHTML = ""
	var prototype = getEntityPrototype(value.TYPENAME)

	var titles = []
	if (value.getTitles != undefined) {
		titles = value.getTitles()
	} else {
		titles = ["lnk"]
	}

	var refName = ""
	for (var j = 0; j < titles.length; j++) {
		refName += titles[j]
		if (j < titles.length - 1) {
			refName += " "
		}
	}

	// I will display the remove button on case of user want to remove the linked entity.
	var content = valueDiv.appendElement({ "tag": "div" }).down()
	var ref = content.appendElement({ "tag": "div", "style": "display: inline; padding-left: 3px; vertical-align: text-bottom;" }).down().appendElement({ "tag": "a", "href": "#", "title": value.TYPENAME, "innerHtml": refName }).down()
	ref.element.id = value.UUID
	var deleteLnkButton = content.appendElement({ "tag": "div", "style": "display: inline; width: 100%;", "class": "row_button" }).down().appendElement({ "tag": "i", "class": "fa fa-trash", "style": "margin-left: 8px;" }).down()

	ref.element.onclick = function (entity, ref, valueDiv, field) {
		return function () {
			// Here I ill display the entity panel...
			if (ref.subEntityPanel == undefined) {
				ref.subEntityPanel = new EntityPanel(valueDiv, entity.TYPENAME,
					// The init callback. 
					function (entityPanel) {
						entityPanel.setEntity(entities[entity.UUID])
					},
					undefined, true, entity, field)
			} else {
				if (ref.subEntityPanel.panel.element.style.display == "none") {
					ref.subEntityPanel.panel.element.style.display = ""
				} else {
					ref.subEntityPanel.panel.element.style.display = "none"
				}
			}
		}
	}(value, ref, valueDiv, field)

	// Now the delete action...
	deleteLnkButton.element.onclick = function (entityUuid, object, field, content) {
		return function () {
			var entity = entities[entityUuid]
			removeObjectValue(entity, field, object)
			server.entityManager.saveEntity(entity,
				function (result, content) {
					// nothing to do here.
					if (content.element.parentNode != null) {
						content.element.parentNode.removeChild(content.element)
					}
				},
				function (result, caller) {
					// nothing to here.
				},
				content)
		}
	}(entity.UUID, value, field, content)
	return content
}

/**
 * Fromat the content of the cell in respect of the value.
 * @param {} value The value to display in the cell.
 */
TableCell.prototype.formatValue = function (value) {

	// Format the cell value before display...
	var fieldType = this.row.table.model.getColumnClass(this.index)
	if (fieldType == undefined) {
		return value
	}

	// In case of a reference string...
	if (fieldType.endsWith("string") && this.row.table.model.constructor === EntityTableModel) {
		if (isObjectReference(value)) {
			// In that case I will create the link object to reach the 
			// reference...
			fieldType = value.split("%")[0] + ":Ref"
		}
	}

	// Binary string here.
	if (fieldType.endsWith("base64Binary")) {
		if (isString(value)) {
			try {
				value = JSON.parse(value);
			} catch (e) {
				/** Nothing here */
			}
		}

		if (isString(value)) {
			fieldType = "xs.string"
		} else {
			if (isArray(value)) {
				for (var i = 0; i < value.length; i++) {
					if (isObject(value[i])) {
						fieldType = "[]" + value[i].TYPENAME + ":Ref"
					} else if (isString(value[i])) {
						fieldType = "xs.string"
					}
					// TODO do other types here...
				}
			} else {
				if (isObject(value)) {
					if (value.TYPENAME != undefined) {
						fieldType = value.TYPENAME + ":Ref"
					}
				}
			}
		}
	}

	// if is extension of a base type.
	var baseType = fieldType
	if (!baseType.startsWith('[]xs.')) {
		baseType = getBaseTypeExtension(fieldType)
	}

	var field = "M_" + this.row.table.model.titles[this.index]
	var isArray_ = fieldType.startsWith("[]")
	var formater = this.row.table.columnFormater
	var isRef = fieldType.endsWith(":Ref")

	fieldType = fieldType.replace("[]", "")

	// if its xs type... 
	var isBaseType = isXsBaseType(fieldType) || isXsBaseType(baseType) || field == "M_valueOf" || field == "M_listOf"
	if (isBaseType) {
		if (!isArray_) {
			if (isXsDate(fieldType) || isXsDate(baseType)) {
				value = formater.formatDate(value);
				this.valueDiv.element.className = "xs_date"
			} else if (isXsNumeric(fieldType) || isXsNumeric(baseType)) {
				this.valueDiv.element.className = "xs_number"
				if (isXsMoney(fieldType) || fieldType.indexOf("Price") != -1) {
					value = formater.formatReal(value, 2)
				} else {
					value = formater.formatReal(value, 3)
				}
			} else if (isXsId(fieldType) || isXsId(baseType)) {
				// here I have an entity now I need to know if is an abstact entity or not.
				var entity = entities[this.row.table.model.entities[this.row.index].UUID]
				if (entity == undefined) {
					// In case of newly created entity.
					entity = this.row.table.model.entities[this.row.index]
				}
				var entityType = entity.__class__
				var entityPrototype = getEntityPrototype(entityType)
				if (entityPrototype.IsAbstract) {
					// Here I will treat it as a reference.
					var field = "M_" + this.row.table.model.titles[this.index]
					server.entityManager.getEntityByUuid(entity.UUID,
						function (result, caller) {
							var entity = caller.entity
							if (entities[entity.UUID] != undefined) {
								entity = entities[entity.UUID]
							}
							caller.createItemLnk(entity, result, caller.field, caller.valueDiv)
						}, function () {

						}, { "valueDiv": this.valueDiv, "field": field, "createItemLnk": createItemLnk, "entity": entity })
					// Here will create an indermine 
					if (entities[entity.UUID] == undefined) {
						return new Element(null, { "tag": "progress", "innerHtml": "Unknown" })// return here...
					}
					return new Element(null, { "tag": "div", "innerHtml": "" })
				} else {
					// simply set the id as input...
					this.valueDiv.element.className = "xs_id"
					value = formater.formatString(value)
				}

			} else if (isXsRef(fieldType)) {
				// Here I will put a text area..
				this.valueDiv.element.className = "xs_ref"
			} else if (isXsInt(fieldType)) {
				this.valueDiv.element.className = "xs_number"
			} else if (isXsBoolean(fieldType)) {
				this.valueDiv.element.className = "xs_bool"
			} else if (isXsString(fieldType)) {
				this.valueDiv.element.className = "xs_string"
				value = formater.formatString(value)
			} else {
				value = formater.formatString(value)
			}
		} else {
			var content = this.valueDiv.getChildById(field + "_content")
			if (content != undefined) {
				return content
			}

			content = new Element(this.valueDiv, { "tag": "div", "style": "display: table-row;", "id": field + "_content" })
			var newLnkButton = content.appendElement({ "tag": "div", "class": "new_row_button row_button", "id": field + "_plus_btn" }).down()
			newLnkButton.appendElement({ "tag": "i", "class": "fa fa-plus" }).down()

			// A generic table.
			var tableModel = new TableModel(["index", "values"])
			tableModel.fields = ["xs.int", baseType.replace("[]", "")]
			tableModel.editable[1] = true
			var entity = entities[this.row.table.model.entities[this.row.index].UUID]
			var itemTable = new Table(randomUUID(), content)

			itemTable.setModel(tableModel, function (table, items, entity, field) {
				return function () {
					for (var i = 0; i < items.length; i++) {
						row = table.appendRow([i + 1, items[i]], i /*entity.UUID*/)
						// The delete row action...
						row.deleteBtn.element.onclick = function (entity, field, row) {
							return function () {
								// Here I will simply remove the element 
								// The entity must contain a list of field...
								if (isString(entity)) {
									entity = entities[entity]
								}
								if (entity[field] != undefined) {
									entity[field].splice(row.index, 1)
									entity.NeedSave = true
									server.entityManager.saveEntity(entity)
								}
							}
						}(entity, field, row)

						// The save row action
						row.saveBtn.element.onclick = function (entity, field, row) {
							return function () {
								// Here I will simply remove the element 
								// The entity must contain a list of field...
								if (isString(entity)) {
									entity = entities[entity]
								}
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
						}(entity, field, row)

					}
					table.refresh()
				}
			}(itemTable, value, this.row.table.model.entities[this.row.index].UUID, field))

			// Here the new value...
			newLnkButton.element.onclick = function (itemTable, entity, field, fieldType) {
				return function () {
					if (isString(entity)) {
						entity = entities[entity]
					}
					var newRow = itemTable.appendRow([entity[field].length + 1, "0"], entity[field].length)
					newRow.saveBtn.element.style.visibility = "visible"
					simulate(newRow.cells[entity[field].length, 1].div.element, "dblclick");

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
					}(entity, field, newRow)

					// The save row action
					newRow.saveBtn.element.onclick = function (entity, field, row) {
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
									server.entityManager.createEntity(entity.ParentUuid, entity.ParentLnk, entity.TYPENAME, "", entity,
										function (result, caller) {
											caller.style.visibility = "hidden"
										},
										function () {

										}, this)
								}
							}
						}
					}(entity, field, newRow)
				}
			}(itemTable, this.row.table.model.entities[this.row.index].UUID, field, fieldType)
			return itemTable
		}
	} else {
		if (isArray_) {
			var entity = entities[this.row.table.model.entities[this.row.index].UUID]
			if (entity == undefined) {
				// In case of newly created entity.
				entity = this.row.table.model.entities[this.row.index]
			}

			if (isRef) {

				// In that case the array contain a list of reference.
				// An entity table...
				var prototype = getEntityPrototype(fieldType.replace("[]", "").replace(":Ref", ""))
				var field = "M_" + this.row.table.model.titles[this.index]

				var content = this.valueDiv.getChildById(field + "_content")
				if (content != undefined) {
					return content
				}

				content = new Element(this.valueDiv, { "tag": "div", "id": field + "_content", "style": "display: block;" })
				// The append link button...
				var newLnkButton = content.appendElement({ "tag": "div", "class": "new_row_button row_button", "id": field + "_plus_btn" }).down()
				newLnkButton.appendElement({ "tag": "i", "class": "fa fa-plus" }).down()

				var valueDiv = content.appendElement({ "tag": "div", "style": "display: table-cell; width: 100%;" }).down().appendElement({ "tag": "div", "style": "display: table;" })
				for (var i = 0; i < value.length; i++) {
					var lnkDiv = valueDiv.appendElement({ "tag": "div", "style": "display: table-row;" }).down()
					var uuid
					if (isObject(value[i])) {
						uuid = value[i].UUID
					} else {
						uuid = value[i]
					}

					var entity = this.row.table.model.entities[this.row.index]
					entity = entities[entity.UUID]

					if (uuid.length > 0 && isObjectReference(uuid)) {
						if (entity["set_" + field + "_" + uuid + "_ref"] == undefined) {
							// In that case the value is a generic entity from arrary of byte so I will 
							// create the function here.
							// Here I must implement the set and reset function 
							if (value[i].UUID != undefined) {
								setRef(entity, field, value[i].UUID, false)
							} else {
								// In that case the reference is a string...
								setRef(entity, field, value[i], false)
							}
						}

						entity["set_" + field + "_" + uuid + "_ref"](
							function (entity, field, valueDiv, createItemLnk) {
								return function (ref) {
									createItemLnk(entity, ref, field, valueDiv)
								}
							}(entity, field, lnkDiv, createItemLnk)
						)
					}
				}

				newLnkButton.element.onclick = function (valueDiv, entity, fieldType, field, cell) {
					return function () {
						var newLnkInput = valueDiv.getChildById("new_" + field + "_input_lnk")

						if (newLnkInput != undefined) {
							// remove from the layout
							if (newLnkInput.element.parentNode) {
								newLnkInput.element.parentNode.removeChild(newLnkInput.element)
							}
						}

						// Now  i will create the new input box...
						newLnkInput = valueDiv.appendElement({ "tag": "div", "style": "display: table-row;" }).down().appendElement({ "tag": "input", "id": "new_" + field + "_input_lnk", "style": "width: auto;" }).down()

						// Now i will set it autocompletion list...
						attachAutoCompleteInput(newLnkInput, fieldType, field, valueDiv, entity.getTitles(),
							function (newLnkInput, entity, field, valueDiv, cell) {
								return function (value) {
									// I will get the value from the entity manager...
									if (value.UUID.length > 0) {
										value = entities[value.UUID]
									}

									if (entities[entity.UUID] != undefined) {
										entity = entities[entity.UUID]
									}

									var lnkDiv = valueDiv.appendElement({ "tag": "div", "style": "display: table-row;" }).down()
									createItemLnk(entity, value, field, lnkDiv)
									newLnkInput.element.parentNode.removeChild(newLnkInput.element)
									appendObjectValue(entity, field, value)

									//cell.row.saveBtn.element.style.visibility = "visible"
									server.entityManager.saveEntity(entity)
								}
							}(newLnkInput, entity, field, valueDiv, cell))

						newLnkInput.element.focus()
						newLnkInput.element.select();
					}
				}(valueDiv, entity, prototype.TypeName, field, this)

				return content
			} else {
				// Here i will append an array inside the array.
				// The append link button...
				var content = this.valueDiv.getChildById(field + "_content")
				if (content != undefined) {
					return content
				}

				content = new Element(this.valueDiv, { "tag": "div", "style": "display: table-row;", "id": field + "_content" })
				var newLnkButton = content.appendElement({ "tag": "div", "class": "new_row_button row_button", "id": field + "_plus_btn" }).down()
				newLnkButton.appendElement({ "tag": "i", "class": "fa fa-plus" }).down()
				var prototype = getEntityPrototype(fieldType.replace("[]", ""))

				if (value.length > 0) {
					// An entity table...
					var itemsTableModel = new EntityTableModel(prototype)
					var itemTable = new Table(randomUUID(), content)

					// Keep the table reference in the entity.
					entity[field + "_table"] = itemTable

					itemTable.setModel(itemsTableModel, function (itemsTable, values, field) {
						return function () {
							itemsTable.init() // connect events...
							for (var i = 0; i < values.length; i++) {
								if (values[i].UUID != undefined) {
									values[i].ParentLnk = field
									itemsTable.appendRow(values[i], values[i].UUID)
								} else {
									itemsTable.model.fields[0] = fieldType
									itemsTable.appendRow(values[i], i)
								}
							}
						}
					}(itemTable, value, field))
				}

				// Here I will append a new entity as a table row.
				newLnkButton.element.onclick = function (itemTable, entity, field, fieldType, content) {
					return function () {
						var item = eval("new " + fieldType + "()")
						item.UUID = fieldType + "%" + randomUUID()
						item.TYPENAME = fieldType

						// I will push the item
						if (entity[field] == undefined) {
							entity[field] = []
						}

						entity[field].push(item)

						// Set the parent uuid.
						item.ParentUuid = entity.UUID
						item.ParentLnk = field

						if (itemTable == undefined) {
							itemTable = entity[field + "_table"]
						}
						if (itemTable == undefined) {
							var itemTable = new Table(randomUUID(), content)
							var itemsTableModel = new EntityTableModel(getEntityPrototype(fieldType))
							itemTable.setModel(itemsTableModel, function (table, item) {
								return function () {
									newRow = itemTable.appendRow(item, item.UUID)
									newRow.saveBtn.element.style.visibility = "visible"
								}
							}(itemTable, item))

						} else {
							newRow = itemTable.appendRow(item, item.UUID)
							newRow.saveBtn.element.style.visibility = "visible"
						}

					}
				}(itemTable, entity, field, fieldType, content)

				return itemTable
			}
		} else {
			if (isRef) {
				// The item panel contain the item lnk and the entity panel.
				var field = "M_" + this.row.table.model.titles[this.index]
				var content = this.valueDiv.getChildById(value)
				if (content == undefined) {
					content = new Element(this.valueDiv, { "tag": "div", "style": "display: block;", "id": value })
				}

				var uuid
				if (isObject(value)) {
					uuid = value.UUID
				} else {
					uuid = value
				}

				var entity = this.row.table.model.entities[this.row.index]
				if (entity != undefined) {
					entity = entities[entity.UUID]
					if (uuid != undefined && entity != undefined) {
						if (uuid.length > 0 && isObjectReference(uuid)) {
							// TODO use setRef(owner, property, refValue, isArray) insted
							if (entity["set_" + field + "_" + uuid + "_ref"] == undefined) {
								// In that case the value is a generic entity from arrary of byte so I will 
								// create the function here.
								// Here I must implement the set and reset function 
								if (value.UUID != undefined) {
									setRef(entity, field, value.UUID, false)
								} else {
									// In that case the reference is a string...
									setRef(entity, field, value, false)
								}
							}

							entity["set_" + field + "_" + uuid + "_ref"](
								function (entity, field, valueDiv, createItemLnk) {
									return function (ref) {
										createItemLnk(entity, ref, field, valueDiv)
									}
								}(entity, field, content, createItemLnk)
							)
						}
					}
				}

				return content
			} else {
				// Here i will display a entity panel.
				return new EntityPanel(this.valueDiv, fieldType.replace("[]", "").replace(":Ref", ""),
					// The init callback. 
					function (entity) {
						return function (panel) {
							panel.setEntity(entity)
						}
					}(value),
					undefined, true, undefined, "")
			}
		}
	}

	return value
}

/**
 * Return the data type of the cell.
 * @returns {string} The data type.
 */
TableCell.prototype.getType = function () {
	return this.row.table.model.getColumnClass(this.index)
}

/**
 * Return the value contain in the cell.
 * @returns {*} Return the value contain in a cell
 */
TableCell.prototype.getValue = function () {
	value = this.row.table.model.getValueAt(this.row.index, this.index)
	return value
}

/**
 * Set the model value with the cell value.
 * @param {} value The value to set.
 */
TableCell.prototype.setValue = function (value) {
	if (this.valueDiv.element.innerHTML == value) {
		// if the value is the same I dont need to change it.
		return
	}

	// Display in the div...
	if (isObject(value)) {
		this.valueDiv.removeAllChilds()
		this.valueDiv.element.innerHTML = ""
		if (value.M_valueOf != undefined) {
			this.valueDiv.element.innerHTML = this.formatValue(value.M_valueOf)
		} else {
			var formated = this.formatValue(value)
			if (formated != undefined) {
				this.valueDiv.appendElement(this.formatValue(formated))
			}
		}
	} else {
		var formated = this.formatValue(value)
		if (isObject(formated)) {
			this.valueDiv.removeAllChilds()
			this.valueDiv.element.innerHTML = ""
			this.valueDiv.appendElement(formated)
		} else {
			this.valueDiv.element.innerHTML = formated
		}

	}

	// Set the value div visible.
	this.valueDiv.element.style.display = ""

	// Set save button visible.
	this.row.saveBtn.element.style.visibility = "visible"

	// Set in the model.
	this.row.table.model.setValueAt(value, this.row.index, this.index)

	// Refresh the table.
	this.row.table.refresh()
}

/**
 * Test if a cell is editable.
 * returns {bool} True if the cell can be edit.
 */
TableCell.prototype.isEditable = function () {
	return this.row.table.model.isCellEditable(this.index)
}

/**
 * Depending of the data type it return the correct cell editor.
 */
TableCell.prototype.appendCellEditor = function (w, h) {
	var column = this.index
	var row = this.row.index
	var width = this.row.table.header.getColumnWidth(column)

	if (this.isEditable() == false) {
		return // The cell is not editable...
	}

	var type = this.getType()
	var value = this.getValue()
	var parent = this.row.table.rowGroup

	// I will get the cell editor related to this column...
	var editor = this.row.table.cellEditors[this.index]

	// One editor at time.
	if (editor != undefined) {
		try {
			editor.element.parentNode.removeChild(editor.element)
		} catch (err) {

		}
		delete this.row.table.cellEditors[this.index]
		editor = null
	}

	var prototype = this.row.table.model.proto

	var entity = null
	var field
	var fieldType

	if (prototype != null) {
		// I will get the field...
		var field = prototype.Fields[this.index + 3]
		var fieldType = prototype.FieldsType[this.index + 3]

		// The value...
		if (value != null) {
			if (value.M_valueOf != undefined) {
				entity = value
				value = value.M_valueOf
				field = "M_valueOf"
				fieldType = getBaseTypeExtension(entity.TYPENAME)
			}
		}
	}

	// Here is the default editor if is undefined...
	if (editor == null) {
		if (isXsString(type)) {
			// Here I will put a text area..
			if (value.length > 50) {
				editor = this.div.appendElement({ "tag": "textarea", "resize": "false" }).down()
			} else {
				editor = this.div.appendElement({ "tag": "input" }).down()
			}
			editor.element.value = value
		} else if (isXsId(type)) {
			// Here I will put a text area..
			editor = this.div.appendElement({ "tag": "input" }).down()
			if (value != undefined) {
				editor.element.value = value
			}
		} else if (isXsRef(type)) {
			// Here I will put a text area..
			console.log("------> ref found!!!!")
		} else if (isXsDate(type)) {
			editor = this.div.appendElement({ "tag": "input", "type": "datetime-local", "name": "date" }).down()
			editor.element.value = moment(value).format('YYYY-MM-DDTHH:mm:ss');
			editor.element.step = 7
		} else if (isXsNumeric(type)) {
			editor = this.div.appendElement({ "tag": "input", "type": "number", "step": "0.01" }).down()
			editor.element.value = value
		} else if (isXsBoolean(type)) {
			editor = this.div.appendElement({ "tag": "input", "type": "checkbox" }).down()
			editor.element.checked = value
		} else if (isXsInt(type)) {
			editor = this.div.appendElement({ "tag": "input", "type": "number", "step": "1" }).down()
			editor.element.value = value
		} else if (prototype != null) {
			// get the rentity reference from the model.
			if (entity == null) {
				if (this.row.table.model.entities[this.row.index].UUID.length > 0) {
					if (entities[this.row.table.model.entities[this.row.index].UUID] != undefined) {
						entity = entities[this.row.table.model.entities[this.row.index].UUID]
					}
				} else {
					entity = this.row.table.model.entities[this.row.index]
				}
			}

			// If is an object...
			var isArray = type.startsWith("[]")
			var isRef = type.endsWith(":Ref")
			type = type.replace("[]", "").replace(":Ref", "")

			if (isRef) {
				// if the field is an array it will be display by another array in the array...
				if (isArray) {
					// nothing todo here...
					//var editor = appendRefEditor(this.div, entity, type, field)
				} else {

					// The editor will be an input box
					var editor = this.div.appendElement({ "tag": "input", "style": "display: inline;" }).down()
					editor.element.value = this.valueDiv.element.innerText

					// hide the lnk...
					this.valueDiv.element.style.display = "none"

					// I will keep the reference of the new and delete button inside the editor itself.
					editor.element.focus()
					editor.element.select();

					// Set clear function
					editor.clear = function () {
						return function () {
							editor.element.value = ""
						}
					}(editor)

					// Now i will set it autocompletion list...
					attachAutoCompleteInput(editor, type, field, entity, [],
						function (tableCell, entity, field, valueDiv, editor) {
							return function (value) {
								// Here I will set the field of the entity...
								if (value != null) {
									if (entity != undefined) {
										// Show the save button
										if (value.UUID.length > 0) {
											value = entities[value.UUID]
										}
										if (document.getElementById(value.UUID + "_" + field + "_lnk") == undefined) {

											appendObjectValue(entity, field, value)

											var lnkDiv = valueDiv.appendElement({ "tag": "div", "id": value.UUID + "_" + field + "_lnk", "style": "display: table-row;" }).down()
											createItemLnk(entity, value, field, lnkDiv)

											delete tableCell.row.table.cellEditors[tableCell.index]

											valueDiv.element.style.display = ""
											tableCell.row.saveBtn.element.style.visibility = "visible"
											entity.ParentLnk = tableCell.row.table.model.entities[tableCell.row.index].ParentLnk
											tableCell.row.table.model.entities[tableCell.row.index] = entity
											try {
												editor.element.parentNode.removeChild(editor.element)
											} catch (err) {
											}
										}
									}
								}
							}
						}(this, entity, field, this.valueDiv, editor))
				}
			} else {
				// Here I will get the prototype for the field type
				var fieldPrototype = getEntityPrototype(type)

				// Here it's an enumeration of value.
				if (fieldPrototype.Restrictions != undefined) {
					if (fieldPrototype.Restrictions.length > 0) {
						editor = this.div.appendElement({ "tag": "select", "id": "" }).down()
						for (var i = 0; i < fieldPrototype.Restrictions.length; i++) {
							var restriction = fieldPrototype.Restrictions[i]
							if (restriction.Type == 1) {
								type = "xs.string"
								editor.appendElement({ "tag": "option", "value": restriction.Value, "innerHtml": restriction.Value })
							}
						}

						// Set to the current value.
						editor.element.value = value

						// Hide the value div.
						this.valueDiv.element.style.display = "none"
					}
				} else {
					// Here the editor is an entity panel.
					editor = new EntityPanel(this.valueDiv, fieldPrototype.TypeName, function (parent, field, typeName) {
						return function (panel) {
							// Here I will set the actual values..
							if (parent[field] != undefined && parent[field] != "") {
								panel.setEntity(parent[field])
							} else {
								// Here the entity dosent already exist so I will create it...
								var entity = eval("new " + typeName + "()")
								entity.ParentLnk = field
								server.entityManager.createEntity(parent.UUID, field, entity.TYPENAME, "", entity,
									function (result, caller) {
										// Set the newly created entity.
										caller.entityPanel.setEntity(result)
										caller.parent[caller.field] = result
										entities[caller.parent.UUID] = caller.parent
									},
									function () {

									}, { "entityPanel": panel, "parent": parent, "field": field })
							}
							panel.maximizeBtn.element.click()
							panel.header.element.style.display = "none"
						}
					}(entity, field, fieldPrototype.TypeName)).panel
				}
			}
		}
		this.row.table.cellEditors[this.index] = editor
	}

	// I will set the editor on the page...
	if (editor != null /*&& type.startsWith("xs.")*/) {
		// When the selection change I will set the save button.
		editor.element.onchange = function (row) {
			return function () {
				row.saveBtn.element.style.visibility = "visible"
			}
		}(this.row)

		this.div.appendElement(editor)
		this.valueDiv.element.style.display = "none"

		editor.element.focus()
		if (editor.element.select != undefined) {
			editor.element.select()
		}

		// Set the editor size...
		if (isXsString(type)) {
			editor.element.style.width = w + "px"
			editor.element.style.height = h + "px"
		}
		editor.element.style.border = "none"
		editor.element.style.padding = "0px"
		editor.element.style.margin = "0px"

		var onblur = function (self, editor, onblur, field, entity) {
			// If the value change...
			return function () {
				var value
				if (this.type == "checkbox") {
					value = this.checked
				} else {
					value = this.value
				}

				if (self.value != value) {
					if (entity != undefined) {
						entity[field] = value
						entity.NeedSave = true
						self.setValue(value)
						// set the table entity.
						if (self.row.table.model.entities[self.row.index].TYPENAME == entity.TYPENAME) {
							self.row.table.model.entities[self.row.index] = entity
						} else if (entity.ParentUuid == self.row.table.model.entities[self.row.index].UUID) {
							self.row.table.model.entities[self.row.index][entity.ParentLnk] = entity
						}
					} else {
						self.setValue(value)
					}

				}
				if (editor.deleteButton != null) {
					if (editor.deleteButton.element.parentNode != null) {
						editor.deleteButton.element.parentNode.removeChild(editor.deleteButton.element)
					}
				}

				if (this.parentNode != null) {
					this.parentNode.removeChild(this)
				}
				self.valueDiv.element.style.display = ""
				try {
					editor.element.parentNode.removeChild(editor.element)
				} catch (err) {

				}

			}
		}(this, editor, onblur, field, entity)
		editor.element.onblur = onblur

		// If the editor is the entity panel I will
		if (editor.element.tagName == "DIV") {
			editor.element.addEventListener("keyup", function (editor, cell, entity, field, fieldType) {
				return function (evt) {
					evt.stopPropagation()
					if (evt.keyCode == 27) {
						editor.element.style.display = "none"
						if (entity[field] == undefined) {
							entity[field] = eval("new " + fieldType + "()")
						}
						cell.setValue(entity[field]) // Set the value.
						cell.valueDiv.element.style.display = ""
					}
				}
			}(editor, this, entity, field, fieldPrototype.TypeName))
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  The column formater...
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Column formater.
 * @constructor
 */
var ColumnFormater = function () {
	return this
}

/**
 * Format a date.
 * @param {date} value The date to format.
 * @param {format} format A string containing the format to apply, YYYY-MM-DD HH:mm:ss is the default.
 */
ColumnFormater.prototype.formatDate = function (value, format) {

	// Try to convert from a unix time.
	var date = new Date(value * 1000)
	if ((date instanceof Date && !isNaN(date.valueOf()))) {
		value = date
	}

	format = typeof format !== 'undefined' ? format : 'YYYY-MM-DD HH:mm:ss'
	value = moment(value).format(format);
	return value
}

/**
 * Display a real number to a given precision.
 * @param {real} value The number to format.
 * @param {int} digits The number of digits after the point.
 */
ColumnFormater.prototype.formatReal = function (value, digits) {
	if (digits == undefined) {
		digits = 2
	}
	value = parseFloat(value).toFixed(digits);
	return value
}

/**
 * Format a string to be display in a input or text area control.
 * @param {string} value The string to format.
 */
ColumnFormater.prototype.formatString = function (value) {
	if (value == null) {
		return ""
	}
	if (value.replace != undefined) {
		value = value.replace(/\r\n|\r|\n/g, "<br />")
	}
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  The column sorter...
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * This class is use to sort values. The sorter are chain whit other sorter, so
 * the sorting can be done between more than one column.
 * @param {TableCell} parent The the corresponding header cell
 * @param {int} index The index of the cell.
 * @param {Table} table The table where the values belong to
 * @constructor
 */
var ColumnSorter = function (index, table) {
	/* The table to sort...  **/
	this.table = table

	/* The column index **/
	this.index = index

	/* Child sorter, recursive... **/
	this.childSorter = null

    /*
     * The state can be 0 nothing, 1 asc or 2 desc
     * @type {number}
     */
	this.state = 0
	this.order = 0

	this.div = new Element(null, { "tag": "div", "class": "column_sorter" })

	// Now I will create the button 
	this.sort = this.div.appendElement({ "tag": "div" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-sort" }).down()

	this.sortAsc = this.div.appendElement({ "tag": "div" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-sort-asc", "style": "display:none;" }).down()

	this.sortDesc = this.div.appendElement({ "tag": "div" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-sort-desc", "style": "display:none;" }).down()

	this.orderDiv = this.div.appendElement({ "tag": "span" }).down()

	// The action...
	this.sort.element.onclick = function (sorter) {
		return function () {
			sorter.sort.element.style.display = "none"
			sorter.sortAsc.element.style.display = ""
			sorter.sortDesc.element.style.display = "none"
			sorter.state = 2
			sorter.setOrder()
		}
	}(this)

	this.sortAsc.element.onclick = function (sorter) {
		return function () {
			sorter.sort.element.style.display = "none"
			sorter.sortAsc.element.style.display = "none"
			sorter.sortDesc.element.style.display = ""
			sorter.state = 1
			sorter.setOrder()
		}
	}(this)

	this.sortDesc.element.onclick = function (sorter) {
		return function () {
			sorter.sort.element.style.display = ""
			sorter.sortAsc.element.style.display = "none"
			sorter.sortDesc.element.style.display = "none"
			sorter.state = 0
			sorter.setOrder()
		}
	}(this)
}

/**
 * Sort values.
 * @params values The values to sort.
 */
ColumnSorter.prototype.sortValues = function (values) {

	// Sort each array...
	values.sort(function (sorter) {
		return function (row1, row2) {
			var colIndex = sorter.index

			var value1 = sorter.table.model.getValueAt(row1.index, colIndex)
			var value2 = sorter.table.model.getValueAt(row2.index, colIndex)

			if (typeof value1 == "string") {
				value1.trim().toUpperCase()
			}
			if (typeof value2 == "string") {
				value2.trim().toUpperCase()
			}
			if (sorter.state == 2) {
				// asc
				if (value1 < value2)
					return -1;
				if (value1 > value2)
					return 1;
				return 0;
			} else {
				// desc
				if (value1 > value2)
					return -1;
				if (value1 < value2)
					return 1;
				return 0;
			}
		}

	}(this))

	// find same values and make it filter by the child sorter...
	if (this.childSorter != null) {
		var sameValueIndex = -1
		for (var i = 1; i < values.length; i++) {
			var value1 = this.table.model.getValueAt(values[i - 1].index, this.index)
			var value2 = this.table.model.getValueAt(values[i].index, this.index)
			if (value1 === value2) {
				if (sameValueIndex == -1) {
					sameValueIndex = i - 1
				}
			} else {
				if (sameValueIndex != -1) {
					// sort the values...
					var slice = values.slice(sameValueIndex, i)
					this.childSorter.sortValues(slice)
					// put back the value in the values...
					for (var j = 0; j < slice.length; j++) {
						values[sameValueIndex + j] = slice[j]
					}
					sameValueIndex = -1
				}
			}
		}
	}
}

/**
 * The the table sorter order.
 */
ColumnSorter.prototype.setOrder = function () {

	// Now I need to update order value...
	var activeSorter = new Array()

	for (var s in this.table.sorters) {
		if (this.table.sorters[s].state != undefined) {
			if (this.table.sorters[s].state != 0) {
				activeSorter[this.table.sorters[s].order] = this.table.sorters[s]
			}
		}
	}

	if (this.state != 0 && this.order == 0) {
		// Here I need to set the index
		this.order = Object.keys(activeSorter).length
		this.orderDiv.element.innerHTML = this.order.toString()
	} else if (this.state == 0) {
		this.order = 0
		this.orderDiv.element.innerHTML = ""
		var index = 1
		for (var i = 1; i < activeSorter.length; i++) {
			if (activeSorter[i] != undefined) {
				activeSorter[i].order = index
				activeSorter[i].orderDiv.element.innerHTML = activeSorter[i].order.toString()
				index++
			}
		}
	}

	// Sort the table...
	this.table.sort()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  The column filter...
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * This class is us to filtering elements of a table.
 *
 * @param {TableCell} parent The cell where the filter is.
 * @param {int} index The index of the cell.
 * @param {Table} table The parent table.
 * @constructor
 */
var ColumnFilter = function (index, table) {

	/* The parent element **/
	this.index = index
	this.table = table
	this.type = table.model.getColumnClass(this.index)

	// Keep the current list of filter...
	this.filters = {}

	// The array of all checkbox
	this.checkboxs = []

	// The column filter div...
	this.div = new Element(null, { "tag": "div", "class": "column_filter" })

	// I will create the button...
	this.filterIcon = this.div.appendElement({ "tag": "i", "class": "fa fa-filter filter" }).down()

	// The panel where the filter options reside...
	this.filterPanelDiv = this.div.appendElement({ "tag": "div", "class": "filter_panel_div" }).down()

	this.filterPanel = this.filterPanelDiv.appendElement({ "tag": "div", "class": "filter_panel" }).down()

	// Now the button...
	var filterPanelButtons = this.filterPanelDiv.appendElement({ "tag": "div", "class": "filter_panel_buttons" }).down()

	this.okBtn = filterPanelButtons.appendElement({ "tag": "div", "innerHtml": "ok" }).down()
	this.cancelBtn = filterPanelButtons.appendElement({ "tag": "div", "innerHtml": "cancel" }).down()

	// Simply close the panel...
	this.cancelBtn.element.onclick = function (filter) {
		return function () {
			filter.filterPanelDiv.element.style.display = "none"
			var checkSelectAll = true

			// I will reset the value to original state...
			for (var i = 1; i < filter.checkboxs.length; i++) {
				if (filter.filters[filter.checkboxs[i].element.name] != undefined) {
					filter.checkboxs[i].element.checked = true
				} else {
					filter.checkboxs[i].element.checked = false
					checkSelectAll = false
				}
			}

			filter.checkboxs[0].element.checked = checkSelectAll
		}
	}(this)

	// if a function is define here it will be called after values will be filer...
	this.filterCallback = null

	// Apply the filter...
	this.okBtn.element.onclick = function (filter) {
		return function () {
			filter.filterPanelDiv.element.style.display = "none"
			filter.filters = {}
			for (var i = 1; i < filter.checkboxs.length; i++) {
				if (filter.checkboxs[i].element.checked == true) {
					filter.filters[filter.checkboxs[i].element.name] = filter.checkboxs[i].element
				}
			}

			// Filter the values...
			filter.table.filterValues()
			if (filter.filterCallback != null) {
				// Call the filter callback.
				var values = []
				for (var i = 0; i < filter.table.rows.length; i++) {
					if (filter.table.rows[i].div.element.style.display != "none") {
						values.push(filter.table.model.getValueAt(i, filter.index))
					}
				}
				filter.filterCallback(values)
			}
		}
	}(this)

	var selectAll = this.appendFilter("(Sélectionner tout)")

	selectAll.element.onclick = function (filter) {
		return function () {
			for (var i = 1; i < filter.checkboxs.length; i++) {
				if (this.checked == true) {
					filter.checkboxs[i].element.checked = true
				} else {
					filter.checkboxs[i].element.checked = false
				}
			}
		}
	}(this)

	// Init the filer panel...
	this.initFilterPanel()

	////////////////////////////////////////////////////////////////////////////////////
	//            Event...
	////////////////////////////////////////////////////////////////////////////////////
	this.filterIcon.element.onclick = function (filter) {
		return function () {
			if (filter.filterPanelDiv.element.style.display != "block") {
				filter.filterPanelDiv.element.style.display = "block"
				/*var rect1 = localToGlobal(filter.filterPanelDiv.element)
				var rect2 = localToGlobal(filter.table.div.element)
				if (rect1.right > rect2.right) {
					filter.filterPanelDiv.element.style.right = "0px"
				}*/
			} else {
				filter.filterPanelDiv.element.style.display = "none"
			}
		}
	}(this)

	return this
}

/**
 * Determine if the fileter id activated or not.
 */
ColumnFilter.prototype.isActive = function () {
	return this.filterIcon.element.className.indexOf("filter-applied") != -1
}

/**
 * Get the list of different values.
 * @returns The list of different value.
 */
ColumnFilter.prototype.getValues = function () {
	var values = new Array()
	for (var i = 0; i < this.table.model.getRowCount(); i++) {
		// Get unique value...
		var value = this.table.model.getValueAt(i, this.index)
		var formater = this.table.columnFormater
		if (isXsDate(this.type)) {
			value = formater.formatDate(value, this.table.filterFormat);
		} else if (isXsNumeric(this.type) || isXsMoney(this.type)) {
			var precision = 2
			if (!isXsMoney(this.type)) {
				precision = 3
			}
			value = parseFloat(formater.formatReal(value, precision))
		} else if (isXsString(this.type)) {
			value = formater.formatString(value)
		} else if (isXsInt(this.type)) {
			value = parseInt(value)
		} else if (isXsBoolean(this.type)) {
			if (value == "true") {
				value = true
			} else {
				value = false
			}
		} else if (isXsBinary(this.type)) {
			console.log("Binary value found")
		} else if (isXsTime(this.type)) {
			value = parseInt(value)
		}

		var j = 0
		for (j = 0; j < values.length; j++) {
			if (values[j] == value) {
				break
			}
		}
		// replace existing value if already there...
		values[j] = value
	}

	if (this.type == "int" || this.type == "real" || this.type == "float") {
		values.sort(function sortNumber(a, b) {
			return a - b;
		})
	} else {
		values.sort()
	}

	return values
}

/**
 * Create the the filter gui.
 */
ColumnFilter.prototype.initFilterPanel = function () {

	// First of all I wil get the list of value...
	var type = this.table.model.getColumnClass(this.index)
	if (!isXsDate(this.type)) {
		var values = this.getValues()
		// So here I will create the list of values with checkbox...
		for (var i = 0; i < values.length; i++) {

			// I will append all the value independently...
			this.appendFilter(values[i])
		}
	} else {

		var years = {} // map of years... 

		// The list of date...
		var values = this.getValues()
		values.sort()
		for (var i = 0; i < values.length; i++) {
			dateValues = values[i].split("-")
			if (years[parseInt(dateValues[0])] == undefined) {
				years[parseInt(dateValues[0])] = { "year": dateValues[0], "months": {} }
				years[parseInt(dateValues[0])].months[parseInt(dateValues[1])] = { "days": {} }
				years[parseInt(dateValues[0])].months[parseInt(dateValues[1])].days[parseInt(dateValues[2])] = parseInt(dateValues[2])
			} else {
				if (years[parseInt(dateValues[0])].months[parseInt(dateValues[1])] == undefined) {
					years[parseInt(dateValues[0])].months[parseInt(dateValues[1])] = { "days": {} }
				}
				years[parseInt(dateValues[0])].months[parseInt(dateValues[1])].days[parseInt(dateValues[2])] = parseInt(dateValues[2])
			}
		}

		// Now I will create the filters for the dates...
		for (var y in years) {
			// So here I will create the years selector...
			var yid = randomUUID()
			var yline = this.filterPanel.appendElement({ "tag": "div", "id": yid, "style": "display: table;" }).down()
			var ymaximizeBtn = yline.appendElement({ "tag": "i", "class": "fa fa-plus", "style": "display: table-cell; padding-left:5px;" }).down()
			var ycheckbox = yline.appendElement({ "tag": "input", "type": "checkbox", "name": "year_" + y, "checked": "true", "style": "display: table-cell; vertical-align: middle;" }).down()
			var ylabel = yline.appendElement({ "tag": "label", "for": yid, "name": "year_" + y, "innerHtml": y, "style": "display: table-cell; vertical-align: middle;" }).down()

			ycheckbox.element.year = y

			// Append the checkbox to the filters...
			this.checkboxs.push(ycheckbox)
			this.filters["year_" + y] = ycheckbox

			// Now the months selector...
			var months = this.filterPanel.appendElement({ "tag": "div", "style": "margin-left:15px; display:none;" }).down()

			// Keep reference to the month  
			ycheckbox.monthsCheckBox = []

			var monthsIsShow = false
			ymaximizeBtn.element.onclick = function (months, ymaximizeBtn, monthsIsShow) {
				return function () {
					if (!monthsIsShow) {
						ymaximizeBtn.element.className = "fa fa-minus"
						months.element.style.display = "block"
					} else {
						ymaximizeBtn.element.className = "fa fa-plus"
						months.element.style.display = "none"
					}
					monthsIsShow = !monthsIsShow
				}
			}(months, ymaximizeBtn, monthsIsShow)

			for (var m in years[y].months) {
				var mid = randomUUID()
				var mline = months.appendElement({ "tag": "div", "id": mid, "style": "display: table;" }).down()
				var mmaximizeBtn = mline.appendElement({ "tag": "i", "class": "fa fa-plus", "style": "display: table-cell; padding-left:5px;" }).down()
				var mcheckbox = mline.appendElement({ "tag": "input", "type": "checkbox", "name": "year_" + y + "_month_" + m, "checked": "true", "style": "display: table-cell; vertical-align: middle;" }).down()
				var mlabel = mline.appendElement({ "tag": "label", "for": mid, "name": "year_" + y + "_month_" + m, "innerHtml": m, "style": "display: table-cell; vertical-align: middle;" }).down()
				server.languageManager.setElementText(mlabel, "month_" + m)
				// Now the Days...
				var days = months.appendElement({ "tag": "div", "style": "margin-left:30px; display:none;" }).down()
				var daysIsShow = false
				mcheckbox.element.month = m
				mcheckbox.element.year = y
				mcheckbox.ycheckbox = ycheckbox
				ycheckbox.monthsCheckBox.push(mcheckbox)
				this.checkboxs.push(mcheckbox)
				this.filters["year_" + y + "_month_" + m] = mcheckbox

				mmaximizeBtn.element.onclick = function (days, mmaximizeBtn, daysIsShow) {
					return function () {
						if (!daysIsShow) {
							mmaximizeBtn.element.className = "fa fa-minus"
							days.element.style.display = "block"
						} else {
							mmaximizeBtn.element.className = "fa fa-plus"
							days.element.style.display = "none"
						}
						daysIsShow = !daysIsShow
					}
				}(days, mmaximizeBtn, daysIsShow)

				mcheckbox.daysCheckBox = []

				for (var d in years[y].months[m].days) {
					var did = randomUUID()
					var dline = days.appendElement({ "tag": "div", "id": did, "style": "display: table;" }).down()
					var dcheckbox = dline.appendElement({ "tag": "input", "type": "checkbox", "name": "year_" + y + "_month_" + m + "_day_" + d, "checked": "true", "style": "display: table-cell; vertical-align: middle;" }).down()
					var dlabel = dline.appendElement({ "tag": "label", "for": did, "name": "year_" + y + "_month_" + m + "_day_" + d, "innerHtml": d, "style": "display: table-cell; vertical-align: middle;" }).down()
					dcheckbox.element.day = d
					dcheckbox.element.month = m
					dcheckbox.element.year = y
					dcheckbox.mcheckbox = mcheckbox
					mcheckbox.daysCheckBox.push(dcheckbox)
					this.checkboxs.push(dcheckbox)
					this.filters["year_" + y + "_month_" + m + "_day_" + d] = dcheckbox

					// Now the day checkbox...
					dcheckbox.element.onclick = function (dcheckbox) {
						return function () {
							// Now the year...
							if (this.checked == true) {
								dcheckbox.mcheckbox.element.checked = true
							} else {
								// Here I will set will uncheck the year if no month are checked...
								var hasDayChecked = false
								for (var i = 0; i < dcheckbox.mcheckbox.daysCheckBox.length && !hasDayChecked; i++) {
									hasDayChecked = dcheckbox.mcheckbox.daysCheckBox[i].element.checked
								}
								if (!hasDayChecked) {
									dcheckbox.mcheckbox.element.checked = false
								}
							}
						}
					}(dcheckbox)
				}

				mcheckbox.element.onclick = function (mcheckbox) {
					return function () {
						for (var i = 0; i < mcheckbox.daysCheckBox.length; i++) {
							mcheckbox.daysCheckBox[i].element.checked = this.checked
						}
						// Now the year...
						if (this.checked == true) {
							mcheckbox.ycheckbox.element.checked = true
						} else {
							// Here I will set will uncheck the year if no month are checked...
							var hasmonthChecked = false
							for (var i = 0; i < mcheckbox.ycheckbox.monthsCheckBox.length && !hasmonthChecked; i++) {
								hasmonthChecked = mcheckbox.ycheckbox.monthsCheckBox[i].element.checked
							}
							if (!hasmonthChecked) {
								mcheckbox.ycheckbox.element.checked = false
							}
						}
					}
				}(mcheckbox)
			}

			ycheckbox.element.onclick = function (ycheckbox) {
				return function () {
					for (var i = 0; i < ycheckbox.monthsCheckBox.length; i++) {
						ycheckbox.monthsCheckBox[i].element.checked = this.checked
						for (var j = 0; j < ycheckbox.monthsCheckBox[i].daysCheckBox.length; j++) {
							ycheckbox.monthsCheckBox[i].daysCheckBox[j].element.checked = this.checked
						}
					}
				}
			}(ycheckbox)

		}
	}
}

/**
 * Append new filter value.
 * @param the filter value.
 */
ColumnFilter.prototype.appendFilter = function (value) {
	if (value != "" && value != undefined) {
		var id = randomUUID()
		var line = this.filterPanel.appendElement({ "tag": "div", "id": id }).down()
		var checkbox = line.appendElement({ "tag": "input", "type": "checkbox", "name": value, "checked": "true" }).down()
		var label = line.appendElement({ "tag": "label", "for": id, "name": value, "innerHtml": value.toString() })
		this.checkboxs.push(checkbox) // Keep reference to the checkbox...
		this.filters[value] = checkbox
		checkbox.element.onclick = function (filter) {
			return function () {
				if (filter.hasFilter()) {
					filter.checkboxs[0].element.checked = false
				} else {
					filter.checkboxs[0].element.checked = true
				}
			}
		}(this)

		return checkbox
	}

	return null
}

/**
 * Tell if a column contain a filter.
 */
ColumnFilter.prototype.hasFilter = function () {
	for (var i = 1; i < this.checkboxs.length; i++) {
		if (this.checkboxs[i].element.checked == false) {
			return true
		}
	}
	return false
}

/**
 * Apply the value as filter...
 */
ColumnFilter.prototype.filterValues = function () {
	var type = this.table.model.getColumnClass(this.index)

	// The list of values...
	if (this.hasFilter()) {
		this.filterIcon.element.className = 'fa fa-filter filter-applied'
	} else {
		this.filterIcon.element.className = 'fa fa-filter filter'
		return
	}

	// Now I will apply the filters on each row of the table...
	for (var i = 0; i < this.table.rows.length; i++) {
		var row = this.table.rows[i]
		var cellValue = this.table.model.getValueAt(i, this.index)
		var isShow = false
		if (isXsDate(this.type)) {
			// Now filters are apply...
			for (var filter in this.filters) {
				var f = this.filters[filter]
				if (f.month != undefined && f.day != undefined) {
					// The month filter
					var filter_value = f.year + "-"
					if (f.month < 10) {
						filter_value += "0"
					}

					filter_value += f.month + "-"

					if (f.day < 10) {
						filter_value += "0"
					}
					filter_value += f.day

					isShow = cellValue.startsWith(filter_value)
					if (isShow) {
						break
					}
				}
			}
		} else {
			for (var filter in this.filters) {
				if (filter == cellValue) {
					isShow = true
					break
				}
			}
		}
		if (isShow == false) {
			row.div.element.style.display = "none"
		}
	}

	this.table.refresh()
}