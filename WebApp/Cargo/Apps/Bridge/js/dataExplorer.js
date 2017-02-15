
/**
 * That class is use to display the structure of information about a given data store.
 */
var DataExplorer = function (parent) {

    // Keep reference to the parent.
    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "data_explorer" }).down()

    // Set the resize event.
    window.addEventListener('resize',
        function (dataExplorer) {
            return function () {
                dataExplorer.resize()
            }
        } (this), true);

    // That contain the map of data aready loaded.
    this.shemasView = {}
    this.configs = {}

    return this
}

DataExplorer.prototype.resize = function () {
    var height = this.parent.element.offsetHeight - this.parent.element.firstChild.offsetHeight;
    this.panel.element.style.height = height - 2 + "px"
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.initDataSchema = function (storeConfig) {
    //this.panel.element.innerHTML = this.schemas[storeId]
    this.configs[storeConfig.M_id] = storeConfig
    this.shemasView[storeConfig.M_id] = new Element(this.panel, { "tag": "div", "class": "shemas_view" })

    // So here I will get the list of all prototype from a give store and
    // create it's relavite information.
    if (storeConfig.M_dataStoreType == 1) {
        // Sql data store.
    } else if (storeConfig.M_dataStoreType == 2) {
        // Entity data store.
        server.entityManager.getEntityPrototypes(storeConfig.M_id,
            // success callback.
            function (results, caller) {
                caller.dataExplorer.generatePrototypesView(caller.storeId, results)
            },
            // error callback.
            function (errMsg, caller) {

            },
            { "dataExplorer": this, "storeId": storeConfig.M_id })
    }
}

/**
 * Generate the prototypes view.
 */
DataExplorer.prototype.generatePrototypesView = function (storeId, prototypes) {
    this.panel.element.style.borderTop = "1px solid grey"
    // Here I will create the prototype views...
    for (var i = 0; i < prototypes.length; i++) {
        // Here I will append the prototype name...
        new PrototypeTreeView(this.shemasView[storeId], prototypes[i])
    }
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.setDataSchema = function (storeId) {

    // here I will calculate the height...
    this.resize()

    for (var id in this.shemasView) {
        this.shemasView[id].element.style.display = "none"
    }
    if (this.shemasView[storeId] != undefined) {
        this.shemasView[storeId].element.style.display = ""
    }
}

/**
 * 
 */
var PrototypeTreeView = function (parent, prototype) {
    this.parent = parent
    this.panel = new Element(parent, { "tag": "div", "class": "data_prototype_tree_view" })
    this.panel.appendElement({ "tag": "div", "innerHtml": prototype.TypeName })
    return this
}