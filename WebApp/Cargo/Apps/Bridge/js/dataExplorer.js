
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
    this.prototypesView = {}
    this.configs = {}

    return this
}

DataExplorer.prototype.resize = function () {
    var height = this.parent.element.offsetHeight - this.parent.element.firstChild.offsetHeight;
    this.panel.element.style.height = height - 10 + "px"
}

/**
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.initDataSchema = function (storeConfig) {
    // init one time.
    if (this.configs[storeConfig.M_id] != undefined) {
        return
    }

    this.configs[storeConfig.M_id] = storeConfig
    this.shemasView[storeConfig.M_id] = new Element(this.panel, { "tag": "div", "class": "shemas_view" })

    // Only display the first panel at first.
    if (Object.keys(this.shemasView).length > 1) {
        this.shemasView[storeConfig.M_id].element.style.display = "none"
    }

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
        if (this.prototypesView[prototypes[i].TypeName] == undefined) {
            this.prototypesView[prototypes[i].TypeName] = new PrototypeTreeView(this.shemasView[storeId], prototypes[i])
        }
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
 * Display the data schema of a given data store.
 */
DataExplorer.prototype.removeDataSchema = function (storeId) {

    // here I will calculate the height...
    for (var id in this.shemasView) {
        this.shemasView[id].element.style.display = "none"
        if (this.shemasView[id] == storeId) {
            this.shemasView[storeId].element.parentNode.removeChild(this.shemasView[storeId].element)
            delete this.shemasView[storeId]
        }
    }

}

/**
 * That view is use to display prototype structures
 */
var PrototypeTreeView = function (parent, prototype) {
    this.parent = parent
    this.panel = new Element(parent, { "tag": "div", "class": "data_prototype_tree_view" })
    this.fieldsView = {}

    // The type name without the schemas name.
    var typeName = prototype.TypeName.substring(prototype.TypeName.indexOf(".") + 1) //.split(".")[1]

    // Display the type name and the expand shrink button.
    var header = this.panel.appendElement({ "tag": "div", "class": "data_prototype_tree_view" }).down().appendElement({ "tag": "div", "class": "data_prototype_tree_view_header" }).down()
    header.appendElement()

    /** The expand button */
    this.expandBtn = header.appendElement({ "tag": "i", "class": "fa fa-caret-right", "style": "display:inline;" }).down()

    /** The shrink button */
    this.shrinkBtn = header.appendElement({ "tag": "i", "class": "fa fa-caret-down", "style": "display:none;" }).down()
    header.appendElement({ "tag": "span", "innerHtml": typeName }).down()

    this.fieldsPanel = this.panel.appendElement({ "tag": "div", "class": "data_prototype_tree_view_fields" }).down()

    // The code for display field of a given type.
    this.expandBtn.element.onclick = function (view) {
        return function () {
            view.fieldsPanel.element.style.display = "table"
            this.style.display = "none"
            view.shrinkBtn.element.style.display = ""
        }
    } (this)

    this.shrinkBtn.element.onclick = function (view) {
        return function () {
            view.fieldsPanel.element.style.display = ""
            this.style.display = "none"
            view.expandBtn.element.style.display = ""
        }
    } (this)

    // Now the fields.
    for (var i = 0; i < prototype.Fields.length; i++) {
        if (prototype.Fields[i].startsWith("M_")) {

            this.fieldsView[prototype.Fields[i]] = new PrototypeTreeViewField(this.fieldsPanel, prototype, prototype.Fields[i], prototype.FieldsType[i], prototype.FieldsVisibility[i], prototype.FieldsNillable[i])
        }
    }

    return this
}

/**
 * The view of a given field...
 */
var PrototypeTreeViewField = function (parent, prototype, fieldName, fieldType, isVisible, isNillable) {
    // if is an id...
    var isKey = contains(prototype.Ids, fieldName)
    var isIndex = contains(prototype.Indexs, fieldName)

    // Not display the M_
    fieldName = fieldName.replace("M_", "")

    // The parent prototype.
    this.prototype = prototype

    // The parent panel.
    this.parent = parent

    // Create the new panel.
    this.panel = new Element(parent, { "tag": "div", "class": "data_prototype_tree_view_field" })

    var visibilityClass = ""
    if (isVisible) {
        visibilityClass = "field_visibility visible"
    }

    var className = ""
    if (isKey) {
        className = "field_id"
    } else if (isIndex) {
        className = "field_index"
    }

    // append the the field name.
    this.panel.appendElement({ "tag": "i", "title": "visibility", "class": "fa fa-lightbulb-o " + visibilityClass }).appendElement({ "tag": "span", "innerHtml": fieldName, "class": className }).down()

    // Now the typename 
    this.panel.appendElement({ "tag": "span", "innerHtml": fieldType }).down()

}