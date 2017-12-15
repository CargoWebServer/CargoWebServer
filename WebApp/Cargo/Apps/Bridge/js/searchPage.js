
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
    var searchInputBar = this.panel.appendElement({ "tag": "div", "style": "display: table-row; width: 100%" }).down()
    this.searchInput = searchInputBar.appendElement({ "tag": "input", "style": "display: table-cell; border: 1px solid;" }).down()
    this.searchBtn = searchInputBar.appendElement({ "tag": "div", "class": "search_btn", "style": "display: table-cell;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-search" }).down()

    /** Now the action. */
    this.searchBtn.element.onclick = function (searchPage) {
        return function () {
            searchPage.search()
        }
    }(this)

    return this
}

/**
 * Fire a search...
 */
SearchPage.prototype.search = function () {

    // Index csv file the file must exist on the server before that method is call.
    var dbpath = "/home/dave/Documents/CargoWebServer/WebApp/Cargo/Data/CargoEntities/CargoEntities.glass"
    // var dbpath = "C:\\Users\\mm006819\\Documents\\CargoWebServer\\WebApp\\Cargo\\Data\\CargoEntities\\CargoEntities.glass"
    xapian.search(
        dbpath,
        this.searchInput.element.value,
        ["XD:data"],
        "en",
        0,
        10,
        // success callback
        function (result, caller) {
            console.log(result)
        },
        // error callback
        function () {

        }, {})
}