// var field = "X" + prototype.TypeName.toLowerCase() + "." +prototype.Fields[j].toLowerCase() + "%:" + prototype.TypeName.split(".")[1] + "_" + prototype.Fields[j].substring(2).toLowerCase()
/**
 * That structure contain the information about a search.
 */
var SearchInfo = function () {
    this.TYPENAME = "SearchInfo"
    this.UUID = randomUUID();
    this.M_id = -1 // Temporary id display in the page
    this.M_name = ""
    return this
}

/**
 * That panel contain information about the datasources, datatypes
 * range and other option of the search.
 */
var SearchOptionsPanel = function (parent) {
    this.panel = parent.appendElement({ "tag": "div", "id": "search_option_panel", "class": "search_options_panel" }).down()

    // The panel will contain the list of all datastore... and each datastore will contain the list 
    // of all it datatypes...
    this.tabPanelHeader = this.panel.appendElement({ "tag": "div", "class": "search_options_panel_tab_panel_header" }).down()

    // Contain the content...
    this.tabPanelBody = this.panel.appendElement({ "tag": "div", "style": "display: table; width: 100%;" }).down()
        .appendElement({ "tag": "div", "class": "search_options_panel_tab_panel_body" }).down()

    // I will now createe the fields information as the form:
    // Xcargoentities.file.m_data:data
    this.tabs = {}

    for (var i = 0; i < server.activeConfigurations.M_dataStoreConfigs.length; ++i) {
        var storeId = server.activeConfigurations.M_dataStoreConfigs[i].M_id;
        // discard some store...
        if (storeId != "xs" && storeId != "sqltypes" && storeId != "XMI_types" && server.activeConfigurations.M_dataStoreConfigs[i].M_dataStoreType == 2) {
            this.tabs[storeId] = new SearchOptionPanelStoreInfo(this, server.activeConfigurations.M_dataStoreConfigs[i])
        }
    }

    // empty element to push tabs to left.
    this.tabPanelHeader.appendElement({ "tag": "div", "style": "display: table-cell; width: 100%;" })

    this.tabs["CargoEntities"].tab.element.click() // set the default selected.

    return this
}

/**
 * That function return the list of all database to look at...
 */
SearchOptionsPanel.prototype.getDataStoreList = function () {
    dbpaths = []
    for (var id in this.tabs) {
        if (this.tabs[id].isSelected()) {
            // var dbpath = server.root + "/Data/" + id + "/" + storeId + ".glass"
            dbpaths = dbpaths.concat(this.tabs[id].getSelectedDbPaths())
        }
    }
    return dbpaths
}

/**
 * That panel contain information to display the datastore information.
 * @param {*} searchPanel 
 * @param {*} dataStoreConfig 
 */
var SearchOptionPanelStoreInfo = function (searchPanel, dataStoreConfig) {
    this.id = dataStoreConfig.M_id
    this.tab = searchPanel.tabPanelHeader.appendElement({ "tag": "div", "style": "display: table-cell; padding-left: 1px; padding-right: 1px;" }).down()
        .appendElement({ "tag": "div", "class": "search_options_panel_tab_panel_header_tab" }).down()
    this.isSelectedBtn = this.tab.appendElement({ "tag": "input", "type": "checkbox", "id": this.id + "_select" }).down()

    this.tab.appendElement({ "tag": "span", "innerHtml": dataStoreConfig.M_storeName })

    // The datasotre is selected by default.
    if (dataStoreConfig.M_dataStoreType == 2) {
        this.isSelectedBtn.element.checked = true
    }

    // Select or unselect types all a once.
    this.isSelectedBtn.element.onclick = function (storeId) {
        return function () {
            var selects = document.getElementsByName(storeId + "_select")
            for (var i = 0; i < selects.length; i++) {
                if (selects[i].checked != this.checked) {
                    selects[i].checked = this.checked

                }
            }
        }
    }(dataStoreConfig.M_id)

    this.searchOptionPanelDataTypeInfo = new SearchOptionPanelDataTypeInfo(searchPanel, dataStoreConfig)

    // Now the actions...
    this.tab.element.onclick = function (searchOptionPanelStoreInfo) {
        return function () {
            var tabs = document.getElementsByClassName("search_options_panel_tab_panel_header_tab")
            for (var i = 0; i < tabs.length; ++i) {
                tabs[i].className = "search_options_panel_tab_panel_header_tab" // remove active if there..
            }
            this.className = "search_options_panel_tab_panel_header_tab active"
            var searchOptionPanelDataTypeInfoPanels = document.getElementsByClassName("search_option_panel_data_type_info")
            for (var i = 0; i < searchOptionPanelDataTypeInfoPanels.length; i++) {
                searchOptionPanelDataTypeInfoPanels[i].style.display = ""
            }
            searchOptionPanelStoreInfo.searchOptionPanelDataTypeInfo.panel.element.style.display = "table"
        }
    }(this)

    return this
}

/**
 * Return true if the datastore is selected.
 */
SearchOptionPanelStoreInfo.prototype.isSelected = function () {
    return this.isSelectedBtn.element.checked
}

/**
 * Return the list of datastore path to query
 */
SearchOptionPanelStoreInfo.prototype.getSelectedDbPaths = function () {
    var dbpaths = []
    for (var id in this.searchOptionPanelDataTypeInfo.isSelectedBtns) {
        if (this.searchOptionPanelDataTypeInfo.isSelectedBtns[id].element.checked) {
            var dbpath = server.root + "/Data/" + id.split(".")[0] + "/" + id + ".glass"
            dbpaths.push(dbpath)
        }
    }
    return dbpaths
}

/**
 * That panel is use to create the query from the datatype.
 * @param {*} searchPanel 
 * @param {*} dataStoreConfig 
 */
var SearchOptionPanelDataTypeInfo = function (searchPanel, dataStoreConfig) {
    this.id = dataStoreConfig.M_id // same id as the tab...
    this.panel = searchPanel.tabPanelBody.appendElement({ "tag": "div", "class": "search_option_panel_data_type_info" }).down()
    this.isSelectedBtns = {}
    this.isSelectedFieldBtns = {}

    // Get the list of entity prototypes for that store.
    server.entityManager.getEntityPrototypes(dataStoreConfig.M_id,
        function (prototypes, searchOptionPanelDataTypeInfo) {
            for (var i = 0; i < prototypes.length; ++i) {
                var prototype = prototypes[i]
                searchOptionPanelDataTypeInfo.appendDataTypeInfos(prototype)
            }
        },
        function () {

        }, this)

    return this;
}

SearchOptionPanelDataTypeInfo.prototype.appendDataTypeInfos = function (prototype) {

    // So here I will 
    var typeInfoDiv = this.panel.appendElement({ "tag": "div", "style": "display: table;" }).down()

    /** The expand button */
    var expandBtn = typeInfoDiv.appendElement({ "tag": "i", "class": "fa fa-caret-right", "style": "display:inline;" }).down()

    /** The shrink button */
    var shrinkBtn = typeInfoDiv.appendElement({ "tag": "i", "class": "fa fa-caret-down", "style": "display:none;" }).down()

    // So here I will display the liste
    typeInfoDiv.appendElement({ "tag": "span", "style": "display: table-cell", "innerHtml": prototype.TypeName.split(".")[prototype.TypeName.split(".").length - 1] })

    var isSelectBtn = typeInfoDiv.appendElement({ "tag": "input", "name": prototype.TypeName.split(".")[0] + "_select", "id": prototype.TypeName + "_select", "type": "checkbox", "style": "display: table-cell" }).down()
    isSelectBtn.element.checked = true
    this.isSelectedBtns[prototype.TypeName] = isSelectBtn

    isSelectBtn.element.onclick = function (typeName, isSelectedBtns) {
        return function () {
            var selects = document.getElementsByName(typeName + "_select")
            for (var i = 0; i < selects.length; i++) {
                selects[i].checked = this.checked
            }
            // Now I will adjust it parent...
            var isSelectBtn = document.getElementById(typeName.split(".")[0] + "_select")
            isSelectBtn.checked = false
            for (var key in isSelectedBtns) {
                if (isSelectedBtns[key].element.checked) {
                    isSelectBtn.checked = true
                    break
                }
            }
        }
    }(prototype.TypeName, this.isSelectedBtns)


    // Now I will create the div where type will be displayed.
    var typeDiv = this.panel.appendElement({ "tag": "div", "style": "display: none; padding-left: 20px; padding-bottom: 5px; border-spacing:2px 2px;" }).down()

    this.isSelectedFieldBtns[prototype.TypeName + "_select"] = []

    // Hew I will append field informations.
    for (var i = 0; i < prototype.FieldsType.length; i++) {
        // Here only xs type can be display...
        var fieldType = prototype.FieldsType[i]
        var field = prototype.Fields[i]
        if (field.startsWith("M_")) {
            if (isXsBaseType(fieldType) || isXsBaseType(getBaseTypeExtension(fieldType))) {
                // console.log(fieldType)
                var fieldDiv = typeDiv.appendElement({ "tag": "div", "style": "display: table-row;" }).down()
                var fieldId = prototype.TypeName.toLowerCase().replace(".", "_") + "_" + field.substring(2).toLowerCase()
                var isSelectFieldBtn = fieldDiv.appendElement({ "tag": "input", "type": "checkbox", "class": "field_checkbox", "id": fieldId, "name": prototype.TypeName + "_select" }).down()
                fieldDiv.appendElement({ "tag": "span", "style": "display: table-cell;", "innerHtml": field.substring(2) })
                this.isSelectedFieldBtns[prototype.TypeName + "_select"].push(isSelectFieldBtn);
                if (this.isSelectedBtns[prototype.TypeName].element.checked) {
                    isSelectFieldBtn.element.checked = true
                }

                // Field to append in the query.
                isSelectFieldBtn.element.onclick = function (isSelectBtn, typeName, isSelectedBtns) {
                    return function () {
                        isSelectBtn.element.checked = false;
                        var selects = document.getElementsByName(typeName + "_select")
                        for (var i = 0; i < selects.length; i++) {
                            if (selects[i].checked == true) {
                                isSelectBtn.element.checked = true
                                break
                            }
                        }
                        // Now the whole type...
                        var isSelectBtn_ = document.getElementById(typeName.split(".")[0] + "_select")
                        isSelectBtn_.checked = false
                        for (var key in isSelectedBtns) {
                            if (isSelectedBtns[key].element.checked) {
                                isSelectBtn_.checked = true
                                break
                            }
                        }
                    }
                }(isSelectBtn, prototype.TypeName, this.isSelectedBtns)
            }
        }
    }

    // Now i will set the 
    expandBtn.element.onclick = function (shrinkBtn, typeDiv) {
        return function () {
            this.style.display = "none"
            shrinkBtn.element.style.display = "inline"
            typeDiv.element.style.display = "table"
        }
    }(shrinkBtn, typeDiv)

    shrinkBtn.element.onclick = function (expandBtn, typeDiv) {
        return function () {
            this.style.display = "none"
            expandBtn.element.style.display = "inline"
            typeDiv.element.style.display = "none"
        }
    }(expandBtn, typeDiv)
}

/**
 * A search page display they interface to do search.
 * @param {*} parent The parent element where the search page is display.
 */
var SearchPage = function (parent, searchInfo) {

    /** That structure has the information to recreate the search page. */
    this.searchInfo = searchInfo

    /** The panel who will display the search result. */
    this.panel = parent.appendElement({ "tag": "div", "class": "entity admin_table", "style": "display: flex; flex-direction: column; top: 0px; bottom: 0px; left: 0px; right: 0px; padding-bottom: 0px; position: absolute; overflow: hidden;" }).down()

    /** The search input where the key words will be written */
    var searchInputBar = this.panel.appendElement({ "tag": "div", "style": "display: table; flex-basis: 30px; vertical-align: middle; position: relative;" }).down()

    this.searchInput = searchInputBar.appendElement({ "tag": "input", "style": "display: table-cell;margin: 2px; border: 1px solid; vertical-align: middle;" }).down()

    this.searchOptionsBtn = searchInputBar.appendElement({ "tag": "div", "class": "search_btn", "style": "display: table-cell;margin: 2px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-caret-down" }).down()

    this.searchOptionPanel = new SearchOptionsPanel(searchInputBar)

    // Here I will display or hide the search option panel.
    this.searchOptionsBtn.element.onclick = function () {
        var searchOptionPanel = document.getElementById("search_option_panel")
        if (searchOptionPanel.style.display == "") {
            searchOptionPanel.style.display = "table"
        } else {
            searchOptionPanel.style.display = ""
        }
    }

    /** The search button */
    this.searchBtn = searchInputBar.appendElement({ "tag": "div", "class": "search_btn", "style": "display: table-cell;margin: 2px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-search" }).down()

    /** The search result row. */
    this.resultPanel = this.panel.appendElement({ "tag": "div", "style": "flex-grow: 1; position: relative; display: flex;" }).down()

    /** The serch result panel */
    this.searchResultsPage = new SearchResultsPage(this.resultPanel,function(searchPage){
        return function(offset, pageSize){
            var fields = []
            // Now I will append the list of selected fields for each selected types.
            for (var tabId in searchPage.searchOptionPanel.tabs) {
                var tab = searchPage.searchOptionPanel.tabs[tabId]
                for (var btnId in tab.searchOptionPanelDataTypeInfo.isSelectedBtns) {
                    var isSelectedBtn = tab.searchOptionPanelDataTypeInfo.isSelectedBtns[btnId]
                    if (isSelectedBtn.element.checked) {
                        var isSelectedFieldBtns = tab.searchOptionPanelDataTypeInfo.isSelectedFieldBtns[isSelectedBtn.element.id]
                        for (var i = 0; i < isSelectedFieldBtns.length; i++) {
                            if (isSelectedFieldBtns[i].element.checked) {
                                fields.push("X" + isSelectedFieldBtns[i].id + "%:" + isSelectedFieldBtns[i].id)
                            }
                        }
                    }
                }
            }

            // Hide the search option panel
            var searchOptionPanel = document.getElementById("search_option_panel")
            searchOptionPanel.style.display = ""
            var dbpaths = searchPage.searchOptionPanel.getDataStoreList()
            searchPage.search(offset, pageSize, fields, dbpaths)
        }
    }(this))

    /** Now the action. */
    this.searchBtn.element.onclick = function (searchPage) {
        return function () {
            var offset = 0;
            var pageSize = 10;
            var fields = []
            // Now I will append the list of selected fields for each selected types.
            for (var tabId in searchPage.searchOptionPanel.tabs) {
                var tab = searchPage.searchOptionPanel.tabs[tabId]
                for (var btnId in tab.searchOptionPanelDataTypeInfo.isSelectedBtns) {
                    var isSelectedBtn = tab.searchOptionPanelDataTypeInfo.isSelectedBtns[btnId]
                    if (isSelectedBtn.element.checked) {
                        var isSelectedFieldBtns = tab.searchOptionPanelDataTypeInfo.isSelectedFieldBtns[isSelectedBtn.element.id]
                        for (var i = 0; i < isSelectedFieldBtns.length; i++) {
                            if (isSelectedFieldBtns[i].element.checked) {
                                fields.push("X" + isSelectedFieldBtns[i].id + "%:" + isSelectedFieldBtns[i].id)
                            }
                        }
                    }
                }
            }

            // searchPage.resultPanel.removeAllChilds()

            // Hide the search option panel
            var searchOptionPanel = document.getElementById("search_option_panel")
            searchOptionPanel.style.display = ""
            var dbpaths = searchPage.searchOptionPanel.getDataStoreList()
            searchPage.search(offset, pageSize, fields, dbpaths)

        }
    }(this)

    return this
}

/**
 * Fire a search...
 */
SearchPage.prototype.search = function (offset, pageSize, fields, dbpath) {
    // First of a ll I will clear the search panel.
    xapian.search(
        dbpath,
        this.searchInput.element.value,
        fields,
        "en",
        offset,
        pageSize,
        // success callback
        function (results, searchPage) {
            // Keep the page in memory so it can be display latter without server call...
            // new SearchResultsPage(searchPage.resultPanel, searchPage.search)
            searchPage.searchResultsPage.displayResults(results)
        },
        // error callback
        function () {

        }, this)

}

var SearchResultsPage = function (parent, search) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "class": "search_results" })

    // Here I will display the result header...
    this.searchResultsHeader = this.panel.appendElement({ "tag": "div", "class": "search_results_header" }).down()

    this.headerMessage = this.searchResultsHeader.appendElement({ "tag": "span" }).down()

    // Now I will display the list of results...
    this.searchResultsPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content" }).down()
        .appendElement({ "tag": "div", "style": "position: absolute; top: 0px; left: 0px; right: 0px; top: 0px; bottom: 0px;" }).down()
        .appendElement({ "tag": "div", "style": "display: table; border-spacing:2px 2px; width: 100%;" }).down()

    // The footer.
    this.searchResultsNavigation = this.panel.appendElement({ "tag": "div", "class": "search_results_navigation" }).down()

    // The page selector.
    this.pageSelector = new PageSelector(this.searchResultsNavigation, search)

    return this;
}

SearchResultsPage.prototype.displayResults = function (results) {
    // results.offset, results.estimate,  results.pagesize,
    var headerText = "No results found for \"" + results.query + "\""
    if (results.estimate > 0) {
        headerText = "About " + results.estimate + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
    }

    // Set the display message.
    this.headerMessage.element.innerHTML = headerText;
    this.searchResultsPanel.removeAllChilds()
    for (var i = 0; i < results.results.length; i++) {
        new SearchResult(this.searchResultsPanel, results.results[i], results.indexs)
    }

    // Now the navigator part.
    this.pageSelector.setResults(results.offset, results.estimate, results.pagesize)
}



/**
 * That class display a single result.
 * @param {*} parent 
 * @param {*} result 
 */
var SearchResult = function (parent, result, indexs) {
    this.panel = parent;

    // The data contain the information to display to the user.
    var data = result.data

    if (result.data.TYPENAME != undefined) {
        // In that case I will get the entity prototype to get hint how to display the data 
        // to the end user.
        server.entityManager.getEntityPrototype(result.data.TYPENAME, result.data.TYPENAME.split(".")[0],
            // success callback 
            function (prototype, caller) {
                caller.searchResult.displayData(caller.result, indexs, result.terms, prototype, result.snippet)
            },
            // error callback
            function (errObj, caller) {

            }, { "result": result, "searchResult": this })
    }


    return this;
}

SearchResult.prototype.displayData = function (result, indexs, terms, prototype, snippet) {

    // If the entity prototype isn't null I will intialyse the entity
    // from the data received.
    if (prototype != undefined) {
        var entity = eval("new " + prototype.TypeName + "()")
        entity.init(result.data)

        // Here I will display the resuls...
        var title = this.panel.appendElement({ "tag": "div", "style": "display: table;border-spacing:2px 2px; margin-top: 5px;" }).down()
        title.appendElement({ "tag": "div", "class": "search_result_rank", "innerHtml": result.rank.toString() })
        var titles = entity.getTitles()
        if (entity.TYPENAME == "CargoEntities.File") {
            this.displayFileResult(entity, title, indexs, terms, snippet)
        } else {
            this.displayEntityResult(entity, title, indexs, terms, snippet)
        }
    }
}

/**
 * Generic entity search result display.
 * @param {*} entity 
 */
SearchResult.prototype.displayEntityResult = function (entity, title, indexs, terms, snippet) {
    var titles = entity.getTitles()
    for (var i = 0; i < titles.length; i++) {
        title.appendElement({ "tag": "div", "class": "search_result_rank", "innerHtml": titles[i] })
    }
    // Now the search informations.
    var founded = this.panel.appendElement({ "tag": "div", "style": "display: table; border-spacing:2px 2px" }).down()

    // So here I will display the field and that contain <b> </b>
    function appendSnippet(panel, propertie, value) {
        // Thats means a snippet is found.
        panel.appendElement({ "tag": "div", "style": "display: table; padding-left: 20px; border-spacing:2px 5px" }).down()
            .appendElement({ "tag": "div", "style": "display: table-cell", "innerHtml": "<b>" + propertie + ": </b>" })
            .appendElement({ "tag": "div", "style": "display: table-cell", "innerHtml": value })
    }

    for (var propertie in snippet) {
        if (isString(snippet[propertie])) {
            if (snippet[propertie].indexOf("<b>") != -1) {
                // Thats means a snippet is found.
                appendSnippet(this.panel, propertie, snippet[propertie])
            }
        } else if (isArray(snippet[propertie])) {
            for (var i = 0; i < snippet[propertie].length; i++) {
                if (isString(snippet[propertie][i])) {
                    if (snippet[propertie][i].indexOf("<b>") != -1) {
                        // Thats means a snippet is found.
                        appendSnippet(this.panel, propertie, snippet[propertie][i])
                    }
                }
            }
        }
    }
}

/**
 * Display the search result for a file.
 * @param {*} file 
 */
SearchResult.prototype.displayFileResult = function (file, title, indexs, terms, snippet) {
    // Here I will display the file link...
    var filePath = file.M_path + "/" + file.M_name
    var fileLnk = title.appendElement({ "tag": "div", "class": "search_result_rank", "innerHtml": filePath }).down()

    fileLnk.element.onclick = function (file) {
        return function () {
            // Here I will generate file open event...
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
            server.eventHandler.broadcastLocalEvent(evt)
        }
    }(file)

    // So here I will try to find the searched text in the resut.
    var text = decode64(file.M_data)

    // var snippets = this.getSnippets(file.UUID, text, terms, 50)
    var snippets = "<pre>" + snippet.M_data + "</pre>"

    // Now I will display the search result.
    this.panel.appendElement({ "tag": "div", "style": "display: table;padding-left: 20px; border-spacing:2px 5px", "innerHtml": snippets }).down()


    // TODO from the text found in snippet.M_data I will set the index of bold field and set the on click event on it to 
    // go at the line in text...
    // I will get element by name
    /*var foundedSpans = document.getElementsByName(file.UUID)
    for (var i = 0; i < foundedSpans.length; i++) {
        foundedSpans[i].onclick = function (file) {
            return function () {
                var values = this.title.split(",")
                var ln = values[0].trim().split(" ")[1]
                var col = values[1].trim().split(" ")[1] - 1
                // Here I will throw an event to open the file and set it current position at the given
                // position.
                evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file, "coord": { "ln": ln, "col": col } } }
                server.eventHandler.broadcastLocalEvent(evt)
            }
        }(file)
    }*/
}