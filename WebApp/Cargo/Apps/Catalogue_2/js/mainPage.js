function deleteOreder(userId){
   var q = new EntityQuery()
    q.TYPENAME = "Server.EntityQuery"
    q.TypeName = "CatalogSchema.OrderType"
    q.Fields = ["M_userId"]
    q.Query = 'CatalogSchema.OrderType.M_userId=="' + userId + '"'
    
    
    server.entityManager.getEntities("CatalogSchema.OrderType", "CatalogSchema", q, 0,-1, [], true,false, 
    function(index,total,caller){
        
    },
    function(orders,caller){
        var deleteOrders = function(oreders){
            var order = oreders.pop()
            console.log("---> delete order ",order.UUID)
            server.entityManager.removeEntity(order.UUID,
                function(result, oreders){
                    if(oreders.length > 0){
                        deleteOrders(orders)
                    }
                }, 
                function(){
                    
                }, oreders);
        }
        
        deleteOrders(oreders)
    },
    function(errObj,orders){

        
    },{"userId" : userId})
}

/** Text contain in the main page **/
var mainPageText = {
    "en": {
        "quickSearchText": " Quick search",
        "welcomeText" : "Welcome to the catalog!",
         "welcomeTipText1" : "To find an item, type a keyword in the search bar, or use the shortcuts ",
        "welcomeTipText2" : " for a category search and ",
        "welcomeTipText3" : " for a filtered search.",
        "advancedSearchText": " Advanced search",
        "navbarSearchText" : "Search in the catalog...",
        "recentSearchText" : "Recent searches",
        "supplierSearchLnk" : "by supplier",
        "specsSearchLnk" : "by properties",
        "generalSearchLnk" : "générale"
    },
    "fr": {
        "quickSearchText": " Recherche rapide",
        "welcomeText" : "Bienvenue dans le catalogue!",
        "welcomeTipText1" : "Pour trouver un item, tapez un mot-clé dans la barre de recherche ou utilisez les raccourcis ",
        "welcomeTipText2" : " pour une recherche par catégorie et ",
        "welcomeTipText3" : " pour une recherche filtrée.",
        "advancedSearchText" : " Recherche avancée",
        "navbarSearchText" : "Recherche dans le catalogue...",
        "recentSearchText" : "Recherches récentes",
        "supplierSearchLnk" : "par fournisseur",
        "specsSearchLnk" : "par propriétés",
        "generalSearchLnk" : "générale"
    },
    
    "es" : {
        "quickSearchText" : " Búsqueda rápida",
        "welcomeText" : "¡Bienvenido al catálogo!",
        "advancedSearchText" : " Búsqueda avanzada",
        "welcomeTipText1" : "Para encontrar un artículo, escriben una palabra clave en la barra de búsqueda o usen las atajos ",
        "welcomeTipText2" : " para una búsqueda de categoría y ",
        "welcomeTipText3" : " para una búsqueda filtrada.",
        "navbarSearchText" : "Buscar en el catálogo ...",
        "recentSearchText" : "Búsquedas recientes",
        "supplierSearchLnk" : "",
        "specsSearchLnk" : "",
        "generalSearchLnk" : ""
    }
};

// Set the text info.
server.languageManager.appendLanguageInfo(mainPageText);


/** Contain the main page layout **/
var MainPage = function () {
       
    //deleteOreder("mtmx7184")
    //deleteOreder("mm006819")
    
    // Set special search context to general (item's).
    this.searchContext = "generalSearchLnk";
    
    // The panel.
    this.panel = new Element(document.body, { "tag": "div", "id": "main-page" });
    
    // Building the navbar with all the content inside
    this.panel.appendElement({"tag" : "nav", "class" : "navbar navbar-expand-md navbar-fixed-top navbar-dark bg-dark", "style" : "z-index : 10;", "id" : "top-navigation"}).down()
    .appendElement({"tag" : "div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "div", "class" : "navbar-header"}).down()
 
    .appendElement({"tag" : "a", "class" : "navbar-brand", "onclick" : "goToHome()", "style" : "cursor : pointer;"}).down()
    .appendElement({"tag" : "img", "src" : "image/safran.png","style" : "width : 24px;", "class" : "nav-item"}).up()

    //.appendElement({"tag" : "a", "class" : "navbar-brand navbar-text", "href" : "#","style" : "font-family: 'Questrial', sans-serif;", "innerHtml" : "Catalogue"})
    .appendElement({"tag" : "button" ,"class" : "btn btn-default text-light", "style" :"font-size : 0.8rem;cursor:pointer;", "onclick" : "toggleNav()"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-navicon"})
    .appendElement({"tag" : "span", "id" : "quickSearchText"}).up().up()
    .appendElement({"tag" : "button", "class" : "navbar-toggler", "type" : "button", "data-toggle": "collapse", "data-target" : "#collapsibleNavbar"}).down()
    .appendElement({"tag" : "span", "class" : "navbar-toggler-icon"}).up()
    .appendElement({"tag" : "div", "class" : "collapse navbar-collapse", "id" : "collapsibleNavbar"}).down()
    
    .appendElement({"tag" : "div", "class": "ml-auto", "id" : "searchbardiv"}).down()
    .appendElement({"tag" : "div", "class" : "navbar-form navbar-center", "id" : "searchbar"}).down()
    .appendElement({"tag" : "div", "class" : "input-group"}).down()
    .appendElement({"tag" : "input", "type" : "text", "class" : "form-control","style" : "font-family: 'Questrial', sans-serif;","id" : "navbarSearchText"})
    .appendElement({"tag" : "div", "class" : "input-group-btn"}).down()

    .appendElement({"tag" : "button", "class" : "btn btn-default text-light", "data-toggle" : "dropdown","type" : "button", "aria-expanded" : "false", "role": "menu", "id" : "advancedSearchButtonNavbar"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-caret-down", "style":"padding-right: 5px;"})
    .appendElement({"tag" : "span" , "id" : "advancedSearchText"}).up()
    .appendElement({"tag" : "div", "class" : "dropdown-menu bg-dark text-light border border-secondary", "id" : "advancedSearchCollapse"}).down()
    .appendElement({"tag":"a", "class":"dropdown-item bg-dark", "id":"generalSearchLnk"})
    .appendElement({"tag":"a", "class":"dropdown-item bg-dark", "id":"supplierSearchLnk"})
    .appendElement({"tag":"a", "class":"dropdown-item bg-dark", "id":"specsSearchLnk"}).up()
    

    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-default","id" :"searchButton"}).down()
    .appendElement({"tag" : "i", "class" :"fa fa-search"}).up().up().up().up().up()
    .appendElement({"tag" : "ul", "class" : "nav navbar-nav ml-auto"}).down()
    .appendElement({"tag" : "li", "class" : "btn btn-outline-secondary","style" : "display:none;", "id" : "orderBtn"}).down()
    .appendElement({"tag" : "span", "id" : "cartCount"})
    .appendElement({"tag" : "i", "class" : "fa fa-shopping-cart", "style" : "margin-left:5px;"}).up()
    .appendElement({"tag" : "li", "id":"sessionPanelDiv", "class" : "navItem"})
    .appendElement({"tag" : "li", "class": "nav-item"}).down()
    .appendElement({"tag" : "div", "id" : "languageSelector", "class" : "navItem"}).up();
    
    new SessionPanel(this.panel.getChildById("sessionPanelDiv"))
     
    new LanguageSelector( this.panel.getChildById("languageSelector"),
        [{ "flag": "ca", "name": "français", "id": "fr" }, { "flag": "us", "name": "english", "id": "en" }, { "flag": "es", "name": "español", "id" : "es" }],
        function (language) {
            // set the language.
            server.languageManager.setLanguage(language.id)
        }
    );
    
    this.panel.appendElement({"tag" : "div", "class" : "row contentrow", "style" : "height : calc(100vh - 56px); width : 100%;"}).down()
    .appendElement({"tag" : "div", "class" : "sidenav", "id" : "mySideNav", "style" : "background-color : #343a40; position:fixed;top:50px; text-align:center;"}).down()
    .appendElement({"tag" : "a" , "href" : "#", "onclick" : "toggleNav()","class" : "closebtn", "innerHtml" : "&times", "style":"position:absolute;"})
    .appendElement({"tag" : "div", "class" : "btn-group", "role" : "group", "style": "margin-top:45px; margin-bottom: 20px;"}).down()
    .appendElement({"tag" : "button", "class" : " d-flex flex-column align-items-center btn btn-outline-secondary justify-content-center", "id" : "categoryQuickSearchButton"}).down()
    .appendElement({"tag" : "span", "class" : "fa fa-book"})
    .appendElement({"tag" : "span", "innerHtml" : "Catégorie"}).up()
    .appendElement({"tag" : "button", "class" : " d-flex flex-column align-items-center btn btn-outline-secondary justify-content-center", "id"  :"supplierQuickSearchButton"}).down()
    .appendElement({"tag" : "span", "class" : "fa fa-industry"})
    .appendElement({"tag" : "span", "innerHtml" : "Fournisseur"}).up()
    .appendElement({"tag" : "button", "class" : " d-flex flex-column align-items-center btn btn-outline-secondary justify-content-center", "id" : "locationQuickSearchButton"}).down()
    .appendElement({"tag" : "i", "class" : "material-icons", "innerHtml" : "place", "style" : "font-size:18px;"})
    .appendElement({"tag" : "span", "innerHtml" : "Localisation"}).up().up()
    .appendElement({"tag" : "div", "id" : "quickSearchCategoryContent", "style" : "color:#868e96"}).down()
    .appendElement({"tag" : "div", "class" : "swoop", "style" : "top: 25%;left:50%;position: absolute;"}).up()
    .appendElement({"tag" : "div", "id" : "quickSearchSupplierContent", "style" : "color:#868e96"}).down()
    .appendElement({"tag" : "div", "class" : "swoop", "style" : "top: 25%;left:50%;position: absolute;"}).up()
    .appendElement({"tag" : "div", "id" : "quickSearchLocationContent", "style" : "color:#868e96"}).down()
    .appendElement({"tag" : "div", "class" : "swoop", "style" : "top: 25%;left:50%;position: absolute;"}).up().up()
    .appendElement({"tag" : "span", "id" : "main", "class" : "col", "style" : "overflow-y : scroll;"}).down()
    .appendElement({"tag" : "div", "class" : "container", "style" : "height:100vh; padding-top : 20px;", "id" : "main-container"}).down()
    .appendElement({"tag" : "div","class" : "jumbotron bg-dark text-light border border-dark", "style" : "margin: 20px; padding : 2rem 1rem;text-align:center;"}).down()
    .appendElement({"tag" : "h1", "class" : "display-5", "id" : "welcomeText", "style" : "text-align : center;"})
    .appendElement({"tag" : "hr", "class" : "my-4 dark"})
    .appendElement({"tag" : "p", "class" : "lead"}).down()
    .appendElement({"tag" : "span","id":"welcomeTipText1"})
    .appendElement({"tag" : "button" ,"class": "btn btn-outline-light disabled", "style":"opacity : 1; font-size : 0.8rem;cursor:default;"}).down()
    .appendElement({"tag" : "i", "class" :"fa fa-navicon"})
    .appendElement({"tag" : "span","id":"quickSearchText"}).up()
    .appendElement({"tag" : "span","id":"welcomeTipText2"})
    .appendElement({"tag":"button", "class": "btn btn-outline-light disabled", "style":"opacity :1;font-size : 0.8rem;cursor:default;"}).down()
    .appendElement({"tag":"i", "class":"fa fa-angle-double-down"})
    .appendElement({"tag" :"span","id":"advancedSearchText"}).up()
    .appendElement({"tag" : "span","id":"welcomeTipText3"}).up()
    .appendElement({"tag" : "p", "style" : "text-align:center;"}).down()
    .appendElement({"tag" : "hr", "class" : "my-4"}).up()
    .appendElement({"tag" : "button", "class" : "btn btn-light", "onclick" : "goToSearches()"}).down()
    .appendElement({"tag"  : "span", "innerHtml" : "Aller aux recherches "})
    .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up().up()
    .appendElement({"tag" : "div","class" : "row", "style" : "justify-content : space-between;"}).down()
    .appendElement({"tag" : "div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "div", "class" : "row contentrow ", "style" : "justify-content:space-between;"}).down()
    .appendElement({"tag" : "div","class" :"col-md-5 card bg-light mb-3 border border-dark align-items-center", "style" : "margin:20px; padding:0;"}).down()
    .appendElement({"tag" : "div", "class" : "card-header bg-dark text-light", "style" : "width:100%;"}).down()
    .appendElement({"tag" : "h1", "class" : "card-title smallertitle", "id" : "recentSearchText", "style" : "text-align:center;"}).up()
    .appendElement({"tag" : "div", "class" : "card-body", "style" : "width:100%;padding:0;"}).down()
    .appendElement({"tag" : "ul", "class" :"list-group list-group-flush"}).down()
    .appendElement({"tag" : "li", "class" : "list-group-item"})
    .appendElement({"tag" : "li", "class" : "list-group-item"})
    .appendElement({"tag" : "li", "class" : "list-group-item"})
    .appendElement({"tag" : "li", "class" : "list-group-item"}).up().up()
    .appendElement({"tag" : "hr", "class" : "my-4"}).up()
    .appendElement({"tag" : "div","class" : "col card border border-dark align-items-center", "style" : "margin:20px;"}).down()
    .appendElement({"tag" : "h1", "class" : "card-body display-5 smallertitle"})
    .appendElement({"tag" : "hr", "class" : "my-4"}).up().up().up().up().up()
    .appendElement({"tag" : "div", "id" : "notificationDiv" , "style" : "z-index : 100;position:fixed; width: 275px; right:2%;top:10%;"}).down()
    .up().up()
     toggleNav();
    this.contentPanel = this.panel.getChildById("main")
    this.personalPage = null
    this.searchResultPage = new SearchResultPage(this.contentPanel)
    this.itemDisplayPage = new ItemDisplayPage(this.contentPanel)
    this.packageDisplayPage = new PackageDisplayPage(this.contentPanel)
    this.orderPage = new OrderPage(this.contentPanel)
    this.welcomePage = this.panel.getChildById("main-container")
    
    catalogMessageHub.attach(this, welcomeEvent, function (evt, mainPage) {
        if(server.sessionId == evt.dataMap.welcomeEvent.sessionId){
            
            var user = evt.dataMap.welcomeEvent.user
            mainPage.showNotification("primary", "Bienvenue "+ user.M_firstName + " !",4000)
            mainPage.panel.getChildById("orderBtn").element.style.display = "flex"
            mainPage.panel.getChildById("orderBtn").element.style["align-items"] = "center"
            
        	var q = new EntityQuery()
            q.TypeName = "CatalogSchema.OrderType"
            q.Fields = ["M_userId"]
            q.Query = 'CatalogSchema.OrderType.M_userId=="' + user.M_id + '"'
            
            server.entityManager.getEntities("CatalogSchema.OrderType", "CatalogSchema", q, 0,-1, [], true,false, 
            function(index,total,caller){
                
            },
            function(orders,caller){
                // If there is an open order I will set it as order to display.
                // otherwize I will create an open a new one.
                var openOrderIndex = -1;
                for(var i=0; i < orders.length; i++){
                    if(orders[i].M_status.M_valueOf == "Open"){
                        openOrderIndex = i;
                        break;
                    }
                }
                
                if(orders.length == 0 || openOrderIndex == -1){
                    var order = new CatalogSchema.OrderType()
                    order.M_userId = caller.userId
                    order.M_creationDate = new Date()
                    order.M_status = new CatalogSchema.OrderStatus()
                    order.M_status.M_valueOf = "Open";
                    
                    server.entityManager.createEntity(catalog.UUID, "M_orders", order, 
                        function (result,caller){
                            caller.orderPage.setOrder(result)
                        },function(errObj,caller){
                            
                        },caller)
                }else{
                    
                    
                    caller.orderPage.setOrder(orders[openOrderIndex])
                }
                
                // I will set the personnal page with the user command.
                mainPage.personalPage = new PersonalPage(caller.userId, orders)
            },
            function(errObj,caller){

                
            },{"userId" : user.M_id, "orderPage" : mainPage.orderPage})
           
        }else{
            var user = evt.dataMap.welcomeEvent.user
            mainPage.showNotification("primary", user.M_firstName + " s'est connecté au catalogue!",4000)

        }
        
    })
    
    
    catalogMessageHub.attach(this, modifiedOrderEvent, function (evt, mainPage) {
        
        
    })
    
    
    this.panel.getChildById("main").element.onclick = function () {
       closeNav()
    }
    
    
    this.panel.getChildById("searchButton").element.onclick = function () {
        try{
            closeNav()
            mainPage.welcomePage.element.style.display = "none"
            }
        catch (error){
            
        }
        spinner.panel.element.style.display = "";
        
        // Here I will display the seach results.
        var searchInput = mainPage.panel.getChildById("navbarSearchText").element
        mainPage.searchItems(searchInput.value)
        searchInput.setSelectionRange(0, searchInput.value.length)
    }
    
    this.panel.getChildById("navbarSearchText").element.onkeyup = function (evt) {
        if (evt.keyCode == 13) {
            try{
                closeNav()
                mainPage.welcomePage.element.style.display = "none"
            }
            catch (error){
            
            }
            spinner.panel.element.style.display = "";
            // Here I will display the seach results.
            var searchInput = mainPage.panel.getChildById("navbarSearchText").element
            mainPage.searchItems(searchInput.value)
            searchInput.setSelectionRange(0, searchInput.value.length)
        }
    }
    
    this.panel.getChildById("categoryQuickSearchButton").element.onclick = function(panel){
        return function(){
             panel.getChildById("quickSearchCategoryContent").element.style.display = ""
            panel.getChildById("quickSearchLocationContent").element.style.display = "none"
            panel.getChildById("quickSearchSupplierContent").element.style.display = "none"
        }
        
    }(this.panel)
    this.panel.getChildById("supplierQuickSearchButton").element.onclick = function (panel) {
        return function(){
            panel.getChildById("quickSearchSupplierContent").element.style.display = ""
            panel.getChildById("quickSearchCategoryContent").element.style.display = "none"
            panel.getChildById("quickSearchLocationContent").element.style.display = "none"
            
            
        	
        }
    }(this.panel)
    this.panel.getChildById("locationQuickSearchButton").element.onclick = function (panel) {
         return function(){
            panel.getChildById("quickSearchLocationContent").element.style.display = ""
            panel.getChildById("quickSearchSupplierContent").element.style.display = "none"
            panel.getChildById("quickSearchCategoryContent").element.style.display = "none"
              
                      
           
            }
    }(this.panel)
    
    this.panel.getChildById("orderBtn").element.onclick = function (){
        mainPage.orderPage.displayOrder()
    }
    
    
    // Advance search.
    var supplierLnk = this.panel.getChildById("supplierSearchLnk");
    var specsLnk = this.panel.getChildById("specsSearchLnk")
    var generalSearchLnk = this.panel.getChildById("generalSearchLnk")
    generalSearchLnk.element.onclick = supplierLnk.element.onclick = specsLnk.element.onclick = function(contentPanel, advancedSearchText){
        return function(){
            mainPage.searchContext = this.id
            if(this.id == "generalSearchLnk"){
                server.languageManager.refresh()
            }else{
                advancedSearchText.element.innerHTML = this.innerHTML
            }
            document.getElementById("navbarSearchText").focus()
        }
    }(this.contentPanel, this.panel.getChildById("advancedSearchText"))
    
    
    loadQuickSearchData(this.panel)

    this.itemDisplayPage.panel.element.style.display = "none";
    this.packageDisplayPage.panel.element.style.display = "none";
    this.orderPage.panel.element.style.display = "none"
    return this;
}

function goToSearches(){
    if(document.getElementById("item_search_result_page") !=null){
        document.getElementById("item_search_result_page").style.display = ""
        document.getElementById("main-container").style.display= "none"
    }
}

function goToBasket(){
   mainPage.orderPage.displayOrder()
}

function loadQuickSearchData(panel){
    panel.getChildById("quickSearchSupplierContent").element.style.display = "none"
    panel.getChildById("quickSearchCategoryContent").element.style.display = "none"
    panel.getChildById("quickSearchLocationContent").element.style.display = "none"
    
   
            
    server.entityManager.getEntities("CatalogSchema.CategoryType", "CatalogSchema", null, 0,-1,[],true,true, 
        function(index,total, caller){
            
        },
        function(categories,caller){
            
            if(caller.panel.getChildById(categories[0].M_id) == null){
                caller.panel.getChildById("quickSearchCategoryContent").element.innerHTML = ""
                categories.sort(function(a,b){
                    if(a.M_name < b.M_name){
                        return -1
                    }
                    if(a.M_name > b.M_name){
                        return 1
                    }
                    return 0
                })
                

                for(var i = 0; i<categories.length; i++){
                    var id = categories[i].getEscapedUuid();
                    caller.panel.getChildById("quickSearchCategoryContent").appendElement({"tag" : "a", "class" : "btn btn-outline-secondary", "data-toggle" : "collapse", "href" : "#" + id, "role" :"button", "aria-expanded" : "false", "aria-controls" :id, "innerHtml" : categories[i].M_id, "style" : "margin-right:10px;margin-left:10px;margin-top: 15px;margin-bottom:15px;border-radius: .5rem;"})
                    var typeNamesDiv = caller.panel.getChildById("quickSearchCategoryContent").appendElement({"tag" : "div", "class" : "collapse", "id" : id}).down()
                    
                    // Here I will apprend the list of item that contain that category...
                    getTypeNames = function (index, items, typeNames, callback ) {
                        
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
                                        if(typeName.trim().length == 0) {
                                             typeName = "generic"
                                        }
                                        
                                        if (caller.typeNames[typeName] == undefined) {
                                            caller.typeNames[typeName] = []
                                        }
                                        caller.typeNames[typeName].push({ "ovmm": results[0][0][1].trim(), "uuid": caller.items[caller.index] })
                                    }
                                }
                                if (caller.index < caller.items.length) {
                                    getTypeNames(caller.index + 1, caller.items, caller.typeNames, caller.callback)
                                } else {
                                    // Render the array here.
                                    caller.callback(caller.typeNames)
                                }
                            },
                            // progress callback
                            function (index, total, caller) {
                
                            },
                            // error callback
                            function (errObj, caller) {
                
                            },
                            { "index": index, "items": items, "typeNames": typeNames, "callback":callback })
                    }
                        
                    getTypeNames(0, categories[i].M_items, {},
                        function(typeNamesDiv){``
                            return function(typeNames){
                                var keys = Object.keys(typeNames).sort()
                               
                                for(var i=0; i < keys.length; i++){
                                    // Here I will append the typeName inside the div
                                    if(keys[i].length > 0){
                                        var lnk = typeNamesDiv.appendElement({"tag":"a", "class":"btn-outline-secondary sideNavCategory", "innerHtml":keys[i] + " (" + typeNames[keys[i]].length + ")"}).down()
                                        lnk.element.onclick = function(items, typeName){
                                            return function(){
                                                console.log("items: ", items)
                                                var results = {}
                                                results.results = []
                                                results.estimate = items.length
                                                
                                                for(var i=0; i < items.length; i++){
                                                    results.results.push({"data" : {"UUID":items[i].uuid, "TYPENAME":"CatalogSchema.ItemType"}})
                                                }
                                                mainPage.searchContext = "generalSearchLnk";
                                                mainPage.welcomePage.element.style.display = "none"
                                                mainPage.searchResultPage.displayResults(results,typeName, "generalSearchLnk")
                                            }
                                        }(typeNames[keys[i]], keys[i])
                                        
                                    }
                                }
                            }
                        }(typeNamesDiv))
                }
            }
           
        },
        function(){
        },{"panel" : panel})
   
    server.entityManager.getEntities("CatalogSchema.SupplierType", "CatalogSchema", null, 0,-1,[],true,true, 
        function(index,total, caller){
            
        },
        function(suppliers,caller){
            if(caller.panel.getChildById(suppliers[0].M_id) == null){
                caller.panel.getChildById("quickSearchSupplierContent").element.innerHTML = ""
            var orderedSuppliers = {}
            suppliers.sort(function(a,b){
                if(a.M_name < b.M_name){
                    return -1
                }
                if(a.M_name > b.M_name){
                    return 1
                }
                return 0
            })
            
            
            
            for(var i = 0; i<suppliers.length; i++){
                if(isNumeric(suppliers[i].M_name[0])){
                    if(orderedSuppliers["numeric"] === undefined){
                        orderedSuppliers["numeric"] = []
                    }
                    orderedSuppliers["numeric"].push(suppliers[i])
                    //console.log(suppliers[i].M_name)
                }else{
                    if(orderedSuppliers[suppliers[i].M_name[0]] === undefined){
                        orderedSuppliers[suppliers[i].M_name[0]] = []
                    }
                    orderedSuppliers[suppliers[i].M_name[0]].push(suppliers[i])
                }
                
            }
            
            caller.panel.getChildById("quickSearchSupplierContent").appendElement({"tag" : "a", "class" : "btn btn-outline-secondary", "data-toggle" : "collapse", "href" : "#numeric", "role" :"button", "aria-expanded" : "false", "aria-controls" :"numeric", "innerHtml" : "#", "style" : "margin-right:10px;margin-left:10px;margin-top: 15px;margin-bottom:15px;border-radius: 0.5rem;"})
            caller.panel.getChildById("quickSearchSupplierContent").appendElement({"tag" : "div", "class" : "collapse", "id" : "numeric"})
            function displaySupplier(supplier){
                
                function getItems(uuids, items, callback, query){
                    var uuid = uuids.pop()
                    if(uuid != undefined){
                        server.entityManager.getEntityByUuid(uuid, false, 
                            function(item, caller){
                                caller.items.push({"data":item})
                                if(uuids.length == 0){
                                    mainPage.searchContext = "supplierSearchLnk";
                                    mainPage.welcomePage.element.style.display = "none"
                                    mainPage.searchResultPage.displayResults({"results":caller.items, "estimate":caller.items.length}, caller.query, "supplierSearchLnk") 
                                }else{
                                    caller.callback(caller.uuids, caller.items, caller.callback, caller.query)
                                }
                            },
                            function(errObj){
                                
                            }, {"uuids":uuids, "callback":callback, "items":items, "query":query})
                    }
                    
                }
                
                getItems(supplier.M_items, [], getItems, supplier.M_id)
            }
            
            for(var i = 0; i<orderedSuppliers["numeric"].length; i++){
               var lnk = caller.panel.getChildById("numeric").appendElement({"tag" : "a","class" :"btn-outline-secondary sideNavCategory", "innerHtml" : orderedSuppliers["numeric"][i].M_name, "style" : "padding:8px; font-size:18px;"}).down()
               lnk.element.onclick = function(supplier){
                    return function(){
                        displaySupplier(supplier)
                    }
                }(orderedSuppliers["numeric"][i])
            }
    
            for(var key in orderedSuppliers){
                 if(key != "numeric"){
                     caller.panel.getChildById("quickSearchSupplierContent").appendElement({"tag" : "a", "class" : "btn btn-outline-secondary", "data-toggle" : "collapse", "href" : "#" + key, "role" :"button", "aria-expanded" : "false", "aria-controls" :key, "innerHtml" : key, "style" : "margin-right:10px;margin-left:10px;margin-top: 15px;margin-bottom:15px;border-radius: .5rem;"})
                    caller.panel.getChildById("quickSearchSupplierContent").appendElement({"tag" : "div", "class" : "collapse", "id" : key})
                     for(var i = 0; i<orderedSuppliers[key].length; i++){
                        var lnk = caller.panel.getChildById(key).appendElement({"tag" : "a","class" : "btn-outline-secondary sideNavCategory", "innerHtml" : orderedSuppliers[key][i].M_name, "style" : "padding:8px; font-size:18px;"}).down()
                        lnk.element.onclick = function(supplier){
                            return function(){
                                displaySupplier(supplier)
                            }
                        }(orderedSuppliers[key][i])
                     }
                 }
            }
            }
            
           
        },
        function(){
        }, {"panel" : panel})
    
 // The query.
	var q = new EntityQuery()
    q.TypeName = "CatalogSchema.LocalisationType"
    q.Fields = ["M_parent"]
    q.Query = 'CatalogSchema.LocalisationType.M_parent==null '

    server.entityManager.getEntities("CatalogSchema.LocalisationType", "CatalogSchema", q, 0,-1,[],true,true, 
        function(index,total, caller){
            
        },
        function(locations,caller){
            if(panel.getChildById(locations[0].M_id) == null){
                caller.panel.getChildById("quickSearchLocationContent").element.innerHTML = ""  
                for(var i = 0; i<locations.length; i++){
                    if(locations[i].M_name != ""){
                            caller.panel.getChildById("quickSearchLocationContent").appendElement({"tag" : "a", "class" : " btn-outline-secondary sideNavCategory", "data-toggle" : "collapse", "href" : "#" + locations[i].M_id, "role" :"button", "aria-expanded" : "false", "aria-controls" :locations[i].M_id, "innerHtml" : locations[i].M_name, "style" : "padding:8px; font-size:18px;", "id": locations[i].M_id + "-trigger"})
                        }else{
                           caller.panel.getChildById("quickSearchLocationContent").appendElement({"tag" : "a", "class" : " btn-outline-secondary sideNavCategory", "data-toggle" : "collapse", "href" : "#" + locations[i].M_id, "role" :"button", "aria-expanded" : "false", "aria-controls" :locations[i].M_id, "innerHtml" : locations[i].M_id, "style" : "padding:8px; font-size:18px;", "id" : locations[i].M_id + "-trigger"})
                        }
                        
                        caller.panel.getChildById("quickSearchLocationContent").appendElement({"tag" : "div", "class" : "collapse", "id" : locations[i].M_id, "style" : "padding: 0px;margin-left: 10px;margin-right:10px;border: 1px solid transparent;border-color: #868e96;border-radius: .25rem;"}).down()
                        
                        caller.panel.getChildById(locations[i].M_id+"-trigger").element.onclick = function(panel, location){
                               return function(evt){
                                   evt.stopPropagation();
                                   appendLocation(panel, location);
                                   
                               }
                           }(caller.panel, locations[i])
                    
                }    
                   
            }
        },
        function(){
        },{"panel" : panel})
    
}


function appendLocation(panel, parent){
    panel.getChildById(parent.M_id).element.innerHTML = ""
    for(var i = 0; i < parent.M_subLocalisations.length; i++){
        server.entityManager.getEntityByUuid(parent.M_subLocalisations[i], false, 
        function(result,caller){
            caller.panel.getChildById(parent.M_id).appendElement({"tag" : "a","class" : "btn-outline-secondary sideNavCategory", "innerHtml" : result.M_name, "style" : "padding:8px; font-size:18px;", "id" : result.M_id + "-trigger", "role" : "button", "href" : "#" + result.M_id, "data-toggle" : "collapse"})
            if(result.M_subLocalisations.length > 0){
                caller.panel.getChildById(parent.M_id).appendElement({"tag" : "div", "class" : "collapse", "id" : result.M_id, "style" : "padding: 0px;margin-left: 10px;margin-right:10px;border: 1px solid transparent;border-color: #868e96;border-radius: .25rem;"})
            }
            caller.panel.getChildById(result.M_id+"-trigger").element.onclick = function(panel, location){
                return function(){
                    if(result.M_subLocalisations.length > 0){
                         appendLocation(panel, location)
                    }
                }
            }(caller.panel, result)
        }, function () {}, {"parent" : parent, "panel": panel})
    }
}

/**
 * The context can be general, supplier or localisation.
 */
MainPage.prototype.searchItems = function (keyword) {
    // Set specific fields here.
    var fields = []
    // First of a ll I will clear the search panel.
    xapian.search(
        dbpaths,
        keyword.toUpperCase(),
        fields,
        "en",
        0,
        1000,
        // success callback
        function (results, caller) {
            if(results != null){
                mainPage.welcomePage.element.style.display = "none"
                mainPage.searchResultPage.displayResults(results, caller.keyword, mainPage.searchContext );
            }else{
                // In that case the connection is lost so I will reconnect.
            }
        },
        // error callback
        function () {

        }, {"keyword":keyword })
}

function goToHome(){
    document.getElementById("main-container").style.display = ""
    if(document.getElementById("item_search_result_page") != null){
        document.getElementById("item_search_result_page").style.display = "none"
    }
    if(document.getElementById("item_display_page_panel") != null){
        document.getElementById("item_display_page_panel").style.display = "none"
    }
    document.getElementById("order_page_panel").style.display = "none"
}

MainPage.prototype.showNotification = function (alertType, message,delay){
    var id = randomUUID();
    var msg =  mainPage.panel.getChildById("notificationDiv").appendElement({"tag" : "div", "class" : "alert alert-"+alertType+" alert-dismissible", "role" : "alert"}).down()
    msg.appendElement({"tag" : "p","class" : "mb-0", "innerHtml" : message})
    .appendElement({"tag" : "button","id" : id, "class" : "close", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;"}).up().up()
    msg.element.style.opacity = 0;
    //msg.element.style.transform = "translateX("+msg.element.offsetWidth+"px)"
    msg.element.style.left = msg.element.offsetWidth + "px";
    var keyframe = "100% {left:0px; opacity:1;}"
    msg.animate(keyframe, .75, 
        function(msg){
            return function(){
                 msg.element.style.opacity = 1;
                 msg.element.style.left = "0px";
               // msg.element.parentNode.removeChild(msg.element);
            }
    }(msg))
    msg.getChildById(id).element.onclick = function(msg){
        return function(){
            var keyframe = "100% {transform: translateX("+msg.element.offsetWidth+"px); opacity:0;}"
            msg.animate(keyframe, .75, 
            function(msg){
                return function(){
                    msg.element.parentNode.removeChild(msg.element);
                }
            }(msg))
        }
    }(msg)
    if(delay != undefined){
        setTimeout(function(closebtn) {
            return function(){
                closebtn.element.click();
            }
        }(msg.getChildById(id)), delay);
    }
}



function toggleNav() {
    if(!document.getElementById("mySideNav").classList.contains("closedNav")){
        document.getElementById("mySideNav").classList.remove("openedNav");
        document.getElementById("mySideNav").classList.add("closedNav");
    }else{
        document.getElementById("mySideNav").classList.remove("closedNav");
        document.getElementById("mySideNav").classList.add("openedNav");
    }
}

function closeNav(){
     document.getElementById("mySideNav").classList.remove("openedNav");
        document.getElementById("mySideNav").classList.add("closedNav");
}

function openNav(){
    document.getElementById("mySideNav").classList.remove("closedNav");
        document.getElementById("mySideNav").classList.add("openedNav");
}
