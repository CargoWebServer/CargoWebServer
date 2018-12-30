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

// I made use of that map to store model reference.
var models = {}

/**
 * This is an abstract class that help to connect the table with the data to display
 * by the Table class.
 * @constructor
 * @param titles The list of table headers
 */
var TableModel = function (titles) {

    this.id = randomUUID()

    // The table that use this model.
    this.table = null
    // The titles to display in the headers, this must follow the fields...
    this.titles = titles
    this.values = []
    this.fields = []
    this.editable = []

    if (this.titles != null) {
        for (var i = 0; i < this.titles.length; i++) {
            this.editable[i] = false
        }
    }

    // keep it in the model map.
    models[this.id] = this

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

    // The append row table button action for generic table...
    // TODO the table is one dimension, make it n dimensions...
    this.table.appendRowBtn.element.onclick = function (table, model) {
        return function () {
            // The value to append.
            var values = []
            var index = model.values.length

            for (var i = 0; i < model.fields.length; i++) {
                var fieldType = model.fields[i]
                if (model.ParentLnk == "M_listOf") {
                    var baseType = getBaseTypeExtension(fieldType)
                    if (baseType.startsWith("xs.")) {
                        fieldType = baseType
                    }
                }

                // Now the value to append.
                if (isXsString(fieldType)) {
                    values.push("")
                } else if (isXsInt(fieldType)) {
                    values.push(0)
                } else if (isXsNumeric(fieldType)) {
                    values.push(0.0)
                } else if( isXsBoolean(fieldType)){
                    values.push(false)
                }
            }

            model.appendRow(values, index)
            if (table.rows.length == 0) {
                table.setHeader()
                table.header.maximizeBtn.element.click()
            }

            var row = new TableRow(table, index, values, index)
            row.table = table
            table.rows.push(row)
            
            // So here I will append the new value in the table. Values are created in 
            // respect of the predefined types.
            simulate(row.cells[index, 0].div.element, "dblclick");
        }
    }(this.table, this)
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
    return true //this.editable[column] TODO uncomment!
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
    if (this.values[row][column] != value) {
        this.values[row][column] = value
        // save the value.
        this.saveValue(row)
    }
}

TableModel.prototype.appendRow = function (values) {
    if(isArray(values)){
        this.values.push(values)
    }else{
        this.values.push([values])
    }
    
    return values
}

TableModel.prototype.removeRow = function (rowIndex) {
    this.values.splice(rowIndex, 1)

    // save the value.
    this.saveValue(rowIndex)
}

TableModel.prototype.removeAllValues = function (rowIndex) {
    this.values = []
}

TableModel.prototype.saveValue = function (rowIndex) {

    // Set the save value function.
    var fieldName = this.ParentLnk
    var entity = entities[this.ParentUuid]
    if (entity != undefined) {
        var values = []
        // In case of M_listOf the values will be one dimensional array
        if(fieldName=="M_listOf"){
            for(var i=0; i < this.values.length; i++){
                values.push(this.values[i][0])
            }
        }else{
            // Tow dimensional array here.
            values = this.values
        }
        entity[fieldName] = values
        server.entityManager.saveEntity(entity,
            // success callback
            function (results, caller) { 
                caller.table.header.maximizeBtn.element.click()
            },
            // error callback
            function () { },
            // caller
            this)
    }
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
var EntityTableModel = function (proto, query) {

    // Here i will display the 
    if (proto != undefined) {
        this.proto = proto
    }

    this.query = query
    var titles = []
    TableModel.call(this, titles);

    /* I will simply get all entities for the given type. **/

    // Here I will intialyse the fields and corresponding title...
    if (this.query == null) {
        for (var i = 0; i < this.proto.Fields.length; i++) {
            if (proto.FieldsVisibility[i] == true) {
                var fieldIndex = proto.FieldsOrder[i]
                titles.push(proto.Fields[fieldIndex].replace("M_", ""))
                this.fields.push(proto.FieldsType[fieldIndex])
                this.editable.push(false)
            }
        }
    } else {
        for (var i = 0; i < this.query.Fields.length; i++) {
            var fieldIndex = proto.getFieldIndex(this.query.Fields[i])
            if (proto.FieldsVisibility[fieldIndex] == true) {
                titles.push(proto.Fields[fieldIndex].replace("M_", ""))
                this.fields.push(proto.FieldsType[fieldIndex])
                this.editable.push(false)
            }
        }
    }

    /*
     * Here I will keep the entities display in the model...
     */
    this.entities = []

    /**
     * If the it's a propertie of a parent entity
     */
    this.ParentLnk = proto.ParentLnk

    return this
}

// inherit from TableModel object...
EntityTableModel.prototype = new TableModel(null);
EntityTableModel.prototype.constructor = EntityTableModel;

/**
 * Initialisation of the table model.
 */
EntityTableModel.prototype.init = function (successCallback, progressCallback, errorCallback, caller) {

    if (this.query != null) {
        server.entityManager.getEntities(this.proto.TypeName, this.proto.TypeName.split(".")[0], this.query, 0, -1, [], true, false,
            // Progress callback
            function (index, total, caller) {
                // nothing to do here
                if (caller.progressCallback != undefined) {
                    caller.progressCallback(index, total, caller.caller)
                }
            },
            // Success callack
            function (results, caller) {

                var table = caller.caller.caller
                for (var i = 0; i < results.length; i++) {
                    table.getModel().appendRow(results[i], results[i].UUID)
                }

                if (caller.successCallback != undefined) {
                    caller.successCallback(results, caller.caller)
                }

                // typeName string, storeId string, queryStr string, offset int, limit int, orderBy []interface{}, asc bool
                if (caller.caller.initCallback != undefined) {
                    // init the table.
                    caller.caller.initCallback()
                    caller.caller.initCallback = undefined
                }
                server.languageManager.setLanguage()
            },
            // Error Callback
            function (errObj, caller) {
                if (caller.errorCallback != undefined) {
                    caller.errorCallback(errObj, caller.caller)
                }
            }, { "successCallback": successCallback, "errorCallback": errorCallback, "progressCallback": progressCallback, "caller": caller })
    } else {
        // typeName string, storeId string, queryStr string, offset int, limit int, orderBy []interface{}, asc bool
        if (caller.initCallback != undefined) {
            for (var i = 0; i < this.entities.length; i++) {
                this.appendRow(this.entities[i], this.entities[i].UUID)
            }
            caller.initCallback()
            caller.initCallback = undefined
            server.languageManager.setLanguage()
        }
    }

    // Now the append row button.
    this.table.appendRowBtn.element.onclick = function (table, model) {
        return function () {
            var entity = eval("new " + model.proto.TypeName + "()")

            entity.ParentLnk = model.ParentLnk
            entity.ParentUuid = model.ParentUuid;

            var data = model.appendRow(entity)

            // The row is not append in the table rows collection, but display.
            var lastRowIndex = table.rows.length
            var row = new TableRow(table, lastRowIndex, data, undefined)
            table.rows.push(row)

            // In that case the interface use to edit entity properties is the 
            // row.
            entity.getPanel = function (row) {
                return function () {
                    return row
                }
            }(row)

            simulate(row.cells[lastRowIndex, 0].div.element, "dblclick");
        }
    }(this.table, this)

    // Now the edit permission...
    /*if (server.sessionId != undefined) {
        server.accountManager.me(server.sessionId,
            function (result, caller) {
                if (result.M_id == "admin") {
                    for (var i = 0; i < caller.editable.length; i++) {
                        caller.editable[i] = true;
                    }
                }
            }, function (errObj, caller) {

            }, this)
    }*/

    ////////////////////////////////////////////////////////////////////////////////////////////////
    // Event listener connection here.
    ///////////////////////////////////////////////////////////////////////////////////////////////
    server.entityManager.attach(this, DeleteEntityEvent, function (evt, model) {
        // So here I will remove the line from the table...
        var toDelete = evt.dataMap["entity"]
        if (entities[toDelete.UUID] != undefined) {
            toDelete = entities[toDelete.UUID]
            // remove it from the map.
            delete entities[toDelete.UUID]
        }

        if (toDelete.TYPENAME == model.proto.TypeName || model.proto.SubstitutionGroup.indexOf(toDelete.TYPENAME) != -1) {
            // So here I will remove the line from the model...
            var orderedRows = []
            for (var i = 0; i < model.table.orderedRows.length; i++) {
                var row = model.table.orderedRows[i]
                if (row.id != toDelete.UUID) {
                    orderedRows.push(row)
                } else {
                    // remove from the display...
                    row.div.element.parentNode.removeChild(row.div.element)
                }
            }

            // Now set the model values and entities.
            var entities_ = []
            var rows = []

            for (var i = 0; i < model.entities.length; i++) {
                var row = model.table.rows[i]
                if (model.entities[i].UUID != toDelete.UUID) {
                    var entity = model.entities[i]
                    if (entities[entity.UUID] != undefined) {
                        entity = model.entities[i] = entities[entity.UUID]
                    }
                    entities_.push(entity)
                    row.id = entity.UUID
                    row.index = rows.length;
                    rows.push(row)
                } else {
                    if (row.div.element.parentNode != null) {
                        row.div.element.parentNode.removeChild(row.div.element)
                    }
                }
            }

            // Set the values...
            model.entities = entities_
            model.table.orderedRows = orderedRows
            model.table.rows = rows
            model.table.refresh()
        }
    })

    // The new entity event...
    server.entityManager.attach(this, NewEntityEvent, function (evt, model) {
        model = models[model.id]
        if (evt.dataMap["entity"] != undefined) {
            var entity = entities[evt.dataMap["entity"].UUID]
            if (entity != undefined) {
                if (entity.TYPENAME == model.proto.TypeName || model.proto.SubstitutionGroup.indexOf(entity.TYPENAME) != -1) {
                    // If the object does not exist in the array.
                    if (!objectPropInArray(model.entities, "UUID", entity.UUID)) {
                        var exist = false
                        for (var i = 0; i < model.entities.length; i++) {
                            if (model.entities[i].UUID == undefined) {
                                model.entities.splice(i, 1)
                            } else if (model.entities[i].UUID == entity.UUID) {
                                exist = true
                            }
                        }
                        if (!exist && model.ParentUuid == entity.ParentUuid) {
                            var row = model.table.appendRow(entity, entity.UUID)
                            model.table.refresh()
                        }
                    }
                }
            }
        }
    })

    // Set the update listener for each row entity...
    server.entityManager.attach(this, UpdateEntityEvent, function (evt, model) {
        model = models[model.id]
        if (evt.dataMap["entity"] != undefined) {
            var entity = entities[evt.dataMap["entity"].UUID]
            if (entity != undefined) {
                for (var i = 0; i < model.entities.length; i++) {
                    if (model.entities[i] != undefined) {
                        if (model.entities[i].TYPENAME == entity.TYPENAME) {
                            if (model.entities[i].UUID == entity.UUID) {
                                model.entities[i] = entity;
                                // Replace the value of the existing row with the entity value.
                                var row = model.table.appendRow(entity, entity.UUID)
                                break;
                            }
                        }
                    }
                }
            }
        }
    })
}


/**
 * Remove all the element from the talbe.
 */
EntityTableModel.prototype.removeAllValues = function () {
    this.entities = []
}

/**
 * Remove a row whit a given index from the model.
 * @param {string} The index of the row to remove.
 */
EntityTableModel.prototype.removeRow = function (rowIndex, callback) {
    var entity = this.entities[rowIndex]
    if (entity == undefined) {
        return
    }

    var parentEntity = entities[entity.ParentUuid]
    var isRef = false
    if (parentEntity != null) {
        // Now I will get the data type for that type.
        var parentPrototype = getEntityPrototype(parentEntity.TYPENAME)
        var field = entity.ParentLnk
        var fieldType = parentPrototype.FieldsType[parentPrototype.getFieldIndex(field)]
        isRef = fieldType.endsWith(":Ref")
    }

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
            var prototype = getEntityPrototype(entity.TYPENAME)
            var id = prototype.Ids[1] // 0 is the uuid...
            if (id == undefined) {
                id = "uuid"
            }

            var index = prototype.Indexs[1] // 0 is the uuid...

            confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Do you want to delete entity " + entity.getTitles()[0] + "?" })
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
                }
            }(confirmDialog, entity, callback, this)
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

    if (!objectPropInArray(this.entities, "UUID", values.UUID)) {
        this.entities.push(values)
    }

    var objectValues = []
    if (values.TYPENAME != undefined) {
        var prototype = values.getPrototype()

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
                this.values.push(objectValues)
            }
        }
    } else {
        console.log("values not entity ", values)
    }

    this.table.setHeader();
    this.table.refresh();

    return objectValues
}

/**
 * Get the value of a cell.
 * @param row The row index.
 * @param column The colum index.
 */
EntityTableModel.prototype.getValueAt = function (row, column) {
    var entity = this.entities[row]
    if (entity == undefined) {
        return null;
    }

    if (entities[entity.UUID] != undefined) {
        entity = this.entities[row] = entities[entity.UUID]
    }
    var field = "M_" + this.titles[column]
    var value = entity[field]
    if (isArray(value)) {
        // remove previously deleted values.
        for (var i = 0; i < value.length; i++) {
            if (isObject(value[i])) {
                if (value[i].UUID != undefined) {
                    if (entities[value[i].UUID] == null) {
                        value.splice(i, 1)
                    }
                }
            }
        }
    }

    // in case of array.
    if (value == null) {
        if (this.proto.FieldsType[this.proto.getFieldIndex(field)].startsWith("[]")) {
            value = [];
        }
    }

    return value
}

/**
 * Set entity propertie value.
 * @param {} value The map of value. Can be an entity or a value.
 * @param {int} row The row index.
 * @param {int} column The column index.
 */
EntityTableModel.prototype.setValueAt = function (value, row, column) {
    var entity = this.entities[row]
    if (entities.UUID != undefined) {
        if (entities[entity.UUID] != undefined) {
            entity = this.entities[row] = entities[entity.UUID]
        }
    }
    var field = "M_" + this.titles[column]
    if (entity != undefined) {
        if (value[field] != undefined) {
            entity[field] = value[field]
        } else {
            entity[field] = value
        }
    }

    // save the value.
    this.saveValue(row)
}

/**
 * Save the value contain at a given row.
 */
EntityTableModel.prototype.saveValue = function (rowIndex) {
    var entity = this.entities[rowIndex]
    // Here I will save the entity...
    if (entity != null) {
        // Always use the entity from the enities map it contain
        // the valid data of the entity.
        if (entities[entity.UUID] != null) {
            entity = this.entities[rowIndex] = entities[entity.UUID]
        }

        if (entity.exist == false) {
            // I must remove the temporary object it will be recreated when the NewEntity event will be received.
            this.entities.pop(rowIndex)
            var row = this.table.rows.pop(rowIndex)
            row.div.element.parentNode.removeChild(row.div.element)

            server.entityManager.createEntity(entity.ParentUuid, entity.ParentLnk, entity,
                // Success callback
                function (entity, row) {

                },
                // Error callback.
                function (result, caller) {

                }, row)
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
            if (caller.successCallback != undefined) {
                caller.successCallback(result[0], caller.caller)
                caller.successCallback = undefined
            }
        },
        // Progress call back
        function (index, total, caller) { // The progress callback
            // Keep track of the file transfert...
            if (caller.progressCallback != undefined) {
                caller.progressCallback(index, total, caller.caller)
                caller.progressCallback = undefined
            }
        },
        function (errMsg, caller) {
            // call the immediate error callback...
            if (caller.errorCallback != undefined) {
                caller.errorCallback(errMsg, caller.caller)
                caller.errorCallback = undefined
            }
            // dispatch the message...
            server.errorManager.onError(errMsg)
        },
        { "tableModel": this, "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback })

    /* The delete row event **/
    server.dataManager.attach(this, DeleteRowEvent, function (evt, table) {
        // So here I will remove the line from the table...
        if (evt.dataMap.tableName == table.id) {
            // So here I will remove the line from the model...F
            table.getModel().removeRow(evt.dataMap.id_0)

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
 * Get the value of a cell.
 * @param row The row index.
 * @param column The colum index.
 */
SqlTableModel.prototype.getValueAt = function (row, column) {
    return this.values[row][column]
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