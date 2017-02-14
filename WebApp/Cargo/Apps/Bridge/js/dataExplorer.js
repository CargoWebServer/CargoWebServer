
/**
 * That class is use to display the structure of information about a given data store.
 */
var DataExplorer = function (parent) {

    // Keep reference to the parent.
    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "style": "display: table-row;" }).down().appendElement({ "tag": "div", "class": "data_explorer" }).down()

    // That contain the map of data aready loaded.
    this.schemas = {}
    this.configs = {}
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.initDataSchema = function (storeConfig) {
    //this.panel.element.innerHTML = this.schemas[storeId]
    this.configs[storeConfig.M_id] = storeConfig

    // So here I will get the list of all prototype from a give store and
    // create it's relavite information.
    if (storeConfig.M_dataStoreType == 0) {
        // Sql data store.
    }else if (storeConfig.M_dataStoreType == 1) {
        // Entity data store.
        server.entityManager.getEntityPrototypes(storeId,
            // success callback.
            function (results, caller) {

            },
            // error callback.
            function (errMsg, caller) {

            },
            this)
    } 

    this.schemas[storeId] = storeId
}

/**
 * Generate the prototypes view.
 */
DataExplorer.prototype.generatePrototypesView = function (storeId) {

}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.setDataSchema = function (storeId) {
    if (this.schemas[storeId] != undefined) {

    }
}