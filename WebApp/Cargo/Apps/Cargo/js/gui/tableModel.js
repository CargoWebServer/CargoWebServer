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

/**
 * This is an abstract class that help to connect the table with the data to display
 * by the Table class.
 * @constructor
 * @param titles The list of table headers
 */
var TableModel = function (titles) {

    // The titles to display in the headers, this must follow the fields...
    this.titles = titles
    this.values = []
    this.fields = []
    this.editable = {}

    if (this.titles != null) {
        for (var i = 0; i < this.titles.length; i++) {
            this.editable[i] = false
        }
    }
    return this
}

/**
 * Initialyse the data of the table.
 * @param {function} callback The function to call when the funtion has finish it's work.
 */
TableModel.prototype.init = function (successCallback, progressCallback, errorCallback, caller) {
    if (successCallback != undefined) {
        successCallback(undefined, caller)
    }
}

/**
 *  Return the number of row of the model.
 *  @returns {int} The number of row of the model.
 */
TableModel.prototype.getRowCount = function () {
    if (this.values == undefined) {
        this.values = []
    }
    return this.values.length
}

/**
 *  Return the number of column of the model
 *  @returns The number of colunm display in the model.
 */
TableModel.prototype.getColumnCount = function () {
    if (this.titles == undefined) {
        this.titles = []
    }
    return this.titles.length
}

/**
 * Get the value of a cell.
 * @param row The row index.
 * @param column The colum index.
 */
TableModel.prototype.getValueAt = function (row, column) {
    return this.values[row][column]
}

/**
 * Get the column header value
 * @param column Return the diplayed value at a given index.
 */
TableModel.prototype.getColumnName = function (column) {
    return this.titles[column]
}

/*
 * Get the data type of a column
 * @param column The datatype of a given column, usefull to know how to display the cell content.
 */
TableModel.prototype.getColumnClass = function (column) {
    return this.fields[column]
}

/**
 * Get the editable state of a cell.
 * @param {int} row The row index
 * @param {int} column The column index
 */
TableModel.prototype.isCellEditable = function (column) {
    return this.editable[column]
}

TableModel.prototype.setIsCellEditable = function (column, val) {
    this.editable[column] = val
}

/**
 * Set the value of a given cell, the change are local util the summit function is call.
 * @param value The value to set
 * @param {int} row The row of index of the cell
 * @param {int} column The column index of the cell
 */
TableModel.prototype.setValueAt = function (value, row, column) {
    // Must be implemented in the derived class.
    this.values[row][column] = value
}

TableModel.prototype.appendRow = function (values) {
    this.values.push(values)
    return values
}

TableModel.prototype.removeRow = function (rowIndex) {

}

TableModel.prototype.removeAllValues = function (rowIndex) {
    this.values = []
}

////////////////////////////////////////////////////////////
// Entity Table model
////////////////////////////////////////////////////////////

/**
 * Implementation of the TableModel for the Entity class. It display
 * the content of entities in tabular form.
 * @extends TableModel
 * @constructor
 * @param proto The entity prototype.
 */
var EntityTableModel = function (proto) {

    // Here i will display the 
    if (proto != undefined) {
        this.proto = proto
    }

    var titles = []
    TableModel.call(this, titles);

    /* I will simply get all entities for the given type. **/

    // Here I will intialyse the fields and corresponding title...
    for (var i = 0; i < this.proto.Fields.length; i++) {
        if (proto.FieldsVisibility[i] == true) {
            var fieldIndex = proto.FieldsOrder[i]
            titles.push(proto.Fields[fieldIndex].replace("M_", ""))
            this.fields.push(proto.FieldsType[fieldIndex])
        }
    }

    /*
     * Here I will keep the entities display in the model...
     */
    this.entities = []

    return this
}

// inherit from TableModel object...
EntityTableModel.prototype = new TableModel(null);
EntityTableModel.prototype.constructor = EntityTableModel;

/**
 * Initialisation of the table model.
 */
EntityTableModel.prototype.init = function (successCallback, progressCallback, errorCallback, caller) {
    if (successCallback != undefined) {
        successCallback(undefined, caller)
    }
}

EntityTableModel.prototype.getParentUuid = function () {
    if (this.entities.length > 0) {
        return this.entities[0].ParentUuid
    }
    return undefined
}

/**
 * Remove all the element from the talbe.
 */
EntityTableModel.prototype.removeAllValues = function () {
    this.entities = []
    this.values = []
}

/**
 * Remove a row whit a given id from the model.
 * @param {string} The id of the row to remove.
 */
EntityTableModel.prototype.removeRow = function (rowIndex) {
    var entity = this.entities[rowIndex]
    if (entity == undefined) {
        return
    }

    var parentEntity = entities[entity.ParentUuid]


    // Now I will get the data type for that type.
    var parentPrototype = entityPrototypes[parentEntity.TYPENAME]
    var field = entity.parentLnk
    var fieldType = parentPrototype.FieldsType[parentPrototype.getFieldIndex(field)]

    var isRef = fieldType.endsWith(":Ref")
    if (isRef) {
        // I here I simple new to remove the reference from the parent 
        // entity.
        //alert("Remove entity reference!")
        removeObjectValue(parentEntity, field, entity)
    } else {
        // Here I need to delete the entity itself.
        if (entity != null) {
            // Here I will ask the user if here realy want to remove the entity...
            var confirmDialog = new Dialog(randomUUID(), undefined, true)
            confirmDialog.div.element.style.maxWidth = "450px"
            confirmDialog.setCentered()
            server.languageManager.setElementText(confirmDialog.title, "delete_dialog_entity_title")
            var prototype = entityPrototypes[entity.TYPENAME]
            var id = prototype.Ids[1] // 0 is the uuid...
            if (id == undefined) {
                id = "uuid"
            }

            var index = prototype.Indexs[1] // 0 is the uuid...
            var label = entity[index]
            if (label == undefined) {
                label = entity[id]
            } else if (label.length == 0) {
                label = entity[id]
            }
            confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete entity " + label + "?" })
            confirmDialog.ok.element.onclick = function (dialog, entity, model) {
                return function () {
                    // I will call delete file
                    server.entityManager.removeEntity(entity.UUID,
                        // Success callback 
                        function (result, caller) {
                            /** The action will be done in the event listener */
                        },
                        // Error callback
                        function (errMsg, caller) {

                        }, entity)
                    dialog.close()

                    // Now set the model values and entities.
                    var values = []
                    var entities = []
                    for (var i = 0; i < model.entities.length; i++) {
                        if (model.entities[i].UUID != entity.UUID) {
                            if (this.entities != undefined) {
                                entities.push(this.entities[i])
                                values.push(this.values[i])
                            }
                        }
                    }
                    model.values = values
                    model.entities = entities

                }
            } (confirmDialog, entity, this)
        }
    }
}

/**
 * Append a new row in the model.
 * @param {values} The entity to append to the model.
 */
EntityTableModel.prototype.appendRow = function (values) {
    if (values == undefined) {
        return []
    }

    this.entities.push(values)
    var isListOf_ = isListOf(this.proto.TypeName)
    var objectValues = []
    var prototype = entityPrototypes[values.TYPENAME]

    for (var j = 0; j < this.titles.length; j++) {
        var field = "M_" + this.titles[j]
        var fieldIndex = prototype.getFieldIndex(field)
        var fieldType = prototype.FieldsType[fieldIndex]
        if (values[field] != undefined) {
            objectValues[j] = values[field]
        } else {
            var isArray = fieldType.startsWith("[]")
            var isRef = fieldType.endsWith(":Ref")
            if (isArray) {
                objectValues[j] = []
            } else if (isRef) {
                objectValues[j] = null
            } else {
                objectValues[j] = null
            }
        }
        if (j == this.titles.length - 1) {
            // In that case the value is a list of base value.
            if (isListOf_) {
                // here I will append the values...
                objectValues.push(values.values)
            }
            this.values.push(objectValues)
        }
    }

    return objectValues
}

/**
 * Set entity propertie value.
 * @param {} value The map of value. Can be an entity or a value.
 * @param {int} row The row index.
 * @param {int} column The column index.
 */
EntityTableModel.prototype.setValueAt = function (value, row, column) {
    var field = "M_" + this.titles[column]
    if (value[field] != undefined) {
        this.entities[row][field] = value[field]
        this.values[row][column] = value[field]
    } else {
        this.entities[row][field] = value
        this.values[row][column] = value
    }
}

/**
 * Save the value contain at a given row.
 */
EntityTableModel.prototype.saveValue = function (row) {
    var entity = row.table.model.entities[row.index]

    // Here I will save the entity...
    if (entity != null) {
        entity.NeedSave = true
        if (entity.exist == false) {
            // Remove the tmp entity...
            server.entityManager.createEntity(entity.ParentUuid, entity.parentLnk, entity.TYPENAME, entity.M_id, entity,
                // Success callback
                function (entity, table) {

                },
                // Error callback.
                function (result, caller) {

                }, row.table)
        } else {
            server.entityManager.saveEntity(entity,
                function (result, row) {

                }, function () {

                }, this)
        }
    }
}

////////////////////////////////////////////////////////////
// Sql Table model
////////////////////////////////////////////////////////////
/**
 * This implement de model for a sql table.
 * @param db The name of the database.
 * @param query The query that generate the value.
 * @param fields  The data type for each column
 * @param params  The parameter of the where close if any
 * @param titles  The column textual title.
 * @returns {*}
 * @constructor
 */
var SqlTableModel = function (db, query, fields, params, titles) {
    TableModel.call(this, titles);

    // Sql base request...
    this.db = db
    this.params = params
    this.fields = fields

    // The query...
    this.query = query.replace(", ", ",") // Read query...

    return this
}

// inherit from TableModel object...
SqlTableModel.prototype = new TableModel(null);
SqlTableModel.prototype.constructor = SqlTableModel;

////////////////////////////////////////////////////////////
// Sql Table model function
////////////////////////////////////////////////////////////
/**
 * Initialyse the sql data...
 * @param callback the function to call when the initialisation is done...
 */
SqlTableModel.prototype.init = function (successCallback, progressCallback, errorCallback, caller) {

    server.dataManager.read(this.db, this.query, this.fields, this.params,
        function (result, caller) {
            caller.tableModel.values = result[0]
            caller.successCallback(result[0], caller.caller)
        },
        // Progress call back
        function (index, total, caller) { // The progress callback
            // Keep track of the file transfert...
            caller.progressCallback(index, total, caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback...
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message...
            server.errorManager.onError(errMsg)
        },
        { "tableModel": this, "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback })
}


/**
 * Remove a row whit a given id from the model.
 * @param {string} The id of the row to remove.
 */
SqlTableModel.prototype.removeRow = function (id) {
    var values = []
    for (var i = 0; i < this.values.length; i++) {
        // The id must be the first value of the row...
        if (this.values[i][0] != id) {
            values.push(this.values[i])
        }
    }
    // Set the new array...
    this.values = values
}

/**
 * Append a new row in the model.
 * @param {values} The array of values to append to the model.
 */
SqlTableModel.prototype.appendRow = function (values) {
    this.values.push(values)
    return values
}

SqlTableModel.prototype.saveValue = function (row) {
    // here i will get the values and save it.
    var queryValues = this.query.split(" ")
    var fields = queryValues[1].split(",")
    var from = queryValues[2]
    var tableName = queryValues[3]

    var params = new Array()
    var updateQuery = ""
    var values = []

    // So here I will create the update query whit the info I know...
    updateQuery += "Update " + tableName + " Set "
    // The first column must be the id
    for (var i = 1; i < fields.length; i++) {
        updateQuery += fields[i] + "=?"
        var val = this.values[row.index][i]

        // format date...
        if (this.getColumnClass(i) == 'date') {
            val = moment(val).format('YYYY-MM-DD HH:mm:ss.SSS');
        }

        values.push(val)

        if (i < fields.length - 1) {
            updateQuery += ","
        }
    }

    // The first row must be the id.
    updateQuery += " WHERE " + fields[0] + "=?"
    params.push(this.values[row.index][0])

    // Send the update query...
    server.dataManager.update(this.db, updateQuery, values, params,
        // Success call back
        function (result, caller) {
            // Update the model value...
        },
        // Error callback
        function () {

        },
        {})
}

/**
 * Set model propertie value.
 * @param {} value The map of value.
 * @param {int} row The row index.
 * @param {int} column The column index.
 */
SqlTableModel.prototype.setValueAt = function (value, row, column) {

    // Check if the value need to be update...
    if (this.getValueAt(row, column) == value) {
        return /* Nothing todo here **/
    }

    // The query values...
    var queryValues = this.query.split(" ")
    var select = queryValues[0]
    var fields = queryValues[1].split(",")
    var from = queryValues[2]
    var tableName = queryValues[3]

    if (this.values == null) {
        return
    }

    /* The the dataManager... **/
    var values = new Array()
    values[0] = value
    if (this.getColumnClass(column) == 'date') {
        values[0] = moment(value).format('YYYY-MM-DD HH:mm:ss');
    }

    var params = new Array()
    var updateQuery = ""

    // So here I will create the update query whit the info I know...
    updateQuery += "Update " + tableName + " Set " + fields[column] + "=? WHERE "

    // The other field will be use has parameters...
    var index = 0
    for (var i = 0; i < /*fields.length*/ 1; i++) {
        var val = this.getValueAt(row, i)
        if (i != column && this.getColumnClass(i) != 'real') {
            if (i < fields.length && i != 0) {
                updateQuery += " AND "
            }
            if (this.getColumnClass(i) == 'string') {
                updateQuery += fields[i] + "LIKE ? "
            } else {
                updateQuery += fields[i] + "=? "
            }

            if (this.getColumnClass(i) == 'date') {
                params[index] = moment(val).format('YYYY-MM-DD HH:mm:ss.SSS');
            } else {
                params[index] = val
            }
            index += 1
        }

    }

    // Send the update query...
    server.dataManager.update(this.db, updateQuery, values, params,
        // Success call back
        function (result, caller) {
            // Update the model value...
            var model = caller.caller
            var rowIndex = caller.rowIndex
            var colIndex = caller.colIndex
            var value = caller.value
            model.values[rowIndex][colIndex] = value
        },
        // Error callback
        function () {

        },
        { "caller": this, "rowIndex": row, "colIndex": column, "value": value })
}