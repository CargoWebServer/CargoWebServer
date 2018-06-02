/** Text contain in the main page **/
var itemSearchResultPageText = {
    "en": {
        "comments_tbl_header": "comments",
        "id_tbl_header": "OVMM",
        "keywords_tbl_header": "keywords",
        "Type_tbl_header": "type",
        "Secteur_tbl_header": "sector",
        "Couleur_tbl_header": "color",
        "Duré_de_vie_tbl_header": "lifespan",
        "Longueur_tbl_header": "length",
        "Largeur_tbl_header": "width",
        "Épaisseur_tbl_header": "tickness",
        "Diamètre_tbl_header": "diameter",
    },
    "fr": {
        "comments_tbl_header": "commentaire(s)",
        "id_tbl_header": "OVMM",
        "keywords_tbl_header": "mots clés",
        "Type_tbl_header": "sorte",
        "Secteur_tbl_header": "secteur",
        "Couleur_tbl_header": "couleur",
        "Duré_de_vie_tbl_header": "duré de vie",
        "Longueur_tbl_header": "longueur",
        "Largeur_tbl_header": "largeur",
        "Épaisseur_tbl_header": "épaisseur",
        "Diamètre_tbl_header": "diamètre",
    }
};

// Set the text info.
server.languageManager.appendLanguageInfo(itemSearchResultPageText);

/**
 *  Display items seach results.
 *  ** Search is the search function with prototype function(offset, pageSize)
 */
var ItemSearchResultPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "item_search_result_page", "class": "search_results" })

    // Now I will display the list of results...
    this.itemDisplayPanel = this.panel.appendElement({ "tag": "div", "class": "item_display_content" }).down()

    // I will made use of a tab panel to display the results.
    this.tabPanel = new TabPanel(this.itemDisplayPanel);
    
    return this;
}

ItemSearchResultPage.prototype.displayResults = function (results, query) {
    
    var headerText = "No results found for \"" + results.query + "\""
    spinner.panel.element.style.display = "none";
    query = query.replace("*", "-");
    
    var packages = []
    var uuids = []
    if (results.estimate > 0) {
        for (var i = 0; i < results.results.length; i++) {
            var result = results.results[i];
            if (result.data.TYPENAME == "CatalogSchema.ItemType") {
                if (uuids.indexOf(results.results[i].data.UUID) == -1) {
                    uuids.push(results.results[i].data.UUID)
                }
            } else if (result.data.TYPENAME == "CatalogSchema.PropertyType") {
                if (uuids.indexOf(results.results[i].data.ParentUuid) == -1) {
                    uuids.push(results.results[i].data.ParentUuid)
                }
            } else if (result.data.TYPENAME == "CatalogSchema.ItemSupplierType") {
                // In case of package type...
                if (packages.indexOf(result.data.M_package) == -1) {
                    packages.push(result.data.M_package)
                }
            }
        }
        headerText = "About " + (uuids.length) + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
    }

    // Display the number of results...
    var footer = document.getElementById("main-page-footer")
    footer.innerText = headerText;
    setTimeout(function (footer) {
        return function () {
            footer.innerText = "";
        }
    }(footer), 4000)

    // So here I will regroup results by type...
    var results_ = [];

    // Properties are asynchrone...
    var getItem = function (uuid, itemSearchResultPage, results, properties, query, getItem, callback) {
        // In that case I will get the entity from the server...
        server.entityManager.getEntityByUuid(uuid, false,
            // success callback
            function (entity, caller) {
                caller.results.push(entity)
                var next = caller.uuids.pop()
                if (caller.uuids.length == 0) {
                    caller.callback(caller.results, caller.query, caller.itemSearchResultPage)
                } else {
                    caller.getItem(next, caller.itemSearchResultPage, caller.results, caller.uuids, caller.query, caller.getItem, caller.callback)
                }
            },
            // error callback
            function () {

            }, { "itemSearchResultPage": itemSearchResultPage, "results": results, "uuids": uuids, "query": query, "getItem": getItem, "callback": callback })
    }

    // Get the properties items if any's
    if (uuids.length > 0) {
        getItem(uuids[0], this, results_, uuids, query, getItem,
            // The callback...
            function (packages) {
                return function (results, query, itemSearchResultPage) {
                    if (packages.length == 0) {
                        // simply display the results.
                        itemSearchResultPage.displayTabResults(results, query)
                    } else {
                        // In that case I will also append items from the packages...
                        var getPackageItems = function (pacakageUuid, packages, itemSearchResultPage, results, query, getPackageItems, callback) {
                            server.entityManager.getEntityByUuid(pacakageUuid, false,
                                // success callback 
                                function (results, caller) {
                                    caller.getPacakageItem = function (items, caller) {
                                        var item = items.pop()
                                        caller.items = items;
                                        // Here I will get the item reference value.
                                        server.entityManager.getEntityByUuid(item, false,
                                            // success callback 
                                            function (results, caller) {
                                                if (!objectPropInArray(caller.results, "UUID", results.UUID)) {
                                                    caller.results.push(results)
                                                }
                                                // Here I will get the item reference value.
                                                if (caller.items.length == 0) {
                                                    if (caller.packages.length == 0) {
                                                        // end of process
                                                        caller.callback(caller.results, caller.query, caller.itemSearchResultPage)
                                                    } else {
                                                        // call another time.
                                                        caller.getPackageItems(caller.packages.pop(0), caller.packages, caller.itemSearchResultPage, caller.results, caller.query, caller.getPackageItems, caller.callback)
                                                    }
                                                } else {
                                                    caller.getPacakageItem(caller.items, caller)
                                                }
                                            },
                                            // error callback
                                            function (errObj, caller) {

                                            }, caller)
                                    }

                                    caller.getPacakageItem(results.M_items, caller)

                                },
                                // error callback
                                function (errObj, caller) {

                                }, { "itemSearchResultPage": itemSearchResultPage, "packages": packages, "results": results, "query": query, "getPackageItems": getPackageItems, "callback": callback })
                        }
                        // process the first item.
                        getPackageItems(packages.pop(0), packages, itemSearchResultPage, results, query, getPackageItems, function (results, query, itemSearchResultPage) {
                            itemSearchResultPage.displayTabResults(results, query)
                        })
                    }
                }
            }(packages)
        )
    }
}

// Slipt camel case word and append space...
function splitCamelCaseToString(s) {
    return s.split(/(?=[A-Z])/).join(' ');
}

ItemSearchResultPage.prototype.displayTabResults = function (results, query) {
    spinner.panel.element.style.display = "";
    var footer = document.getElementById("main-page-footer")

    headerText = "About " + results.length + " results"

    footer.innerText = headerText;
    setTimeout(function (footer) {
        return function () {
            footer.innerText = "";
        }
    }(footer), 4000)

    // Display the results.
    this.panel.element.style.display = ""

    // Display the selector.
    document.getElementById("search_context_selector").style.display = "none"
    if (Object.keys(mainPage.itemDisplayPage.tabPanel.tabs).length > 0) {
        document.getElementById("item_context_selector").style.display = ""
    } else {
        document.getElementById("item_context_selector").style.display = "none"
    }

    // display the panel.
    document.getElementById("item_search_result_page").style.display = ""
    document.getElementById("item_display_page_panel").style.display = "none"

    // Here I will create the tab for one panel.
    var tab = this.tabPanel.appendTab(query)
    tab.closeBtn.element.addEventListener("click", function (itemSearchResultPage) {
        return function () {
            if (Object.keys(itemSearchResultPage.tabPanel.tabs).length == 0) {
                document.getElementById("search_context_selector").style.display = "none"
            }
        }
    }(this));

    tab.title.element.innerHTML = query + "(" + results.length + ") ";
    tab.content.element.style.backgroundColor = "white";
    tab.content.element.style.position = "relative";

    // Now I will append the found element in a table of entities.
    var proto = getEntityPrototype("CatalogSchema.ItemType")

    // Set fields visibility.
    proto.FieldsVisibility[10] = false;
    proto.FieldsVisibility[11] = false;

    var model = new EntityTableModel(proto);
    model.editable[2] = true;
    model.editable[3] = true;

    // Set the list of entities in the model.
    model.entities = results;

    var table = new Table(+ "_search_result", tab.content)

    // Here I will set the tow string function for the properties type
    // that function is use by the column filter to help filtering the 
    // table values.
    for (var i = 0; i < results.length; i++) {
        for (var j = 0; j < results[i].M_properties.length; j++) {
            var property = results[i].M_properties[j]
            if (isObject(property.M_dimensionValue)) {
                // In that case I will display the value...
                property.toString = function (dimension) {
                    return function () {
                        // Return the dimension value.
                        return dimension.M_valueOf + " " + dimension.M_unitOfMeasure.M_valueOf
                    }
                }(property.M_dimensionValue)
            } else if (property.M_stringValue != undefined) {
                property.toString = function (str) {
                    return function () {
                        return str
                    }
                }(property.M_stringValue)
            } else if (property.M_numericValue != undefined) {
                property.toString = function (val) {
                    return function () {
                        if (isInt(val)) {
                            return formatValue(val, "xs.int")
                        } else {
                            return formatValue(val, "xs.decimal")
                        }
                    }
                }(property.M_numericValue)
            } else if (property.M_booleanValue != undefined) {
                property.toString = function (val) {
                    return function () {
                        if (val == true) {
                            return "true"
                        } else {
                            return "false"
                        }
                    }
                }(property.M_booleanValue)
            }
        }
    }

    // Here I will set the renderer for the type CatalogSchema.DimensionType
    // for the table.
    table.renderFcts["CatalogSchema.DimensionType"] = function (value) {
        // So here I will return the string that contain the unit of measure.
        var str = value.M_valueOf

        if (value.M_unitOfMeasure.M_valueOf == "IN (inch)") {
            str += '"'
        } else if (value.M_unitOfMeasure.M_valueOf == "FT (feet)") {
            str += "'"
        } else {
            str += " " + value.M_unitOfMeasure.M_valueOf
        }
        return new Element(null, { "tag": "div", "innerHtml": str });
    }

    // The property display function...
    table.renderFcts["[]CatalogSchema.PropertyType"] = function (properties) {
        // So here I will return the string that contain the unit of measure.
        var panel = new Element(null, { "tag": "div", "style": "display: table;width: 100%; border-collapse: collapse;margin-top: 2px;margin-bottom: 2px;" })
        for (var i = 0; i < properties.length; i++) {
            var property = properties[i]
            var row = panel.appendElement({ "tag": "div", "style": "display:table-row; width: 100%;" }).down()

            // The propertie name.
            row.appendElement({ "tag": "div", "style": "display:table-cell; padding-right: 5px;", "innerHtml": property.M_name + ": " })

            var valueDiv = row.appendElement({ "tag": "div", "style": "display:table-cell; width: 100%;" }).down()
            // The propertie value...
            if (isObject(property.M_dimensionValue)) {
                // Render the Dimension type with the function....
                valueDiv.appendElement(table.renderFcts["CatalogSchema.DimensionType"](property.M_dimensionValue))
            } else if (property.M_stringValue != undefined) {
                valueDiv.element.innerHTML = property.M_stringValue
            }
        }
        return panel;
    }

    // I will append information to the model
    table.setModel(model, function (table) {
        return function () {
            // Set hidden column here.
            // init the table.
            table.parent.element.style.position = "relative"
            table.init()
            table.header.maximizeBtn.element.click();
            spinner.panel.element.style.display = "none";

            // I will set the action on the first cell...
            for (var i = 0; i < table.rows.length; i++) {
                var cell = table.rows[i].cells[0]
                cell.div.element.style.textDecoration = "underline"
                cell.div.element.onmouseover = function () {
                    this.style.color = "steelblue"
                    this.style.cursor = "pointer"
                }

                cell.div.element.onmouseleave = function () {
                    this.style.color = ""
                    this.style.cursor = "default"
                }

                // Now the onclick... 
                var entity = table.getModel().entities[i]
                cell.div.element.onclick = function (entity) {
                    return function () {
                        mainPage.searchResultPage.element.parentNode.removeChild(mainPage.searchResultPage.element)
                        mainPage.itemDisplayPage.displayTabItem(entity)
                    }
                }(entity)
            }
        }
    }(table))


}