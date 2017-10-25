/**
 * The query editor is use to edit query in EQL and SQL language.
 * It also contain a table to display it results.
 */

var QueryEditor = function (parent, file, initCallback) {

    this.initCallback = initCallback
    this.parent = parent
    this.file = file
    this.dataConfigs = {}
    this.activeDataConfig = null
    this.queryToolBar = null
    this.dataSelect = null

    // result navigation.
    this.moveFirstBtn = null
    this.moveLeftBtn = null
    this.nagivatorState = null
    this.moveRightBtn = null
    this.moveLastBtn = null

    // Keep the number of row that can be displayed.
    this.countRows = 0
    this.maxRows = 5000 // Must be set in config.

    // Those info came from the query itself.
    this.fromRows = undefined
    this.nbRows = 0

    // Sql data source.
    this.isSql = file.M_name.endsWith(".sql") || file.M_name.endsWith(".SQL")

    // Entity data source.
    this.isEql = file.M_name.endsWith(".eql") || file.M_name.endsWith(".EQL")

    this.panel = parent.appendElement({ "tag": "div", "class": "query_editor" }).down()
    this.mainArea = this.panel.appendElement({ "tag": "div", "style": "display: table; width:100%; height:100%" }).down()

    // So the panel will be divide in tow parts...
    // The query panel.
    var splitArea1 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-row; position: relative; with:100%; height: auto;" }).down()

    // The edition panel.
    this.editQueryPanel = new Element(splitArea1, { "tag": "div", "class": "edit_query_panel" })

    // The splitter.
    var queryEditorSplitor = this.mainArea.appendElement({ "tag": "div", "class": "splitter horizontal", "id": "query_editor_splitor" }).down()

    // The result panel.
    var splitArea2 = this.mainArea.appendElement({ "tag": "div", "style": "display: table-row; position: relative; width: 100%; height:100%;" }).down()

    this.resultQueryPanel = new Element(splitArea2, { "tag": "div", "class": "result_query_panel", "style": "position: relative;" })

    // Init the splitter action.
    initSplitter(queryEditorSplitor, this.editQueryPanel)

    var filePanel = this.editQueryPanel.appendElement({ "tag": "div", "class": "filePanel", "id": file.M_id + "_query_editor", "innerHtml": decode64(file.M_data) }).down()
    this.editor = ace.edit(file.M_id + "_query_editor");

    // In case of sql query..
    if (this.isSql) {
        this.editor.getSession().setMode("ace/mode/sql");

    } else if (this.isEql) {
        // The file mode of the edior is simple javascript.
        this.editor.getSession().setMode("ace/mode/javascript");
    }

    this.editor.getSession().on('change', function (fileId, fileUUID, editor) {
        return function () {
            var evt = { "code": ChangeFileEvent, "name": FileEvent, "dataMap": { "fileId": fileId } }
            var file = entities[fileUUID]
            file.M_data = encode64(editor.getSession().getValue())
            server.eventHandler.broadcastLocalEvent(evt)
        }
    }(file.M_id, file.UUID, this.editor));

    return this
}

/**
 * Initialyse the query editor.
 */
QueryEditor.prototype.init = function () {
    // Now I will get the list of datastore from the server for the given type.
    server.entityManager.getEntities("Config.DataStoreConfiguration", "Config", "", 0, -1, [], true,
        // progress
        function () {
            // nothing here
        },
        // success
        function (results, caller) {
            var dataConfigs = []
            for (var i = 0; i < results.length; i++) {
                if (results[i].M_dataStoreType == 1 && caller.isSql) {
                    // Sql datatype
                    dataConfigs.push(results[i])
                } else if (results[i].M_dataStoreType == 2 && caller.isEql) {
                    // Eql datatype
                    dataConfigs.push(results[i])
                }
            }

            // Now I will set the data configs...
            caller.setDataConfigs(dataConfigs)

        },
        // error
        function () {
            // nothing here
        }, this)
}

/**
 * Set the data configuration.
 */
QueryEditor.prototype.setDataConfigs = function (configs) {
    // Keep ref to the data configuration.
    for (var i = 0; i < configs.length; i++) {
        this.dataConfigs[configs[i].M_id] = configs[i]
    }

    // So here I will create the tool bar for the query editor.
    this.queryToolBar = new Element(homepage.toolbarDiv, { "tag": "div", "class": "toolbar", "id": this.file.M_id + "_toolbar" })

    // The datasource selection.
    this.dataSelect = this.queryToolBar.appendElement({ "tag": "div", "style": "display: table-cell; height: 100%; vertical-align: middle;" }).down()
        .appendElement({ "tag": "select" }).down()

    for (var configId in this.dataConfigs) {
        this.dataSelect.appendElement({ "tag": "option", "innerHtml": configId, "value": configId })
        if (this.activeDataConfig == undefined) {
            this.setActiveDataConfig(configId)
        }
    }

    // Here I will set the current data source
    this.dataSelect.element.onchange = function (queryEditor) {
        return function () {
            queryEditor.setActiveDataConfig(this.value)
        }
    }(this)

    // The query button.
    var playQueryBtn = this.queryToolBar.appendElement({ "tag": "div", "class": "toolbarButton" }).down()
    playQueryBtn.appendElement({ "tag": "i", "id": randomUUID() + "_play_query_btn", "class": "fa fa-bolt" })

    playQueryBtn.element.onclick = function (queryEditor) {
        return function () {
            // Here I will call the query.
            queryEditor.runQuery()
        }
    }(this)

    // Now the pages navigation button to use in case of query that contain limits
    this.queryToolBar.appendElement({ "tag": "div", "class": "toolbarButton", "id": "moveFirstBtn", "style": "display: none;" }).down()
        .appendElement({ "tag": "i", "id": randomUUID() + "_play_query_btn", "class": "fa fa-angle-double-left" }).up()
        .appendElement({ "tag": "div", "class": "toolbarButton", "id": "moveLeftBtn", "style": "display: none;" }).down()
        .appendElement({ "tag": "i", "id": randomUUID() + "_play_query_btn", "class": "fa fa-angle-left" }).up()
        .appendElement({ "tag": "div", "style": "display: none; vertical-align: middle; color: #657383; padding: 0px 10px 0px 10px;", "id": "nagivatorState" })
        .appendElement({ "tag": "div", "class": "toolbarButton", "id": "moveRightBtn", "style": "display: none;" }).down()
        .appendElement({ "tag": "i", "id": randomUUID() + "_play_query_btn", "class": "fa fa-angle-right " }).up()
        .appendElement({ "tag": "div", "class": "toolbarButton", "id": "moveLastBtn", "style": "display: none;" }).down()
        .appendElement({ "tag": "i", "id": randomUUID() + "_play_query_btn", "class": "fa fa-angle-double-right " })

    // Now the action button
    this.moveFirstBtn = this.queryToolBar.getChildById("moveFirstBtn")
    this.moveLeftBtn = this.queryToolBar.getChildById("moveLeftBtn")
    this.nagivatorState = this.queryToolBar.getChildById("nagivatorState")
    this.moveRightBtn = this.queryToolBar.getChildById("moveRightBtn")
    this.moveLastBtn = this.queryToolBar.getChildById("moveLastBtn")

    // The move first button.
    this.moveFirstBtn.element.onclick = function (queryEditor) {
        return function () {
            // So here I will set the from to the last...
            queryEditor.fromRows = 0

            var total = parseInt((queryEditor.countRows / queryEditor.nbRows) + .5)
            queryEditor.nagivatorState.element.innerHTML = "0 | " + total
            // Hide left buttons...
            queryEditor.moveFirstBtn.element.style.display = "none"
            queryEditor.moveLeftBtn.element.style.display = "none"

            // display the right buttons
            queryEditor.moveRightBtn.element.style.display = ""
            queryEditor.moveLastBtn.element.style.display = ""

            queryEditor.runQuery()
        }
    }(this)

    // The move last button.
    this.moveLastBtn.element.onclick = function (queryEditor) {
        return function () {
            // So here I will set the from to the last...
            queryEditor.fromRows = parseInt((queryEditor.countRows / queryEditor.nbRows)) * queryEditor.nbRows
            var total = parseInt((queryEditor.countRows / queryEditor.nbRows) + .5)
            queryEditor.nagivatorState.element.innerHTML = total + " | " + total
            // display right buttons...
            queryEditor.moveFirstBtn.element.style.display = ""
            queryEditor.moveLeftBtn.element.style.display = ""

            // hide the right buttons
            queryEditor.moveRightBtn.element.style.display = "none"
            queryEditor.moveLastBtn.element.style.display = "none"

            queryEditor.runQuery()
        }
    }(this)

    // The move right button.
    this.moveRightBtn.element.onclick = function (queryEditor) {
        return function () {
            // move to nb rows
            queryEditor.fromRows += queryEditor.nbRows

            var total = parseInt((queryEditor.countRows / queryEditor.nbRows) + .5)
            var index = parseInt((queryEditor.fromRows / queryEditor.nbRows) + .5)

            // Hide the button if the index is 
            if (index == total) {
                queryEditor.moveRightBtn.element.style.display = "none"
                queryEditor.moveLastBtn.element.style.display = "none"
            }

            queryEditor.runQuery()
        }
    }(this)

    // Move left.
    this.moveLeftBtn.element.onclick = function (queryEditor) {
        return function () {
            // move to nb rows
            queryEditor.fromRows -= queryEditor.nbRows

            var total = parseInt((queryEditor.countRows / queryEditor.nbRows) + .5)
            var index = parseInt((queryEditor.fromRows / queryEditor.nbRows) + .5)

            // Hide the button if the index is 
            if (index == 0) {
                queryEditor.moveLeftBtn.element.style.display = "none"
                queryEditor.moveLeftBtn.element.style.display = "none"
            }

            queryEditor.runQuery()
        }
    }(this)

    // Call the init callback.
    this.initCallback(this)
}

QueryEditor.prototype.setResultsNavigatorSate = function () {
    if (this.nbRows > 0) {
        if (this.countRows > this.nbRows) {
            var total = parseInt((this.countRows / this.nbRows) + .5)
            var index = this.fromRows / this.nbRows

            if (this.fromRows == 0) {
                this.moveFirstBtn.element.style.display = "none"
                this.moveLeftBtn.element.style.display = "none"
                this.moveRightBtn.element.style.display = ""
                this.moveLastBtn.element.style.display = ""
            } else if (index == total) {
                this.moveRightBtn.element.style.display = "none"
                this.moveLastBtn.element.style.display = "none"
                this.moveFirstBtn.element.style.display = ""
                this.moveLeftBtn.element.style.display = ""
            } else {
                this.moveFirstBtn.element.style.display = ""
                this.moveLeftBtn.element.style.display = ""
                this.moveRightBtn.element.style.display = ""
                this.moveLastBtn.element.style.display = ""
            }

            this.nagivatorState.element.style.display = "table-cell"
            this.nagivatorState.element.innerHTML = index + " | " + total

        } else {
            // In that case there no need for the navigation
            this.moveFirstBtn.element.style.display = "none"
            this.moveLeftBtn.element.style.display = "none"
            this.nagivatorState.element.style.display = "none"
            this.moveRightBtn.element.style.display = "none"
            this.moveLastBtn.element.style.display = "none"
        }
    }
}

/**
 * Set the current database.
 */
QueryEditor.prototype.setActiveDataConfig = function (configId) {
    // Keep ref to the data configuration.
    this.activeDataConfig = this.dataConfigs[configId]
    for (var i = 0; i < this.dataSelect.element.childNodes.length; i++) {
        if (this.dataSelect.element.childNodes[i].value == this.activeDataConfig.M_id) {
            this.dataSelect.element.selectedIndex = i
            break
        }
    }
}

/**
 * Run the query and set the result.
 */
QueryEditor.prototype.runQuery = function () {
    // Keep ref to the data configuration.
    if (this.isSql) {
        // Here I will execute sql query...
        var query = this.editor.getSession().getValue().replaceAll("[", "").replaceAll("]", "").replaceAll(";", "")

        var ast = simpleSqlParser.sql2ast(query)
        // in case of where
        var params = []

        // list of type.
        var fields = []

        // Here I will keep the list of prototype.
        var prototypes = {}

        // Query type...
        var type = "" // Can be CREATE, READ, UPDATE, DELETE

        // Now I will parse 
        if (ast.status == true) {
            // The list of prototypes...
            for (var i = 0; i < ast.value.from.length; i++) {
                var prototype = ast.value.from[i].table
                var values = prototype.split(".")
                if (values.length == 2) {
                    // Tow value can be database_name.table_name or schema_name.table_name
                    // only schema will be keep.
                    for (var sourceId in this.dataConfigs) {
                        if (sourceId == values[0]) {
                            this.setActiveDataConfig(values[0])
                            prototype = values[1]
                            break
                        }
                    }
                }

                // the final prototype id will be database_name.(schema_name.)table_name schema is optional.
                prototype = this.activeDataConfig.M_id + "." + prototype
                prototypes[prototype] = prototype
            }

            // Here There is no syntax error...
            for (var i = 0; i < ast.value.select.length; i++) {
                type = "READ"
                fields.push(ast.value.select[i])
            }
        }

        // Set the row to start and the number of row to display from start.
        if (ast.value.limit != null) {
            if (this.fromRows != undefined) {
                ast.value.limit.from = this.fromRows
            } else {
                this.fromRows = ast.value.limit.from
            }

            this.nbRows = ast.value.limit.nb
        }

        // The query can be formated...
        var query_ = simpleSqlParser.ast2sql(ast)
        if (query_.length > 0) {
            if (query != query_) {
                this.editor.getSession().setValue(query_)
                query = query_
            }
        }

        // Now I will retreive the prototypes from prototypes names and exectute the query...
        for (var prototype in prototypes) {
            // In that case the prototype contain the schema id and table id.
            server.entityManager.getEntityPrototype(prototype, "sql_info",
                // success callback
                function (result, caller) {
                    //console.log(result)
                    var prototypes = caller.prototypes
                    prototypes[result.TypeName] = result

                    var done = true
                    for (var prototype in prototypes) {
                        if (isString(prototypes[prototype])) {
                            done = false
                        }
                    }
                    if (done) {
                        // here I will get the list of field type.
                        var fieldsType = []
                        var fields = []
                        for (var i = 0; i < caller.fields.length; i++) {
                            var field = caller.fields[i]
                            if (field.column != null) {
                                var values = field.column.split(".")
                                if (values.length == 1 && Object.keys(prototypes).length == 1) {
                                    // That's means there's only one table in use for the query.
                                    if (field.column != "*") {
                                        var fieldType = result.FieldsType[result.getFieldIndex("M_" + field.column)]
                                        fieldsType.push(fieldType)
                                        if (field.alias != null) {
                                            fields.push(field.alias)
                                        } else {
                                            fields.push(field.column)
                                        }
                                    } else {
                                        // In that case the user want all the value from the table.
                                        for (var i = 0; i < result.FieldsType.length; i++) {
                                            if (result.FieldsType[i].startsWith("sqltypes.")) {
                                                fieldsType.push(result.FieldsType[i])
                                                fields.push(result.Fields[i].replace("M_", ""))
                                            }
                                        }
                                    }

                                } else {
                                    var tableId = values[0]
                                    var fieldName = values[1]
                                }
                            } else {
                                if (field.alias != null) {
                                    fields.push(field.alias)
                                }

                                if (field.expression != null) {
                                    // The count expression...
                                    if (field.expression.startsWith("COUNT(")) {
                                        fieldsType.push("sqltypes.int")
                                    }
                                }
                            }
                        }
                    }

                    // The list of field types.
                    caller.fieldsType = fieldsType
                    caller.fields = fields

                    // Here I will execute the query...
                    if (caller.type == "READ") {
                        // Here I will get the number of expected result before execute the final query...
                        var ast = simpleSqlParser.sql2ast(caller.query)
                        ast.value.limit = null // remove the limits
                        var query_ = simpleSqlParser.ast2sql(ast)
                        var countQuery = "select count(*) as numRow from (" + query_ + ") src";
                        server.dataManager.read(caller.queryEditor.activeDataConfig.M_id, countQuery, ["sqltypes.int"], [],
                            // success callback
                            function (result, caller) {
                                caller.queryEditor.countRows = result[0][0]
                                caller.queryEditor.setResult(caller.query, caller.fields, caller.fieldsType, [], caller.type)
                            },
                            // progress callback
                            function (index, total, caller) {

                            }
                            , //error callback 
                            function (errMsg, caller) {

                            },
                            caller)
                    }

                },
                // error callback
                function (errObj, caller) {

                }, { "prototypes": prototypes, "type": type, "fields": fields, "query": query, "queryEditor": this })

        }
    } else if (this.isEql) {
        // The code cotain the query to execute.
        var querySrc = this.editor.getSession().getValue()
        var queryVal = "query"

        // (variable name).TypeName must be given by each query...
        for (var i = 0; i < querySrc.split("\n").length; i++) {
            if (querySrc.split("\n")[i].indexOf(".TypeName") != -1) {
                queryVal = querySrc.split("\n")[i].split(".")[0] // The 
                break
            }
        }

        eval(querySrc)
        this.setResult(JSON.stringify(eval(queryVal)), eval(queryVal + ".Fields"), eval(queryVal + ".FieldsType"), [],"READ")
    }
}

/**
 * Set the reuslt table.
 */
QueryEditor.prototype.setResult = function (query, fields, fieldsType, param, type) {
    if (this.isSql) {
        if (type == "READ") {
            // Here I will create a sql table.
            this.resultQueryPanel.removeAllChilds()
            var table = new Table(this.activeDataConfig.M_id, this.resultQueryPanel)
            var model = new SqlTableModel(this.activeDataConfig.M_id, query, fieldsType, [], fields)

            // Set the results navigator state.
            this.setResultsNavigatorSate()

            table.setModel(model, function (table, queryEditor) {
                return function () {
                    // init the table.
                    table.init()
                }
            }(table, this))
        }
    } else if (this.isEql) {
        // In that case 
        if (type == "READ") {
            this.resultQueryPanel.removeAllChilds()
            var table = new Table(this.activeDataConfig.M_id, this.resultQueryPanel)
            var model = new SqlTableModel(this.activeDataConfig.M_id, query, fieldsType, [], fields)
            table.setModel(model, function (table, queryEditor) {
                return function () {
                    // init the table.
                    table.init()
                }
            }(table, this))
        }
    }
}