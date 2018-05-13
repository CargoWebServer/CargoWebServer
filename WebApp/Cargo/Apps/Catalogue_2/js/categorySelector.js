/** Class to display all categories and theire respective items. */
var CatergorySelector = function (parent) {
    this.parent = parent;
    this.panel = parent.appendElement({ "tag": "div", "class": "category_selector" }).down()

    this.categories = {}

    // So here  I wil get the list of all category.
    server.entityManager.getEntities("CatalogSchema.CategoryType", "CatalogSchema", '', 0, -1, [], true, false,
        // Progress...
        function (index, total) {

        },
        // success
        function (results, caller) {
            for (var i = 0; i < results.length; i++) {
                caller.appendCategory(results[i])
            }
        },
        // error
        function () {

        }, this)

    return this
}

/**
 * Append a new category to the panel
 * @param {*} category 
 */
CatergorySelector.prototype.appendCategory = function (category) {
    this.categories[category.M_id] = new CategoryPanel(this.panel)
    this.categories[category.M_id].init(category)
}


var CategoryPanel = function (parent, category) {
    this.panel = parent.appendElement({ "tag": "div", "class": "category_panel" }).down()

    //console.log(category)
    this.header = this.panel.appendElement({ "tag": "div", "class": "category_panel_header" }).down()
    this.expandBtn = this.header.appendElement({ "tag": "div", "class": "category_panel_header_button" }).down()
    this.expandBtn.appendElement({ "tag": "i", "class": "fa fa-caret-right" }).down()
    this.shrinkBtn = this.header.appendElement({ "tag": "div", "class": "category_panel_header_button", "style": "display: none;" }).down()
    this.shrinkBtn.appendElement({ "tag": "i", "class": "fa fa-caret-down" }).down()
    this.title = this.header.appendElement({ "tag": "div", "class": "category_panel_header_title" }).down()

    this.items = this.panel.appendElement({ "tag": "div", "class": "category_items_panel", "style": "display: none;" }).down()

    // Now the event...
    this.expandBtn.element.onclick = function (header, itemPanel) {
        return function (evt) {
            evt.stopPropagation(true)
            header.expandBtn.element.style.display = "none"
            header.shrinkBtn.element.style.display = ""
            itemPanel.element.style.display = "table"
        }
    }(this, this.items)

    this.shrinkBtn.element.onclick = function (header, itemPanel) {
        return function (evt) {
            evt.stopPropagation(true)
            header.expandBtn.element.style.display = ""
            header.shrinkBtn.element.style.display = "none"
            itemPanel.element.style.display = "none"
        }
    }(this, this.items)

    // Now I will get the list of item by type...

    // So now I need to get information about items contain in the category...
    this.typeNames = {}


    return this
}

CategoryPanel.prototype.init = function (category) {
    this.title.element.innerText = category.M_name
    
    getTypeName = function (index, items, categoryPanel) {
        // In that case I will get the ids with the data manager.
        var query = {}
        query.TypeName = "CatalogSchema.ItemType"
        query.Fields = ["M_name", "M_id"]
        query.Query = 'CatalogSchema.ItemType.UUID =="' + items[index] + '"'
        server.dataManager.read("CatalogSchema", JSON.stringify(query), [], [],
            // success callback
            function (results, caller) {
                // return the results.
                if (results[0].length == 1) {
                    if (results[0][0].length == 2) {
                        var typeName = results[0][0][0]
                        if (typeName == null) {
                            typeName = "generic"
                        }
                        typeName.trim()
                        if (caller.categoryPanel.typeNames[typeName] == undefined) {
                            caller.categoryPanel.typeNames[typeName] = []
                        }
                        caller.categoryPanel.typeNames[typeName].push({ "ovmm": results[0][0][1].trim(), "uuid": caller.items[caller.index] })
                    }
                }
                if (caller.index < caller.items.length) {
                    caller.callback(caller.index + 1, caller.items, caller.categoryPanel)
                } else {
                    // Render the array here.
                    caller.categoryPanel.displayItemsByTypeName()
                }
            },
            // progress callback
            function (index, total, caller) {

            },
            // error callback
            function (errObj, caller) {

            },
            { "index": index, "items": items, "callback": getTypeName, "categoryPanel": categoryPanel })
    }
    getTypeName(0, category.M_items, this)
}

/**
 * 
 */
CategoryPanel.prototype.displayItemsByTypeName = function () {
    var items = []
    for(var typeName in this.typeNames){
        items = items.concat(this.typeNames[typeName])
    }

    this.title.element.onclick = function (items, title) {
        return function () {
            spinner.panel.element.style.display = "";
            var getItems = function (items, values, callback) {
                server.entityManager.getEntityByUuid(items[values.length].uuid, false,
                    // success callback
                    function (results, caller) {
                        values.push(results)
                        if(caller.values.length < caller.items.length){
                            caller.callback(caller.items, values, callback)
                        }else{
                            if (mainPage.searchResultPage == undefined) {
                                mainPage.searchResultPage = new ItemSearchResultPage(mainPage.content, mainPage.searchItems, caller.title)
                            }
                            mainPage.searchResultPage.displayTabResults(values, caller.title);
                        }
                    },
                    // error callback
                    function (errObj, caller) {

                    }, {"items":items, "values":values, "callback":callback, "title":title})
            }

            getItems(items, [], getItems)
        }
    }(items, this.title.element.innerText)

    // Here I will display the list of items by alphabetic order.
    Object.keys(this.typeNames).sort().forEach(
        function (typeNames, categoryPanel) {
            return function (typeName) {
                var lnk = categoryPanel.items.appendElement({ "tag": "div", "class": "type_name_items", "innerHtml": typeName + "  (" + typeNames[typeName].length + ")" }).down()
                // Now I will set the action when the link is clicke...
                lnk.element.onclick = function (items) {
                    return function () {
                        spinner.panel.element.style.display = "";
                        var getItems = function (items, values, callback) {
                            server.entityManager.getEntityByUuid(items[values.length].uuid, false,
                                // success callback
                                function (results, caller) {
                                    values.push(results)
                                    if(caller.values.length < caller.items.length){
                                        caller.callback(caller.items, values, callback)
                                    }else{
                                        if (mainPage.searchResultPage == undefined) {
                                            mainPage.searchResultPage = new ItemSearchResultPage(mainPage.content, mainPage.searchItems, mainPage.searchInput.value)
                                        }
                                    
                                        mainPage.searchResultPage.displayTabResults(values, results.M_name);
                                    }
                                },
                                // error callback
                                function (errObj, caller) {

                                }, {"items":items, "values":values, "callback":callback})
                        }

                        getItems(items, [], getItems)
                    }
                }(typeNames[typeName])
            }
        }(this.typeNames, this)
    );
}