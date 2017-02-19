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
    this.tool
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

    this.resultQueryPanel = new Element(splitArea2, { "tag": "div", "class": "result_query_panel" })

    // Init the splitter action.
    initSplitter(queryEditorSplitor, this.editQueryPanel)

    var filePanel = this.editQueryPanel.appendElement({ "tag": "div", "class": "filePanel", "id": file.M_id + "_editor", "innerHtml": decode64(file.M_data) }).down()
    var editor = ace.edit(file.M_id + "_editor");

    // In case of sql query..
    if (this.isSql) {
        editor.getSession().setMode("ace/mode/sql");
        editor.getSession().on('change', function (fileId, fileUUID, editor) {
            return function () {
                var evt = { "code": ChangeFileEvent, "name": FileEvent, "dataMap": { "fileId": fileId } }
                var file = server.entityManager.entities[fileUUID]
                file.M_data = encode64(editor.getSession().getValue())
                server.eventHandler.BroadcastEvent(evt)
            }
        } (file.M_id, file.UUID, editor));
    } else if (this.isEql) {
        // TODO implement the syntax highlight for EQL...
    }

    return this
}

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