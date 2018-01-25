
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


var tableTextInfo = {
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
server.languageManager.appendLanguageInfo(tableTextInfo)

/**
 * A table to display tabular data, extend the html table functinality.
 * 
 * @constructor
 * @param id The table id
 * @param parent The container.
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

	// The body
	this.rowGroup = null

	// The array of row...
	this.rows = []

	// The column width...
	this.columnsWidth = [];

	/* the order of row **/
	this.orderedRows = []

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

	// Link both side of the relation.
	this.model = model
	model.table = this

	// Initialyse the table model.
	this.model.init(
		// Success callback...
		function (results, caller) {
			//caller.caller.init()
			var table = caller.caller

			if (caller.initCallback != undefined) {
				caller.initCallback()
				caller.initCallback = undefined
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
				this.header.numberOfRowLabel.element.innerHTML = this.model.getRowCount();
			}
			var value = this.model.getValueAt(i, j)
			data.push(value)
		}

		// If the model contain entities I will set the row id and set map entry.
		this.rows[i] = new TableRow(this, i, data)
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
	// Function that dermine if the rows array already contain a given row.
	var hasRow = function (id) {
		return function (row) {
			if (row.id === id) {
				return row;
			}
			return null;
		}
	}(id)
	return this.rows.find(hasRow)
}

/**
 * Append a new row inside the table. If the row already exist it update it value.
 */
Table.prototype.appendRow = function (values, id) {
	if (id == undefined) {
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
		this.rows.push(row)
	} else {
		// Here i will update the values...
		for (var i = 0; i < row.cells.length; i++) {
			var cell = row.cells[i]
			this.model.setValueAt(values, this.rows.indexOf(row), i)
			cell.setValue(this.model.values[this.rows.indexOf(row)][i])
		}
	}

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

	this.numberOfRowLabel = this.buttonDiv.appendElement({ "tag": "div", "class": "number_of_row_label" }).down()

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
				rowGroup.element.style.display = ""
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
	this.numberOfRowLabel.element.style.display = ""

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

	this.saveBtn = this.div.appendElement({ "tag": "div", "class": "row_button save_row_btn", "style": "visibility: hidden;", "id": this.id + "_save_btn" }).down()
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
	this.deleteBtn = this.div.appendElement({ "tag": "div", "style": "visibility: hidden;", "class": "row_button delete_row_btn", "id": this.id + "_delete_btn" }).down()
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
	this.valueDiv = null

	// The object that display the data value in the cell
	this.renderer = new TableCellRenderer(this);

	// The object to 
	this.editor = new TableCellEditor(this, function (cell) {
		return function () {
			// Here I will remove the cell editor from the cell div.
			cell.div.element.removeChild(cell.editor.editor.element);
			cell.valueDiv.element.style.display = "";
			cell.setValue(cell.editor.getValue());
		}
	}(this));

	if (value != null) {
		// get the formated value
		var fieldType = this.row.table.model.getColumnClass(this.index);
		var formatedValue = this.renderer.render(value, fieldType);
		if (formatedValue != undefined) {
			if (formatedValue.element.tagName == "IMG") {
				this.div.element.style.textAlign = "center";
			}
			this.valueDiv = this.div.appendElement(formatedValue).down()
		}
	}

	// The click event is use to edit array ...
	if (this.getType().startsWith("[]")) {
		// On mouse enter I will append the functionality to edit cells...
		this.div.element.onmouseenter = function (cell) {
			return function (e) {
				e.stopPropagation()
				cell.setArrayCellEditor()
			}
		}(this)

		this.div.element.onmouseleave = function (cell) {
			return function (e) {
				e.stopPropagation()
				cell.resetArrayCellEditor()
			}
		}(this)
	} else {
		// Now the double click event...
		this.div.element.ondblclick = function (cell) {
			return function (e) {
				e.stopPropagation()
				cell.setCellEditor()
			}
		}(this)
	}
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
	if (this.valueDiv == undefined) {
		return
	}

	if (!isArray(value)) {
		if (this.valueDiv.element.innerHTML == value) {
			// if the value is the same I dont need to change it.
			return
		}
	}

	// Get the formated value.
	var fieldType = this.row.table.model.getColumnClass(this.index);
	var formated = this.renderer.render(value, fieldType)
	if (formated != undefined) {
		// Special rule for IMG element...
		if (formated.element.tagName == "IMG") {
			this.div.element.style.textAlign = "center";
		}

		// Remove the content of the div.
		this.div.removeAllChilds();

		// Set it value.
		this.valueDiv = this.div.appendElement(formated).down();
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
TableCell.prototype.setCellEditor = function (index) {
	// Set the editor...
	if (this.isEditable() == false) {
		return // The cell is not editable...
	}

	if (index == undefined) {
		// Here the cell is editable so I will display the cell editor.
		this.editor.edit(this.getValue(), this.getType());

		// Hide the cell renderer div...
		this.valueDiv.element.style.display = "none"
		// Append the editor...
		this.div.appendElement(this.editor.editor);
	} else {
		// So the the cell editor is in the row...
		var row = this.valueDiv.childs[Object.keys(this.valueDiv.childs)[index]]
		if (row == undefined) {
			row = this.valueDiv.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
		}
		// I will hide the actual value...
		var valueDiv = row.element.firstChild;

		// Here the cell is editable so I will display the cell editor.
		this.editor.edit(this.getValue()[index], this.getType().replace("[]", ""), function (row, cell, index) {
			return function () {
				// get the value from the editor and set it in the cell
				var values = cell.getValue()
				// must not use the closure directly...
				values[index] = cell.editor.getValue();
				cell.setValue(values);
				row.element.removeChild(row.element.firstChild)
				if (row.element.firstChild != undefined) {
					row.element.firstChild.style.display = "";
				}
				cell.row.table.refresh();
			}
		}(row, this, index));

		row.prependElement(this.editor.editor);
	}
	if (this.editor.editor != null) {
		if (this.editor.editor.element.type != "number") {
			this.editor.editor.element.setSelectionRange(0, this.editor.editor.element.value.length)
		}
		if (valueDiv != undefined) {
			valueDiv.style.display = "none";
		}
		this.editor.editor.element.focus();
	}
}

/**
 * That function is use to display the delete button and append button for array.
 */
TableCell.prototype.setArrayCellEditor = function () {
	if (this.div.getChildById("appendRowBtn") == undefined) {
		// First of all I will append the append row button
		this.div.element.style.position = "relative";
		this.div.element.style.display = "table-cell";
		this.div.element.style.paddingRight = "17px";

		this.div.appendElement({ "tag": "div", "class": "row_button", "id": "appendRowBtn", "style": "position: absolute; top: 0px; right: 0px;" }).down()
			.appendElement({ "tag": "i", "class": "fa fa-plus" });
		var appendRowBtn = this.div.getChildById("appendRowBtn");

		appendRowBtn.element.onclick = function (cell) {
			return function (evt) {
				evt.stopPropagation();
				if (cell.valueDiv == undefined) {
					cell.valueDiv = cell.div.appendElement({ "tag": "div", "style": "display: table; width: 100%;" }).down();
				}

				// Here I will append a new cell and set it editor.
				var values = cell.getValue();
				var typeName = cell.getType().replace("[]", "");
				// Depending of the values type I will append different value...
				if (isXsString(typeName)) {
					values.push("");
				} else if (isXsInt(typeName)) {
					values.push(0);
				} else if (isXsNumeric(typeName)) {
					values.push(0.0);
				} else if (isXsBoolean(typeName)) {
					values.push(false);
				} else if (typeName.endsWith(":Ref")) {
					// Here I will push an empty string...
					values.push("");
				} else {
					// Asynch here...
					var parentEntity = cell.row.table.model.entities[cell.row.index];
					var parentLnk = "M_" + cell.row.table.model.titles[cell.index]

					// Here I will test if is an entity prototype and if it is I will create an new object of that type.
					server.entityManager.getEntityPrototype(typeName, typeName.split(".")[0],
						// success callback.
						function (prototype, caller) {
							// so here I will create an new entity inside it parent.
							var entity = eval("new " + prototype.TypeName + "()")
							entity.ParentLnk = caller.parentLnk
							var entity = eval("new " + typeName + "()");
							entity.ParentLnk = caller.parentLnk;
							entity.ParentUuid = parentEntity.UUID;
							var cell = caller.cell;
							var values = caller.values;
							values.push(entity);
							cell.setValue(values);
							cell.row.table.refresh();
							cell.setCellEditor(values.length - 1);

							/*server.entityManager.createEntity(caller.parentEntity.UUID, entity.ParentLnk , entity.TYPENAME, "", entity,
								function (result, caller) {
									console.log("---> entity was created ", result)
								},
								function () {

								}, {})*/
						},
						// error callback.
						function (errObj, caller) {

						}, { "parentEntity": parentEntity, "parentLnk": parentLnk, "cell": cell, "values": values })

				}

				// Set the cell editor as needed...
				if (value.length > 0) {
					cell.setValue(values);
					cell.row.table.refresh();
					cell.setCellEditor(values.length - 1);
				}
			}
		}(this)

		// Now I will append the delete button to each existing row.
		if (this.valueDiv != undefined) {
			if (this.valueDiv.element.className == "entity_sub-table") {
				// in that case the value div contain a sub table of entitie so I will will simply display their 
				// delete button...
				var deleteBtns = this.valueDiv.element.getElementsByClassName("delete_row_btn");
				for (var i = 0; i < deleteBtns.length; i++) {
					deleteBtns[i].style.visibility = "visible";
				}
			} else {
				var index = 0;
				for (var id in this.valueDiv.childs) {
					var row = this.valueDiv.childs[id]
					var deleteBtn = row.appendElement({ "tag": "div", "class": "row_button delete_row_btn", "style": "vertical-align: top;" }).down()
						.appendElement({ "tag": "i", "class": "fa fa-trash" }).down();
					deleteBtn.element.onclick = function (cell, index, table, row) {
						return function () {
							// remove the row from the interface.
							table.removeElement(row)
							var values = cell.getValue();
							values.splice(index, 1);
							cell.setValue(values);
							cell.row.table.refresh();
						}
					}(this, index, this.valueDiv, row)
					index++;
				}
			}
		}
	}


}

/**
 * Remove the array cell edition control.
 */
TableCell.prototype.resetArrayCellEditor = function () {
	// So here I will remove the array edition control...
	var appendRowBtn = this.div.getChildById("appendRowBtn")
	if (appendRowBtn != undefined) {
		this.div.removeElement(appendRowBtn);
	}

	if (this.valueDiv != null) {
		if (this.valueDiv.element.className == "entity_sub-table") {
			var deleteBtns = this.valueDiv.element.getElementsByClassName("delete_row_btn");
			for (var i = 0; i < deleteBtns.length; i++) {
				deleteBtns[i].style.visibility = "hidden";
			}
		} else {
			var deleteBtns = this.div.getChildsByClassName("delete_row_btn")
			for (var i = 0; i < deleteBtns.length; i++) {
				deleteBtns[i].parentElement.removeElement(deleteBtns[i]);
			}
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  Various value formating function.
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Format a date.
 * @param {date} value The date to format.
 * @param {format} format A string containing the format to apply, YYYY-MM-DD HH:mm:ss is the default.
 */
function formatDate(value) {
	// Try to convert from a unix time.
	var date = new Date(value * 1000)
	if ((date instanceof Date && !isNaN(date.valueOf()))) {
		value = date
	}

	// Here I will use the browser 
	var format = 'YYYY-MM-DD HH:mm:ss';

	value = moment(value).format(format);
	return value
}

/**
 * Display a real number to a given precision.
 * @param {real} value The number to format.
 * @param {int} digits The number of digits after the point.
 */
function formatReal(value, digits) {
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
function formatString(value) {
	if (value == null) {
		return ""
	}
	if (value.replace != undefined) {
		value = value.replace(/\r\n|\r|\n/g, "<br />")
	}
	return value
}

/**
 * Format boolean value.
 * @param {*} value 
 */
function formatBoolean(value) {
	if (value == 1) {
		return "true";
	} else if (value == 0) {
		return "false";
	}
	return value.toString();
}

// Format the value
function formatValue(value, typeName) {
	// Here I will display basic types.
	if (isString(value)) {
		// remove empty values.
		if (value.startsWith("data:image/")) {
			// In that case I got a picture so I will create an image from it.
			var img = new Element(null, { "tag": "img", "src": value })
			return img;
		}
		formatedValue = formatString(value);

	} else if (isBoolean(value)) {
		formatedValue = formatBoolean(value);
	} else if (isInt(value)) {
		// In case of a date.
		if (isXsDate(typeName)) {
			formatedValue = formatDate(value);
		} else {
			// Int are numeric value with 0 digit.
			formatedValue = formatReal(value, 0);
		}
	} else if (isNumeric(value)) {
		if (isXsMoney(typeName) || typeName.indexOf("Price") != -1) {
			formatedValue = formatReal(value, 2)
		} else {
			formatedValue = formatReal(value, 3)
		}
	}

	return formatedValue
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  The cell editor...
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * That function will overide default edit function. It take a value as parameter and create 
 * the editor that can edit it content.
 */
var editFcts = {}

var TableCellEditor = function (cell, onblur) {
	// The parent cell.
	this.cell = cell

	// The underlying element.
	this.editor = null;

	// The on blur event listener.
	this.onblur = onblur

	return this
}

TableCellEditor.prototype.edit = function (value, typeName, onblur) {

	if (editFcts[typeName] != null) {
		// Here the function will create a custom editor for a given type.
		this.editor = editFcts[typeName](value)
		if (this.editor != null) {
			this.editor.element.onblur = this.onblur;
		}
		return;
	}

	// If the editor does not already exist I will initialyse it.
	if (isString(value)) {
		if (isXsString(typeName)) {
			// Here I will put a text area..
			if (this.editor == null) {
				if (value.length > 50) {
					this.editor = new Element(null, { "tag": "textarea", "resize": "false" })
				} else {
					this.editor = new Element(null, { "tag": "input" })
				}
			}
			// Set the string value.
			this.editor.element.value = formatString(value);

		} else if (isXsId(typeName)) {
			// Here I will put a text area..
			if (this.editor == null) {
				this.editor = new Element(null, { "tag": "input" })
			}
			if (value != undefined) {
				this.editor.element.value = value;
			}
		} else if (typeName.endsWith(":Ref") || isXsRef(typeName)) {
			// In that case is a reference.
			// The editor will be an input box
			if (this.editor == null) {
				this.editor = new Element(null, { "tag": "input", "style": "display: inline;", "id": randomUUID() });
			}

			// Here I will use the data manager to get the list of id, index and uuid...
			var typeName = typeName.split(":")[0] // Remove the :Ref
			server.entityManager.getEntityPrototype(typeName, typeName.split(".")[0],
				// success callback.
				function (prototype, caller) {
					var q = {};
					q.TypeName = prototype.TypeName;
					q.Fields = [];
					var fieldsType = [];
					for (var i = 0; i < prototype.Ids.length; i++) {
						q.Fields.push(prototype.Ids[i])
						fieldsType.push(prototype.FieldsType[prototype.getFieldIndex(prototype.Ids[i])]);
					}
					server.dataManager.read(prototype.TypeName.split(".")[0], JSON.stringify(q), fieldsType, [],
						// success callback
						function (results, caller) {
							var results = results[0];
							var elementLst = [];
							var idUuid = {};
							for (var i = 0; i < results.length; i++) {
								elementLst.push(results[i][1])
								idUuid[results[i][1]] = results[i][0]
								if (caller.value == results[i][0]) {
									caller.editor.editor.element.value = results[i][1];
								}
							}
							// I will attach the autocomplete box.
							attachAutoComplete(caller.editor.editor, elementLst, true,
								function (caller, idUuid) {
									return function (id) {
										caller.editor.editor.element.value = id;
										caller.editor.getValue = function (idUuid) {
											return function () {
												return idUuid[id];
											}
										}(idUuid)
										caller.onblur(); // Call the onblur function
									}
								}(caller, idUuid));
						},
						// progress callback
						function (index, total, caller) {

						},
						// error callback
						function (errObj, caller) {

						}, caller
					)

				},
				// error callback
				function (errObj, caller) {

				}, { "editor": this, "onblur": onblur, "value": value })
			// I will cancel the onblur event and call it when the data will be select.
			onblur = null;
			this.onblur = null;
		}
	} else if (isBoolean(value)) {
		if (this.editor == null) {
			this.editor = new Element(null, { "tag": "input", "type": "checkbox" });
		}
		this.editor.element.checked = value;
	} else if (isInt(value)) {
		// In case of a date.
		if (isXsDate(typeName)) {
			if (this.editor == null) {
				this.editor = new Element(null, { "tag": "input", "type": "datetime-local", "name": "date" });
			}
			this.editor.element.value = moment(value).format('YYYY-MM-DDTHH:mm:ss');
			this.editor.element.step = 7
		} else {
			if (this.editor == null) {
				this.editor = new Element(null, { "tag": "input", "type": "number", "step": "1" });
			}
			this.editor.element.value = value;
		}
	} else if (isNumeric(value)) {
		if (this.editor == null) {
			this.editor = new Element(null, { "tag": "input", "type": "number", "step": "0.01" });
		}
		this.editor.element.value = value;
	} else if (isObject(value)) {
		console.log("---> edit object!")
	}

	// Set the on blur event.
	if (this.editor != null) {
		if (onblur != undefined) {
			this.editor.element.onblur = onblur;
		} else {
			this.editor.element.onblur = this.onblur;
		}
	}
}

TableCellEditor.prototype.getValue = function () {
	return this.editor.element.value;
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//  The cell renderer...
//////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Column formater.
 * @constructor
 */
var TableCellRenderer = function (cell) {
	// The parent cell.
	this.cell = cell;
	return this
}

// You can overide default render function by setting the rendering 
// function for a given type. The function must take a value as parameter and
// return an element.
// Here is an example to render the type xs.ID differently...
// renderFcts["xs.ID"] = function(value){
//	return new Element(null, {"tag":"div", "innerHtml":"Test " + value});
// }
var renderFcts = {};

/**
 * Fromat the content of the cell in respect of the value.
 * @param {} value The value to display in the cell.
 */
TableCellRenderer.prototype.render = function (value, fieldType) {
	// Depending of the data type I will call the appropriate formater...
	var formatedValue = null;

	if (renderFcts[fieldType] != undefined) {
		// In that case I will use the overide function to render the cell.
		return renderFcts[fieldType](value)
	}


	// I will us Javasript type to determine how I will display the data...
	if (isArray(value)) {
		// Format array create it own element.
		var div = this.renderArray(value, fieldType);
		return div;
	} else if (isObject(value)) {
		formatedValue = this.renderEntity(value, fieldType);
	} else {
		formatedValue = formatValue(value, fieldType);
		if (isObjectReference(formatedValue)) {
			// Here the value represent an object reference.
			var lnk = new Element(null, { "tag": "a", "href": "#", "class": fieldType.replaceAll(".", "_") });
			server.entityManager.getEntityByUuid(formatedValue,
				function (entity, caller) {
					var titles = entity.getTitles();
					if (titles.length > 0) {
						caller.element.innerHTML = titles[0]
						caller.element.onclick = function (entity) {
							return function (evt) {
								evt.stopPropagation();
							}
						}(entity)
					}
				},
				function (errObj, lnk) {

				}, lnk)

			return lnk;
		} else {
			// Here I will create the 
			if (formatedValue != null) {
				return new Element(null, { "tag": "div", "class": fieldType.replaceAll(".", "_"), "innerHtml": formatedValue });
			}
		}
	}

	return null;
}

/**
 * Render array.
 */
TableCellRenderer.prototype.renderArray = function (values, typeName) {
	// So here I will create an array inside an array...
	if (values.length > 0) {
		// I will use the first element of the array to determine how i will 
		// create that new array.
		var div = new Element(null, { "tag": "div" });
		typeName = typeName.replace("[]", "")
		if (typeName.startsWith("xs.") || typeName.endsWith(":Ref")) {
			div = new Element(null, { "tag": "div", "style": "display: table; width: 100%;" });
			for (var i = 0; i < values.length; i++) {
				var row = div.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down()
				var formatedValue = this.render(values[i], typeName.replace("[]", ""));
				if (formatedValue != undefined) {
					formatedValue.element.style.display = "table-cell"
					formatedValue.element.style.verticalAlign = "middle";
					formatedValue.element.style.width = "100%";
					row.appendElement(formatedValue)
					// Now I will set the double click event for the subvalue.
					formatedValue.element.ondblclick = function (renderer, index) {
						return function (e) {
							e.stopPropagation()
							renderer.cell.setCellEditor(index)
						}
					}(this, i)
				}
			}
		} else {
			// Here I will asynchronously get all items of that types.
			server.entityManager.getEntityPrototype(typeName, typeName.split(".")[0],
				/** The success callback */
				function (prototype, caller) {
					// Here I got an array of entities.
					var model = null;
					model = new EntityTableModel(prototype);
					model.entities = values;
					div.element.className = "entity_sub-table";
					var table = new Table(randomUUID(), div)
					table.setModel(model, function (table) {
						return function () {
							// init the table.
							table.init()
						}
					}(table))
				},
				/** The error callva */
				function (errObj, caller) {
					// Here the div contain a table of values.
				}, div)
		}
		return div;
	}
}

/**
 * Format entity.
 */
TableCellRenderer.prototype.renderEntity = function (value, typeName) {
	// Here I will get 
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

	this.filterPanel = this.filterPanelDiv.appendElement({ "tag": "div", "class": "filter_panel_scroll" }).down()
		.appendElement({ "tag": "div", "class": "filter_panel" }).down()

	// Now the button...
	var filterPanelButtons = this.filterPanelDiv.appendElement({ "tag": "div", "class": "filter_panel_buttons" }).down()

	this.okBtn = filterPanelButtons.appendElement({ "tag": "div", "innerHtml": "ok", "style": "border-right: 1px solid;" }).down()
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
		// single value
		if (!this.type.startsWith("[]")) {
			// Get unique value...
			var value = this.table.model.getValueAt(i, this.index)
			value = formatValue(value, this.type)
			var j = 0
			for (j = 0; j < values.length; j++) {
				if (values[j] == value) {
					break
				}
			}
			// replace existing value if already there...
			values[j] = value
		} else {
			// multiple values
			var values_ = this.table.model.getValueAt(i, this.index)
			for (var j = 0; j < values_.length; j++) {
				value = formatValue(values_[j], this.type)
				var k = 0
				for (k = 0; k < values.length; k++) {
					if (values[k] == value) {
						break
					}
				}
				// replace existing value if already there...
				values[k] = value
			}
		}
		// In case of integer number.
		if (this.type == "int" || this.type == "real" || this.type == "float") {
			values.sort(function sortNumber(a, b) {
				return a - b;
			})
		} else {
			values.sort()
		}
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
			this.appendFilter(values[i].toString())
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