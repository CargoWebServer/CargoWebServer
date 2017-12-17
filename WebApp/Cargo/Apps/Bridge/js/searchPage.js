
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
 * A search page display they interface to do search.
 * @param {*} parent The parent element where the search page is display.
 */
var SearchPage = function (parent, searchInfo) {

    /** That structure has the information to recreate the search page. */
    this.searchInfo = searchInfo

    /** The panel who will display the search result. */
    this.panel = parent.appendElement({ "tag": "div", "class": "entity admin_table", "style": "top: 0px; bottom: 0px; left: 0px; right: 0px; position: absolute;" }).down()

    /** The search input where the key words will be written */
    var searchInputBar = this.panel.appendElement({ "tag": "div", "style": "display: table;" }).down()
    this.searchInput = searchInputBar.appendElement({ "tag": "input", "style": "display: table-cell;margin: 2px; border: 1px solid;" }).down()
    this.searchBtn = searchInputBar.appendElement({ "tag": "div", "class": "search_btn", "style": "display: table-cell;margin: 2px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-search" }).down()

    /** Now the action. */
    this.searchBtn.element.onclick = function (searchPage) {
        return function () {
            var offset = 0;
            var pageSize = 10;
            var fields = ["Xcargoentities.file.m_data:data"];
            var dbpath = "/home/dave/Documents/CargoWebServer/WebApp/Cargo/Data/CargoEntities/CargoEntities.glass"

            // Clear previous search
            searchPage.resultsPages = []
            searchPage.resultPanel.removeAllChilds()

            // Set the new search.
            searchPage.search(offset, pageSize, fields, dbpath)
        }
    }(this)

    /** The search result row. */
    this.resultPanel = this.panel.appendElement({ "tag": "div", "style": "display: table-row;" }).down()

    // This will hold the results for the time of navigation.
    this.resultsPages = [];



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
            searchPage.resultsPages[results.offset] = new SearchResultsPage(searchPage.resultPanel, results)

        },
        // error callback
        function () {

        }, this)

}

/**
 * That class contain the reuslt of search.
 */
var SearchResultsPage = function (parent, results) {
    this.panel = parent.appendElement({ "tag": "div", "class": "search_results" }).down()

    // Here I will display the result header...
    this.searchResultsHeader = this.panel.appendElement({ "tag": "div", "class": "search_results_header" }).down()

    // So here I will display a line with general seach info...
    var headerText = "No results found for \"" + results.query + "\""
    if (results.estimate > 0) {
        headerText = "About " + results.estimate + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
    }
    this.searchResultsHeader.appendElement({ "tag": "span", "innerHtml": headerText })

    // Now I will display the list of results...
    this.searchResultsPanel = this.panel.appendElement({ "tag": "div", "style": "display: table; border-spacing:2px 2px;" }).down()

    this.searchResults = []

    for (var i = 0; i < results.results.length; i++) {
        this.searchResults[i] = new SearchResult(this.searchResultsPanel, results.results[i], results.indexs)
    }

    return this;
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
                caller.searchResult.displayData(caller.result, indexs, result.terms, prototype)
            },
            // error callback
            function (errObj, caller) {

            }, { "result": result, "searchResult": this })
    }


    return this;
}

SearchResult.prototype.displayData = function (result, indexs, terms, prototype) {

    console.log(result)

    // If the entity prototype isn't null I will intialyse the entity
    // from the data received.
    if (prototype != undefined) {
        var entity = eval("new " + prototype.TypeName + "()")
        entity.init(result.data)

        // Here I will display the resuls...
        var title = this.panel.appendElement({ "tag": "div", "style": "display: table;border-spacing:2px 2px;  padding-top: 10px;" }).down()
        title.appendElement({ "tag": "div", "class": "search_result_rank", "innerHtml": result.rank.toString() })
        var titles = entity.getTitles()
        if (entity.TYPENAME == "CargoEntities.File") {
            this.displayFileResult(entity, title, indexs, terms)
        } else {
            this.displayEntityResult(entity, title, indexs, terms)
        }
    }
}

/**
 * Generic entity search result display.
 * @param {*} entity 
 */
SearchResult.prototype.displayEntityResult = function (entity, title, indexs, terms) {
    for (var i = 0; i < title.length; i++) {
        title.appendElement({ "tag": "div", "class": "search_result_rank", "innerHtml": titles[i] })
    }
    // Now the search informations.
    var founded = this.panel.appendElement({ "tag": "div", "style": "display: table; border-spacing:2px 2px" }).down()
}

/**
 * Display the search result for a file.
 * @param {*} file 
 */
SearchResult.prototype.displayFileResult = function (file, title, indexs, terms) {
    // Here I will display the file link...
    var filePath = file.M_path + "/" + file.M_name
    var fileLnk = title.appendElement({ "tag": "div", "class": "search_result_rank", "innerHtml": filePath }).down()
    
    fileLnk.element.onclick = function(file){
        return function(){
            // Here I will generate file open event...
            evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
            server.eventHandler.broadcastLocalEvent(evt)
        }
    }(file)

    // So here I will try to find the searched text in the resut.
    var text = decode64(file.M_data)

    var snippets = this.getSnippets(file.UUID, text, terms, 50)

    // Now the search informations.
    var founded = this.panel.appendElement({ "tag": "div", "style": "display: table; border-spacing:2px 5px", "innerHtml": snippets }).down()

    // I will get element by name
    var foundedSpans = document.getElementsByName(file.UUID)
    for (var i = 0; i < foundedSpans.length; i++) {
        foundedSpans[i].onclick = function (file){
            return function () {
                var values = this.title.split(",")
                var ln = values[0].trim().split(" ")[1]
                var col = values[1].trim().split(" ")[1]-1
                // Here I will throw an event to open the file and set it current position at the given
                // position.
                evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file, "coord":{"ln":ln, "col":col} } }
                server.eventHandler.broadcastLocalEvent(evt)
            }
        } (file)
    }
}

/**
 * Return the text arround the found term.
 * @param {*} text 
 * @param {*} terms 
 */
SearchResult.prototype.getSnippets = function (fileUuid, text, terms, size) {
    // First of all I will find the position in text of all term...
    var positions = []

    for (var i = 0; i < terms.length; i++) {
        var term = terms[i]
        // I will remove the prefix if there is some.
        if (term.match(/^[A-Z]{1}/)) {
            term = term.substring(1)
        }
        // Not case sensitve...
        positions_ = text.toLowerCase().indices(term)
        for (var j = 0; j < positions_.length; j++) {
            // start index
            positions.push(positions_[j])
            // end index...
            positions.push(positions_[j] + term.length)
        }
    }

    // sort array by number and not string value.
    function sortNumber(a, b) {
        return a - b;
    }

    // sort the array with all it values.
    positions.sort(sortNumber)

    // Now from the fouded position in text i will generate 
    // a condensate result...
    var offset = size / 2
    var index = 0
    var startLine = 0;
    var index = 0
    var col = 0
    var ln = 0
    var snippet = ""

    for (var i = 0; i < text.length; i++) {
        startLine = i
        col = 0; // col
        var textLine = ""
        var isSnippet = false
        while (text[i] != "\n" && i < text.length) {
            if (positions[index] == i) {
                textLine += "<span name='" + fileUuid + "' title='Ln " + (ln+1) + ", col " + (col+1) + "' class='founded_reusult' style='vertical-align: text-bottom; padding: 0px;'>"
            } else if (positions[index + 1] == i) {
                textLine += "</span>"
                index += 2
                isSnippet = true
            }
            textLine += text[i]
            col++
            i++
        }
        ln++
        if (isSnippet) {
            snippet += "<div style='display: inline; overflow-x: hidden; padding: 0px 0px 10px 10px;'>" + textLine.trim() + "</div></br>"
        }
    }
    return snippet
}