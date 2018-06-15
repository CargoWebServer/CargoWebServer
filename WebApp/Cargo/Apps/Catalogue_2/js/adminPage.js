var AdminPage = function (parent){
    this.panel = new Element(parent, { "tag": "div", "id": "admin_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px; display: none;"})
    this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs querynav"}).down()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link active show", "innerHtml" : "Fournisseur", "href" : "#supplier-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "supplier-adminContent", "id":"supplier-adminTab"}).up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link", "innerHtml" : "Localisation", "href" : "#localisation-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "localisation-adminContent", "id":"localisation-adminTab"}).up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "innerHtml" : "Fabricant", "href" : "#manufacturer-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "manufacturer-adminContent", "id":"manufacturer-adminTab"}).up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "innerHtml" : "Catégorie", "href" : "#category-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "category-adminContent", "id":"category-adminTab"}).up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "innerHtml" : "Item", "href" : "#item-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "item-adminContent", "id":"item-adminTab"}).up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "innerHtml" : "Paquet", "href" : "#package-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "package-adminContent", "id":"package-adminTab"}).up().up()
    .appendElement({"tag" : "div", "class" : "tab-content"}).down()
    .appendElement({"tag" : "div", "class" : "tab-pane active show result-tab-content","id" : "supplier-adminContent", "role" : "tabpanel", "aria-labelledby" : "supplier-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane result-tab-content","id" : "localisation-adminContent", "role" : "tabpanel", "aria-labelledby" : "localisation-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "manufacturer-adminContent", "role" : "tabpanel", "aria-labelledby" : "manufacturer-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "category-adminContent", "role" : "tabpanel", "aria-labelledby" : "category-adminTab" }).down()
    .appendElement({"tag" : "div", "class" : "col-md-4", "id" : "categoryAdminNavigation"})
    .appendElement({"tag" : "div", "class" : "col-md-8", "id" : "categoryAdminControl"}).up()
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "item-adminContent", "role" : "tabpanel", "aria-labelledby" : "item-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "package-adminContent", "role" : "tabpanel", "aria-labelledby" : "package-adminTab" }).down()
    .appendElement({"tag" : "div", "class" : "col-md-4", "id" : "packageAdminNavigation"})
    .appendElement({"tag" : "div", "class" : "col-md-8", "id" : "packageAdminControl"}).up()
  
    this.adminItemPage = new AdminItemPage(this.panel.getChildById("item-adminContent"));
    
    this.adminSupplierPage = new AdminSupplierPage(this.panel.getChildById("supplier-adminContent"))
    
    this.adminLocalisationPage = new AdminLocalisationPage(this.panel.getChildById("localisation-adminContent"))
  
    this.adminManufacturerPage = new AdminManufacturerPage(this.panel.getChildById("manufacturer-adminContent"))
  
    return this
}

AdminPage.prototype.displayAdminPage = function(){
    document.getElementById("main-container").style.display = "none"
    document.getElementById("admin_page_panel").style.display = ""
    if(document.getElementById("item_search_result_page") != null){
        document.getElementById("item_search_result_page").style.display = "none"
    }
    if(document.getElementById("item_display_page_panel") != null){
        document.getElementById("item_display_page_panel").style.display = "none"
    }
    document.getElementById("order_page_panel").style.display = "none"
    
    if(document.getElementById("order_history_page_panel") != null){
        document.getElementById("order_history_page_panel").style.display = "none"
    }
    document.getElementById("package_display_page_panel").style.display = "none"
}


