/**
 * Display the item page.
 * @param {*} parent 
 */
var ItemDisplayPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "item_display_page_panel", "class": "item_display" })

    // Now I will display the list of results...
    this.searchResultsPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content" }).down()

    // I will made use of a tab panel to display the results.
    this.tabPanel = new TabPanel(this.searchResultsPanel);

    return this
}
  
/**
 * Display an item information.
 * @param {*} item 
 */
ItemDisplayPage.prototype.displayTabItem = function (item) {
    var footer = document.getElementById("main-page-footer")
    footerText = "Display " + item.M_id + " results"

    footer.innerText = footerText;
    setTimeout(function (footer) {
        return function () {
            footer.innerText = "";
        }
    }(footer), 4000)

    // Display the results.
    document.getElementById("item_context_selector").style.display = "none"
    if (Object.keys(mainPage.searchResultPage.tabPanel.tabs).length > 0) {
        document.getElementById("search_context_selector").style.display = ""
    } else {
        document.getElementById("search_context_selector").style.display = "none"
    }

    document.getElementById("item_display_page_panel").style.display = ""
    document.getElementById("item_search_result_page").style.display = "none"

    // Here I will create the tab for one panel.
    var tab = this.tabPanel.appendTab(item.M_id)
    tab.closeBtn.element.addEventListener("click", function (itemSearchResultPage) {
        return function () {
            if (Object.keys(itemSearchResultPage.tabPanel.tabs).length == 0) {
                document.getElementById("item_context_selector").style.display = "none"
            }
        }
    }(this));
    // Display the tab.
    tab.title.element.innerHTML = item.M_id;
    tab.content.element.style.backgroundColor = "white";
    tab.content.element.style.position = "relative";

    // Now I will create the entity panel for the item...
    var itemPanel = new EntityPanel(tab.content, "CatalogSchema.ItemType", function (entity) {
        return function (panel) {
            panel.setEntity(entity)
        }
    }(item))
}