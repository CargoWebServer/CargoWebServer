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
        this.editor.getSession().on('change', function (fileId, fileUUID, editor) {
            return function () {
                var evt = { "code": ChangeFileEvent, "name": FileEvent, "dataMap": { "fileId": fileId } }
                var file = server.entityManager.entities[fileUUID]
                file.M_data = encode64(editor.getSession().getValue())
                server.eventHandler.BroadcastEvent(evt)
            }
        } (file.M_id, file.UUID, this.editor));
    } else if (this.isEql) {
        // TODO implement the syntax highlight for EQL...
    }

    return this
}

/**
 * Initialyse the query editor.
 */
QueryEditor.prototype.init = function () {
    // Now I will get the list of datastore from the server for the given type.
    server.entityManager.getObjectsByType("Config.DataStoreConfiguration", "Config", "",
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

    // Call the init callback.
    this.initCallback(this)
}

/**
 * Set the current database.
 */
QueryEditor.prototype.setActiveDataConfig = function (configId) {
    // Keep ref to the data configuration.
    this.activeDataConfig = this.dataConfigs[configId]
}

/**
 * Run the query and set the result.
 */
QueryEditor.prototype.runQuery = function () {
    // Keep ref to the data configuration.
    if (this.isSql) {
        // Here I will execute sql query...
        var query = this.editor.getSession().getValue().replaceAll("[", "").replaceAll("]", "")

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
                            this.activeDataConfig = this.dataConfigs[values[0]]
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
                fields.push(ast.value.select[i].column)
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
                            var values = field.split(".")
                            if (values.length == 1 && Object.keys(prototypes).length == 1) {
                                // That's means there's only one table in use for the query.
                                if(field != "*"){
                                    var fieldType = result.FieldsType[result.getFieldIndex("M_" + field)]
                                    fieldsType.push(fieldType)
                                    fields.push(field)
                                }else{
                                    // In that case the user want all the value from the table.
                                    for(var i=0; i < result.FieldsType.length; i++){
                                        if(result.FieldsType[i].startsWith("sqltypes.")){
                                            fieldsType.push(result.FieldsType[i])
                                            fields.push(result.Fields[i])
                                        }
                                    }
                                }
                            } else {
                                var tableId = values[0]
                                var fieldName = values[1]
                            }
                        }

                        // Here I will execute the query...
                        if (caller.type == "READ") {
                            caller.queryEditor.setResult(caller.query, fields, fieldsType, [], caller.type)
                        }
                    }
                },
                // error callback
                function (errObj, caller) {

                }, { "prototypes": prototypes, "type": type, "fields": fields, "query": query, "queryEditor": this })

        }
        console.log(ast)
    }
}

/**
 * Set the reuslt table.
 */
QueryEditor.prototype.setResult = function (query, fields, fieldsType, param, type) {
    if (this.isSql) {
        if (type == "READ") {
            // Here I will create a sql table.
            var table = new Table(this.activeDataConfig.M_id, this.resultQueryPanel)

            // Set the table model.
            var model = new SqlTableModel(this.activeDataConfig.M_id, query, fieldsType, [], fields)

            table.setModel(model, function (table, queryEditor) {
                return function () {
                    // init the table.
                    table.init()
                    table.header.maximizeBtn.element.click()

                    // Because html table suck with the header position vs scroll 
                    // I do some little hack to fix it without weird stuff like 
                    // copy another header etc...
                    var widths = []
                    // Now I will wrote the code for the layout...
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
                    table.rowGroup.element.style.top = table.header.div.element.offsetHeight + 2 + "px"

                    // Now the height of the panel...
                    table.rowGroup.element.style.height = queryEditor.resultQueryPanel.element.offsetHeight - table.header.div.element.offsetHeight + "px"

                    table.rowGroup.element.onscroll = function(header){
                        return function(){
                        var position = this.scrollTop;
                        if(this.scrollTop > 0){
                            header.className = "table_header scrolling"
                        }else{
                            header.className = "table_header"
                        }
                        }
                    }(table.header.div.element)

                    // Now the resize event.
                    window.addEventListener('resize',
                        function (queryEditor, table) {
                            return function () {
                                table.rowGroup.element.style.height = queryEditor.resultQueryPanel.element.offsetHeight - table.header.div.element.offsetHeight + "px"
                            }
                        } (queryEditor, table), true);
                }
            } (table, this))
        }
    }
}